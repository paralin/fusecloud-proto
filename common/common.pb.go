// Code generated by protoc-gen-gogo.
// source: github.com/synrobo/proto/common/common.proto
// DO NOT EDIT!

/*
	Package common is a generated protocol buffer package.

	It is generated from these files:
		github.com/synrobo/proto/common/common.proto

	It has these top-level messages:
		IPRange
		IPAddress
		CertChain
		GeoLocation
		CommitState
*/
package common

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// IP address CIDR
type IPRange struct {
	// IP data
	Ip *IPAddress `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	// Prefix length: 0->32
	Plen uint32 `protobuf:"varint,2,opt,name=plen,proto3" json:"plen,omitempty"`
}

func (m *IPRange) Reset()                    { *m = IPRange{} }
func (m *IPRange) String() string            { return proto.CompactTextString(m) }
func (*IPRange) ProtoMessage()               {}
func (*IPRange) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{0} }

func (m *IPRange) GetIp() *IPAddress {
	if m != nil {
		return m.Ip
	}
	return nil
}

type IPAddress struct {
	Address []uint32 `protobuf:"varint,1,rep,packed,name=address" json:"address,omitempty"`
}

func (m *IPAddress) Reset()                    { *m = IPAddress{} }
func (m *IPAddress) String() string            { return proto.CompactTextString(m) }
func (*IPAddress) ProtoMessage()               {}
func (*IPAddress) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{1} }

// Certificate chain
type CertChain struct {
	// Cert chain, idx 0 should be last cert.
	Cert       []string `protobuf:"bytes,1,rep,name=cert" json:"cert,omitempty"`
	ValidUntil int64    `protobuf:"varint,2,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
	Csr        string   `protobuf:"bytes,3,opt,name=csr,proto3" json:"csr,omitempty"`
	CsrWaiting bool     `protobuf:"varint,4,opt,name=csr_waiting,json=csrWaiting,proto3" json:"csr_waiting,omitempty"`
}

func (m *CertChain) Reset()                    { *m = CertChain{} }
func (m *CertChain) String() string            { return proto.CompactTextString(m) }
func (*CertChain) ProtoMessage()               {}
func (*CertChain) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{2} }

type GeoLocation struct {
	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// optional, if using gps coords
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// optional, if using a map
	Zoomlevel int32 `protobuf:"varint,4,opt,name=zoomlevel,proto3" json:"zoomlevel,omitempty"`
}

func (m *GeoLocation) Reset()                    { *m = GeoLocation{} }
func (m *GeoLocation) String() string            { return proto.CompactTextString(m) }
func (*GeoLocation) ProtoMessage()               {}
func (*GeoLocation) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{3} }

type CommitState struct {
	CurrentIndex int32 `protobuf:"varint,1,opt,name=current_index,json=currentIndex,proto3" json:"current_index,omitempty"`
	ConsulIndex  int32 `protobuf:"varint,2,opt,name=consul_index,json=consulIndex,proto3" json:"consul_index,omitempty"`
}

func (m *CommitState) Reset()                    { *m = CommitState{} }
func (m *CommitState) String() string            { return proto.CompactTextString(m) }
func (*CommitState) ProtoMessage()               {}
func (*CommitState) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{4} }

func init() {
	proto.RegisterType((*IPRange)(nil), "common.IPRange")
	proto.RegisterType((*IPAddress)(nil), "common.IPAddress")
	proto.RegisterType((*CertChain)(nil), "common.CertChain")
	proto.RegisterType((*GeoLocation)(nil), "common.GeoLocation")
	proto.RegisterType((*CommitState)(nil), "common.CommitState")
}
func (m *IPRange) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *IPRange) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ip != nil {
		data[i] = 0xa
		i++
		i = encodeVarintCommon(data, i, uint64(m.Ip.Size()))
		n1, err := m.Ip.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Plen != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintCommon(data, i, uint64(m.Plen))
	}
	return i, nil
}

func (m *IPAddress) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *IPAddress) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		data3 := make([]byte, len(m.Address)*10)
		var j2 int
		for _, num := range m.Address {
			for num >= 1<<7 {
				data3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			data3[j2] = uint8(num)
			j2++
		}
		data[i] = 0xa
		i++
		i = encodeVarintCommon(data, i, uint64(j2))
		i += copy(data[i:], data3[:j2])
	}
	return i, nil
}

