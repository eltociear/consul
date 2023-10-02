// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package topoutil

import (
	"testing"

	"github.com/hashicorp/consul/testing/deployer/topology"
	"github.com/stretchr/testify/require"
)

// DisableNode is a no-op if the node is already disabled.
func DisableNode(t *testing.T, cfg *topology.Config, clusterName string, nid topology.NodeID) *topology.Config {
	nodes := cfg.Cluster(clusterName).Nodes
	var found bool
	for _, n := range nodes {
		if n.ID() == nid {
			found = true
			if n.Disabled {
				return cfg
			}
			t.Logf("disabling node %s in cluster %s", nid.String(), clusterName)
			n.Disabled = true
			break
		}
	}
	require.True(t, found, "expected to find nodeID %q in cluster %q", nid.String(), clusterName)
	return cfg
}

// EnableNode is a no-op if the node is already enabled.
func EnableNode(t *testing.T, cfg *topology.Config, clusterName string, nid topology.NodeID) *topology.Config {
	nodes := cfg.Cluster(clusterName).Nodes
	var found bool
	for _, n := range nodes {
		if n.ID() == nid {
			found = true
			if !n.Disabled {
				return cfg
			}
			t.Logf("enabling node %s in cluster %s", nid.String(), clusterName)
			n.Disabled = false
			break
		}
	}
	require.True(t, found, "expected to find nodeID %q in cluster %q", nid.String(), clusterName)
	return cfg
}
