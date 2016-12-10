// Code generated by protoc-gen-gogo.
// source: github.com/fuserobotics/proto/device/device-service.proto
// DO NOT EDIT!

/*
Package device is a generated protocol buffer package.

It is generated from these files:
	github.com/fuserobotics/proto/device/device-service.proto
	github.com/fuserobotics/proto/device/device.proto

It has these top-level messages:
	ListDevicesRequest
	ListDevicesResponse
	CreateDeviceRequest
	CreateDeviceResponse
	SaveDeviceRequest
	SaveDeviceResponse
	GetDeviceRequest
	GetDeviceResponse
	GetConsulPeersRequest
	GetConsulPeersResponse
	CheckBootstrapperExistsRequest
	CheckBootstrapperExistsResponse
	Device
	DeviceNetworkTemplate
*/
package device

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/fuserobotics/proto/common"
import identify "github.com/fuserobotics/proto/identify"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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
	Region   string             `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
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

type SaveDeviceRequest struct {
	Identify        *identify.Identify            `protobuf:"bytes,1,opt,name=identify" json:"identify,omitempty"`
	Hostname        string                        `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Region          string                        `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	NetworkSettings *Device_DeviceNetworkSettings `protobuf:"bytes,4,opt,name=network_settings,json=networkSettings" json:"network_settings,omitempty"`
}

func (m *SaveDeviceRequest) Reset()                    { *m = SaveDeviceRequest{} }
func (m *SaveDeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*SaveDeviceRequest) ProtoMessage()               {}
func (*SaveDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceService, []int{4} }

func (m *SaveDeviceRequest) GetIdentify() *identify.Identify {
	if m != nil {
		return m.Identify
	}
	return nil
}

func (m *SaveDeviceRequest) GetNetworkSettings() *Device_DeviceNetworkSettings {
	if m != nil {
		return m.NetworkSettings
	}
	return nil
}

type SaveDeviceResponse struct {
	Changed bool `protobuf:"varint,1,opt,name=changed,proto3" json:"changed,omitempty"`
}

func (m *SaveDeviceResponse) Reset()                    { *m = SaveDeviceResponse{} }
func (m *SaveDeviceResponse) String() string            { return proto.CompactTextString(m) }
func (*SaveDeviceResponse) ProtoMessage()               {}
func (*SaveDeviceResponse) Descriptor() ([]byte, []int) { return fileDescriptorDeviceService, []int{5} }

type GetDeviceRequest struct {
	Identify *identify.Identify `protobuf:"bytes,1,opt,name=identify" json:"identify,omitempty"`
	Hostname string             `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Region   string             `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
}

func (m *GetDeviceRequest) Reset()                    { *m = GetDeviceRequest{} }
func (m *GetDeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDeviceRequest) ProtoMessage()               {}
func (*GetDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptorDeviceService, []int{6} }

func (m *GetDeviceRequest) GetIdentify() *identify.Identify {
	if m != nil {
		return m.Identify
	}
	return nil
}

type GetDeviceResponse struct {
	Device *Device `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
}

func (m *GetDeviceResponse) Reset()                    { *m = GetDeviceResponse{} }
func (m *GetDeviceResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDeviceResponse) ProtoMessage()               {}
func (*GetDeviceResponse) Descriptor() ([]byte, []int) { return fileDescriptorDeviceService, []int{7} }

func (m *GetDeviceResponse) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type GetConsulPeersRequest struct {
	Identify *identify.Identify `protobuf:"bytes,1,opt,name=identify" json:"identify,omitempty"`
}

func (m *GetConsulPeersRequest) Reset()         { *m = GetConsulPeersRequest{} }
func (m *GetConsulPeersRequest) String() string { return proto.CompactTextString(m) }
func (*GetConsulPeersRequest) ProtoMessage()    {}
func (*GetConsulPeersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorDeviceService, []int{8}
}

func (m *GetConsulPeersRequest) GetIdentify() *identify.Identify {
	if m != nil {
		return m.Identify
	}
	return nil
}

