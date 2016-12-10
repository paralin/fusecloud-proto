// Code generated by protoc-gen-gogo.
// source: github.com/fuserobotics/proto/permissions/permissions.proto
// DO NOT EDIT!

/*
Package permissions is a generated protocol buffer package.

It is generated from these files:
	github.com/fuserobotics/proto/permissions/permissions.proto

It has these top-level messages:
	SystemPermissions
*/
package permissions

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/fuserobotics/proto/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

//
// SystemPermissions: full list of possible permissions
type SystemPermissions struct {
	// all permissions
	AllPermissions bool `protobuf:"varint,1,opt,name=all_permissions,json=allPermissions,proto3" json:"all_permissions,omitempty"`
	// Create / delete regions
	ManageRegions bool `protobuf:"varint,2,opt,name=manage_regions,json=manageRegions,proto3" json:"manage_regions,omitempty"`
	// Create / delete devices
	ManageDevices bool `protobuf:"varint,3,opt,name=manage_devices,json=manageDevices,proto3" json:"manage_devices,omitempty"`
	// View / general basic access
	Access bool `protobuf:"varint,4,opt,name=access,proto3" json:"access,omitempty"`
}

func (m *SystemPermissions) Reset()                    { *m = SystemPermissions{} }
func (m *SystemPermissions) String() string            { return proto.CompactTextString(m) }
func (*SystemPermissions) ProtoMessage()               {}
func (*SystemPermissions) Descriptor() ([]byte, []int) { return fileDescriptorPermissions, []int{0} }

func init() {
	proto.RegisterType((*SystemPermissions)(nil), "permissions.SystemPermissions")
}

var fileDescriptorPermissions = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2b, 0x2d, 0x4e, 0x2d, 0xca, 0x4f, 0xca, 0x2f,
	0xc9, 0x4c, 0x2e, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x48, 0x2d, 0xca, 0xcd, 0x2c,
	0x2e, 0xce, 0xcc, 0xcf, 0x2b, 0x46, 0x66, 0xeb, 0x81, 0x65, 0x85, 0xb8, 0x91, 0x84, 0xa4, 0x0c,
	0xf1, 0x9b, 0x04, 0x14, 0xce, 0xcd, 0xcf, 0x83, 0x52, 0x10, 0xfd, 0x4a, 0x8b, 0x18, 0xb9, 0x04,
	0x83, 0x2b, 0x8b, 0x4b, 0x52, 0x73, 0x03, 0x10, 0x06, 0x09, 0xa9, 0x73, 0xf1, 0x27, 0xe6, 0xe4,
	0xc4, 0x23, 0x99, 0x2d, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x11, 0xc4, 0x07, 0x14, 0x46, 0x56, 0xa8,
	0xca, 0xc5, 0x97, 0x9b, 0x98, 0x97, 0x98, 0x9e, 0x1a, 0x5f, 0x94, 0x9a, 0x0e, 0x56, 0xc7, 0x04,
	0x56, 0xc7, 0x0b, 0x11, 0x0d, 0x82, 0x08, 0x22, 0x29, 0x4b, 0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0x2d,
	0x96, 0x60, 0x46, 0x56, 0xe6, 0x02, 0x11, 0x14, 0x12, 0xe3, 0x62, 0x4b, 0x4c, 0x06, 0x32, 0x8a,
	0x25, 0x58, 0xc0, 0xd2, 0x50, 0x5e, 0x12, 0x1b, 0xd8, 0xad, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfa, 0x31, 0x88, 0xe1, 0x2a, 0x01, 0x00, 0x00,
}
