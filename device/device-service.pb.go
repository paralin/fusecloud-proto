// Code generated by protoc-gen-gogo.
// source: github.com/synrobo/proto/device/device-service.proto
// DO NOT EDIT!

/*
	Package device is a generated protocol buffer package.

	It is generated from these files:
		github.com/synrobo/proto/device/device-service.proto
		github.com/synrobo/proto/device/device.proto

	It has these top-level messages:
		ListDevicesRequest
		ListDevicesResponse
		CreateDeviceRequest
		CreateDeviceResponse
		Device
*/
package device

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

type ListDevicesRequest struct {
	Identify *identify.Identify `protobuf:"bytes,1,opt,name=identify" json:"identify,omitempty"`
	Region   string             `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
}

func (m *ListDevicesRequest) Reset()                    { *m = ListDevicesRequest{} }
func (m *ListDevicesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDevicesRequest) ProtoMessage()               {}
func (*ListDevicesRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceService, []int{0} }

func (m *ListDevicesRequest) GetIdentify() *identify.Identify {
	if m != nil {
		return m.Identify
	}
	return nil
}

type ListDevicesResponse struct {
	Device []*Device `protobuf:"bytes,1,rep,name=device" json:"device,omitempty"`
}

func (m *ListDevicesResponse) Reset()                    { *m = ListDevicesResponse{} }
func (m *ListDevicesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDevicesResponse) ProtoMessage()               {}
func (*ListDevicesResponse) Descriptor() ([]byte, []int) { return fileDescriptorDeviceService, []int{1} }

func (m *ListDevicesResponse) GetDevice() []*Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type CreateDeviceRequest struct {
	Identify *identify.Identify `protobuf:"bytes,1,opt,name=identify" json:"identify,omitempty"`
	Device   *Device            `protobuf:"bytes,2,opt,name=device" json:"device,omitempty"`
}

func (m *CreateDeviceRequest) Reset()                    { *m = CreateDeviceRequest{} }
func (m *CreateDeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDeviceRequest) ProtoMessage()               {}
func (*CreateDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceService, []int{2} }

func (m *CreateDeviceRequest) GetIdentify() *identify.Identify {
	if m != nil {
		return m.Identify
	}
	return nil
}

func (m *CreateDeviceRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type CreateDeviceResponse struct {
	CreatedDevice *Device `protobuf:"bytes,1,opt,name=created_device,json=createdDevice" json:"created_device,omitempty"`
}

func (m *CreateDeviceResponse) Reset()         { *m = CreateDeviceResponse{} }
func (m *CreateDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDeviceResponse) ProtoMessage()    {}
func (*CreateDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorDeviceService, []int{3}
}

func (m *CreateDeviceResponse) GetCreatedDevice() *Device {
	if m != nil {
		return m.CreatedDevice
	}
	return nil
}

func init() {
	proto.RegisterType((*ListDevicesRequest)(nil), "device.ListDevicesRequest")
	proto.RegisterType((*ListDevicesResponse)(nil), "device.ListDevicesResponse")
	proto.RegisterType((*CreateDeviceRequest)(nil), "device.CreateDeviceRequest")
	proto.RegisterType((*CreateDeviceResponse)(nil), "device.CreateDeviceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for GlobalDeviceService service

type GlobalDeviceServiceClient interface {
	ListDevices(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error)
	CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceResponse, error)
}

type globalDeviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewGlobalDeviceServiceClient(cc *grpc.ClientConn) GlobalDeviceServiceClient {
	return &globalDeviceServiceClient{cc}
}

func (c *globalDeviceServiceClient) ListDevices(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error) {
	out := new(ListDevicesResponse)
	err := grpc.Invoke(ctx, "/device.GlobalDeviceService/ListDevices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalDeviceServiceClient) CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceResponse, error) {
	out := new(CreateDeviceResponse)
	err := grpc.Invoke(ctx, "/device.GlobalDeviceService/CreateDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GlobalDeviceService service

type GlobalDeviceServiceServer interface {
	ListDevices(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error)
	CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceResponse, error)
}

func RegisterGlobalDeviceServiceServer(s *grpc.Server, srv GlobalDeviceServiceServer) {
	s.RegisterService(&_GlobalDeviceService_serviceDesc, srv)
}

func _GlobalDeviceService_ListDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalDeviceServiceServer).ListDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.GlobalDeviceService/ListDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalDeviceServiceServer).ListDevices(ctx, req.(*ListDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalDeviceService_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalDeviceServiceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.GlobalDeviceService/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalDeviceServiceServer).CreateDevice(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GlobalDeviceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "device.GlobalDeviceService",
	HandlerType: (*GlobalDeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDevices",
			Handler:    _GlobalDeviceService_ListDevices_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _GlobalDeviceService_CreateDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorDeviceService,
}

func (m *ListDevicesRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ListDevicesRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Identify != nil {
		data[i] = 0xa
		i++
		i = encodeVarintDeviceService(data, i, uint64(m.Identify.Size()))
		n1, err := m.Identify.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Region) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintDeviceService(data, i, uint64(len(m.Region)))
		i += copy(data[i:], m.Region)
	}
	return i, nil
}

func (m *ListDevicesResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ListDevicesResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Device) > 0 {
		for _, msg := range m.Device {
			data[i] = 0xa
			i++
			i = encodeVarintDeviceService(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CreateDeviceRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CreateDeviceRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Identify != nil {
		data[i] = 0xa
		i++
		i = encodeVarintDeviceService(data, i, uint64(m.Identify.Size()))
		n2, err := m.Identify.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Device != nil {
		data[i] = 0x12
		i++
		i = encodeVarintDeviceService(data, i, uint64(m.Device.Size()))
		n3, err := m.Device.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *CreateDeviceResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CreateDeviceResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CreatedDevice != nil {
		data[i] = 0xa
		i++
		i = encodeVarintDeviceService(data, i, uint64(m.CreatedDevice.Size()))
		n4, err := m.CreatedDevice.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeFixed64DeviceService(data []byte, offset int, v uint64) int {
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
func encodeFixed32DeviceService(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDeviceService(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ListDevicesRequest) Size() (n int) {
	var l int
	_ = l
	if m.Identify != nil {
		l = m.Identify.Size()
		n += 1 + l + sovDeviceService(uint64(l))
	}
	l = len(m.Region)
	if l > 0 {
		n += 1 + l + sovDeviceService(uint64(l))
	}
	return n
}

func (m *ListDevicesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Device) > 0 {
		for _, e := range m.Device {
			l = e.Size()
			n += 1 + l + sovDeviceService(uint64(l))
		}
	}
	return n
}

func (m *CreateDeviceRequest) Size() (n int) {
	var l int
	_ = l
	if m.Identify != nil {
		l = m.Identify.Size()
		n += 1 + l + sovDeviceService(uint64(l))
	}
	if m.Device != nil {
		l = m.Device.Size()
		n += 1 + l + sovDeviceService(uint64(l))
	}
	return n
}

func (m *CreateDeviceResponse) Size() (n int) {
	var l int
	_ = l
	if m.CreatedDevice != nil {
		l = m.CreatedDevice.Size()
		n += 1 + l + sovDeviceService(uint64(l))
	}
	return n
}

func sovDeviceService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDeviceService(x uint64) (n int) {
	return sovDeviceService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListDevicesRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceService
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
			return fmt.Errorf("proto: ListDevicesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListDevicesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identify", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceService
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
				return ErrInvalidLengthDeviceService
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
					return ErrIntOverflowDeviceService
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
				return ErrInvalidLengthDeviceService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Region = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceService
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
func (m *ListDevicesResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceService
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
			return fmt.Errorf("proto: ListDevicesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListDevicesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceService
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
				return ErrInvalidLengthDeviceService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = append(m.Device, &Device{})
			if err := m.Device[len(m.Device)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceService
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
func (m *CreateDeviceRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceService
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
			return fmt.Errorf("proto: CreateDeviceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateDeviceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identify", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceService
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
				return ErrInvalidLengthDeviceService
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
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceService
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
				return ErrInvalidLengthDeviceService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Device == nil {
				m.Device = &Device{}
			}
			if err := m.Device.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceService
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
func (m *CreateDeviceResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceService
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
			return fmt.Errorf("proto: CreateDeviceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateDeviceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedDevice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceService
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
				return ErrInvalidLengthDeviceService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedDevice == nil {
				m.CreatedDevice = &Device{}
			}
			if err := m.CreatedDevice.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceService
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
func skipDeviceService(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceService
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
					return 0, ErrIntOverflowDeviceService
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
					return 0, ErrIntOverflowDeviceService
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
				return 0, ErrInvalidLengthDeviceService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeviceService
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
				next, err := skipDeviceService(data[start:])
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
	ErrInvalidLengthDeviceService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceService   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorDeviceService = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xae, 0xcc, 0x2b, 0xca, 0x4f, 0xca, 0xd7, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0x85, 0x52, 0xba, 0xc5, 0xa9,
	0x45, 0x20, 0x5a, 0x0f, 0x2c, 0x27, 0xc4, 0x06, 0x11, 0x95, 0xd2, 0x45, 0xd2, 0x9d, 0x9e, 0x9f,
	0x0e, 0xd5, 0x9a, 0x54, 0x9a, 0x06, 0xe6, 0x41, 0xcc, 0x01, 0xb1, 0x20, 0xda, 0xa4, 0x0c, 0x70,
	0x5a, 0x96, 0x99, 0x92, 0x9a, 0x57, 0x92, 0x99, 0x56, 0x09, 0x67, 0x40, 0x75, 0xe8, 0x10, 0xe7,
	0x3c, 0x88, 0x6a, 0xa5, 0x18, 0x2e, 0x21, 0x9f, 0xcc, 0xe2, 0x12, 0x17, 0xb0, 0x58, 0x71, 0x50,
	0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x1e, 0x17, 0x07, 0xcc, 0x54, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0x6e, 0x23, 0x21, 0x3d, 0xb8, 0x35, 0x9e, 0x50, 0x46, 0x10, 0x5c, 0x8d, 0x90, 0x18, 0x17,
	0x5b, 0x51, 0x6a, 0x7a, 0x66, 0x7e, 0x9e, 0x04, 0x13, 0x50, 0x35, 0x67, 0x10, 0x94, 0xa7, 0x64,
	0xcb, 0x25, 0x8c, 0x62, 0x7a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x1a, 0x17, 0x34, 0x34,
	0x80, 0x86, 0x33, 0x03, 0x0d, 0xe7, 0xd3, 0x83, 0xba, 0x09, 0xa2, 0x30, 0x08, 0x2a, 0xab, 0x94,
	0xcb, 0x25, 0xec, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x0a, 0x15, 0x27, 0xd3, 0x75, 0x08, 0xeb, 0x98,
	0xc0, 0xaa, 0x71, 0x59, 0xe7, 0xcb, 0x25, 0x82, 0x6a, 0x1d, 0xd4, 0xb9, 0xa6, 0x5c, 0x7c, 0xc9,
	0x60, 0xf1, 0x94, 0x78, 0xb8, 0xb3, 0xb1, 0x99, 0xc3, 0x0b, 0x55, 0x05, 0xe1, 0x1a, 0xad, 0x60,
	0xe4, 0x12, 0x76, 0xcf, 0xc9, 0x4f, 0x4a, 0xcc, 0x81, 0x08, 0x04, 0x43, 0xd2, 0x83, 0x90, 0x1b,
	0x17, 0x37, 0x52, 0xa0, 0x08, 0x49, 0xc1, 0x4c, 0xc1, 0x8c, 0x07, 0x29, 0x69, 0xac, 0x72, 0x50,
	0x67, 0x79, 0x72, 0xf1, 0x20, 0x3b, 0x57, 0x08, 0xae, 0x18, 0x4b, 0x98, 0x49, 0xc9, 0x60, 0x97,
	0x84, 0x18, 0xe5, 0xc4, 0x73, 0xe2, 0x91, 0x1c, 0xe3, 0x05, 0x20, 0x7e, 0x00, 0xc4, 0x49, 0x6c,
	0xe0, 0xa4, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x1d, 0xf8, 0x81, 0xe9, 0x02, 0x00,
	0x00,
}
