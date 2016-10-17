// Code generated by protoc-gen-gogo.
// source: github.com/fuserobotics/proto/auth/auth.proto
// DO NOT EDIT!

package auth

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import common "github.com/fuserobotics/proto/common"
import permissions "github.com/fuserobotics/proto/permissions"

import errors "errors"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User_UserRole int32

const (
	// Nothing special
	User_NONE User_UserRole = 0
	// Access to everything
	User_SUPERUSER User_UserRole = 1
)

var User_UserRole_name = map[int32]string{
	0: "NONE",
	1: "SUPERUSER",
}
var User_UserRole_value = map[string]int32{
	"NONE":      0,
	"SUPERUSER": 1,
}

func (x User_UserRole) String() string {
	return proto.EnumName(User_UserRole_name, int32(x))
}
func (User_UserRole) EnumDescriptor() ([]byte, []int) { return fileDescriptorAuth, []int{0, 0} }

//
// User: a standard user with username, email, password
type User struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Encrypted password for database
	Password string             `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Meta     *User_UserMetadata `protobuf:"bytes,3,opt,name=meta" json:"meta,omitempty"`
	// certificates, latest at index 0
	Cert []*User_UserCert `protobuf:"bytes,4,rep,name=cert" json:"cert,omitempty"`
	// processed first
	GlobalRole *User_UserRoleSet `protobuf:"bytes,5,opt,name=global_role,json=globalRole" json:"global_role,omitempty"`
	// processed second
	RegionRole map[string]*User_UserRoleSet `protobuf:"bytes,6,rep,name=region_role,json=regionRole" json:"region_role,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// extra permissions
	GlobalExtraPermission *permissions.SystemPermissions `protobuf:"bytes,7,opt,name=global_extra_permission,json=globalExtraPermission" json:"global_extra_permission,omitempty"`
	// per region extra permissions
	RegionExtraPermission map[string]*permissions.SystemPermissions `protobuf:"bytes,8,rep,name=region_extra_permission,json=regionExtraPermission" json:"region_extra_permission,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorAuth, []int{0} }

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

func (m *User) GetGlobalRole() *User_UserRoleSet {
	if m != nil {
		return m.GlobalRole
	}
	return nil
}

func (m *User) GetRegionRole() map[string]*User_UserRoleSet {
	if m != nil {
		return m.RegionRole
	}
	return nil
}

func (m *User) GetGlobalExtraPermission() *permissions.SystemPermissions {
	if m != nil {
		return m.GlobalExtraPermission
	}
	return nil
}

func (m *User) GetRegionExtraPermission() map[string]*permissions.SystemPermissions {
	if m != nil {
		return m.RegionExtraPermission
	}
	return nil
}

type User_UserMetadata struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (m *User_UserMetadata) Reset()                    { *m = User_UserMetadata{} }
func (m *User_UserMetadata) String() string            { return proto.CompactTextString(m) }
func (*User_UserMetadata) ProtoMessage()               {}
func (*User_UserMetadata) Descriptor() ([]byte, []int) { return fileDescriptorAuth, []int{0, 2} }

type User_UserCert struct {
	Chain     *common.CertChain `protobuf:"bytes,1,opt,name=chain" json:"chain,omitempty"`
	Pkey      string            `protobuf:"bytes,2,opt,name=pkey,proto3" json:"pkey,omitempty"`
	PublicKey string            `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (m *User_UserCert) Reset()                    { *m = User_UserCert{} }
func (m *User_UserCert) String() string            { return proto.CompactTextString(m) }
func (*User_UserCert) ProtoMessage()               {}
func (*User_UserCert) Descriptor() ([]byte, []int) { return fileDescriptorAuth, []int{0, 3} }

func (m *User_UserCert) GetChain() *common.CertChain {
	if m != nil {
		return m.Chain
	}
	return nil
}

type User_UserRoleSet struct {
	Role []User_UserRole `protobuf:"varint,1,rep,name=role,enum=auth.User_UserRole" json:"role,omitempty"`
}

func (m *User_UserRoleSet) Reset()                    { *m = User_UserRoleSet{} }
func (m *User_UserRoleSet) String() string            { return proto.CompactTextString(m) }
func (*User_UserRoleSet) ProtoMessage()               {}
func (*User_UserRoleSet) Descriptor() ([]byte, []int) { return fileDescriptorAuth, []int{0, 4} }