type GetConsulPeersResponse struct {
	ServerPeer []*common.IPAddress `protobuf:"bytes,1,rep,name=server_peer,json=serverPeer" json:"server_peer,omitempty"`
	AgentPeer  []*common.IPAddress `protobuf:"bytes,2,rep,name=agent_peer,json=agentPeer" json:"agent_peer,omitempty"`
	WanPeer    []*common.IPAddress `protobuf:"bytes,3,rep,name=wan_peer,json=wanPeer" json:"wan_peer,omitempty"`
}

func (m *GetConsulPeersResponse) Reset()         { *m = GetConsulPeersResponse{} }
func (m *GetConsulPeersResponse) String() string { return proto.CompactTextString(m) }
func (*GetConsulPeersResponse) ProtoMessage()    {}
func (*GetConsulPeersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorDeviceService, []int{9}
}

func (m *GetConsulPeersResponse) GetServerPeer() []*common.IPAddress {
	if m != nil {
		return m.ServerPeer
	}
	return nil
}

func (m *GetConsulPeersResponse) GetAgentPeer() []*common.IPAddress {
	if m != nil {
		return m.AgentPeer
	}
	return nil
}

func (m *GetConsulPeersResponse) GetWanPeer() []*common.IPAddress {
	if m != nil {
		return m.WanPeer
	}
	return nil
}

