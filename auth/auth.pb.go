// Code generated by protoc-gen-go.
// source: github.com/fuserobotics/proto/auth/auth.proto
// DO NOT EDIT!

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	github.com/fuserobotics/proto/auth/auth.proto

It has these top-level messages:
	UserPublicIdentity
	UserPrivateData
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/fuserobotics/proto/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserPrivateData_PrivateKeyStrategy int32

const (
	UserPrivateData_PKEY_AES256 UserPrivateData_PrivateKeyStrategy = 0
)

var UserPrivateData_PrivateKeyStrategy_name = map[int32]string{
	0: "PKEY_AES256",
}
var UserPrivateData_PrivateKeyStrategy_value = map[string]int32{
	"PKEY_AES256": 0,
}

func (x UserPrivateData_PrivateKeyStrategy) String() string {
	return proto.EnumName(UserPrivateData_PrivateKeyStrategy_name, int32(x))
}
func (UserPrivateData_PrivateKeyStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

// *
// UserPublicIdentity: public identity for a user.
type UserPublicIdentity struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
}

func (m *UserPublicIdentity) Reset()                    { *m = UserPublicIdentity{} }
func (m *UserPublicIdentity) String() string            { return proto.CompactTextString(m) }
func (*UserPublicIdentity) ProtoMessage()               {}
func (*UserPublicIdentity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserPublicIdentity) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserPublicIdentity) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