func init() {
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*User_UserMetadata)(nil), "auth.User.UserMetadata")
	proto.RegisterType((*User_UserCert)(nil), "auth.User.UserCert")
	proto.RegisterType((*User_UserRoleSet)(nil), "auth.User.UserRoleSet")
	proto.RegisterEnum("auth.User_UserRole", User_UserRole_name, User_UserRole_value)
}
func (m *User) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *User) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintAuth(data, i, uint64(len(m.Username)))
		i += copy(data[i:], m.Username)
	}
	if len(m.Password) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintAuth(data, i, uint64(len(m.Password)))
		i += copy(data[i:], m.Password)
	}
	if m.Meta != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintAuth(data, i, uint64(m.Meta.Size()))
		n1, err := m.Meta.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Cert) > 0 {
		for _, msg := range m.Cert {
			data[i] = 0x22
			i++
			i = encodeVarintAuth(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.GlobalRole != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintAuth(data, i, uint64(m.GlobalRole.Size()))
		n2, err := m.GlobalRole.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.RegionRole) > 0 {
		for k, _ := range m.RegionRole {
			data[i] = 0x32
			i++
			v := m.RegionRole[k]
			if v == nil {
				return 0, errors.New("proto: map has nil element")
			}
			msgSize := v.Size()
			mapSize := 1 + len(k) + sovAuth(uint64(len(k))) + 1 + msgSize + sovAuth(uint64(msgSize))
			i = encodeVarintAuth(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintAuth(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintAuth(data, i, uint64(v.Size()))
			n3, err := v.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n3
		}
	}
	if m.GlobalExtraPermission != nil {
		data[i] = 0x3a
		i++
		i = encodeVarintAuth(data, i, uint64(m.GlobalExtraPermission.Size()))
		n4, err := m.GlobalExtraPermission.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.RegionExtraPermission) > 0 {
		for k, _ := range m.RegionExtraPermission {
			data[i] = 0x42
			i++
			v := m.RegionExtraPermission[k]
			if v == nil {
				return 0, errors.New("proto: map has nil element")
			}
			msgSize := v.Size()
			mapSize := 1 + len(k) + sovAuth(uint64(len(k))) + 1 + msgSize + sovAuth(uint64(msgSize))
			i = encodeVarintAuth(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintAuth(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintAuth(data, i, uint64(v.Size()))
			n5, err := v.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n5
		}
	}
	return i, nil
}

func (m *User_UserMetadata) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *User_UserMetadata) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintAuth(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if len(m.Email) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintAuth(data, i, uint64(len(m.Email)))
		i += copy(data[i:], m.Email)
	}
	return i, nil
}

func (m *User_UserCert) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *User_UserCert) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Chain != nil {
		data[i] = 0xa
		i++
		i = encodeVarintAuth(data, i, uint64(m.Chain.Size()))
		n6, err := m.Chain.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.Pkey) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintAuth(data, i, uint64(len(m.Pkey)))
		i += copy(data[i:], m.Pkey)
	}
	if len(m.PublicKey) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintAuth(data, i, uint64(len(m.PublicKey)))
		i += copy(data[i:], m.PublicKey)
	}
	return i, nil
}

func (m *User_UserRoleSet) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *User_UserRoleSet) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Role) > 0 {
		for _, num := range m.Role {
			data[i] = 0x8
			i++
			i = encodeVarintAuth(data, i, uint64(num))
		}
	}
	return i, nil
}

