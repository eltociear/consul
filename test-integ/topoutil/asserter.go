// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package topoutil

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"github.com/hashicorp/consul/sdk/testutil/retry"
	libassert "github.com/hashicorp/consul/test/integration/consul-container/libs/assert"
	"github.com/hashicorp/consul/test/integration/consul-container/libs/utils"
	"github.com/hashicorp/consul/testing/deployer/topology"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Asserter is a utility to help in reducing boilerplate in invoking test
// assertions against consul-topology Sprawl components.
//
// The methods should largely take in *topology.Service instances in lieu of
// ip/ports if there is only one port that makes sense for the assertion (such
// as use of the envoy admin port 19000).
//
// If it's up to the test (like picking an upstream) leave port as an argument
// but still take the service and use that to grab the local ip from the
// topology.Node.
type Asserter struct {
	sp SprawlLite
}

// *sprawl.Sprawl satisfies this. We don't need anything else.
type SprawlLite interface {
	HTTPClientForCluster(clusterName string) (*http.Client, error)
	APIClientForNode(clusterName string, nid topology.NodeID, token string) (*api.Client, error)
	APIClientForCluster(clusterName string, token string) (*api.Client, error)
	ResourceServiceClientForCluster(clusterName string) pbresource.ResourceServiceClient
	Topology() *topology.Topology
}

// NewAsserter creates a new assertion helper for the provided sprawl.
func NewAsserter(sp SprawlLite) *Asserter {
	return &Asserter{
		sp: sp,
	}
}

func (a *Asserter) mustGetHTTPClient(t *testing.T, cluster string) *http.Client {
	client, err := a.httpClientFor(cluster)
	require.NoError(t, err)
	return client
}

func (a *Asserter) mustGetAPIClient(t *testing.T, cluster string) *api.Client {
	clu := a.sp.Topology().Clusters[cluster]
	cl, err := a.sp.APIClientForCluster(clu.Name, "")
	require.NoError(t, err)
	return cl
}

