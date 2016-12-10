// Code generated by protoc-gen-gogo.
// source: github.com/fuserobotics/proto/consul/consul.proto
// DO NOT EDIT!

package consul

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/fuserobotics/proto/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsulNode struct {
	Name    string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address *common.IPAddress `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *ConsulNode) Reset()                    { *m = ConsulNode{} }
func (m *ConsulNode) String() string            { return proto.CompactTextString(m) }
func (*ConsulNode) ProtoMessage()               {}
func (*ConsulNode) Descriptor() ([]byte, []int) { return fileDescriptorConsul, []int{0} }

func (m *ConsulNode) GetAddress() *common.IPAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

type ConsulHealthCheck struct {
	Node        string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	CheckId     string `protobuf:"bytes,2,opt,name=check_id,json=checkId,proto3" json:"check_id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Status      string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Notes       string `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	Output      string `protobuf:"bytes,6,opt,name=output,proto3" json:"output,omitempty"`
	ServiceId   string `protobuf:"bytes,7,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ServiceName string `protobuf:"bytes,8,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (m *ConsulHealthCheck) Reset()                    { *m = ConsulHealthCheck{} }
func (m *ConsulHealthCheck) String() string            { return proto.CompactTextString(m) }
func (*ConsulHealthCheck) ProtoMessage()               {}
func (*ConsulHealthCheck) Descriptor() ([]byte, []int) { return fileDescriptorConsul, []int{1} }

type ConsulAgentService struct {
	Id                string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceName       string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Tags              []string `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
	Port              uint32   `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Address           string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	EnableTagOverride bool     `protobuf:"varint,6,opt,name=enable_tag_override,json=enableTagOverride,proto3" json:"enable_tag_override,omitempty"`
}

func (m *ConsulAgentService) Reset()                    { *m = ConsulAgentService{} }
func (m *ConsulAgentService) String() string            { return proto.CompactTextString(m) }
func (*ConsulAgentService) ProtoMessage()               {}
func (*ConsulAgentService) Descriptor() ([]byte, []int) { return fileDescriptorConsul, []int{2} }

type ConsulServiceHealth struct {
	Node   *ConsulNode          `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	Checks []*ConsulHealthCheck `protobuf:"bytes,2,rep,name=checks" json:"checks,omitempty"`
	Svc    *ConsulAgentService  `protobuf:"bytes,3,opt,name=svc" json:"svc,omitempty"`
}

func (m *ConsulServiceHealth) Reset()                    { *m = ConsulServiceHealth{} }
func (m *ConsulServiceHealth) String() string            { return proto.CompactTextString(m) }
func (*ConsulServiceHealth) ProtoMessage()               {}
func (*ConsulServiceHealth) Descriptor() ([]byte, []int) { return fileDescriptorConsul, []int{3} }

func (m *ConsulServiceHealth) GetNode() *ConsulNode {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *ConsulServiceHealth) GetChecks() []*ConsulHealthCheck {
	if m != nil {
		return m.Checks
	}
	return nil
}

func (m *ConsulServiceHealth) GetSvc() *ConsulAgentService {
	if m != nil {
		return m.Svc
	}
	return nil
}

type ConsulQueryMeta struct {
	LastIndex   uint64 `protobuf:"varint,1,opt,name=last_index,json=lastIndex,proto3" json:"last_index,omitempty"`
	LastContact int64  `protobuf:"varint,2,opt,name=last_contact,json=lastContact,proto3" json:"last_contact,omitempty"`
	KnownLeader bool   `protobuf:"varint,3,opt,name=known_leader,json=knownLeader,proto3" json:"known_leader,omitempty"`
	RequestTime int64  `protobuf:"varint,4,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
}

func (m *ConsulQueryMeta) Reset()                    { *m = ConsulQueryMeta{} }
func (m *ConsulQueryMeta) String() string            { return proto.CompactTextString(m) }
func (*ConsulQueryMeta) ProtoMessage()               {}
func (*ConsulQueryMeta) Descriptor() ([]byte, []int) { return fileDescriptorConsul, []int{4} }

func init() {
	proto.RegisterType((*ConsulNode)(nil), "consul.ConsulNode")
	proto.RegisterType((*ConsulHealthCheck)(nil), "consul.ConsulHealthCheck")
	proto.RegisterType((*ConsulAgentService)(nil), "consul.ConsulAgentService")
	proto.RegisterType((*ConsulServiceHealth)(nil), "consul.ConsulServiceHealth")
	proto.RegisterType((*ConsulQueryMeta)(nil), "consul.ConsulQueryMeta")
}

var fileDescriptorConsul = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xda, 0x2e, 0x6d, 0x4f, 0xf8, 0x51, 0x3d, 0x84, 0xb2, 0x49, 0x48, 0x90, 0x0b, 0x84,
	0x04, 0x4a, 0xb5, 0xf2, 0x04, 0x53, 0x6f, 0xa8, 0xc4, 0x06, 0x98, 0xdd, 0x47, 0x6e, 0x72, 0x68,
	0xa3, 0xa5, 0x71, 0xb1, 0x9d, 0x02, 0x8f, 0xc2, 0x0d, 0xef, 0xc1, 0xcb, 0xf0, 0x2c, 0xd8, 0xc7,
	0xee, 0xd6, 0x69, 0x17, 0xbb, 0x8a, 0xcf, 0xf7, 0x7d, 0xe7, 0xf3, 0xf9, 0x71, 0xe0, 0x6c, 0x55,
	0x9b, 0x75, 0xb7, 0xcc, 0x4b, 0xb9, 0x99, 0x7e, 0xeb, 0x34, 0x2a, 0xb9, 0x94, 0xa6, 0x2e, 0xf5,
	0x74, 0xab, 0xa4, 0x91, 0xd3, 0x52, 0xb6, 0xba, 0x6b, 0xc2, 0x27, 0x27, 0x8c, 0xc5, 0x3e, 0x3a,
	0x7d, 0x30, 0x75, 0xb3, 0x91, 0x6d, 0xf8, 0xf8, 0xd4, 0xec, 0x02, 0x60, 0x4e, 0xc9, 0x97, 0xb2,
	0x42, 0xc6, 0x60, 0xd0, 0x8a, 0x0d, 0xa6, 0xd1, 0xcb, 0xe8, 0xcd, 0x98, 0xd3, 0x99, 0xbd, 0x85,
	0xa1, 0xa8, 0x2a, 0x85, 0x5a, 0xa7, 0x3d, 0x0b, 0x27, 0xb3, 0x49, 0x1e, 0x1c, 0x16, 0x9f, 0xcf,
	0x3d, 0xc1, 0xf7, 0x8a, 0xec, 0x5f, 0x04, 0x13, 0xef, 0xf7, 0x01, 0x45, 0x63, 0xd6, 0xf3, 0x35,
	0x96, 0xd7, 0x64, 0x6b, 0xed, 0x6f, 0x6c, 0xdd, 0x55, 0x27, 0x30, 0x2a, 0x1d, 0x59, 0xd4, 0x15,
	0xf9, 0x8e, 0xf9, 0x90, 0xe2, 0x45, 0x75, 0x53, 0x45, 0xff, 0xa0, 0x8a, 0xe7, 0x10, 0x6b, 0x23,
	0x4c, 0xa7, 0xd3, 0x01, 0xa1, 0x21, 0x62, 0xcf, 0xe0, 0xa8, 0x95, 0x06, 0x75, 0x7a, 0x44, 0xb0,
	0x0f, 0x9c, 0x5a, 0x76, 0x66, 0xdb, 0x99, 0x34, 0xf6, 0x6a, 0x1f, 0xb1, 0x17, 0x00, 0x76, 0x2a,
	0xbb, 0xba, 0x44, 0x77, 0xed, 0x90, 0xb8, 0x71, 0x40, 0xec, 0xc5, 0xaf, 0xe0, 0xd1, 0x9e, 0xa6,
	0x02, 0x46, 0x24, 0x48, 0x02, 0x76, 0x69, 0xa1, 0xec, 0x6f, 0x04, 0xcc, 0x37, 0x78, 0xbe, 0xc2,
	0xd6, 0x7c, 0xf5, 0x14, 0x7b, 0x02, 0x3d, 0x6b, 0xe8, 0xfb, 0xb3, 0xa7, 0x7b, 0x4e, 0xbd, 0x7b,
	0x4e, 0xae, 0x4b, 0x23, 0x56, 0xda, 0x76, 0xd9, 0x77, 0x5d, 0xba, 0xb3, 0xc3, 0xb6, 0x52, 0x19,
	0xea, 0xf1, 0x31, 0xa7, 0x33, 0x4b, 0x6f, 0xe7, 0xef, 0x7b, 0xdc, 0x87, 0x2c, 0x87, 0x63, 0x6c,
	0xc5, 0xb2, 0xc1, 0xc2, 0x26, 0x17, 0x72, 0x87, 0x4a, 0xd5, 0x76, 0xca, 0xae, 0xe5, 0x11, 0x9f,
	0x78, 0xea, 0x4a, 0xac, 0x3e, 0x05, 0x22, 0xfb, 0x13, 0xc1, 0xb1, 0xaf, 0x3d, 0x94, 0xed, 0x77,
	0xc4, 0x5e, 0x1f, 0xac, 0x27, 0x99, 0xb1, 0x3c, 0xbc, 0xad, 0xdb, 0x77, 0x11, 0x56, 0x76, 0x06,
	0x31, 0xad, 0xc8, 0x3d, 0x84, 0xbe, 0x55, 0x9e, 0xdc, 0x55, 0x1e, 0x6c, 0x9c, 0x07, 0x21, 0x7b,
	0x07, 0x7d, 0xbd, 0x2b, 0x69, 0x93, 0xc9, 0xec, 0xf4, 0xae, 0xfe, 0x70, 0x80, 0xdc, 0xc9, 0xb2,
	0xdf, 0x11, 0x3c, 0xf5, 0xdc, 0x97, 0x0e, 0xd5, 0xaf, 0x0b, 0x34, 0xc2, 0xad, 0xac, 0x11, 0xda,
	0x14, 0x75, 0x5b, 0xe1, 0x4f, 0x2a, 0x71, 0xc0, 0xc7, 0x0e, 0x59, 0x38, 0xc0, 0x0d, 0x9a, 0x68,
	0xeb, 0x6c, 0x44, 0x69, 0x68, 0xd0, 0x7d, 0x9e, 0x38, 0x6c, 0xee, 0x21, 0x27, 0xb9, 0x6e, 0xe5,
	0x8f, 0xb6, 0x68, 0x50, 0x54, 0xa8, 0xa8, 0x98, 0x11, 0x4f, 0x08, 0xfb, 0x48, 0x90, 0x93, 0x28,
	0xfc, 0xde, 0xa1, 0x35, 0x32, 0xb5, 0x5d, 0xd7, 0xc0, 0xbb, 0x04, 0xec, 0xca, 0x42, 0xcb, 0x98,
	0xfe, 0x97, 0xf7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x72, 0x69, 0x1c, 0x9f, 0x03, 0x00,
	0x00,
}
