// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/health_check.proto

package v1alpha1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// HealthCheck defines configuration for health checking.
type HealthCheck struct {
	// List of selectors to match dataplanes that should be configured to do
	// health checks.
	Sources []*Selector `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	// List of selectors to match services that need to be health checked.
	Destinations []*Selector `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
	// Configuration for various types of health checking.
	Conf                 *HealthCheck_Conf `protobuf:"bytes,3,opt,name=conf,proto3" json:"conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4f9382814224e98, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetSources() []*Selector {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *HealthCheck) GetDestinations() []*Selector {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func (m *HealthCheck) GetConf() *HealthCheck_Conf {
	if m != nil {
		return m.Conf
	}
	return nil
}

// Conf defines configuration for various types of health checking.
type HealthCheck_Conf struct {
	// Configuration for active health checking.
	ActiveChecks *HealthCheck_Conf_Active `protobuf:"bytes,1,opt,name=active_checks,json=activeChecks,proto3" json:"active_checks,omitempty"`
	// Configuration for passive health checking.
	PassiveChecks        *HealthCheck_Conf_Passive `protobuf:"bytes,2,opt,name=passive_checks,json=passiveChecks,proto3" json:"passive_checks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *HealthCheck_Conf) Reset()         { *m = HealthCheck_Conf{} }
func (m *HealthCheck_Conf) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_Conf) ProtoMessage()    {}
func (*HealthCheck_Conf) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4f9382814224e98, []int{0, 0}
}

func (m *HealthCheck_Conf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_Conf.Unmarshal(m, b)
}
func (m *HealthCheck_Conf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_Conf.Marshal(b, m, deterministic)
}
func (m *HealthCheck_Conf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_Conf.Merge(m, src)
}
func (m *HealthCheck_Conf) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_Conf.Size(m)
}
func (m *HealthCheck_Conf) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_Conf.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_Conf proto.InternalMessageInfo

func (m *HealthCheck_Conf) GetActiveChecks() *HealthCheck_Conf_Active {
	if m != nil {
		return m.ActiveChecks
	}
	return nil
}

func (m *HealthCheck_Conf) GetPassiveChecks() *HealthCheck_Conf_Passive {
	if m != nil {
		return m.PassiveChecks
	}
	return nil
}

// Active defines configuration for active health checking.
type HealthCheck_Conf_Active struct {
	// Interval between consecutive health checks.
	Interval *duration.Duration `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	// Maximum time to wait for a health check response.
	Timeout *duration.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Number of consecutive unhealthy checks before considering a host
	// unhealthy.
	UnhealthyThreshold uint32 `protobuf:"varint,3,opt,name=unhealthy_threshold,json=unhealthyThreshold,proto3" json:"unhealthy_threshold,omitempty"`
	// Number of consecutive healthy checks before considering a host healthy.
	HealthyThreshold     uint32   `protobuf:"varint,4,opt,name=healthy_threshold,json=healthyThreshold,proto3" json:"healthy_threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_Conf_Active) Reset()         { *m = HealthCheck_Conf_Active{} }
func (m *HealthCheck_Conf_Active) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_Conf_Active) ProtoMessage()    {}
func (*HealthCheck_Conf_Active) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4f9382814224e98, []int{0, 0, 0}
}

func (m *HealthCheck_Conf_Active) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_Conf_Active.Unmarshal(m, b)
}
func (m *HealthCheck_Conf_Active) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_Conf_Active.Marshal(b, m, deterministic)
}
func (m *HealthCheck_Conf_Active) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_Conf_Active.Merge(m, src)
}
func (m *HealthCheck_Conf_Active) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_Conf_Active.Size(m)
}
func (m *HealthCheck_Conf_Active) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_Conf_Active.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_Conf_Active proto.InternalMessageInfo

func (m *HealthCheck_Conf_Active) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *HealthCheck_Conf_Active) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *HealthCheck_Conf_Active) GetUnhealthyThreshold() uint32 {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return 0
}

func (m *HealthCheck_Conf_Active) GetHealthyThreshold() uint32 {
	if m != nil {
		return m.HealthyThreshold
	}
	return 0
}

