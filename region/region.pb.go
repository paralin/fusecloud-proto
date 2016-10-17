// Code generated by protoc-gen-gogo.
// source: github.com/fuserobotics/proto/region/region.proto
// DO NOT EDIT!

package region

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import common "github.com/fuserobotics/proto/common"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//
// Get all regions
type Region struct {
	// slug / ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Human readable name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// IP range for devices
	IpRange *common.IPRange `protobuf:"bytes,3,opt,name=ip_range,json=ipRange" json:"ip_range,omitempty"`
	// State
	State *Region_RegionState `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	// Location
	Location *common.GeoLocation `protobuf:"bytes,6,opt,name=location" json:"location,omitempty"`
	// Map zoom level default
	Zoomlevel int32 `protobuf:"varint,7,opt,name=zoomlevel,proto3" json:"zoomlevel,omitempty"`
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

func (m *Region) GetLocation() *common.GeoLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

type Region_RegionState struct {
	// reserved 1
	ConsulBootstrapperCreated bool `protobuf:"varint,2,opt,name=consul_bootstrapper_created,json=consulBootstrapperCreated,proto3" json:"consul_bootstrapper_created,omitempty"`
	DbCreated                 bool `protobuf:"varint,3,opt,name=db_created,json=dbCreated,proto3" json:"db_created,omitempty"`
	MetaTableCreated          bool `protobuf:"varint,4,opt,name=meta_table_created,json=metaTableCreated,proto3" json:"meta_table_created,omitempty"`
	StreamsTableCreated       bool `protobuf:"varint,5,opt,name=streams_table_created,json=streamsTableCreated,proto3" json:"streams_table_created,omitempty"`
}

func (m *Region_RegionState) Reset()                    { *m = Region_RegionState{} }
func (m *Region_RegionState) String() string            { return proto.CompactTextString(m) }
func (*Region_RegionState) ProtoMessage()               {}
func (*Region_RegionState) Descriptor() ([]byte, []int) { return fileDescriptorRegion, []int{0, 0} }

type RegionDataTable struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Created bool   `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	// defaults to {default: 1}
	Replication map[string]int32 `protobuf:"bytes,3,rep,name=replication" json:"replication,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// defaults to "id"
	PrimaryKey string `protobuf:"bytes,4,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	// defaults to the first key in replication
	PrimaryReplica string `protobuf:"bytes,5,opt,name=primary_replica,json=primaryReplica,proto3" json:"primary_replica,omitempty"`
}

func (m *RegionDataTable) Reset()                    { *m = RegionDataTable{} }
func (m *RegionDataTable) String() string            { return proto.CompactTextString(m) }
func (*RegionDataTable) ProtoMessage()               {}
func (*RegionDataTable) Descriptor() ([]byte, []int) { return fileDescriptorRegion, []int{1} }

func (m *RegionDataTable) GetReplication() map[string]int32 {
	if m != nil {
		return m.Replication
	}
	return nil
}

