// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package pbproxystate

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using Route within kubernetes types, where deepcopy-gen is used.
func (in *Route) DeepCopyInto(out *Route) {
	p := proto.Clone(in).(*Route)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route. Required by controller-gen.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Route. Required by controller-gen.
func (in *Route) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using VirtualHost within kubernetes types, where deepcopy-gen is used.
func (in *VirtualHost) DeepCopyInto(out *VirtualHost) {
	p := proto.Clone(in).(*VirtualHost)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualHost. Required by controller-gen.
func (in *VirtualHost) DeepCopy() *VirtualHost {
	if in == nil {
		return nil
	}
	out := new(VirtualHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new VirtualHost. Required by controller-gen.
func (in *VirtualHost) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RouteRule within kubernetes types, where deepcopy-gen is used.
func (in *RouteRule) DeepCopyInto(out *RouteRule) {
	p := proto.Clone(in).(*RouteRule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteRule. Required by controller-gen.
func (in *RouteRule) DeepCopy() *RouteRule {
	if in == nil {
		return nil
	}
	out := new(RouteRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RouteRule. Required by controller-gen.
func (in *RouteRule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RouteMatch within kubernetes types, where deepcopy-gen is used.
func (in *RouteMatch) DeepCopyInto(out *RouteMatch) {
	p := proto.Clone(in).(*RouteMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMatch. Required by controller-gen.
func (in *RouteMatch) DeepCopy() *RouteMatch {
	if in == nil {
		return nil
	}
	out := new(RouteMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RouteMatch. Required by controller-gen.
func (in *RouteMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PathMatch within kubernetes types, where deepcopy-gen is used.
func (in *PathMatch) DeepCopyInto(out *PathMatch) {
	p := proto.Clone(in).(*PathMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PathMatch. Required by controller-gen.
func (in *PathMatch) DeepCopy() *PathMatch {
	if in == nil {
		return nil
	}
	out := new(PathMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PathMatch. Required by controller-gen.
func (in *PathMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using QueryParameterMatch within kubernetes types, where deepcopy-gen is used.
func (in *QueryParameterMatch) DeepCopyInto(out *QueryParameterMatch) {
	p := proto.Clone(in).(*QueryParameterMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryParameterMatch. Required by controller-gen.
func (in *QueryParameterMatch) DeepCopy() *QueryParameterMatch {
	if in == nil {
		return nil
	}
	out := new(QueryParameterMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new QueryParameterMatch. Required by controller-gen.
func (in *QueryParameterMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HeaderMatch within kubernetes types, where deepcopy-gen is used.
func (in *HeaderMatch) DeepCopyInto(out *HeaderMatch) {
	p := proto.Clone(in).(*HeaderMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderMatch. Required by controller-gen.
func (in *HeaderMatch) DeepCopy() *HeaderMatch {
	if in == nil {
		return nil
	}
	out := new(HeaderMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HeaderMatch. Required by controller-gen.
func (in *HeaderMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RouteDestination within kubernetes types, where deepcopy-gen is used.
func (in *RouteDestination) DeepCopyInto(out *RouteDestination) {
	p := proto.Clone(in).(*RouteDestination)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteDestination. Required by controller-gen.
func (in *RouteDestination) DeepCopy() *RouteDestination {
	if in == nil {
		return nil
	}
	out := new(RouteDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RouteDestination. Required by controller-gen.
func (in *RouteDestination) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using DestinationConfiguration within kubernetes types, where deepcopy-gen is used.
func (in *DestinationConfiguration) DeepCopyInto(out *DestinationConfiguration) {
	p := proto.Clone(in).(*DestinationConfiguration)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationConfiguration. Required by controller-gen.
func (in *DestinationConfiguration) DeepCopy() *DestinationConfiguration {
	if in == nil {
		return nil
	}
	out := new(DestinationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DestinationConfiguration. Required by controller-gen.
func (in *DestinationConfiguration) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RetryPolicy within kubernetes types, where deepcopy-gen is used.
func (in *RetryPolicy) DeepCopyInto(out *RetryPolicy) {
	p := proto.Clone(in).(*RetryPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryPolicy. Required by controller-gen.
func (in *RetryPolicy) DeepCopy() *RetryPolicy {
	if in == nil {
		return nil
	}
	out := new(RetryPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RetryPolicy. Required by controller-gen.
func (in *RetryPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TimeoutConfig within kubernetes types, where deepcopy-gen is used.
func (in *TimeoutConfig) DeepCopyInto(out *TimeoutConfig) {
	p := proto.Clone(in).(*TimeoutConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeoutConfig. Required by controller-gen.
func (in *TimeoutConfig) DeepCopy() *TimeoutConfig {
	if in == nil {
		return nil
	}
	out := new(TimeoutConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TimeoutConfig. Required by controller-gen.
func (in *TimeoutConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadBalancerHashPolicy within kubernetes types, where deepcopy-gen is used.
func (in *LoadBalancerHashPolicy) DeepCopyInto(out *LoadBalancerHashPolicy) {
	p := proto.Clone(in).(*LoadBalancerHashPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerHashPolicy. Required by controller-gen.
func (in *LoadBalancerHashPolicy) DeepCopy() *LoadBalancerHashPolicy {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerHashPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerHashPolicy. Required by controller-gen.
func (in *LoadBalancerHashPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using CookiePolicy within kubernetes types, where deepcopy-gen is used.
func (in *CookiePolicy) DeepCopyInto(out *CookiePolicy) {
	p := proto.Clone(in).(*CookiePolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CookiePolicy. Required by controller-gen.
func (in *CookiePolicy) DeepCopy() *CookiePolicy {
	if in == nil {
		return nil
	}
	out := new(CookiePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CookiePolicy. Required by controller-gen.
func (in *CookiePolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HeaderPolicy within kubernetes types, where deepcopy-gen is used.
func (in *HeaderPolicy) DeepCopyInto(out *HeaderPolicy) {
	p := proto.Clone(in).(*HeaderPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderPolicy. Required by controller-gen.
func (in *HeaderPolicy) DeepCopy() *HeaderPolicy {
	if in == nil {
		return nil
	}
	out := new(HeaderPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HeaderPolicy. Required by controller-gen.
func (in *HeaderPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using QueryParameterPolicy within kubernetes types, where deepcopy-gen is used.
func (in *QueryParameterPolicy) DeepCopyInto(out *QueryParameterPolicy) {
	p := proto.Clone(in).(*QueryParameterPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryParameterPolicy. Required by controller-gen.
func (in *QueryParameterPolicy) DeepCopy() *QueryParameterPolicy {
	if in == nil {
		return nil
	}
	out := new(QueryParameterPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new QueryParameterPolicy. Required by controller-gen.
func (in *QueryParameterPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConnectionPropertiesPolicy within kubernetes types, where deepcopy-gen is used.
func (in *ConnectionPropertiesPolicy) DeepCopyInto(out *ConnectionPropertiesPolicy) {
	p := proto.Clone(in).(*ConnectionPropertiesPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPropertiesPolicy. Required by controller-gen.
func (in *ConnectionPropertiesPolicy) DeepCopy() *ConnectionPropertiesPolicy {
	if in == nil {
		return nil
	}
	out := new(ConnectionPropertiesPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPropertiesPolicy. Required by controller-gen.
func (in *ConnectionPropertiesPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