// *
// UserPrivateData: private data for a user.
type UserPrivateData struct {
	// AES256 encrypted private key, aes key is scrypt result.
	// IV is the same scrypt operation, output length 16
	PrivateKey   []byte                             `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	Strategy     UserPrivateData_PrivateKeyStrategy `protobuf:"varint,2,opt,name=strategy,enum=auth.UserPrivateData_PrivateKeyStrategy" json:"strategy,omitempty"`
	ScryptParams *UserPrivateData_ScryptParams      `protobuf:"bytes,3,opt,name=scrypt_params,json=scryptParams" json:"scrypt_params,omitempty"`
}

func (m *UserPrivateData) Reset()                    { *m = UserPrivateData{} }
func (m *UserPrivateData) String() string            { return proto.CompactTextString(m) }
func (*UserPrivateData) ProtoMessage()               {}
func (*UserPrivateData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserPrivateData) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *UserPrivateData) GetStrategy() UserPrivateData_PrivateKeyStrategy {
	if m != nil {
		return m.Strategy
	}
	return UserPrivateData_PKEY_AES256
}

func (m *UserPrivateData) GetScryptParams() *UserPrivateData_ScryptParams {
	if m != nil {
		return m.ScryptParams
	}
	return nil
}

type UserPrivateData_ScryptParams struct {
	// Difficulty, power of 2. Default is 16384
	N uint32 `protobuf:"varint,1,opt,name=n" json:"n,omitempty"`
	// r * p < 2^30
	R uint32 `protobuf:"varint,2,opt,name=r" json:"r,omitempty"`
	// r * p < 2^30
	P uint32 `protobuf:"varint,3,opt,name=p" json:"p,omitempty"`
}

func (m *UserPrivateData_ScryptParams) Reset()                    { *m = UserPrivateData_ScryptParams{} }
func (m *UserPrivateData_ScryptParams) String() string            { return proto.CompactTextString(m) }
func (*UserPrivateData_ScryptParams) ProtoMessage()               {}
func (*UserPrivateData_ScryptParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *UserPrivateData_ScryptParams) GetN() uint32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *UserPrivateData_ScryptParams) GetR() uint32 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *UserPrivateData_ScryptParams) GetP() uint32 {
	if m != nil {
		return m.P
	}
	return 0
}

func init() {
	proto.RegisterType((*UserPublicIdentity)(nil), "auth.UserPublicIdentity")
	proto.RegisterType((*UserPrivateData)(nil), "auth.UserPrivateData")
	proto.RegisterType((*UserPrivateData_ScryptParams)(nil), "auth.UserPrivateData.ScryptParams")
	proto.RegisterEnum("auth.UserPrivateData_PrivateKeyStrategy", UserPrivateData_PrivateKeyStrategy_name, UserPrivateData_PrivateKeyStrategy_value)
}

func init() { proto.RegisterFile("github.com/fuserobotics/proto/auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xcd, 0x14, 0xd9, 0xde, 0x5a, 0x27, 0x39, 0x8d, 0x81, 0x38, 0x0a, 0xc2, 0x2e, 0xb6,
	0x58, 0x51, 0xbc, 0x0a, 0x1b, 0x22, 0x3d, 0x58, 0x52, 0x3c, 0x78, 0x2a, 0x69, 0x8d, 0x5b, 0xd1,
	0x36, 0x21, 0x49, 0x85, 0xfe, 0x3b, 0xfe, 0xa5, 0xd2, 0xd7, 0x52, 0x86, 0x8a, 0x97, 0x24, 0xdf,
	0x7b, 0xdf, 0xef, 0xcb, 0x4b, 0xe0, 0x72, 0x5b, 0xd8, 0x5d, 0x9d, 0xf9, 0xb9, 0x2c, 0x83, 0xb7,
	0xda, 0x08, 0x2d, 0x33, 0x69, 0x8b, 0xdc, 0x04, 0x4a, 0x4b, 0x2b, 0x03, 0x5e, 0xdb, 0x1d, 0x2e,
	0x3e, 0x6a, 0x7a, 0xd4, 0x9e, 0x17, 0x57, 0xff, 0x43, 0xb9, 0x2c, 0x4b, 0x59, 0xf5, 0x5b, 0x07,
	0x7a, 0x4f, 0x40, 0x9f, 0x8d, 0xd0, 0x71, 0x9d, 0x7d, 0x14, 0xf9, 0xe3, 0xab, 0xa8, 0x6c, 0x61,
	0x1b, 0xba, 0x80, 0x71, 0x8b, 0x57, 0xbc, 0x14, 0x73, 0xb2, 0x24, 0xab, 0x09, 0x1b, 0x34, 0x3d,
	0x03, 0x50, 0xe8, 0x4e, 0xdf, 0x45, 0x33, 0x1f, 0x61, 0x77, 0xd2, 0x55, 0x22, 0xd1, 0x78, 0x5f,
	0x23, 0x98, 0x61, 0xa2, 0x2e, 0x3e, 0xb9, 0x15, 0x6b, 0x6e, 0x39, 0x3d, 0x87, 0xa9, 0xea, 0x24,
	0x32, 0x6d, 0xa2, 0xc3, 0xa0, 0x2f, 0x45, 0xa2, 0xa1, 0x6b, 0x18, 0x1b, 0xab, 0xb9, 0x15, 0xdb,
	0x2e, 0xf1, 0x24, 0x5c, 0xf9, 0xf8, 0xba, 0x1f, 0x49, 0x7e, 0x3c, 0x30, 0x49, 0xef, 0x67, 0x03,
	0x49, 0x1f, 0xc0, 0x35, 0xb9, 0x6e, 0x94, 0x4d, 0x15, 0xd7, 0xbc, 0x34, 0xf3, 0xc3, 0x25, 0x59,
	0x4d, 0x43, 0xef, 0xef, 0xa8, 0x04, 0xad, 0x31, 0x3a, 0x99, 0x63, 0xf6, 0xd4, 0xe2, 0x0e, 0x9c,
	0xfd, 0x2e, 0x75, 0x80, 0x54, 0x38, 0xb5, 0xcb, 0x48, 0xd5, 0x2a, 0x8d, 0x53, 0xba, 0x8c, 0xe8,
	0x56, 0x29, 0xbc, 0xc8, 0x65, 0x44, 0x79, 0x17, 0x40, 0x7f, 0x8f, 0x48, 0x67, 0x30, 0x8d, 0xa3,
	0xcd, 0x4b, 0x7a, 0xbf, 0x49, 0xc2, 0x9b, 0xdb, 0xd3, 0x83, 0xec, 0x18, 0x3f, 0xff, 0xfa, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x11, 0x72, 0x42, 0xbb, 0xe6, 0x01, 0x00, 0x00,
}
