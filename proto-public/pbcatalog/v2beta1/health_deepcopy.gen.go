// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package catalogv2beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using HealthStatus within kubernetes types, where deepcopy-gen is used.
func (in *HealthStatus) DeepCopyInto(out *HealthStatus) {
	p := proto.Clone(in).(*HealthStatus)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthStatus. Required by controller-gen.
func (in *HealthStatus) DeepCopy() *HealthStatus {
	if in == nil {
		return nil
	}
	out := new(HealthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HealthStatus. Required by controller-gen.
func (in *HealthStatus) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HealthChecks within kubernetes types, where deepcopy-gen is used.
func (in *HealthChecks) DeepCopyInto(out *HealthChecks) {
	p := proto.Clone(in).(*HealthChecks)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthChecks. Required by controller-gen.
func (in *HealthChecks) DeepCopy() *HealthChecks {
	if in == nil {
		return nil
	}
	out := new(HealthChecks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HealthChecks. Required by controller-gen.
func (in *HealthChecks) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HealthCheck within kubernetes types, where deepcopy-gen is used.
func (in *HealthCheck) DeepCopyInto(out *HealthCheck) {
	p := proto.Clone(in).(*HealthCheck)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck. Required by controller-gen.
func (in *HealthCheck) DeepCopy() *HealthCheck {
	if in == nil {
		return nil
	}
	out := new(HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck. Required by controller-gen.
func (in *HealthCheck) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPCheck within kubernetes types, where deepcopy-gen is used.
func (in *HTTPCheck) DeepCopyInto(out *HTTPCheck) {
	p := proto.Clone(in).(*HTTPCheck)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPCheck. Required by controller-gen.
func (in *HTTPCheck) DeepCopy() *HTTPCheck {
	if in == nil {
		return nil
	}
	out := new(HTTPCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPCheck. Required by controller-gen.
func (in *HTTPCheck) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TCPCheck within kubernetes types, where deepcopy-gen is used.
func (in *TCPCheck) DeepCopyInto(out *TCPCheck) {
	p := proto.Clone(in).(*TCPCheck)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPCheck. Required by controller-gen.
func (in *TCPCheck) DeepCopy() *TCPCheck {
	if in == nil {
		return nil
	}
	out := new(TCPCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TCPCheck. Required by controller-gen.
func (in *TCPCheck) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using UDPCheck within kubernetes types, where deepcopy-gen is used.
func (in *UDPCheck) DeepCopyInto(out *UDPCheck) {
	p := proto.Clone(in).(*UDPCheck)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDPCheck. Required by controller-gen.
func (in *UDPCheck) DeepCopy() *UDPCheck {
	if in == nil {
		return nil
	}
	out := new(UDPCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new UDPCheck. Required by controller-gen.
func (in *UDPCheck) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GRPCCheck within kubernetes types, where deepcopy-gen is used.
func (in *GRPCCheck) DeepCopyInto(out *GRPCCheck) {
	p := proto.Clone(in).(*GRPCCheck)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCCheck. Required by controller-gen.
func (in *GRPCCheck) DeepCopy() *GRPCCheck {
	if in == nil {
		return nil
	}
	out := new(GRPCCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GRPCCheck. Required by controller-gen.
func (in *GRPCCheck) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using OSServiceCheck within kubernetes types, where deepcopy-gen is used.
func (in *OSServiceCheck) DeepCopyInto(out *OSServiceCheck) {
	p := proto.Clone(in).(*OSServiceCheck)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSServiceCheck. Required by controller-gen.
func (in *OSServiceCheck) DeepCopy() *OSServiceCheck {
	if in == nil {
		return nil
	}
	out := new(OSServiceCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new OSServiceCheck. Required by controller-gen.
func (in *OSServiceCheck) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using CheckTLSConfig within kubernetes types, where deepcopy-gen is used.
func (in *CheckTLSConfig) DeepCopyInto(out *CheckTLSConfig) {
	p := proto.Clone(in).(*CheckTLSConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckTLSConfig. Required by controller-gen.
func (in *CheckTLSConfig) DeepCopy() *CheckTLSConfig {
	if in == nil {
		return nil
	}
	out := new(CheckTLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CheckTLSConfig. Required by controller-gen.
func (in *CheckTLSConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
