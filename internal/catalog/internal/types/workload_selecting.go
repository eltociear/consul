package types

import (
	"google.golang.org/protobuf/proto"

	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
)

// WorkloadSelecting denotes a resource type that uses workload selectors.
type WorkloadSelecting interface {
	proto.Message
	GetWorkloads() *pbcatalog.WorkloadSelector
}