func (m *CertChain) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CertChain) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Cert) > 0 {
		for _, s := range m.Cert {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if m.ValidUntil != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintCommon(data, i, uint64(m.ValidUntil))
	}
	if len(m.Csr) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintCommon(data, i, uint64(len(m.Csr)))
		i += copy(data[i:], m.Csr)
	}
	if m.CsrWaiting {
		data[i] = 0x20
		i++
		if m.CsrWaiting {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *GeoLocation) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GeoLocation) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Latitude != 0 {
		data[i] = 0x9
		i++
		i = encodeFixed64Common(data, i, uint64(math.Float64bits(float64(m.Latitude))))
	}
	if m.Longitude != 0 {
		data[i] = 0x11
		i++
		i = encodeFixed64Common(data, i, uint64(math.Float64bits(float64(m.Longitude))))
	}
	if m.Timestamp != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintCommon(data, i, uint64(m.Timestamp))
	}
	if m.Zoomlevel != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintCommon(data, i, uint64(m.Zoomlevel))
	}
	return i, nil
}

func (m *CommitState) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CommitState) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CurrentIndex != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintCommon(data, i, uint64(m.CurrentIndex))
	}
	if m.ConsulIndex != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintCommon(data, i, uint64(m.ConsulIndex))
	}
	return i, nil
}