// Passive defines configuration for passive health checking.
type HealthCheck_Conf_Passive struct {
	// Number of consecutive failed requests before considering a host
	// unhealthy.
	UnhealthyThreshold uint32 `protobuf:"varint,1,opt,name=unhealthy_threshold,json=unhealthyThreshold,proto3" json:"unhealthy_threshold,omitempty"`
	// Interval a host should be considred unhealthy after reaching unhealthy
	// threshold.
	PenaltyInterval      *duration.Duration `protobuf:"bytes,2,opt,name=penalty_interval,json=penaltyInterval,proto3" json:"penalty_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HealthCheck_Conf_Passive) Reset()         { *m = HealthCheck_Conf_Passive{} }
func (m *HealthCheck_Conf_Passive) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_Conf_Passive) ProtoMessage()    {}
func (*HealthCheck_Conf_Passive) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4f9382814224e98, []int{0, 0, 1}
}

func (m *HealthCheck_Conf_Passive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_Conf_Passive.Unmarshal(m, b)
}
func (m *HealthCheck_Conf_Passive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_Conf_Passive.Marshal(b, m, deterministic)
}
func (m *HealthCheck_Conf_Passive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_Conf_Passive.Merge(m, src)
}
func (m *HealthCheck_Conf_Passive) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_Conf_Passive.Size(m)
}
func (m *HealthCheck_Conf_Passive) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_Conf_Passive.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_Conf_Passive proto.InternalMessageInfo

func (m *HealthCheck_Conf_Passive) GetUnhealthyThreshold() uint32 {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return 0
}

func (m *HealthCheck_Conf_Passive) GetPenaltyInterval() *duration.Duration {
	if m != nil {
		return m.PenaltyInterval
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "kuma.mesh.v1alpha1.HealthCheck")
	proto.RegisterType((*HealthCheck_Conf)(nil), "kuma.mesh.v1alpha1.HealthCheck.Conf")
	proto.RegisterType((*HealthCheck_Conf_Active)(nil), "kuma.mesh.v1alpha1.HealthCheck.Conf.Active")
	proto.RegisterType((*HealthCheck_Conf_Passive)(nil), "kuma.mesh.v1alpha1.HealthCheck.Conf.Passive")
}

func init() { proto.RegisterFile("mesh/v1alpha1/health_check.proto", fileDescriptor_a4f9382814224e98) }

var fileDescriptor_a4f9382814224e98 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6e, 0xd4, 0x30,
	0x14, 0x87, 0xeb, 0x4c, 0x98, 0x19, 0xde, 0xcc, 0x40, 0x31, 0x0b, 0x86, 0xa8, 0x42, 0x11, 0x62,
	0x51, 0x15, 0xe4, 0xa8, 0x45, 0x42, 0x88, 0x5d, 0x33, 0x2c, 0x40, 0x62, 0x51, 0xa5, 0xac, 0xd8,
	0x8c, 0xdc, 0xc4, 0xd3, 0x44, 0xf5, 0xc4, 0x91, 0xed, 0x04, 0x75, 0xcf, 0x09, 0xd8, 0x71, 0x05,
	0xc4, 0x09, 0x58, 0x71, 0x04, 0xae, 0xc1, 0x2d, 0x50, 0x6c, 0x67, 0xd4, 0xbf, 0x52, 0xba, 0x4b,
	0xf4, 0xde, 0xf7, 0xf9, 0xe7, 0xf7, 0x0c, 0xe1, 0x9a, 0xa9, 0x3c, 0x6a, 0xf6, 0x29, 0xaf, 0x72,
	0xba, 0x1f, 0xe5, 0x8c, 0x72, 0x9d, 0x2f, 0xd3, 0x9c, 0xa5, 0x67, 0xa4, 0x92, 0x42, 0x0b, 0x8c,
	0xcf, 0xea, 0x35, 0x25, 0x6d, 0x1b, 0xe9, 0xda, 0x82, 0x9d, 0xcb, 0x94, 0x62, 0x9c, 0xa5, 0x5a,
	0x48, 0x4b, 0x04, 0xcf, 0x4e, 0x85, 0x38, 0xe5, 0x2c, 0x32, 0x7f, 0x27, 0xf5, 0x2a, 0xca, 0x6a,
	0x49, 0x75, 0x21, 0xca, 0xdb, 0xea, 0x5f, 0x25, 0xad, 0x2a, 0x26, 0x95, 0xab, 0x3f, 0x69, 0x28,
	0x2f, 0x32, 0xaa, 0x59, 0xd4, 0x7d, 0xd8, 0xc2, 0xf3, 0x5f, 0x43, 0x98, 0x7c, 0x30, 0x09, 0x17,
	0x6d, 0x40, 0x1c, 0xc3, 0x48, 0x89, 0x5a, 0xa6, 0x4c, 0xcd, 0x51, 0x38, 0xd8, 0x9d, 0x1c, 0xec,
	0x90, 0xeb, 0x61, 0xc9, 0xb1, 0x4b, 0x17, 0xc3, 0xef, 0x7f, 0x7f, 0x06, 0xf7, 0xbe, 0x23, 0x6f,
	0x8c, 0x92, 0x0e, 0xc4, 0x9f, 0x60, 0x9a, 0x31, 0xa5, 0x8b, 0xd2, 0x24, 0x54, 0x73, 0xef, 0x8e,
	0xa2, 0x4b, 0x34, 0x7e, 0x0b, 0x7e, 0x2a, 0xca, 0xd5, 0x7c, 0x10, 0xa2, 0xdd, 0xc9, 0xc1, 0x8b,
	0x9b, 0x2c, 0x17, 0x2e, 0x40, 0x16, 0xa2, 0x5c, 0x25, 0x86, 0x08, 0xfe, 0xfa, 0xe0, 0xb7, 0xbf,
	0xf8, 0x08, 0x66, 0x34, 0xd5, 0x45, 0xc3, 0xec, 0x16, 0xda, 0xab, 0xb5, 0xae, 0x97, 0x7d, 0x5c,
	0xe4, 0xd0, 0x90, 0xc9, 0xd4, 0x1a, 0x4c, 0x41, 0xe1, 0x63, 0x78, 0x50, 0x51, 0xa5, 0x2e, 0x28,
	0x3d, 0xa3, 0x7c, 0xd5, 0x4b, 0x79, 0x64, 0xd1, 0x64, 0xe6, 0x1c, 0x56, 0x1a, 0x7c, 0xf3, 0x60,
	0x68, 0x4f, 0xc3, 0x0b, 0x18, 0x17, 0xa5, 0x66, 0xb2, 0xa1, 0xdc, 0x85, 0x7d, 0x4a, 0xec, 0x8a,
	0x49, 0xb7, 0x62, 0xf2, 0xde, 0x3d, 0x81, 0x78, 0xda, 0xce, 0x6e, 0xf4, 0x13, 0xf9, 0x63, 0xb4,
	0xb7, 0x95, 0x6c, 0x40, 0x7c, 0x08, 0x23, 0x5d, 0xac, 0x99, 0xa8, 0xb5, 0x4b, 0xd7, 0xdb, 0xd1,
	0x71, 0xf8, 0x1d, 0x3c, 0xae, 0x4b, 0xfb, 0x82, 0xcf, 0x97, 0x3a, 0x97, 0x4c, 0xe5, 0x82, 0x67,
	0x66, 0x17, 0xb3, 0xf8, 0x7e, 0xcb, 0xf8, 0x7b, 0x5e, 0xb8, 0x95, 0xe0, 0x4d, 0xd7, 0xe7, 0xae,
	0x09, 0xbf, 0x81, 0x47, 0xd7, 0x49, 0xff, 0x2a, 0xb9, 0x7d, 0x95, 0x0b, 0x7e, 0x20, 0x18, 0xb9,
	0x09, 0xdd, 0x76, 0x3e, 0xea, 0x73, 0x7e, 0x02, 0xdb, 0x15, 0x2b, 0x29, 0xd7, 0xe7, 0xcb, 0xcd,
	0x2c, 0xef, 0x38, 0x87, 0x87, 0x4e, 0xf0, 0xd1, 0xf1, 0x31, 0x7c, 0x19, 0x77, 0x6b, 0x3d, 0x19,
	0x1a, 0xfa, 0xf5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xb0, 0xd5, 0xa4, 0xf0, 0x03, 0x00,
	0x00,
}