func encodeFixed64Auth(data []byte, offset int, v uint64) int {
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
func encodeFixed32Auth(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAuth(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *User) Size() (n int) {
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	if len(m.Cert) > 0 {
		for _, e := range m.Cert {
			l = e.Size()
			n += 1 + l + sovAuth(uint64(l))
		}
	}
	if m.GlobalRole != nil {
		l = m.GlobalRole.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	if len(m.RegionRole) > 0 {
		for k, v := range m.RegionRole {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
			}
			mapEntrySize := 1 + len(k) + sovAuth(uint64(len(k))) + 1 + l + sovAuth(uint64(l))
			n += mapEntrySize + 1 + sovAuth(uint64(mapEntrySize))
		}
	}
	if m.GlobalExtraPermission != nil {
		l = m.GlobalExtraPermission.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	if len(m.RegionExtraPermission) > 0 {
		for k, v := range m.RegionExtraPermission {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
			}
			mapEntrySize := 1 + len(k) + sovAuth(uint64(len(k))) + 1 + l + sovAuth(uint64(l))
			n += mapEntrySize + 1 + sovAuth(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *User_UserMetadata) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	return n
}

func (m *User_UserCert) Size() (n int) {
	var l int
	_ = l
	if m.Chain != nil {
		l = m.Chain.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.Pkey)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	return n
}

func (m *User_UserRoleSet) Size() (n int) {
	var l int
	_ = l
	if len(m.Role) > 0 {
		for _, e := range m.Role {
			n += 1 + sovAuth(uint64(e))
		}
	}
	return n
}

func sovAuth(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAuth(x uint64) (n int) {
	return sovAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &User_UserMetadata{}
			}
			if err := m.Meta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cert", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cert = append(m.Cert, &User_UserCert{})
			if err := m.Cert[len(m.Cert)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalRole", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GlobalRole == nil {
				m.GlobalRole = &User_UserRoleSet{}
			}
			if err := m.GlobalRole.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionRole", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthAuth
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapmsglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				mapmsglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if mapmsglen < 0 {
				return ErrInvalidLengthAuth
			}
			postmsgIndex := iNdEx + mapmsglen
			if mapmsglen < 0 {
				return ErrInvalidLengthAuth
			}
			if postmsgIndex > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := &User_UserRoleSet{}
			if err := mapvalue.Unmarshal(data[iNdEx:postmsgIndex]); err != nil {
				return err
			}
			iNdEx = postmsgIndex
			if m.RegionRole == nil {
				m.RegionRole = make(map[string]*User_UserRoleSet)
			}
			m.RegionRole[mapkey] = mapvalue
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalExtraPermission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GlobalExtraPermission == nil {
				m.GlobalExtraPermission = &permissions.SystemPermissions{}
			}
			if err := m.GlobalExtraPermission.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionExtraPermission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthAuth
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapmsglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				mapmsglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if mapmsglen < 0 {
				return ErrInvalidLengthAuth
			}
			postmsgIndex := iNdEx + mapmsglen
			if mapmsglen < 0 {
				return ErrInvalidLengthAuth
			}
			if postmsgIndex > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := &permissions.SystemPermissions{}
			if err := mapvalue.Unmarshal(data[iNdEx:postmsgIndex]); err != nil {
				return err
			}
			iNdEx = postmsgIndex
			if m.RegionExtraPermission == nil {
				m.RegionExtraPermission = make(map[string]*permissions.SystemPermissions)
			}
			m.RegionExtraPermission[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
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
func (m *User_UserMetadata) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: UserMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
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
func (m *User_UserCert) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: UserCert: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserCert: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Chain == nil {
				m.Chain = &common.CertChain{}
			}
			if err := m.Chain.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pkey = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
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
func (m *User_UserRoleSet) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: UserRoleSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserRoleSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var v User_UserRole
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (User_UserRole(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Role = append(m.Role, v)
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
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
func skipAuth(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
				return 0, ErrInvalidLengthAuth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuth
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
				next, err := skipAuth(data[start:])
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
	ErrInvalidLengthAuth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuth   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorAuth = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x26, 0x6b, 0xba, 0xb5, 0xd7, 0x01, 0xc5, 0x30, 0x1a, 0x45, 0x62, 0x9a, 0x86, 0x10, 0x48,
	0x40, 0x2b, 0x06, 0x82, 0x89, 0xbd, 0x31, 0xe5, 0x09, 0x31, 0x26, 0x57, 0xe5, 0x0d, 0x55, 0x4e,
	0xe6, 0xa6, 0x11, 0x49, 0x1c, 0x39, 0x0e, 0xd0, 0x47, 0xfe, 0x1d, 0x8f, 0xfc, 0x04, 0xc4, 0x2f,
	0xc1, 0x3e, 0x67, 0x6b, 0x68, 0xb7, 0xf1, 0x90, 0xe4, 0xee, 0xbe, 0xef, 0xbb, 0xbb, 0x9c, 0xcf,
	0xf0, 0x3c, 0x4e, 0xd4, 0xbc, 0x0a, 0x87, 0x91, 0xc8, 0x46, 0xb3, 0xaa, 0xe4, 0x52, 0x84, 0x42,
	0x25, 0x51, 0x39, 0x2a, 0xa4, 0x50, 0x62, 0xc4, 0x2a, 0x35, 0xc7, 0xd7, 0x10, 0x7d, 0xe2, 0x1a,
	0xdb, 0x6f, 0x8a, 0x62, 0x11, 0x0b, 0x4b, 0x0e, 0xab, 0x19, 0x7a, 0x56, 0x69, 0x2c, 0x2b, 0xf2,
	0x5f, 0x5c, 0x5f, 0x43, 0x87, 0x33, 0x91, 0xd7, 0x9f, 0x5a, 0x72, 0x74, 0xbd, 0xa4, 0xe0, 0x32,
	0x4b, 0xca, 0x32, 0x11, 0x79, 0xd9, 0xb4, 0xad, 0x78, 0xff, 0xc7, 0x16, 0xb8, 0x13, 0xad, 0x21,
	0x3e, 0x74, 0x8c, 0x36, 0x67, 0x19, 0xf7, 0x9c, 0x3d, 0xe7, 0x49, 0x97, 0x5e, 0xf8, 0x06, 0x2b,
	0x58, 0x59, 0x7e, 0x13, 0xf2, 0xcc, 0xdb, 0xb0, 0xd8, 0xb9, 0x4f, 0x9e, 0x82, 0x9b, 0x71, 0xc5,
	0xbc, 0x96, 0x8e, 0xf7, 0x0e, 0x06, 0x43, 0x1c, 0x80, 0xc9, 0x88, 0xaf, 0x0f, 0x1a, 0x3a, 0x63,
	0x8a, 0x51, 0x24, 0x91, 0xc7, 0xe0, 0x46, 0x5c, 0x2a, 0xcf, 0xdd, 0x6b, 0x69, 0xf2, 0xdd, 0x15,
	0xf2, 0xb1, 0x86, 0x28, 0x12, 0xc8, 0x1b, 0xe8, 0xc5, 0xa9, 0x08, 0x59, 0x3a, 0x95, 0x22, 0xe5,
	0x5e, 0x1b, 0x93, 0xdf, 0x5f, 0xe1, 0x53, 0x0d, 0x8d, 0xb9, 0xa2, 0x60, 0xa9, 0xc6, 0x25, 0x47,
	0xd0, 0x93, 0x3c, 0xd6, 0x3f, 0x68, 0x85, 0x9b, 0x58, 0xc8, 0x6f, 0x08, 0x29, 0xa2, 0x86, 0x1b,
	0xe4, 0x4a, 0x2e, 0x28, 0xc8, 0x8b, 0x00, 0xf9, 0x04, 0x83, 0xba, 0x2a, 0xff, 0xae, 0x24, 0x9b,
	0x2e, 0xc7, 0xe5, 0x6d, 0x61, 0x07, 0xbb, 0xc3, 0xe6, 0x04, 0xc7, 0x8b, 0x52, 0xf1, 0xec, 0x74,
	0x19, 0xa1, 0x3b, 0x56, 0x1e, 0x18, 0xf5, 0x32, 0x4e, 0x3e, 0xc3, 0xa0, 0x6e, 0x6a, 0x2d, 0x6f,
	0x07, 0x1b, 0x7c, 0xb4, 0xd6, 0xe0, 0x4a, 0x0a, 0xdb, 0xeb, 0x8e, 0xbc, 0x0c, 0xf3, 0x27, 0x70,
	0x7b, 0xe5, 0xaf, 0x48, 0x1f, 0x5a, 0x5f, 0xf8, 0xa2, 0x3e, 0x48, 0x63, 0x92, 0x67, 0xd0, 0xfe,
	0xca, 0xd2, 0x8a, 0xe3, 0x01, 0x5e, 0x3d, 0x4b, 0x4b, 0x7a, 0xbb, 0x71, 0xe8, 0xf8, 0x73, 0xf0,
	0xaf, 0xee, 0xe5, 0x92, 0x0a, 0xaf, 0xfe, 0xad, 0xf0, 0xbf, 0x59, 0x35, 0x2a, 0x1d, 0xc2, 0x76,
	0x73, 0x59, 0x08, 0x01, 0xb7, 0xb1, 0x87, 0x68, 0x93, 0x7b, 0xd0, 0xe6, 0x19, 0x4b, 0xd2, 0x7a,
	0x01, 0xad, 0xe3, 0xcf, 0xa0, 0x73, 0xbe, 0x39, 0x7a, 0xb9, 0xda, 0xd1, 0x9c, 0x25, 0x39, 0xca,
	0x7a, 0x07, 0x77, 0x86, 0xf5, 0x2d, 0x31, 0xe0, 0xb1, 0x01, 0xa8, 0xc5, 0x4d, 0xfa, 0xc2, 0xf4,
	0x6e, 0x33, 0xa1, 0x4d, 0x1e, 0x00, 0x14, 0x55, 0x98, 0x26, 0xd1, 0xd4, 0x20, 0x2d, 0x44, 0xba,
	0x36, 0xf2, 0x9e, 0x2f, 0xfc, 0xd7, 0xd0, 0x6b, 0x4c, 0xc9, 0xec, 0x31, 0xae, 0x97, 0xa3, 0x4f,
	0xef, 0xd6, 0xda, 0x1e, 0x1b, 0x16, 0x45, 0xc2, 0xfe, 0x43, 0xdb, 0x1f, 0x6e, 0x57, 0x07, 0xdc,
	0x93, 0x8f, 0x27, 0x41, 0xff, 0x06, 0xb9, 0x09, 0xdd, 0xf1, 0xe4, 0x34, 0xa0, 0x93, 0x71, 0x40,
	0xfb, 0xce, 0xbb, 0xed, 0x9f, 0x7f, 0x76, 0x9d, 0x5f, 0xfa, 0xf9, 0xad, 0x9f, 0x70, 0x13, 0x2f,
	0xe6, 0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x68, 0xcc, 0xd1, 0x33, 0x6e, 0x04, 0x00, 0x00,
}
