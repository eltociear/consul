package types

import (
	"testing"

	catalogtesthelpers "github.com/hashicorp/consul/internal/catalog/catalogtest/helpers"
	"github.com/hashicorp/consul/internal/resource"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v2beta1"
)

func TestDestinationsConfigurationACLs(t *testing.T) {
	catalogtesthelpers.RunWorkloadSelectingTypeACLsTests[*pbmesh.DestinationsConfiguration](t, pbmesh.DestinationsConfigurationType,
		func(selector *pbcatalog.WorkloadSelector) *pbmesh.DestinationsConfiguration {
			return &pbmesh.DestinationsConfiguration{Workloads: selector}
		},
		func(registry resource.Registry) {
			RegisterDestinationsConfiguration(registry)
		})
}