func encodeFixed64Common(data []byte, offset int, v uint64) int {
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
func encodeFixed32Common(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCommon(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *IPRange) Size() (n int) {
	var l int
	_ = l
	if m.Ip != nil {
		l = m.Ip.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.Plen != 0 {
		n += 1 + sovCommon(uint64(m.Plen))
	}
	return n
}

func (m *IPAddress) Size() (n int) {
	var l int
	_ = l
	if len(m.Address) > 0 {
		l = 0
		for _, e := range m.Address {
			l += sovCommon(uint64(e))
		}
		n += 1 + sovCommon(uint64(l)) + l
	}
	return n
}

func (m *CertChain) Size() (n int) {
	var l int
	_ = l
	if len(m.Cert) > 0 {
		for _, s := range m.Cert {
			l = len(s)
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	if m.ValidUntil != 0 {
		n += 1 + sovCommon(uint64(m.ValidUntil))
	}
	l = len(m.Csr)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.CsrWaiting {
		n += 2
	}
	return n
}

func (m *GeoLocation) Size() (n int) {
	var l int
	_ = l
	if m.Latitude != 0 {
		n += 9
	}
	if m.Longitude != 0 {
		n += 9
	}
	if m.Timestamp != 0 {
		n += 1 + sovCommon(uint64(m.Timestamp))
	}
	if m.Zoomlevel != 0 {
		n += 1 + sovCommon(uint64(m.Zoomlevel))
	}
	return n
}

func (m *CommitState) Size() (n int) {
	var l int
	_ = l
	if m.CurrentIndex != 0 {
		n += 1 + sovCommon(uint64(m.CurrentIndex))
	}
	if m.ConsulIndex != 0 {
		n += 1 + sovCommon(uint64(m.ConsulIndex))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IPRange) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: IPRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IPRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ip == nil {
				m.Ip = &IPAddress{}
			}
			if err := m.Ip.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plen", wireType)
			}
			m.Plen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Plen |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *IPAddress) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: IPAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IPAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthCommon
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := data[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Address = append(m.Address, v)
				}
			} else if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Address = append(m.Address, v)
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *CertChain) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: CertChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cert = append(m.Cert, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUntil", wireType)
			}
			m.ValidUntil = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ValidUntil |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Csr = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CsrWaiting", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
			m.CsrWaiting = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *GeoLocation) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: GeoLocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GeoLocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			m.Latitude = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			m.Longitude = float64(math.Float64frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zoomlevel", wireType)
			}
			m.Zoomlevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Zoomlevel |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *CommitState) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: CommitState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommitState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentIndex", wireType)
			}
			m.CurrentIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.CurrentIndex |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsulIndex", wireType)
			}
			m.ConsulIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ConsulIndex |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCommon
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
				next, err := skipCommon(data[start:])
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
	ErrInvalidLengthCommon = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorCommon = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x91, 0xcb, 0x4e, 0xeb, 0x30,
	0x10, 0x86, 0x95, 0xa4, 0xb7, 0x38, 0xad, 0xd4, 0xe3, 0x55, 0x54, 0x1d, 0x9d, 0x4b, 0xd8, 0x14,
	0x09, 0x5a, 0x09, 0x5e, 0x00, 0xda, 0x05, 0xaa, 0xc4, 0xa2, 0x32, 0xaa, 0x58, 0x56, 0x4e, 0x62,
	0x52, 0x4b, 0x8e, 0x1d, 0x25, 0x4e, 0xb9, 0x6c, 0x79, 0x39, 0x96, 0x3c, 0x02, 0xe2, 0x49, 0xb0,
	0x27, 0xa1, 0x65, 0x61, 0xf9, 0x9f, 0xef, 0xff, 0x67, 0x3a, 0x8d, 0xd1, 0x59, 0xc6, 0xf5, 0xae,
	0x8e, 0x67, 0x89, 0xca, 0xe7, 0xd5, 0xb3, 0x2c, 0x55, 0xac, 0xe6, 0x45, 0xa9, 0xb4, 0x9a, 0x1b,
	0x92, 0x2b, 0xd9, 0x5e, 0x33, 0x60, 0xb8, 0xd7, 0x54, 0x93, 0xf3, 0x1f, 0x5d, 0x99, 0xca, 0xda,
	0x96, 0xb8, 0x7e, 0x80, 0xaa, 0xe9, 0xb7, 0xaa, 0x69, 0x8b, 0xae, 0x50, 0x7f, 0xb5, 0x26, 0x54,
	0x66, 0x0c, 0xff, 0x47, 0x2e, 0x2f, 0x42, 0xe7, 0x9f, 0x33, 0x0d, 0x2e, 0x7e, 0xcd, 0xda, 0xe1,
	0xab, 0xf5, 0x75, 0x9a, 0x96, 0xac, 0xaa, 0x88, 0x31, 0x31, 0x46, 0x9d, 0x42, 0x30, 0x19, 0xba,
	0x26, 0x34, 0x22, 0xa0, 0xa3, 0x53, 0xe4, 0x1f, 0x42, 0xf8, 0x37, 0xea, 0xd3, 0x46, 0x9a, 0x41,
	0xde, 0x74, 0xb4, 0x70, 0xc7, 0x0e, 0xf9, 0x46, 0x51, 0x85, 0xfc, 0x25, 0x2b, 0xf5, 0x72, 0x47,
	0xb9, 0xb4, 0xb3, 0x12, 0x53, 0x40, 0xce, 0x27, 0xa0, 0xf1, 0x5f, 0x14, 0xec, 0xa9, 0xe0, 0xe9,
	0xb6, 0x96, 0x9a, 0x0b, 0xf8, 0x19, 0x8f, 0x20, 0x40, 0x1b, 0x4b, 0xf0, 0x18, 0x79, 0x49, 0x55,
	0x86, 0x9e, 0x31, 0x7c, 0x62, 0xa5, 0x6d, 0x31, 0xd7, 0xf6, 0x91, 0x72, 0xcd, 0x65, 0x16, 0x76,
	0x8c, 0x33, 0x20, 0xc8, 0xa0, 0xfb, 0x86, 0x44, 0xaf, 0x0e, 0x0a, 0x6e, 0x98, 0xba, 0x55, 0x09,
	0xd5, 0x5c, 0x49, 0x3c, 0x41, 0x03, 0x61, 0x94, 0xae, 0x53, 0x06, 0x7f, 0xd6, 0x21, 0x87, 0xda,
	0xac, 0xef, 0x0b, 0x25, 0xb3, 0xc6, 0x74, 0xc1, 0x3c, 0x02, 0xeb, 0x6a, 0x9e, 0xb3, 0x4a, 0xd3,
	0xbc, 0x80, 0x15, 0x3c, 0x72, 0x04, 0xd6, 0x7d, 0x51, 0x2a, 0x17, 0x6c, 0xcf, 0x04, 0xac, 0xd1,
	0x25, 0x47, 0x10, 0x6d, 0x50, 0xb0, 0x34, 0x5f, 0x94, 0xeb, 0x3b, 0x4d, 0x35, 0xc3, 0x27, 0x68,
	0x94, 0xd4, 0x65, 0xc9, 0xa4, 0xde, 0x72, 0x99, 0xb2, 0x27, 0xd8, 0xa4, 0x4b, 0x86, 0x2d, 0x5c,
	0x59, 0x66, 0x1e, 0x64, 0x98, 0x28, 0x59, 0xd5, 0xa2, 0xcd, 0xb8, 0x90, 0x09, 0x1a, 0x06, 0x91,
	0xc5, 0xf0, 0xed, 0xf3, 0x8f, 0xf3, 0x6e, 0xce, 0x87, 0x39, 0x71, 0x0f, 0xde, 0xf4, 0xf2, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0xa0, 0x96, 0x8b, 0xa9, 0x3a, 0x02, 0x00, 0x00,
}
