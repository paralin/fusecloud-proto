// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	User
	AuthWithPassword
	AuthWithPasswordSuccess
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Username string             `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string             `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Meta     *User_UserMetadata `protobuf:"bytes,3,opt,name=meta" json:"meta,omitempty"`
	Cert     []*User_UserCert   `protobuf:"bytes,4,rep,name=cert" json:"cert,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetMeta() *User_UserMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *User) GetCert() []*User_UserCert {
	if m != nil {
		return m.Cert
	}
	return nil
}

type User_UserMetadata struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *User_UserMetadata) Reset()                    { *m = User_UserMetadata{} }
func (m *User_UserMetadata) String() string            { return proto.CompactTextString(m) }
func (*User_UserMetadata) ProtoMessage()               {}
func (*User_UserMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type User_UserCert struct {
	Cert string `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
	Pkey string `protobuf:"bytes,2,opt,name=pkey" json:"pkey,omitempty"`
}

func (m *User_UserCert) Reset()                    { *m = User_UserCert{} }
func (m *User_UserCert) String() string            { return proto.CompactTextString(m) }
func (*User_UserCert) ProtoMessage()               {}
func (*User_UserCert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type AuthWithPassword struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Timestamp int32  `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *AuthWithPassword) Reset()                    { *m = AuthWithPassword{} }
func (m *AuthWithPassword) String() string            { return proto.CompactTextString(m) }
func (*AuthWithPassword) ProtoMessage()               {}
func (*AuthWithPassword) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AuthWithPasswordSuccess struct {
	Cert *User_UserCert     `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
	Meta *User_UserMetadata `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
}

func (m *AuthWithPasswordSuccess) Reset()                    { *m = AuthWithPasswordSuccess{} }
func (m *AuthWithPasswordSuccess) String() string            { return proto.CompactTextString(m) }
func (*AuthWithPasswordSuccess) ProtoMessage()               {}
func (*AuthWithPasswordSuccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthWithPasswordSuccess) GetCert() *User_UserCert {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *AuthWithPasswordSuccess) GetMeta() *User_UserMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*User_UserMetadata)(nil), "auth.User.UserMetadata")
	proto.RegisterType((*User_UserCert)(nil), "auth.User.UserCert")
	proto.RegisterType((*AuthWithPassword)(nil), "auth.AuthWithPassword")
	proto.RegisterType((*AuthWithPasswordSuccess)(nil), "auth.AuthWithPasswordSuccess")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x9a, 0x98, 0xb8, 0x58, 0x42,
	0x8b, 0x53, 0x8b, 0x84, 0xa4, 0xb8, 0x38, 0x4a, 0x81, 0x74, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0x0f, 0x92, 0x2b, 0x48, 0x2c, 0x2e, 0x2e, 0xcf, 0x2f, 0x4a,
	0x91, 0x60, 0x82, 0xc8, 0xc1, 0xf8, 0x42, 0xda, 0x5c, 0x2c, 0xb9, 0xa9, 0x25, 0x89, 0x12, 0xcc,
	0x40, 0x71, 0x6e, 0x23, 0x71, 0x3d, 0xb0, 0x0d, 0x20, 0x13, 0xc1, 0x84, 0x2f, 0x50, 0x2a, 0x25,
	0xb1, 0x24, 0x31, 0x08, 0xac, 0x48, 0x48, 0x9d, 0x8b, 0x25, 0x39, 0xb5, 0xa8, 0x44, 0x82, 0x45,
	0x81, 0x19, 0xa8, 0x58, 0x18, 0x4d, 0xb1, 0x33, 0x50, 0x2a, 0x08, 0xac, 0x40, 0xca, 0x82, 0x8b,
	0x07, 0x59, 0xbb, 0x90, 0x10, 0x17, 0x0b, 0x92, 0xcb, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xd4,
	0xdc, 0xc4, 0xcc, 0x1c, 0xa8, 0x93, 0x20, 0x1c, 0x29, 0x23, 0x2e, 0x0e, 0x98, 0x59, 0x20, 0x5d,
	0x60, 0xeb, 0xa0, 0xba, 0x92, 0xa1, 0x62, 0x05, 0xd9, 0xa9, 0x95, 0x50, 0x4d, 0x60, 0xb6, 0x52,
	0x06, 0x97, 0x80, 0x23, 0xd0, 0x25, 0xe1, 0x99, 0x25, 0x19, 0x01, 0x30, 0x7f, 0x91, 0x1b, 0x1e,
	0x32, 0x5c, 0x9c, 0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0xe0, 0x40, 0x61, 0x0d,
	0x42, 0x08, 0x28, 0xe5, 0x73, 0x89, 0xa3, 0xdb, 0x14, 0x5c, 0x9a, 0x9c, 0x9c, 0x5a, 0x5c, 0x0c,
	0x0f, 0x1b, 0x46, 0x70, 0x40, 0xe2, 0x0e, 0x1b, 0x78, 0x88, 0x33, 0x11, 0x11, 0xe2, 0x49, 0x6c,
	0xe0, 0xc8, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x20, 0x6d, 0xa6, 0xa2, 0xfa, 0x01, 0x00,
	0x00,
}
