package types

import (
	"testing"

	catalogtesthelpers "github.com/hashicorp/consul/internal/catalog/catalogtest/helpers"
	"github.com/hashicorp/consul/internal/resource"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v2beta1"
)

func TestProxyConfigurationACLs(t *testing.T) {
	catalogtesthelpers.RunWorkloadSelectingTypeACLsTests[*pbmesh.ProxyConfiguration](t, pbmesh.ProxyConfigurationType,
		func(selector *pbcatalog.WorkloadSelector) *pbmesh.ProxyConfiguration {
			return &pbmesh.ProxyConfiguration{Workloads: selector}
		},
		func(registry resource.Registry) {
			RegisterProxyConfiguration(registry)
		})
}