// httpClientFor returns a pre-configured http.Client that proxies requests
// through the embedded squid instance in each LAN.
//
// Use this in methods below to magically pick the right proxied http client
// given the home of each node being checked.
func (a *Asserter) httpClientFor(cluster string) (*http.Client, error) {
	client, err := a.sp.HTTPClientForCluster(cluster)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// UpstreamEndpointStatus validates that proxy was configured with provided clusterName in the healthStatus
//
// Exposes libassert.UpstreamEndpointStatus for use against a Sprawl.
//
// NOTE: this doesn't take a port b/c you always want to use the envoy admin port.
func (a *Asserter) UpstreamEndpointStatus(
	t *testing.T,
	service *topology.Service,
	clusterName string,
	healthStatus string,
	count int,
) {
	t.Helper()
	node := service.Node
	ip := node.LocalAddress()
	port := service.EnvoyAdminPort
	addr := fmt.Sprintf("%s:%d", ip, port)

	client := a.mustGetHTTPClient(t, node.Cluster)
	libassert.AssertUpstreamEndpointStatusWithClient(t, client, addr, clusterName, healthStatus, count)
}

// HTTPServiceEchoes verifies that a post to the given ip/port combination
// returns the data in the response body. Optional path can be provided to
// differentiate requests.
//
// Exposes libassert.HTTPServiceEchoes for use against a Sprawl.
//
// NOTE: this takes a port b/c you may want to reach this via your choice of upstream.
func (a *Asserter) HTTPServiceEchoes(
	t *testing.T,
	service *topology.Service,
	port int,
	path string,
) {
	t.Helper()
	require.True(t, port > 0)

	node := service.Node
	ip := node.LocalAddress()
	addr := fmt.Sprintf("%s:%d", ip, port)

	client := a.mustGetHTTPClient(t, node.Cluster)
	libassert.HTTPServiceEchoesWithClient(t, client, addr, path)
}

// HTTPServiceEchoesResHeader verifies that a post to the given ip/port combination
// returns the data in the response body with expected response headers.
// Optional path can be provided to differentiate requests.
//
// Exposes libassert.HTTPServiceEchoes for use against a Sprawl.
//
// NOTE: this takes a port b/c you may want to reach this via your choice of upstream.
func (a *Asserter) HTTPServiceEchoesResHeader(
	t *testing.T,
	service *topology.Service,
	port int,
	path string,
	expectedResHeader map[string]string,
) {
	t.Helper()
	require.True(t, port > 0)

	node := service.Node
	ip := node.LocalAddress()
	addr := fmt.Sprintf("%s:%d", ip, port)

	client := a.mustGetHTTPClient(t, node.Cluster)
	libassert.HTTPServiceEchoesResHeaderWithClient(t, client, addr, path, expectedResHeader)
}

func (a *Asserter) HTTPStatus(
	t *testing.T,
	service *topology.Service,
	port int,
	status int,
) {
	t.Helper()
	require.True(t, port > 0)

	node := service.Node
	ip := node.LocalAddress()
	addr := fmt.Sprintf("%s:%d", ip, port)

	client := a.mustGetHTTPClient(t, node.Cluster)

	url := "http://" + addr

	retry.RunWith(&retry.Timer{Timeout: 30 * time.Second, Wait: 500 * time.Millisecond}, t, func(r *retry.R) {
		resp, err := client.Get(url)
		if err != nil {
			r.Fatalf("could not make request to %q: %v", url, err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != status {
			r.Fatalf("expected status %d, got %d", status, resp.StatusCode)
		}
	})
}

// asserts that the service sid in cluster and exported by peer localPeerName is passing health checks,
func (a *Asserter) HealthyWithPeer(t *testing.T, cluster string, sid topology.ServiceID, peerName string) {
	t.Helper()
	cl := a.mustGetAPIClient(t, cluster)
	retry.RunWith(&retry.Timer{Timeout: time.Minute * 1, Wait: time.Millisecond * 500}, t, func(r *retry.R) {
		svcs, _, err := cl.Health().Service(
			sid.Name,
			"",
			true,
			utils.CompatQueryOpts(&api.QueryOptions{
				Partition: sid.Partition,
				Namespace: sid.Namespace,
				Peer:      peerName,
			}),
		)
		require.NoError(r, err)
		assert.GreaterOrEqual(r, len(svcs), 1)
	})
}

func (a *Asserter) UpstreamEndpointHealthy(t *testing.T, svc *topology.Service, upstream *topology.Upstream) {
	t.Helper()
	node := svc.Node
	ip := node.LocalAddress()
	port := svc.EnvoyAdminPort
	addr := fmt.Sprintf("%s:%d", ip, port)

	client := a.mustGetHTTPClient(t, node.Cluster)
	libassert.AssertUpstreamEndpointStatusWithClient(t,
		client,
		addr,
		// TODO: what is default? namespace? partition?
		fmt.Sprintf("%s.default.%s.external", upstream.ID.Name, upstream.Peer),
		"HEALTHY",
		1,
	)
}

type testingT interface {
	require.TestingT
	Helper()
}

// does a fortio /fetch2 to the given fortio service, targetting the given upstream. Returns
// the body, and response with response.Body already Closed.
//
// We treat 400, 503, and 504s as retryable errors
func (a *Asserter) fortioFetch2Upstream(
	t testingT,
	client *http.Client,
	addr string,
	upstream *topology.Upstream,
	path string,
) (body []byte, res *http.Response) {
	t.Helper()

	url := fmt.Sprintf("http://%s/fortio/fetch2?url=%s", addr,
		url.QueryEscape(fmt.Sprintf("http://localhost:%d/%s", upstream.LocalPort, path)),
	)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	require.NoError(t, err)

	res, err = client.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()
	// not sure when these happen, suspect it's when the mesh gateway in the peer is not yet ready
	require.NotEqual(t, http.StatusServiceUnavailable, res.StatusCode)
	require.NotEqual(t, http.StatusGatewayTimeout, res.StatusCode)
	// not sure when this happens, suspect it's when envoy hasn't configured the local upstream yet
	require.NotEqual(t, http.StatusBadRequest, res.StatusCode)
	body, err = io.ReadAll(res.Body)
	require.NoError(t, err)

	return body, res
}

// uses the /fortio/fetch2 endpoint to do a header echo check against an
// upstream fortio
func (a *Asserter) FortioFetch2HeaderEcho(t *testing.T, fortioSvc *topology.Service, upstream *topology.Upstream) {
	const kPassphrase = "x-passphrase"
	const passphrase = "hello"
	path := (fmt.Sprintf("/?header=%s:%s", kPassphrase, passphrase))

	var (
		node   = fortioSvc.Node
		addr   = fmt.Sprintf("%s:%d", node.LocalAddress(), fortioSvc.PortOrDefault("http"))
		client = a.mustGetHTTPClient(t, node.Cluster)
	)

	retry.RunWith(&retry.Timer{Timeout: 60 * time.Second, Wait: time.Millisecond * 500}, t, func(r *retry.R) {
		_, res := a.fortioFetch2Upstream(r, client, addr, upstream, path)
		require.Equal(r, http.StatusOK, res.StatusCode)
		v := res.Header.Get(kPassphrase)
		require.Equal(r, passphrase, v)
	})
}

// similar to libassert.AssertFortioName,
// uses the /fortio/fetch2 endpoint to hit the debug endpoint on the upstream,
// and assert that the FORTIO_NAME == name
func (a *Asserter) FortioFetch2FortioName(t *testing.T, fortioSvc *topology.Service, upstream *topology.Upstream, clusterName string, sid topology.ServiceID) {
	t.Helper()

	var (
		node   = fortioSvc.Node
		addr   = fmt.Sprintf("%s:%d", node.LocalAddress(), fortioSvc.PortOrDefault("http"))
		client = a.mustGetHTTPClient(t, node.Cluster)
	)

	var fortioNameRE = regexp.MustCompile(("\nFORTIO_NAME=(.+)\n"))
	path := "/debug?env=dump"

	retry.RunWith(&retry.Timer{Timeout: 60 * time.Second, Wait: time.Millisecond * 500}, t, func(r *retry.R) {
		body, res := a.fortioFetch2Upstream(r, client, addr, upstream, path)
		require.Equal(r, http.StatusOK, res.StatusCode)

		// TODO: not sure we should retry these?
		m := fortioNameRE.FindStringSubmatch(string(body))
		require.GreaterOrEqual(r, len(m), 2)
		// TODO: dedupe from NewFortioService
		require.Equal(r, fmt.Sprintf("%s::%s", clusterName, sid.String()), m[1])
	})
}

// CatalogServiceExists is the same as libassert.CatalogServiceExists, except that it uses
// a proxied API client
func (a *Asserter) CatalogServiceExists(t *testing.T, cluster string, svc string, opts *api.QueryOptions) {
	t.Helper()
	cl := a.mustGetAPIClient(t, cluster)
	libassert.CatalogServiceExists(t, cl, svc, opts)
}
