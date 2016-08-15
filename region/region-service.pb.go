// Code generated by protoc-gen-gogo.
// source: github.com/fuserobotics/proto/region/region-service.proto
// DO NOT EDIT!

/*
	Package region is a generated protocol buffer package.

	It is generated from these files:
		github.com/fuserobotics/proto/region/region-service.proto
		github.com/fuserobotics/proto/region/region.proto

	It has these top-level messages:
		ListRegionsRequest
		ListRegionsResponse
		CreateRegionRequest
		CreateRegionResponse
		SaveRegionRequest
		SaveRegionResponse
		Region
*/
package region

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import identify "github.com/fuserobotics/proto/identify"
import common "github.com/fuserobotics/proto/common"

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

type ListRegionsRequest struct {
	Identify *identify.Identify `protobuf:"bytes,1,opt,name=identify" json:"identify,omitempty"`
}

func (m *ListRegionsRequest) Reset()                    { *m = ListRegionsRequest{} }
func (m *ListRegionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRegionsRequest) ProtoMessage()               {}
func (*ListRegionsRequest) Descriptor() ([]byte, []int) { return fileDescriptorRegionService, []int{0} }

func (m *ListRegionsRequest) GetIdentify() *identify.Identify {
	if m != nil {
		return m.Identify
	}
	return nil
}

type ListRegionsResponse struct {
	Region []*Region `protobuf:"bytes,1,rep,name=region" json:"region,omitempty"`
}

func (m *ListRegionsResponse) Reset()                    { *m = ListRegionsResponse{} }
func (m *ListRegionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListRegionsResponse) ProtoMessage()               {}
func (*ListRegionsResponse) Descriptor() ([]byte, []int) { return fileDescriptorRegionService, []int{1} }

func (m *ListRegionsResponse) GetRegion() []*Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type CreateRegionRequest struct {
	Region *Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
}

func (m *CreateRegionRequest) Reset()                    { *m = CreateRegionRequest{} }
func (m *CreateRegionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRegionRequest) ProtoMessage()               {}
func (*CreateRegionRequest) Descriptor() ([]byte, []int) { return fileDescriptorRegionService, []int{2} }

func (m *CreateRegionRequest) GetRegion() *Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type CreateRegionResponse struct {
	CreatedRegion *Region `protobuf:"bytes,1,opt,name=created_region,json=createdRegion" json:"created_region,omitempty"`
}

func (m *CreateRegionResponse) Reset()         { *m = CreateRegionResponse{} }
func (m *CreateRegionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRegionResponse) ProtoMessage()    {}
func (*CreateRegionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorRegionService, []int{3}
}

func (m *CreateRegionResponse) GetCreatedRegion() *Region {
	if m != nil {
		return m.CreatedRegion
	}
	return nil
}

type SaveRegionRequest struct {
	Identify     *identify.Identify           `protobuf:"bytes,1,opt,name=identify" json:"identify,omitempty"`
	Id           string                       `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                       `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	DataKeyspace []*Region_RegionDataKeyspace `protobuf:"bytes,5,rep,name=data_keyspace,json=dataKeyspace" json:"data_keyspace,omitempty"`
	Location     *common.GeoLocation          `protobuf:"bytes,6,opt,name=location" json:"location,omitempty"`
	// If set, if data_keyspace is null will delete all of them.
	AllowClearDataKeyspaces bool `protobuf:"varint,7,opt,name=allow_clear_data_keyspaces,json=allowClearDataKeyspaces,proto3" json:"allow_clear_data_keyspaces,omitempty"`
}

func (m *SaveRegionRequest) Reset()                    { *m = SaveRegionRequest{} }
func (m *SaveRegionRequest) String() string            { return proto.CompactTextString(m) }
func (*SaveRegionRequest) ProtoMessage()               {}
func (*SaveRegionRequest) Descriptor() ([]byte, []int) { return fileDescriptorRegionService, []int{4} }

func (m *SaveRegionRequest) GetIdentify() *identify.Identify {
	if m != nil {
		return m.Identify
	}
	return nil
}