type CheckBootstrapperExistsRequest struct {
	Identify *identify.Identify `protobuf:"bytes,1,opt,name=identify" json:"identify,omitempty"`
	Region   string             `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
}

func (m *CheckBootstrapperExistsRequest) Reset()         { *m = CheckBootstrapperExistsRequest{} }
func (m *CheckBootstrapperExistsRequest) String() string { return proto.CompactTextString(m) }
func (*CheckBootstrapperExistsRequest) ProtoMessage()    {}
func (*CheckBootstrapperExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorDeviceService, []int{10}
}

func (m *CheckBootstrapperExistsRequest) GetIdentify() *identify.Identify {
	if m != nil {
		return m.Identify
	}
	return nil
}

type CheckBootstrapperExistsResponse struct {
	ChangesMade bool `protobuf:"varint,1,opt,name=changes_made,json=changesMade,proto3" json:"changes_made,omitempty"`
}

func (m *CheckBootstrapperExistsResponse) Reset()         { *m = CheckBootstrapperExistsResponse{} }
func (m *CheckBootstrapperExistsResponse) String() string { return proto.CompactTextString(m) }
func (*CheckBootstrapperExistsResponse) ProtoMessage()    {}
func (*CheckBootstrapperExistsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorDeviceService, []int{11}
}

func init() {
	proto.RegisterType((*ListDevicesRequest)(nil), "device.ListDevicesRequest")
	proto.RegisterType((*ListDevicesResponse)(nil), "device.ListDevicesResponse")
	proto.RegisterType((*CreateDeviceRequest)(nil), "device.CreateDeviceRequest")
	proto.RegisterType((*CreateDeviceResponse)(nil), "device.CreateDeviceResponse")
	proto.RegisterType((*SaveDeviceRequest)(nil), "device.SaveDeviceRequest")
	proto.RegisterType((*SaveDeviceResponse)(nil), "device.SaveDeviceResponse")
	proto.RegisterType((*GetDeviceRequest)(nil), "device.GetDeviceRequest")
	proto.RegisterType((*GetDeviceResponse)(nil), "device.GetDeviceResponse")
	proto.RegisterType((*GetConsulPeersRequest)(nil), "device.GetConsulPeersRequest")
	proto.RegisterType((*GetConsulPeersResponse)(nil), "device.GetConsulPeersResponse")
	proto.RegisterType((*CheckBootstrapperExistsRequest)(nil), "device.CheckBootstrapperExistsRequest")
	proto.RegisterType((*CheckBootstrapperExistsResponse)(nil), "device.CheckBootstrapperExistsResponse")
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
	SaveDevice(ctx context.Context, in *SaveDeviceRequest, opts ...grpc.CallOption) (*SaveDeviceResponse, error)
	GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error)
	// Methods called by devices
	GetConsulPeers(ctx context.Context, in *GetConsulPeersRequest, opts ...grpc.CallOption) (*GetConsulPeersResponse, error)
	CheckBootstrapperExists(ctx context.Context, in *CheckBootstrapperExistsRequest, opts ...grpc.CallOption) (*CheckBootstrapperExistsResponse, error)
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

func (c *globalDeviceServiceClient) SaveDevice(ctx context.Context, in *SaveDeviceRequest, opts ...grpc.CallOption) (*SaveDeviceResponse, error) {
	out := new(SaveDeviceResponse)
	err := grpc.Invoke(ctx, "/device.GlobalDeviceService/SaveDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalDeviceServiceClient) GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error) {
	out := new(GetDeviceResponse)
	err := grpc.Invoke(ctx, "/device.GlobalDeviceService/GetDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalDeviceServiceClient) GetConsulPeers(ctx context.Context, in *GetConsulPeersRequest, opts ...grpc.CallOption) (*GetConsulPeersResponse, error) {
	out := new(GetConsulPeersResponse)
	err := grpc.Invoke(ctx, "/device.GlobalDeviceService/GetConsulPeers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalDeviceServiceClient) CheckBootstrapperExists(ctx context.Context, in *CheckBootstrapperExistsRequest, opts ...grpc.CallOption) (*CheckBootstrapperExistsResponse, error) {
	out := new(CheckBootstrapperExistsResponse)
	err := grpc.Invoke(ctx, "/device.GlobalDeviceService/CheckBootstrapperExists", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GlobalDeviceService service

type GlobalDeviceServiceServer interface {
	ListDevices(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error)
	CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceResponse, error)
	SaveDevice(context.Context, *SaveDeviceRequest) (*SaveDeviceResponse, error)
	GetDevice(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error)
	// Methods called by devices
	GetConsulPeers(context.Context, *GetConsulPeersRequest) (*GetConsulPeersResponse, error)
	CheckBootstrapperExists(context.Context, *CheckBootstrapperExistsRequest) (*CheckBootstrapperExistsResponse, error)
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

func _GlobalDeviceService_SaveDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalDeviceServiceServer).SaveDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.GlobalDeviceService/SaveDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalDeviceServiceServer).SaveDevice(ctx, req.(*SaveDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalDeviceService_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalDeviceServiceServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.GlobalDeviceService/GetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalDeviceServiceServer).GetDevice(ctx, req.(*GetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalDeviceService_GetConsulPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsulPeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalDeviceServiceServer).GetConsulPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.GlobalDeviceService/GetConsulPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalDeviceServiceServer).GetConsulPeers(ctx, req.(*GetConsulPeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalDeviceService_CheckBootstrapperExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckBootstrapperExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalDeviceServiceServer).CheckBootstrapperExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.GlobalDeviceService/CheckBootstrapperExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalDeviceServiceServer).CheckBootstrapperExists(ctx, req.(*CheckBootstrapperExistsRequest))
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
		{
			MethodName: "SaveDevice",
			Handler:    _GlobalDeviceService_SaveDevice_Handler,
		},
		{
			MethodName: "GetDevice",
			Handler:    _GlobalDeviceService_GetDevice_Handler,
		},
		{
			MethodName: "GetConsulPeers",
			Handler:    _GlobalDeviceService_GetConsulPeers_Handler,
		},
		{
			MethodName: "CheckBootstrapperExists",
			Handler:    _GlobalDeviceService_CheckBootstrapperExists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorDeviceService,
}

var fileDescriptorDeviceService = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x55, 0xe1, 0x6e, 0xd3, 0x30,
	0x10, 0x56, 0x57, 0xd4, 0xb5, 0xd7, 0x51, 0x56, 0x17, 0x46, 0x67, 0x60, 0x40, 0x84, 0x06, 0x3f,
	0x20, 0x65, 0x45, 0xfb, 0x81, 0x10, 0x12, 0xd0, 0x41, 0x55, 0x89, 0xb1, 0x29, 0xfd, 0x8b, 0x54,
	0xa5, 0x89, 0xd7, 0x86, 0xb5, 0x71, 0x89, 0xdd, 0x0e, 0x1e, 0x80, 0xb7, 0xe1, 0x41, 0x78, 0x18,
	0x1e, 0x02, 0x37, 0xbe, 0x66, 0xc9, 0x9a, 0x14, 0x31, 0x10, 0xbf, 0x1c, 0xdf, 0x7d, 0xdf, 0x77,
	0x77, 0xf6, 0xf9, 0x02, 0xcf, 0x07, 0x9e, 0x1c, 0x4e, 0xfb, 0xa6, 0xc3, 0xc7, 0x8d, 0x93, 0xa9,
	0x60, 0x01, 0xef, 0x73, 0xe9, 0x39, 0xa2, 0x31, 0x09, 0xb8, 0xe4, 0x0d, 0x97, 0xcd, 0x3c, 0x87,
	0xe1, 0xf2, 0x44, 0xb9, 0xe7, 0xab, 0x19, 0xfa, 0x48, 0x41, 0x5b, 0xe9, 0xde, 0x6a, 0x09, 0x65,
	0x1e, 0x73, 0x1f, 0x17, 0x4d, 0xa5, 0xfb, 0xab, 0x29, 0x9e, 0xcb, 0x7c, 0xe9, 0x9d, 0x7c, 0x8d,
	0x3e, 0x90, 0xb6, 0xf7, 0x07, 0xc9, 0x6a, 0x8a, 0xf1, 0x11, 0xc8, 0x7b, 0x4f, 0xc8, 0x83, 0xd0,
	0x26, 0x2c, 0xf6, 0x79, 0xca, 0x84, 0x24, 0x26, 0x14, 0x17, 0xd2, 0xf5, 0xdc, 0xbd, 0xdc, 0xa3,
	0x72, 0x93, 0x98, 0x51, 0xac, 0x0e, 0x7e, 0x58, 0x11, 0x86, 0x6c, 0x41, 0x21, 0x60, 0x03, 0x8f,
	0xfb, 0xf5, 0x35, 0x85, 0x2e, 0x59, 0xb8, 0x33, 0x5e, 0x42, 0x2d, 0xa1, 0x2e, 0x26, 0xdc, 0x17,
	0x8c, 0xec, 0x02, 0x9e, 0x8d, 0x12, 0xcf, 0x2b, 0xf1, 0x8a, 0x89, 0x39, 0x69, 0xa0, 0x85, 0x5e,
	0xe3, 0x5b, 0x0e, 0x6a, 0xad, 0x80, 0xd9, 0x92, 0xa1, 0xe3, 0x92, 0xe9, 0x9d, 0xc7, 0x5b, 0x0b,
	0xd1, 0x19, 0xf1, 0x62, 0x65, 0xe4, 0x13, 0x65, 0x1c, 0xc2, 0xf5, 0x64, 0x1a, 0x58, 0xc7, 0x3e,
	0x54, 0x9c, 0xd0, 0xee, 0xf6, 0xa2, 0x7a, 0xd2, 0xf4, 0xaf, 0x22, 0x4a, 0x6f, 0x8d, 0x1f, 0x39,
	0xa8, 0x76, 0xed, 0xd9, 0x5f, 0x16, 0x45, 0xa1, 0x38, 0xe4, 0x42, 0xfa, 0xf6, 0x98, 0xe1, 0xa9,
	0x47, 0xfb, 0xac, 0x42, 0xc8, 0x11, 0x6c, 0xfa, 0x4c, 0x9e, 0xf1, 0xe0, 0xb4, 0x27, 0x98, 0x94,
	0x9e, 0x3f, 0x10, 0xf5, 0x2b, 0x61, 0xac, 0x07, 0xc9, 0x94, 0x71, 0xf9, 0xa0, 0xc1, 0x5d, 0xc4,
	0x5a, 0xd7, 0xfc, 0xa4, 0xc1, 0x30, 0x81, 0xc4, 0x2b, 0xc1, 0x73, 0xa9, 0xc3, 0xba, 0x33, 0xb4,
	0xfd, 0x01, 0x73, 0xc3, 0x4a, 0x8a, 0xd6, 0x62, 0x6b, 0xcc, 0x60, 0xb3, 0xcd, 0xe4, 0x7f, 0x2f,
	0xdc, 0x78, 0x01, 0xd5, 0x58, 0xdc, 0x94, 0x36, 0x5c, 0xd1, 0x16, 0x46, 0x1b, 0x6e, 0x28, 0x72,
	0x4b, 0x71, 0xa6, 0xa3, 0x63, 0xc6, 0x82, 0xcb, 0x3e, 0x13, 0xe3, 0x7b, 0x0e, 0xb6, 0x2e, 0x2a,
	0x61, 0x2e, 0x4d, 0x28, 0xcf, 0xa7, 0x07, 0x0b, 0x7a, 0x13, 0x65, 0xc7, 0x77, 0x51, 0x35, 0x71,
	0x2a, 0x74, 0x8e, 0x5f, 0xbb, 0x6e, 0xc0, 0x84, 0xb0, 0x40, 0xa3, 0xe6, 0x64, 0xf2, 0x14, 0xc0,
	0x1e, 0x28, 0x69, 0x4d, 0x59, 0xcb, 0xa2, 0x94, 0x42, 0x50, 0xc8, 0x78, 0x0c, 0xc5, 0x33, 0xdb,
	0xd7, 0xf8, 0x7c, 0x16, 0x7e, 0x5d, 0x41, 0xe6, 0x68, 0x63, 0x08, 0x3b, 0xad, 0x21, 0x73, 0x4e,
	0xdf, 0x70, 0x2e, 0x85, 0x0c, 0xec, 0xc9, 0x84, 0x05, 0x6f, 0xbf, 0xa8, 0xf7, 0xfc, 0xcf, 0xe7,
	0xc4, 0x01, 0xdc, 0xcd, 0x8c, 0x84, 0x07, 0x74, 0x1f, 0x36, 0x74, 0x13, 0x89, 0xde, 0xd8, 0x76,
	0x19, 0x36, 0x56, 0x19, 0x6d, 0x87, 0xca, 0xd4, 0xfc, 0x99, 0x87, 0x5a, 0x7b, 0xc4, 0xfb, 0xf6,
	0x48, 0x5f, 0x60, 0x57, 0x8f, 0x63, 0xf2, 0x0e, 0xca, 0xb1, 0x29, 0x44, 0xe8, 0xe2, 0x9a, 0x97,
	0x07, 0x1f, 0xbd, 0x95, 0xea, 0xc3, 0x14, 0x3a, 0xb0, 0x11, 0x1f, 0x03, 0x24, 0x02, 0xa7, 0xcc,
	0x28, 0x7a, 0x3b, 0xdd, 0x89, 0x52, 0x2d, 0x80, 0xf3, 0x77, 0x43, 0xb6, 0x17, 0xd8, 0xa5, 0xa9,
	0x40, 0x69, 0x9a, 0x0b, 0x45, 0x5e, 0x41, 0x29, 0x6a, 0x6a, 0x52, 0x5f, 0x00, 0x2f, 0xbe, 0x2f,
	0xba, 0x9d, 0xe2, 0x41, 0x85, 0x23, 0xa8, 0x24, 0xfb, 0x91, 0xdc, 0x89, 0x81, 0x97, 0x3b, 0x9e,
	0xee, 0x64, 0xb9, 0x51, 0xf0, 0x13, 0xdc, 0xcc, 0xb8, 0x48, 0xb2, 0x1b, 0x1d, 0xc8, 0xca, 0x9e,
	0xa2, 0x0f, 0x7f, 0x8b, 0xd3, 0xb1, 0xfa, 0x85, 0xf0, 0x0f, 0xf6, 0xec, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0xd3, 0x9b, 0x93, 0xa3, 0x07, 0x00, 0x00,
}
