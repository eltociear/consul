// Code generated by protoc-json-shim. DO NOT EDIT.
package pbproxystate

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for Cluster
func (this *Cluster) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Cluster
func (this *Cluster) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for FailoverGroup
func (this *FailoverGroup) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FailoverGroup
func (this *FailoverGroup) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for FailoverGroupConfig
func (this *FailoverGroupConfig) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FailoverGroupConfig
func (this *FailoverGroupConfig) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for EndpointGroup
func (this *EndpointGroup) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for EndpointGroup
func (this *EndpointGroup) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for DynamicEndpointGroup
func (this *DynamicEndpointGroup) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DynamicEndpointGroup
func (this *DynamicEndpointGroup) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for PassthroughEndpointGroup
func (this *PassthroughEndpointGroup) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PassthroughEndpointGroup
func (this *PassthroughEndpointGroup) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for DNSEndpointGroup
func (this *DNSEndpointGroup) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DNSEndpointGroup
func (this *DNSEndpointGroup) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for StaticEndpointGroup
func (this *StaticEndpointGroup) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for StaticEndpointGroup
func (this *StaticEndpointGroup) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for DestinationCluster
func (this *DestinationCluster) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DestinationCluster
func (this *DestinationCluster) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for L4WeightedClusterGroup
func (this *L4WeightedClusterGroup) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for L4WeightedClusterGroup
func (this *L4WeightedClusterGroup) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for L7WeightedClusterGroup
func (this *L7WeightedClusterGroup) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for L7WeightedClusterGroup
func (this *L7WeightedClusterGroup) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for L4WeightedDestinationCluster
func (this *L4WeightedDestinationCluster) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for L4WeightedDestinationCluster
func (this *L4WeightedDestinationCluster) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for L7WeightedDestinationCluster
func (this *L7WeightedDestinationCluster) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for L7WeightedDestinationCluster
func (this *L7WeightedDestinationCluster) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for DynamicEndpointGroupConfig
func (this *DynamicEndpointGroupConfig) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DynamicEndpointGroupConfig
func (this *DynamicEndpointGroupConfig) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for LBPolicyLeastRequest
func (this *LBPolicyLeastRequest) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LBPolicyLeastRequest
func (this *LBPolicyLeastRequest) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for LBPolicyRoundRobin
func (this *LBPolicyRoundRobin) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LBPolicyRoundRobin
func (this *LBPolicyRoundRobin) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for LBPolicyRandom
func (this *LBPolicyRandom) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LBPolicyRandom
func (this *LBPolicyRandom) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for LBPolicyRingHash
func (this *LBPolicyRingHash) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LBPolicyRingHash
func (this *LBPolicyRingHash) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for LBPolicyMaglev
func (this *LBPolicyMaglev) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LBPolicyMaglev
func (this *LBPolicyMaglev) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for CircuitBreakers
func (this *CircuitBreakers) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CircuitBreakers
func (this *CircuitBreakers) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for UpstreamLimits
func (this *UpstreamLimits) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for UpstreamLimits
func (this *UpstreamLimits) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for OutlierDetection
func (this *OutlierDetection) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OutlierDetection
func (this *OutlierDetection) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for UpstreamConnectionOptions
func (this *UpstreamConnectionOptions) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for UpstreamConnectionOptions
func (this *UpstreamConnectionOptions) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for PassthroughEndpointGroupConfig
func (this *PassthroughEndpointGroupConfig) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PassthroughEndpointGroupConfig
func (this *PassthroughEndpointGroupConfig) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for DNSEndpointGroupConfig
func (this *DNSEndpointGroupConfig) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DNSEndpointGroupConfig
func (this *DNSEndpointGroupConfig) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for StaticEndpointGroupConfig
func (this *StaticEndpointGroupConfig) MarshalJSON() ([]byte, error) {
	str, err := ClusterMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for StaticEndpointGroupConfig
func (this *StaticEndpointGroupConfig) UnmarshalJSON(b []byte) error {
	return ClusterUnmarshaler.Unmarshal(b, this)
}

var (
	ClusterMarshaler   = &protojson.MarshalOptions{}
	ClusterUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)
