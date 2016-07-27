// Code generated by protoc-gen-gogo.
// source: github.com/synrobo/proto/consul/consul-service.proto
// DO NOT EDIT!

/*
	Package consul is a generated protocol buffer package.

	It is generated from these files:
		github.com/synrobo/proto/consul/consul-service.proto
		github.com/synrobo/proto/consul/consul.proto

	It has these top-level messages:
		ListNodesRequest
		ListNodesResponse
		ConsulNode
*/
package consul

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import identify "github.com/synrobo/proto/identify"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type ListNodesRequest struct {
	Identify *identify.Identify `protobuf:"bytes,1,opt,name=identify" json:"identify,omitempty"`
	Region   string             `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	Near     string             `protobuf:"bytes,3,opt,name=near,proto3" json:"near,omitempty"`
}

func (m *ListNodesRequest) Reset()                    { *m = ListNodesRequest{} }
func (m *ListNodesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListNodesRequest) ProtoMessage()               {}
func (*ListNodesRequest) Descriptor() ([]byte, []int) { return fileDescriptorConsulService, []int{0} }

func (m *ListNodesRequest) GetIdentify() *identify.Identify {
	if m != nil {
		return m.Identify
	}
	return nil
}

type ListNodesResponse struct {
	Node []*ConsulNode `protobuf:"bytes,1,rep,name=node" json:"node,omitempty"`
}

func (m *ListNodesResponse) Reset()                    { *m = ListNodesResponse{} }
func (m *ListNodesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListNodesResponse) ProtoMessage()               {}
func (*ListNodesResponse) Descriptor() ([]byte, []int) { return fileDescriptorConsulService, []int{1} }

func (m *ListNodesResponse) GetNode() []*ConsulNode {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*ListNodesRequest)(nil), "consul.ListNodesRequest")
	proto.RegisterType((*ListNodesResponse)(nil), "consul.ListNodesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ConsulService service

type ConsulServiceClient interface {
	// Devices are called nodes in Consul
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error)
}

type consulServiceClient struct {
	cc *grpc.ClientConn
}

func NewConsulServiceClient(cc *grpc.ClientConn) ConsulServiceClient {
	return &consulServiceClient{cc}
}

func (c *consulServiceClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	out := new(ListNodesResponse)
	err := grpc.Invoke(ctx, "/consul.ConsulService/ListNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConsulService service

type ConsulServiceServer interface {
	// Devices are called nodes in Consul
	ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error)
}

func RegisterConsulServiceServer(s *grpc.Server, srv ConsulServiceServer) {
	s.RegisterService(&_ConsulService_serviceDesc, srv)
}

func _ConsulService_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsulServiceServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consul.ConsulService/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsulServiceServer).ListNodes(ctx, req.(*ListNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsulService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consul.ConsulService",
	HandlerType: (*ConsulServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNodes",
			Handler:    _ConsulService_ListNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorConsulService,
}

func (m *ListNodesRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ListNodesRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Identify != nil {
		data[i] = 0xa
		i++
		i = encodeVarintConsulService(data, i, uint64(m.Identify.Size()))
		n1, err := m.Identify.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Region) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintConsulService(data, i, uint64(len(m.Region)))
		i += copy(data[i:], m.Region)
	}
	if len(m.Near) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintConsulService(data, i, uint64(len(m.Near)))
		i += copy(data[i:], m.Near)
	}
	return i, nil
}

func (m *ListNodesResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ListNodesResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Node) > 0 {
		for _, msg := range m.Node {
			data[i] = 0xa
			i++
			i = encodeVarintConsulService(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64ConsulService(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ConsulService(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConsulService(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ListNodesRequest) Size() (n int) {
	var l int
	_ = l
	if m.Identify != nil {
		l = m.Identify.Size()
		n += 1 + l + sovConsulService(uint64(l))
	}
	l = len(m.Region)
	if l > 0 {
		n += 1 + l + sovConsulService(uint64(l))
	}
	l = len(m.Near)
	if l > 0 {
		n += 1 + l + sovConsulService(uint64(l))
	}
	return n
}

func (m *ListNodesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Node) > 0 {
		for _, e := range m.Node {
			l = e.Size()
			n += 1 + l + sovConsulService(uint64(l))
		}
	}
	return n
}

func sovConsulService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConsulService(x uint64) (n int) {
	return sovConsulService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListNodesRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsulService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListNodesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListNodesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identify", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsulService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConsulService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Identify == nil {
				m.Identify = &identify.Identify{}
			}
			if err := m.Identify.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Region", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsulService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConsulService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Region = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Near", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsulService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConsulService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Near = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsulService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConsulService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListNodesResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsulService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListNodesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListNodesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsulService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConsulService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = append(m.Node, &ConsulNode{})
			if err := m.Node[len(m.Node)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsulService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConsulService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConsulService(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConsulService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConsulService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConsulService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConsulService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConsulService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConsulService(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConsulService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConsulService   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorConsulService = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xae, 0xcc, 0x2b, 0xca, 0x4f, 0xca, 0xd7, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0x2b, 0x2e, 0xcd, 0x81, 0x52, 0xba, 0xc5, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7a, 0x60, 0x39, 0x21, 0x36, 0x88, 0xa8, 0x94, 0x2e, 0x92, 0xee,
	0xf4, 0xfc, 0x74, 0xa8, 0xd6, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0x62, 0x0e, 0x88, 0x05, 0xd1, 0x26,
	0xa5, 0x43, 0x9c, 0x65, 0x50, 0xd5, 0x06, 0x38, 0x55, 0x67, 0xa6, 0xa4, 0xe6, 0x95, 0x64, 0xa6,
	0x55, 0xc2, 0x19, 0x10, 0x1d, 0x4a, 0x79, 0x5c, 0x02, 0x3e, 0x99, 0xc5, 0x25, 0x7e, 0xf9, 0x29,
	0xa9, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x7a, 0x5c, 0x1c, 0x30, 0x55, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x42, 0x7a, 0x70, 0x6d, 0x9e, 0x50, 0x46, 0x10, 0x5c, 0x8d,
	0x90, 0x18, 0x17, 0x5b, 0x51, 0x6a, 0x7a, 0x66, 0x7e, 0x9e, 0x04, 0x13, 0x50, 0x35, 0x67, 0x10,
	0x94, 0x27, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x9a, 0x58, 0x24, 0xc1, 0x0c, 0x16, 0x05, 0xb3, 0x95,
	0xac, 0xb9, 0x04, 0x91, 0xec, 0x2b, 0x2e, 0x00, 0xba, 0x3e, 0x55, 0x48, 0x0d, 0xa8, 0x10, 0x28,
	0x00, 0xb4, 0x8c, 0x19, 0x6c, 0x19, 0xd4, 0x4f, 0xce, 0x60, 0x0a, 0xa4, 0x34, 0x08, 0x2c, 0x6f,
	0x14, 0xc8, 0xc5, 0x0b, 0x11, 0x0b, 0x86, 0x04, 0xad, 0x90, 0x03, 0x17, 0x27, 0xdc, 0x34, 0x21,
	0x09, 0x98, 0x3e, 0x74, 0x0f, 0x49, 0x49, 0x62, 0x91, 0x81, 0x58, 0xed, 0xc4, 0x73, 0xe2, 0x91,
	0x1c, 0xe3, 0x05, 0x20, 0x7e, 0x00, 0xc4, 0x49, 0x6c, 0xe0, 0x40, 0x31, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x13, 0xc4, 0xcb, 0x1d, 0xe3, 0x01, 0x00, 0x00,
}