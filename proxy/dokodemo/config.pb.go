package dokodemo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_net "v2ray.com/core/common/net"
import v2ray_core_common_net1 "v2ray.com/core/common/net"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	Address        *v2ray_core_common_net.IPOrDomain   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Port           uint32                              `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	NetworkList    *v2ray_core_common_net1.NetworkList `protobuf:"bytes,3,opt,name=network_list,json=networkList" json:"network_list,omitempty"`
	Timeout        uint32                              `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
	FollowRedirect bool                                `protobuf:"varint,5,opt,name=follow_redirect,json=followRedirect" json:"follow_redirect,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetAddress() *v2ray_core_common_net.IPOrDomain {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Config) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Config) GetNetworkList() *v2ray_core_common_net1.NetworkList {
	if m != nil {
		return m.NetworkList
	}
	return nil
}

func (m *Config) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Config) GetFollowRedirect() bool {
	if m != nil {
		return m.FollowRedirect
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.proxy.dokodemo.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/dokodemo/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0x49, 0xad, 0x6d, 0x49, 0xfd, 0x03, 0x39, 0xa5, 0x82, 0x50, 0x7b, 0x69, 0xf1, 0x90,
	0x40, 0x3d, 0x7a, 0xb3, 0x15, 0x11, 0x44, 0x97, 0x1c, 0x3c, 0x78, 0x29, 0x6b, 0x36, 0x95, 0xd0,
	0x4d, 0x66, 0x99, 0x8d, 0xd6, 0x7d, 0x25, 0x5f, 0xcb, 0x17, 0x11, 0xb3, 0xbb, 0x28, 0x42, 0xbd,
	0xcd, 0x4c, 0xbe, 0x7c, 0x33, 0xfc, 0xe8, 0xf9, 0xdb, 0x1c, 0xd3, 0x4a, 0x68, 0x70, 0x52, 0x03,
	0x1a, 0x59, 0x20, 0xbc, 0x57, 0x32, 0x83, 0x0d, 0x64, 0xc6, 0x81, 0xd4, 0xe0, 0xd7, 0xf6, 0x45,
	0x14, 0x08, 0x01, 0xd8, 0xa8, 0x65, 0xd1, 0x88, 0xc8, 0x89, 0x96, 0x3b, 0x99, 0xfe, 0xd1, 0x68,
	0x70, 0x0e, 0xbc, 0xf4, 0x26, 0xc8, 0x34, 0xcb, 0xd0, 0x94, 0x65, 0xed, 0xf8, 0x0f, 0xf4, 0x26,
	0x6c, 0x01, 0x37, 0x35, 0x38, 0xf9, 0x24, 0xb4, 0xb7, 0x88, 0xdb, 0xd9, 0x25, 0xed, 0x37, 0x12,
	0x4e, 0xc6, 0x64, 0x36, 0x9c, 0x9f, 0x89, 0x5f, 0x97, 0xd4, 0x06, 0xe1, 0x4d, 0x10, 0xb7, 0xc9,
	0x03, 0x2e, 0xc1, 0xa5, 0xd6, 0xab, 0xf6, 0x07, 0x63, 0xb4, 0x5b, 0x00, 0x06, 0xde, 0x19, 0x93,
	0xd9, 0xa1, 0x8a, 0x35, 0xbb, 0xa6, 0x07, 0xcd, 0xb2, 0x55, 0x6e, 0xcb, 0xc0, 0xf7, 0xa2, 0x75,
	0xb2, 0xc3, 0x7a, 0x5f, 0xa3, 0x77, 0xb6, 0x0c, 0x6a, 0xe8, 0x7f, 0x1a, 0xc6, 0x69, 0x3f, 0x58,
	0x67, 0xe0, 0x35, 0xf0, 0x6e, 0xb4, 0xb7, 0x2d, 0x9b, 0xd2, 0xe3, 0x35, 0xe4, 0x39, 0x6c, 0x57,
	0x68, 0x32, 0x8b, 0x46, 0x07, 0xbe, 0x3f, 0x26, 0xb3, 0x81, 0x3a, 0xaa, 0xc7, 0xaa, 0x99, 0x5e,
	0xdd, 0xd0, 0x53, 0x0d, 0x4e, 0xec, 0x0c, 0x36, 0x21, 0x4f, 0x83, 0xb6, 0xfe, 0xe8, 0x8c, 0x1e,
	0xe7, 0x2a, 0xad, 0xc4, 0xe2, 0x9b, 0x4b, 0x22, 0xb7, 0x6c, 0xde, 0x9e, 0x7b, 0x31, 0xb5, 0x8b,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x1c, 0xb0, 0xd9, 0xd0, 0x01, 0x00, 0x00,
}
