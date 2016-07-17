// Code generated by protoc-gen-gogo.
// source: github.com/synrobo/proto/region/region.proto
// DO NOT EDIT!

/*
	Package region is a generated protocol buffer package.

	It is generated from these files:
		github.com/synrobo/proto/region/region.proto

	It has these top-level messages:
		Region
*/
package region

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import common "github.com/synrobo/proto/common"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

//
// Region: a self contained, isolated, installation with IP ranges.
type Region struct {
	// slug / ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Human readable name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// IP range for devices
	IpRange *common.IPRange `protobuf:"bytes,3,opt,name=ip_range,json=ipRange" json:"ip_range,omitempty"`
	// State
	State *Region_RegionState `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	// Keyspace name
	Keyspace string `protobuf:"bytes,5,opt,name=keyspace,proto3" json:"keyspace,omitempty"`
}

func (m *Region) Reset()                    { *m = Region{} }
func (m *Region) String() string            { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()               {}
func (*Region) Descriptor() ([]byte, []int) { return fileDescriptorRegion, []int{0} }

func (m *Region) GetIpRange() *common.IPRange {
	if m != nil {
		return m.IpRange
	}
	return nil
}

func (m *Region) GetState() *Region_RegionState {
	if m != nil {
		return m.State
	}
	return nil
}

type Region_RegionState struct {
	KeyspaceCreated bool `protobuf:"varint,1,opt,name=keyspace_created,json=keyspaceCreated,proto3" json:"keyspace_created,omitempty"`
	ConsulInited    bool `protobuf:"varint,2,opt,name=consul_inited,json=consulInited,proto3" json:"consul_inited,omitempty"`
	IpRangeVerified bool `protobuf:"varint,3,opt,name=ip_range_verified,json=ipRangeVerified,proto3" json:"ip_range_verified,omitempty"`
}

func (m *Region_RegionState) Reset()                    { *m = Region_RegionState{} }
func (m *Region_RegionState) String() string            { return proto.CompactTextString(m) }
func (*Region_RegionState) ProtoMessage()               {}
func (*Region_RegionState) Descriptor() ([]byte, []int) { return fileDescriptorRegion, []int{0, 0} }

func init() {
	proto.RegisterType((*Region)(nil), "region.Region")
	proto.RegisterType((*Region_RegionState)(nil), "region.Region.RegionState")
}
func (m *Region) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Region) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintRegion(data, i, uint64(len(m.Id)))
		i += copy(data[i:], m.Id)
	}
	if len(m.Name) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintRegion(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if m.IpRange != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintRegion(data, i, uint64(m.IpRange.Size()))
		n1, err := m.IpRange.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.State != nil {
		data[i] = 0x22
		i++
		i = encodeVarintRegion(data, i, uint64(m.State.Size()))
		n2, err := m.State.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Keyspace) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintRegion(data, i, uint64(len(m.Keyspace)))
		i += copy(data[i:], m.Keyspace)
	}
	return i, nil
}

func (m *Region_RegionState) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Region_RegionState) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.KeyspaceCreated {
		data[i] = 0x8
		i++
		if m.KeyspaceCreated {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.ConsulInited {
		data[i] = 0x10
		i++
		if m.ConsulInited {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.IpRangeVerified {
		data[i] = 0x18
		i++
		if m.IpRangeVerified {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Region(data []byte, offset int, v uint64) int {
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
func encodeFixed32Region(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRegion(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Region) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovRegion(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRegion(uint64(l))
	}
	if m.IpRange != nil {
		l = m.IpRange.Size()
		n += 1 + l + sovRegion(uint64(l))
	}
	if m.State != nil {
		l = m.State.Size()
		n += 1 + l + sovRegion(uint64(l))
	}
	l = len(m.Keyspace)
	if l > 0 {
		n += 1 + l + sovRegion(uint64(l))
	}
	return n
}

func (m *Region_RegionState) Size() (n int) {
	var l int
	_ = l
	if m.KeyspaceCreated {
		n += 2
	}
	if m.ConsulInited {
		n += 2
	}
	if m.IpRangeVerified {
		n += 2
	}
	return n
}

func sovRegion(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRegion(x uint64) (n int) {
	return sovRegion(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Region) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegion
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
			return fmt.Errorf("proto: Region: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Region: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegion
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
				return ErrInvalidLengthRegion
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegion
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
				return ErrInvalidLengthRegion
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegion
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
				return ErrInvalidLengthRegion
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IpRange == nil {
				m.IpRange = &common.IPRange{}
			}
			if err := m.IpRange.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegion
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
				return ErrInvalidLengthRegion
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.State == nil {
				m.State = &Region_RegionState{}
			}
			if err := m.State.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keyspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegion
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
				return ErrInvalidLengthRegion
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keyspace = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegion(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegion
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
func (m *Region_RegionState) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegion
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
			return fmt.Errorf("proto: RegionState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegionState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyspaceCreated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegion
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
			m.KeyspaceCreated = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsulInited", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegion
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
			m.ConsulInited = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpRangeVerified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegion
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
			m.IpRangeVerified = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRegion(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegion
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
func skipRegion(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRegion
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
					return 0, ErrIntOverflowRegion
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
					return 0, ErrIntOverflowRegion
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
				return 0, ErrInvalidLengthRegion
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRegion
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
				next, err := skipRegion(data[start:])
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
	ErrInvalidLengthRegion = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRegion   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorRegion = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4e, 0x84, 0x30,
	0x14, 0x86, 0x03, 0x33, 0x83, 0xd8, 0x19, 0x1d, 0xed, 0x8a, 0xb0, 0x30, 0x46, 0x37, 0x6a, 0x14,
	0x8c, 0xde, 0x40, 0x57, 0xb3, 0x33, 0x35, 0x71, 0x4b, 0x00, 0x3b, 0xb5, 0x51, 0x5a, 0x52, 0xc0,
	0x64, 0x8e, 0xe0, 0x75, 0x3c, 0x85, 0x4b, 0x8f, 0x60, 0x3c, 0x89, 0xe5, 0xbd, 0x62, 0xd8, 0xcc,
	0xa2, 0x79, 0xef, 0x7d, 0xfc, 0x7f, 0x1f, 0x7f, 0xc9, 0xa5, 0x90, 0xed, 0x4b, 0x57, 0x24, 0xa5,
	0xae, 0xd2, 0x66, 0xa3, 0x8c, 0x2e, 0x74, 0x5a, 0x1b, 0xdd, 0xea, 0xd4, 0x70, 0x21, 0xb5, 0x72,
	0x25, 0x01, 0x46, 0x03, 0x9c, 0xe2, 0xab, 0x91, 0x4b, 0x68, 0xe1, 0x2c, 0x45, 0xb7, 0x86, 0x09,
	0xfd, 0x7d, 0x87, 0xb6, 0x78, 0xfb, 0x12, 0x4b, 0x2a, 0xbb, 0x04, 0x0b, 0xaa, 0x4f, 0x3e, 0x7d,
	0x12, 0x30, 0xd8, 0x43, 0xf7, 0x89, 0x2f, 0x9f, 0x23, 0xef, 0xd8, 0x3b, 0xdb, 0x65, 0xb6, 0xa3,
	0x94, 0x4c, 0x55, 0x5e, 0xf1, 0xc8, 0x07, 0x02, 0x3d, 0xbd, 0x20, 0xa1, 0xac, 0x33, 0x93, 0x2b,
	0xc1, 0xa3, 0x89, 0xe5, 0xf3, 0x9b, 0x65, 0xe2, 0xee, 0x5b, 0x3d, 0xb0, 0x1e, 0xb3, 0x1d, 0x59,
	0x43, 0x43, 0xaf, 0xc9, 0xac, 0x69, 0xf3, 0x96, 0x47, 0x53, 0x10, 0xc6, 0x89, 0x4b, 0xc7, 0xc6,
	0xe5, 0xb1, 0x57, 0x30, 0x14, 0xd2, 0x98, 0x84, 0xaf, 0x7c, 0xd3, 0xd4, 0x79, 0xc9, 0xa3, 0x19,
	0x6c, 0xfd, 0x9f, 0xe3, 0x0f, 0x8f, 0xcc, 0x47, 0x16, 0x7a, 0x4e, 0x0e, 0x86, 0x6f, 0x59, 0x69,
	0xb8, 0x45, 0xf8, 0xef, 0x21, 0x5b, 0x0e, 0xfc, 0x1e, 0x31, 0x3d, 0x25, 0x7b, 0xa5, 0x56, 0x4d,
	0xf7, 0x96, 0x49, 0x25, 0x7b, 0x9d, 0x0f, 0xba, 0x05, 0xc2, 0x15, 0x30, 0x9b, 0xec, 0x70, 0x48,
	0x96, 0xbd, 0x73, 0x23, 0xd7, 0xd2, 0x0a, 0x27, 0x78, 0xa1, 0x4b, 0xf4, 0xe4, 0xf0, 0xdd, 0xe2,
	0xeb, 0xf7, 0xc8, 0xfb, 0xb6, 0xe7, 0xc7, 0x9e, 0x22, 0x80, 0x97, 0xbc, 0xfd, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x51, 0x61, 0xb9, 0xed, 0xde, 0x01, 0x00, 0x00,
}