func init() {
	proto.RegisterType((*Region)(nil), "region.Region")
	proto.RegisterType((*Region_RegionState)(nil), "region.Region.RegionState")
	proto.RegisterType((*RegionDataTable)(nil), "region.RegionDataTable")
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
	if m.Location != nil {
		data[i] = 0x32
		i++
		i = encodeVarintRegion(data, i, uint64(m.Location.Size()))
		n3, err := m.Location.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Zoomlevel != 0 {
		data[i] = 0x38
		i++
		i = encodeVarintRegion(data, i, uint64(m.Zoomlevel))
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
	if m.ConsulBootstrapperCreated {
		data[i] = 0x10
		i++
		if m.ConsulBootstrapperCreated {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.DbCreated {
		data[i] = 0x18
		i++
		if m.DbCreated {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.MetaTableCreated {
		data[i] = 0x20
		i++
		if m.MetaTableCreated {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.StreamsTableCreated {
		data[i] = 0x28
		i++
		if m.StreamsTableCreated {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *RegionDataTable) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RegionDataTable) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintRegion(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if m.Created {
		data[i] = 0x10
		i++
		if m.Created {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if len(m.Replication) > 0 {
		for k, _ := range m.Replication {
			data[i] = 0x1a
			i++
			v := m.Replication[k]
			mapSize := 1 + len(k) + sovRegion(uint64(len(k))) + 1 + sovRegion(uint64(v))
			i = encodeVarintRegion(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintRegion(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x10
			i++
			i = encodeVarintRegion(data, i, uint64(v))
		}
	}
	if len(m.PrimaryKey) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintRegion(data, i, uint64(len(m.PrimaryKey)))
		i += copy(data[i:], m.PrimaryKey)
	}
	if len(m.PrimaryReplica) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintRegion(data, i, uint64(len(m.PrimaryReplica)))
		i += copy(data[i:], m.PrimaryReplica)
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
	if m.Location != nil {
		l = m.Location.Size()
		n += 1 + l + sovRegion(uint64(l))
	}
	if m.Zoomlevel != 0 {
		n += 1 + sovRegion(uint64(m.Zoomlevel))
	}
	return n
}

func (m *Region_RegionState) Size() (n int) {
	var l int
	_ = l
	if m.ConsulBootstrapperCreated {
		n += 2
	}
	if m.DbCreated {
		n += 2
	}
	if m.MetaTableCreated {
		n += 2
	}
	if m.StreamsTableCreated {
		n += 2
	}
	return n
}

func (m *RegionDataTable) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRegion(uint64(l))
	}
	if m.Created {
		n += 2
	}
	if len(m.Replication) > 0 {
		for k, v := range m.Replication {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovRegion(uint64(len(k))) + 1 + sovRegion(uint64(v))
			n += mapEntrySize + 1 + sovRegion(uint64(mapEntrySize))
		}
	}
	l = len(m.PrimaryKey)
	if l > 0 {
		n += 1 + l + sovRegion(uint64(l))
	}
	l = len(m.PrimaryReplica)
	if l > 0 {
		n += 1 + l + sovRegion(uint64(l))
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
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
			if m.Location == nil {
				m.Location = &common.GeoLocation{}
			}
			if err := m.Location.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zoomlevel", wireType)
			}
			m.Zoomlevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegion
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsulBootstrapperCreated", wireType)
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
			m.ConsulBootstrapperCreated = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DbCreated", wireType)
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
			m.DbCreated = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaTableCreated", wireType)
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
			m.MetaTableCreated = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamsTableCreated", wireType)
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
			m.StreamsTableCreated = bool(v != 0)
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
func (m *RegionDataTable) Unmarshal(data []byte) error {
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
			return fmt.Errorf("proto: RegionDataTable: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegionDataTable: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
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
			m.Created = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replication", wireType)
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
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegion
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
					return ErrIntOverflowRegion
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
				return ErrInvalidLengthRegion
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
					return ErrIntOverflowRegion
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
			var mapvalue int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				mapvalue |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if m.Replication == nil {
				m.Replication = make(map[string]int32)
			}
			m.Replication[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimaryKey", wireType)
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
			m.PrimaryKey = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimaryReplica", wireType)
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
			m.PrimaryReplica = string(data[iNdEx:postIndex])
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
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0x95, 0xa6, 0x69, 0x9b, 0x29, 0xda, 0x56, 0x5e, 0x90, 0x42, 0xf8, 0x5a, 0xed, 0x85,
	0x15, 0x82, 0x16, 0x96, 0x0b, 0xe2, 0xb0, 0x87, 0x05, 0x84, 0xf8, 0x38, 0x20, 0xc3, 0x3d, 0x72,
	0xd2, 0xd9, 0x10, 0x91, 0xc4, 0x91, 0xe3, 0xac, 0x14, 0xde, 0x83, 0x97, 0xe0, 0x49, 0x38, 0x21,
	0x1e, 0x01, 0xf1, 0x24, 0x38, 0xb6, 0x13, 0xb2, 0xbd, 0x70, 0x88, 0x3c, 0xf3, 0x9f, 0xdf, 0x7f,
	0x46, 0x13, 0x1b, 0x9e, 0xa4, 0x99, 0xfc, 0xdc, 0xc4, 0x9b, 0x84, 0x17, 0xdb, 0x8b, 0xa6, 0x46,
	0xc1, 0x63, 0x2e, 0xb3, 0xa4, 0xde, 0x56, 0x82, 0x4b, 0xbe, 0x15, 0x98, 0x66, 0xbc, 0xb4, 0xc7,
	0x46, 0x6b, 0x64, 0x66, 0xb2, 0xf0, 0xd1, 0xc8, 0x9a, 0xf2, 0x94, 0x1b, 0x4b, 0xdc, 0x5c, 0xe8,
	0xcc, 0xf8, 0xbb, 0xc8, 0xd8, 0xc2, 0xff, 0x4c, 0x52, 0x72, 0xa1, 0x26, 0x99, 0xc3, 0x58, 0x8e,
	0xbf, 0xbb, 0x30, 0xa3, 0x7a, 0x18, 0x39, 0x80, 0x49, 0xb6, 0x0b, 0x9c, 0x23, 0xe7, 0xc4, 0xa7,
	0x2a, 0x22, 0x04, 0xa6, 0x25, 0x2b, 0x30, 0x98, 0x68, 0x45, 0xc7, 0xe4, 0x01, 0x2c, 0xb2, 0x2a,
	0x12, 0xac, 0x4c, 0x31, 0x70, 0x95, 0xbe, 0x3c, 0x5d, 0x6d, 0x6c, 0xbf, 0x37, 0x1f, 0x68, 0x27,
	0xd3, 0x79, 0x56, 0xe9, 0x80, 0x3c, 0x06, 0xaf, 0x96, 0x4c, 0x62, 0x30, 0xd5, 0x60, 0xb8, 0xb1,
	0x2b, 0xd2, 0xf1, 0xf1, 0xb1, 0x23, 0xa8, 0x01, 0xc9, 0x16, 0x16, 0x39, 0x4f, 0x98, 0x54, 0x7a,
	0x30, 0xd3, 0xa6, 0xc3, 0xbe, 0xfb, 0x6b, 0xe4, 0xef, 0x6d, 0x89, 0x0e, 0x10, 0xb9, 0x0d, 0xfe,
	0x57, 0xce, 0x8b, 0x1c, 0x2f, 0x31, 0x0f, 0xe6, 0xca, 0xe1, 0xd1, 0x7f, 0x42, 0xf8, 0xd3, 0x81,
	0xe5, 0x68, 0x0a, 0x39, 0x83, 0x5b, 0x09, 0x2f, 0xeb, 0x26, 0x8f, 0x62, 0xce, 0x65, 0x2d, 0x05,
	0xab, 0x2a, 0x14, 0x51, 0x22, 0x50, 0x55, 0x77, 0x7a, 0xcf, 0x05, 0xbd, 0x69, 0x90, 0xf3, 0x11,
	0xf1, 0xc2, 0x00, 0xe4, 0x0e, 0xc0, 0x2e, 0x1e, 0x70, 0x57, 0xe3, 0xfe, 0x2e, 0xee, 0xcb, 0x0f,
	0x81, 0x14, 0x28, 0x59, 0x24, 0x59, 0x9c, 0xe3, 0x80, 0x4d, 0x35, 0xb6, 0xee, 0x2a, 0x9f, 0xba,
	0x42, 0x4f, 0x9f, 0xc2, 0x0d, 0xd5, 0x1f, 0x59, 0x51, 0xef, 0x19, 0x3c, 0x6d, 0x38, 0xb4, 0xc5,
	0xb1, 0xe7, 0xf8, 0xdb, 0x04, 0x56, 0x66, 0xa1, 0x97, 0xcc, 0xb6, 0x1b, 0x6e, 0xc9, 0x19, 0xdd,
	0x52, 0x00, 0xf3, 0xab, 0x4b, 0xf5, 0x29, 0x79, 0x0b, 0x4b, 0x81, 0x55, 0x9e, 0xd9, 0x9f, 0xec,
	0x1e, 0xb9, 0xea, 0x27, 0x9f, 0x5c, 0xbd, 0x99, 0xa1, 0xb7, 0xca, 0x07, 0xf4, 0x55, 0x29, 0x45,
	0x4b, 0xc7, 0x66, 0x72, 0x0f, 0x96, 0x95, 0xc8, 0x0a, 0x26, 0xda, 0xe8, 0x0b, 0xb6, 0x7a, 0x51,
	0x9f, 0x82, 0x95, 0xde, 0x61, 0x4b, 0xee, 0xc3, 0xaa, 0x07, 0xac, 0x4f, 0x2f, 0xe7, 0xd3, 0x03,
	0x2b, 0xdb, 0xfe, 0xe1, 0x19, 0xac, 0xf7, 0x47, 0x91, 0x35, 0xb8, 0x5d, 0x57, 0xb3, 0x56, 0x17,
	0x92, 0xeb, 0xe0, 0x5d, 0xb2, 0xbc, 0x31, 0x0f, 0xd2, 0xa3, 0x26, 0x79, 0x3e, 0x79, 0xe6, 0x9c,
	0x5f, 0xfb, 0xf1, 0xe7, 0xae, 0xf3, 0x4b, 0x7d, 0xbf, 0xd5, 0x17, 0xcf, 0xf4, 0xcb, 0x7e, 0xfa,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x98, 0x9b, 0x50, 0x78, 0x03, 0x00, 0x00,
}