func (m *SaveRegionRequest) GetDataKeyspace() []*Region_RegionDataKeyspace {
	if m != nil {
		return m.DataKeyspace
	}
	return nil
}

func (m *SaveRegionRequest) GetLocation() *common.GeoLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

type SaveRegionResponse struct {
	Changed bool `protobuf:"varint,1,opt,name=changed,proto3" json:"changed,omitempty"`
}

func (m *SaveRegionResponse) Reset()                    { *m = SaveRegionResponse{} }
func (m *SaveRegionResponse) String() string            { return proto.CompactTextString(m) }
func (*SaveRegionResponse) ProtoMessage()               {}
func (*SaveRegionResponse) Descriptor() ([]byte, []int) { return fileDescriptorRegionService, []int{5} }

func init() {
	proto.RegisterType((*ListRegionsRequest)(nil), "region.ListRegionsRequest")
	proto.RegisterType((*ListRegionsResponse)(nil), "region.ListRegionsResponse")
	proto.RegisterType((*CreateRegionRequest)(nil), "region.CreateRegionRequest")
	proto.RegisterType((*CreateRegionResponse)(nil), "region.CreateRegionResponse")
	proto.RegisterType((*SaveRegionRequest)(nil), "region.SaveRegionRequest")
	proto.RegisterType((*SaveRegionResponse)(nil), "region.SaveRegionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for RegionService service

type RegionServiceClient interface {
	ListRegions(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*ListRegionsResponse, error)
	CreateRegion(ctx context.Context, in *CreateRegionRequest, opts ...grpc.CallOption) (*CreateRegionResponse, error)
	SaveRegion(ctx context.Context, in *SaveRegionRequest, opts ...grpc.CallOption) (*SaveRegionResponse, error)
}

type regionServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegionServiceClient(cc *grpc.ClientConn) RegionServiceClient {
	return &regionServiceClient{cc}
}

func (c *regionServiceClient) ListRegions(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*ListRegionsResponse, error) {
	out := new(ListRegionsResponse)
	err := grpc.Invoke(ctx, "/region.RegionService/ListRegions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) CreateRegion(ctx context.Context, in *CreateRegionRequest, opts ...grpc.CallOption) (*CreateRegionResponse, error) {
	out := new(CreateRegionResponse)
	err := grpc.Invoke(ctx, "/region.RegionService/CreateRegion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) SaveRegion(ctx context.Context, in *SaveRegionRequest, opts ...grpc.CallOption) (*SaveRegionResponse, error) {
	out := new(SaveRegionResponse)
	err := grpc.Invoke(ctx, "/region.RegionService/SaveRegion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RegionService service

type RegionServiceServer interface {
	ListRegions(context.Context, *ListRegionsRequest) (*ListRegionsResponse, error)
	CreateRegion(context.Context, *CreateRegionRequest) (*CreateRegionResponse, error)
	SaveRegion(context.Context, *SaveRegionRequest) (*SaveRegionResponse, error)
}

func RegisterRegionServiceServer(s *grpc.Server, srv RegionServiceServer) {
	s.RegisterService(&_RegionService_serviceDesc, srv)
}

func _RegionService_ListRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).ListRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.RegionService/ListRegions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).ListRegions(ctx, req.(*ListRegionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_CreateRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).CreateRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.RegionService/CreateRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).CreateRegion(ctx, req.(*CreateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_SaveRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).SaveRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.RegionService/SaveRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).SaveRegion(ctx, req.(*SaveRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "region.RegionService",
	HandlerType: (*RegionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRegions",
			Handler:    _RegionService_ListRegions_Handler,
		},
		{
			MethodName: "CreateRegion",
			Handler:    _RegionService_CreateRegion_Handler,
		},
		{
			MethodName: "SaveRegion",
			Handler:    _RegionService_SaveRegion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorRegionService,
}

func (m *ListRegionsRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ListRegionsRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Identify != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRegionService(data, i, uint64(m.Identify.Size()))
		n1, err := m.Identify.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *ListRegionsResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ListRegionsResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Region) > 0 {
		for _, msg := range m.Region {
			data[i] = 0xa
			i++
			i = encodeVarintRegionService(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CreateRegionRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CreateRegionRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Region != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRegionService(data, i, uint64(m.Region.Size()))
		n2, err := m.Region.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *CreateRegionResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CreateRegionResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CreatedRegion != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRegionService(data, i, uint64(m.CreatedRegion.Size()))
		n3, err := m.CreatedRegion.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *SaveRegionRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SaveRegionRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Identify != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRegionService(data, i, uint64(m.Identify.Size()))
		n4, err := m.Identify.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.Id) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintRegionService(data, i, uint64(len(m.Id)))
		i += copy(data[i:], m.Id)
	}
	if len(m.Name) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintRegionService(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if len(m.DataKeyspace) > 0 {
		for _, msg := range m.DataKeyspace {
			data[i] = 0x2a
			i++
			i = encodeVarintRegionService(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Location != nil {
		data[i] = 0x32
		i++
		i = encodeVarintRegionService(data, i, uint64(m.Location.Size()))
		n5, err := m.Location.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.AllowClearDataKeyspaces {
		data[i] = 0x38
		i++
		if m.AllowClearDataKeyspaces {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *SaveRegionResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SaveRegionResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Changed {
		data[i] = 0x8
		i++
		if m.Changed {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64RegionService(data []byte, offset int, v uint64) int {
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
func encodeFixed32RegionService(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRegionService(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ListRegionsRequest) Size() (n int) {
	var l int
	_ = l
	if m.Identify != nil {
		l = m.Identify.Size()
		n += 1 + l + sovRegionService(uint64(l))
	}
	return n
}

func (m *ListRegionsResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Region) > 0 {
		for _, e := range m.Region {
			l = e.Size()
			n += 1 + l + sovRegionService(uint64(l))
		}
	}
	return n
}

func (m *CreateRegionRequest) Size() (n int) {
	var l int
	_ = l
	if m.Region != nil {
		l = m.Region.Size()
		n += 1 + l + sovRegionService(uint64(l))
	}
	return n
}

func (m *CreateRegionResponse) Size() (n int) {
	var l int
	_ = l
	if m.CreatedRegion != nil {
		l = m.CreatedRegion.Size()
		n += 1 + l + sovRegionService(uint64(l))
	}
	return n
}

func (m *SaveRegionRequest) Size() (n int) {
	var l int
	_ = l
	if m.Identify != nil {
		l = m.Identify.Size()
		n += 1 + l + sovRegionService(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovRegionService(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRegionService(uint64(l))
	}
	if len(m.DataKeyspace) > 0 {
		for _, e := range m.DataKeyspace {
			l = e.Size()
			n += 1 + l + sovRegionService(uint64(l))
		}
	}
	if m.Location != nil {
		l = m.Location.Size()
		n += 1 + l + sovRegionService(uint64(l))
	}
	if m.AllowClearDataKeyspaces {
		n += 2
	}
	return n
}

func (m *SaveRegionResponse) Size() (n int) {
	var l int
	_ = l
	if m.Changed {
		n += 2
	}
	return n
}

func sovRegionService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRegionService(x uint64) (n int) {
	return sovRegionService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListRegionsRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegionService
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
			return fmt.Errorf("proto: ListRegionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListRegionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identify", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegionService
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
				return ErrInvalidLengthRegionService
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
		default:
			iNdEx = preIndex
			skippy, err := skipRegionService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegionService
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
func (m *ListRegionsResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegionService
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
			return fmt.Errorf("proto: ListRegionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListRegionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Region", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegionService
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
				return ErrInvalidLengthRegionService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Region = append(m.Region, &Region{})
			if err := m.Region[len(m.Region)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegionService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegionService
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
func (m *CreateRegionRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegionService
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
			return fmt.Errorf("proto: CreateRegionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRegionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Region", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegionService
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
				return ErrInvalidLengthRegionService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Region == nil {
				m.Region = &Region{}
			}
			if err := m.Region.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegionService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegionService
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
func (m *CreateRegionResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegionService
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
			return fmt.Errorf("proto: CreateRegionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRegionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedRegion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegionService
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
				return ErrInvalidLengthRegionService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedRegion == nil {
				m.CreatedRegion = &Region{}
			}
			if err := m.CreatedRegion.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegionService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegionService
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
func (m *SaveRegionRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegionService
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
			return fmt.Errorf("proto: SaveRegionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SaveRegionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identify", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegionService
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
				return ErrInvalidLengthRegionService
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegionService
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
				return ErrInvalidLengthRegionService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegionService
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
				return ErrInvalidLengthRegionService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataKeyspace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegionService
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
				return ErrInvalidLengthRegionService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataKeyspace = append(m.DataKeyspace, &Region_RegionDataKeyspace{})
			if err := m.DataKeyspace[len(m.DataKeyspace)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegionService
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
				return ErrInvalidLengthRegionService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Location == nil {
				m.Location = &common.GeoLocation{}
			}
			if err := m.Location.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowClearDataKeyspaces", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegionService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowClearDataKeyspaces = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRegionService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegionService
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
func (m *SaveRegionResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegionService
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
			return fmt.Errorf("proto: SaveRegionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SaveRegionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegionService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Changed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRegionService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegionService
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
func skipRegionService(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRegionService
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
					return 0, ErrIntOverflowRegionService
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
					return 0, ErrIntOverflowRegionService
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
				return 0, ErrInvalidLengthRegionService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRegionService
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
				next, err := skipRegionService(data[start:])
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
	ErrInvalidLengthRegionService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRegionService   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorRegionService = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0xb6, 0xd1, 0x95, 0xbb, 0xb6, 0x12, 0x1e, 0x12, 0xc1, 0x20, 0x04, 0x7d, 0x40, 0xbc,
	0x2c, 0x15, 0x43, 0x7b, 0x40, 0x88, 0x17, 0x3a, 0x81, 0x26, 0xc6, 0x8b, 0xf7, 0x01, 0x95, 0xeb,
	0xb8, 0x99, 0x45, 0x5b, 0x97, 0xd8, 0x19, 0xda, 0xc7, 0xf0, 0x3f, 0x3c, 0xf2, 0x09, 0x88, 0x17,
	0x7e, 0x03, 0xc7, 0xbe, 0x09, 0x09, 0x8b, 0x2a, 0xf6, 0x10, 0xf9, 0xfa, 0xde, 0x73, 0x8e, 0x8f,
	0x6f, 0xae, 0xe1, 0x75, 0xa6, 0xec, 0x65, 0x31, 0x4f, 0x84, 0x5e, 0x4d, 0x16, 0x85, 0x91, 0xb9,
	0x9e, 0x6b, 0xab, 0x84, 0x99, 0x6c, 0x72, 0x6d, 0xf5, 0x24, 0x97, 0x99, 0xd2, 0x6b, 0x5c, 0x8e,
	0x5c, 0xf9, 0x4a, 0x09, 0x99, 0xf8, 0x1a, 0xe9, 0x85, 0x2c, 0x3d, 0x6a, 0x48, 0x64, 0x3a, 0xd3,
	0x81, 0x3a, 0x2f, 0x16, 0x7e, 0x17, 0x74, 0xca, 0x28, 0xd0, 0xe8, 0xc9, 0xf6, 0x13, 0x55, 0x2a,
	0xd7, 0x56, 0x2d, 0xae, 0xeb, 0x00, 0x69, 0x2f, 0x6f, 0x61, 0xf4, 0xff, 0x28, 0x2e, 0xbd, 0x72,
	0x94, 0xb0, 0x04, 0xca, 0xf8, 0x14, 0xc8, 0xb9, 0x32, 0x96, 0x79, 0x19, 0xc3, 0xe4, 0x97, 0x42,
	0x1a, 0x4b, 0x12, 0xe8, 0x57, 0x6e, 0xe2, 0xe8, 0x69, 0xf4, 0xe2, 0xe0, 0x98, 0x24, 0xb5, 0xbd,
	0x33, 0x0c, 0x58, 0x8d, 0x19, 0xbf, 0x85, 0xc3, 0x96, 0x8a, 0xd9, 0xb8, 0x45, 0x92, 0xe7, 0x80,
	0x2d, 0x73, 0x22, 0xbb, 0x4e, 0x64, 0x94, 0xa0, 0xdd, 0x00, 0x64, 0x58, 0x2d, 0xe9, 0xd3, 0x5c,
	0x72, 0x2b, 0x31, 0x8f, 0x2e, 0x9a, 0xf4, 0x68, 0x0b, 0xfd, 0x13, 0xdc, 0x6f, 0xd3, 0xf1, 0xf8,
	0x13, 0x18, 0x09, 0x9f, 0x4f, 0x67, 0x5b, 0x75, 0x86, 0x88, 0x0a, 0xdb, 0xf1, 0xb7, 0x1d, 0xb8,
	0x77, 0xc1, 0xaf, 0xfe, 0x31, 0x73, 0xcb, 0x96, 0x90, 0x11, 0xec, 0xa8, 0x34, 0xde, 0x75, 0xc8,
	0xbb, 0xcc, 0x45, 0x84, 0xc0, 0xde, 0x9a, 0xaf, 0x64, 0xbc, 0xe7, 0x33, 0x3e, 0x26, 0xef, 0x61,
	0x98, 0x72, 0xcb, 0x67, 0x9f, 0xe5, 0xb5, 0xd9, 0x70, 0x21, 0xe3, 0x3b, 0xbe, 0x4d, 0xcf, 0xda,
	0xfe, 0x70, 0x39, 0x75, 0xc8, 0x8f, 0x08, 0x64, 0x83, 0xb4, 0xb1, 0x23, 0x13, 0xe8, 0x2f, 0xb5,
	0xe0, 0xb6, 0xbc, 0x62, 0xcf, 0x7b, 0x3b, 0x4c, 0xf0, 0x2f, 0x7f, 0x90, 0xfa, 0x1c, 0x4b, 0xac,
	0x06, 0x91, 0x37, 0x40, 0xf9, 0x72, 0xa9, 0xbf, 0xce, 0xc4, 0x52, 0xf2, 0x7c, 0xd6, 0x32, 0x61,
	0xe2, 0x7d, 0x27, 0xd1, 0x67, 0x0f, 0x3c, 0x62, 0x5a, 0x02, 0x9a, 0x47, 0x9b, 0x71, 0x02, 0xa4,
	0xd9, 0x1e, 0x6c, 0x76, 0x0c, 0xfb, 0xe2, 0x92, 0xaf, 0x33, 0x99, 0xfa, 0xf6, 0xf4, 0x59, 0xb5,
	0x3d, 0xfe, 0x1d, 0xc1, 0x30, 0x80, 0x2f, 0xc2, 0x73, 0x72, 0xf7, 0x3e, 0x68, 0x8c, 0x0b, 0xa1,
	0xd5, 0x7d, 0x6f, 0x4e, 0x22, 0x7d, 0xd4, 0x59, 0xc3, 0x33, 0xcf, 0x60, 0xd0, 0xfc, 0xf1, 0xa4,
	0x06, 0x77, 0x4c, 0x13, 0x7d, 0xdc, 0x5d, 0x44, 0xa9, 0x29, 0xc0, 0xdf, 0x4b, 0x91, 0x87, 0x15,
	0xf6, 0xc6, 0x1c, 0x50, 0xda, 0x55, 0x0a, 0x22, 0xef, 0x06, 0xdf, 0x7f, 0x3d, 0x89, 0x7e, 0xb8,
	0xef, 0xa7, 0xfb, 0xe6, 0x3d, 0xff, 0xc2, 0x5e, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xea, 0xe2,
	0xb8, 0xca, 0x72, 0x04, 0x00, 0x00,
}
