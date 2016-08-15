// Code generated by protoc-gen-gogo.
// source: github.com/fuserobotics/proto/metric/metric.proto
// DO NOT EDIT!

/*
	Package metric is a generated protocol buffer package.

	It is generated from these files:
		github.com/fuserobotics/proto/metric/metric.proto

	It has these top-level messages:
		MetricIdentifier
		MetricSeries
		MetricDatapoint
*/
package metric

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type MetricSeries_MetricDataType int32

const (
	MetricSeries_NUMBER MetricSeries_MetricDataType = 0
)

var MetricSeries_MetricDataType_name = map[int32]string{
	0: "NUMBER",
}
var MetricSeries_MetricDataType_value = map[string]int32{
	"NUMBER": 0,
}

func (x MetricSeries_MetricDataType) String() string {
	return proto.EnumName(MetricSeries_MetricDataType_name, int32(x))
}
func (MetricSeries_MetricDataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorMetric, []int{1, 0}
}

// Used to identify a stored metric
type MetricIdentifier struct {
	// Use a string identifier for now.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MetricIdentifier) Reset()                    { *m = MetricIdentifier{} }
func (m *MetricIdentifier) String() string            { return proto.CompactTextString(m) }
func (*MetricIdentifier) ProtoMessage()               {}
func (*MetricIdentifier) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{0} }

type MetricSeries struct {
	Id             string                               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title          string                               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description    string                               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	KairosId       string                               `protobuf:"bytes,4,opt,name=kairos_id,json=kairosId,proto3" json:"kairos_id,omitempty"`
	DataType       MetricSeries_MetricDataType          `protobuf:"varint,5,opt,name=data_type,json=dataType,proto3,enum=metric.MetricSeries_MetricDataType" json:"data_type,omitempty"`
	TagDescription []*MetricSeries_MetricTagDescription `protobuf:"bytes,6,rep,name=tag_description,json=tagDescription" json:"tag_description,omitempty"`
}

func (m *MetricSeries) Reset()                    { *m = MetricSeries{} }
func (m *MetricSeries) String() string            { return proto.CompactTextString(m) }
func (*MetricSeries) ProtoMessage()               {}
func (*MetricSeries) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{1} }

func (m *MetricSeries) GetTagDescription() []*MetricSeries_MetricTagDescription {
	if m != nil {
		return m.TagDescription
	}
	return nil
}

type MetricSeries_MetricTagDescription struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *MetricSeries_MetricTagDescription) Reset()         { *m = MetricSeries_MetricTagDescription{} }
func (m *MetricSeries_MetricTagDescription) String() string { return proto.CompactTextString(m) }
func (*MetricSeries_MetricTagDescription) ProtoMessage()    {}
func (*MetricSeries_MetricTagDescription) Descriptor() ([]byte, []int) {
	return fileDescriptorMetric, []int{1, 0}
}

type MetricDatapoint struct {
	SensorSet  string              `protobuf:"bytes,1,opt,name=sensor_set,json=sensorSet,proto3" json:"sensor_set,omitempty"`
	SensorId   int32               `protobuf:"varint,2,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
	MetricName string              `protobuf:"bytes,3,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	Timestamp  int64               `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Location   *common.GeoLocation `protobuf:"bytes,5,opt,name=location" json:"location,omitempty"`
}

func (m *MetricDatapoint) Reset()                    { *m = MetricDatapoint{} }
func (m *MetricDatapoint) String() string            { return proto.CompactTextString(m) }
func (*MetricDatapoint) ProtoMessage()               {}
func (*MetricDatapoint) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{2} }

func (m *MetricDatapoint) GetLocation() *common.GeoLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricIdentifier)(nil), "metric.MetricIdentifier")
	proto.RegisterType((*MetricSeries)(nil), "metric.MetricSeries")
	proto.RegisterType((*MetricSeries_MetricTagDescription)(nil), "metric.MetricSeries.MetricTagDescription")
	proto.RegisterType((*MetricDatapoint)(nil), "metric.MetricDatapoint")
	proto.RegisterEnum("metric.MetricSeries_MetricDataType", MetricSeries_MetricDataType_name, MetricSeries_MetricDataType_value)
}
func (m *MetricIdentifier) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MetricIdentifier) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintMetric(data, i, uint64(len(m.Id)))
		i += copy(data[i:], m.Id)
	}
	return i, nil
}

func (m *MetricSeries) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MetricSeries) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintMetric(data, i, uint64(len(m.Id)))
		i += copy(data[i:], m.Id)
	}
	if len(m.Title) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintMetric(data, i, uint64(len(m.Title)))
		i += copy(data[i:], m.Title)
	}
	if len(m.Description) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintMetric(data, i, uint64(len(m.Description)))
		i += copy(data[i:], m.Description)
	}
	if len(m.KairosId) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintMetric(data, i, uint64(len(m.KairosId)))
		i += copy(data[i:], m.KairosId)
	}
	if m.DataType != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintMetric(data, i, uint64(m.DataType))
	}
	if len(m.TagDescription) > 0 {
		for _, msg := range m.TagDescription {
			data[i] = 0x32
			i++
			i = encodeVarintMetric(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *MetricSeries_MetricTagDescription) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MetricSeries_MetricTagDescription) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintMetric(data, i, uint64(len(m.Id)))
		i += copy(data[i:], m.Id)
	}
	if len(m.Title) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintMetric(data, i, uint64(len(m.Title)))
		i += copy(data[i:], m.Title)
	}
	if len(m.Description) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintMetric(data, i, uint64(len(m.Description)))
		i += copy(data[i:], m.Description)
	}
	return i, nil
}

func (m *MetricDatapoint) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MetricDatapoint) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SensorSet) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintMetric(data, i, uint64(len(m.SensorSet)))
		i += copy(data[i:], m.SensorSet)
	}
	if m.SensorId != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintMetric(data, i, uint64(m.SensorId))
	}
	if len(m.MetricName) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintMetric(data, i, uint64(len(m.MetricName)))
		i += copy(data[i:], m.MetricName)
	}
	if m.Timestamp != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintMetric(data, i, uint64(m.Timestamp))
	}
	if m.Location != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintMetric(data, i, uint64(m.Location.Size()))
		n1, err := m.Location.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeFixed64Metric(data []byte, offset int, v uint64) int {
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
func encodeFixed32Metric(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMetric(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *MetricIdentifier) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	return n
}

func (m *MetricSeries) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	l = len(m.KairosId)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.DataType != 0 {
		n += 1 + sovMetric(uint64(m.DataType))
	}
	if len(m.TagDescription) > 0 {
		for _, e := range m.TagDescription {
			l = e.Size()
			n += 1 + l + sovMetric(uint64(l))
		}
	}
	return n
}

func (m *MetricSeries_MetricTagDescription) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	return n
}

func (m *MetricDatapoint) Size() (n int) {
	var l int
	_ = l
	l = len(m.SensorSet)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.SensorId != 0 {
		n += 1 + sovMetric(uint64(m.SensorId))
	}
	l = len(m.MetricName)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovMetric(uint64(m.Timestamp))
	}
	if m.Location != nil {
		l = m.Location.Size()
		n += 1 + l + sovMetric(uint64(l))
	}
	return n
}

func sovMetric(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetric(x uint64) (n int) {
	return sovMetric(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetricIdentifier) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
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
			return fmt.Errorf("proto: MetricIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
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
func (m *MetricSeries) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
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
			return fmt.Errorf("proto: MetricSeries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricSeries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KairosId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KairosId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataType", wireType)
			}
			m.DataType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DataType |= (MetricSeries_MetricDataType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagDescription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TagDescription = append(m.TagDescription, &MetricSeries_MetricTagDescription{})
			if err := m.TagDescription[len(m.TagDescription)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
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
func (m *MetricSeries_MetricTagDescription) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
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
			return fmt.Errorf("proto: MetricTagDescription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricTagDescription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
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
func (m *MetricDatapoint) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
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
			return fmt.Errorf("proto: MetricDatapoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricDatapoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SensorSet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SensorSet = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SensorId", wireType)
			}
			m.SensorId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.SensorId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetricName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
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
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
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
func skipMetric(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetric
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
					return 0, ErrIntOverflowMetric
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
					return 0, ErrIntOverflowMetric
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
				return 0, ErrInvalidLengthMetric
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetric
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
				next, err := skipMetric(data[start:])
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
	ErrInvalidLengthMetric = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetric   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorMetric = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x6e, 0xd4, 0x30,
	0x14, 0x66, 0x3a, 0x4c, 0x94, 0xbc, 0x54, 0x69, 0x65, 0xba, 0x88, 0x86, 0x02, 0xa3, 0xb0, 0x29,
	0x0b, 0x12, 0x31, 0x5c, 0x00, 0x55, 0x45, 0xa8, 0x12, 0xed, 0xc2, 0x2d, 0x5b, 0x22, 0x27, 0xf1,
	0x04, 0x8b, 0x26, 0x8e, 0xec, 0x37, 0x0b, 0x4e, 0xc0, 0xa9, 0xd8, 0xb3, 0xe4, 0x08, 0x88, 0x93,
	0xe0, 0xd8, 0x16, 0x93, 0x41, 0x08, 0x36, 0x2c, 0x2c, 0xfb, 0x7d, 0x3f, 0xf6, 0xe7, 0x67, 0xc3,
	0x8b, 0x56, 0xe0, 0x87, 0x6d, 0x95, 0xd7, 0xb2, 0x2b, 0x36, 0x5b, 0xcd, 0x95, 0xac, 0x24, 0x8a,
	0x5a, 0x17, 0x83, 0x92, 0x28, 0x8b, 0x8e, 0xa3, 0x12, 0xb5, 0x9f, 0x72, 0x8b, 0x91, 0xc0, 0x55,
	0xcb, 0xe7, 0x13, 0x6b, 0x2b, 0x5b, 0xe9, 0x2c, 0xd5, 0x76, 0x63, 0x2b, 0xe7, 0x1f, 0x57, 0xce,
	0xb6, 0xfc, 0xc7, 0x49, 0x06, 0xee, 0x64, 0xef, 0x27, 0x67, 0xc9, 0x32, 0x38, 0xbe, 0xb2, 0x67,
	0x5d, 0x36, 0xbc, 0x47, 0xb1, 0x11, 0x5c, 0x91, 0x04, 0x0e, 0x44, 0x93, 0xce, 0x56, 0xb3, 0xb3,
	0x88, 0x9a, 0x55, 0xf6, 0x79, 0x0e, 0x87, 0x4e, 0x74, 0xc3, 0x95, 0xe0, 0xfa, 0x77, 0x01, 0x39,
	0x81, 0x05, 0x0a, 0xbc, 0xe3, 0xe9, 0x81, 0x85, 0x5c, 0x41, 0x56, 0x10, 0x37, 0x5c, 0xd7, 0x4a,
	0x0c, 0x28, 0x64, 0x9f, 0xce, 0x2d, 0x37, 0x85, 0xc8, 0x43, 0x88, 0x3e, 0x32, 0xa1, 0xa4, 0x2e,
	0xcd, 0x76, 0xf7, 0x2d, 0x1f, 0x3a, 0xe0, 0xb2, 0x21, 0xaf, 0x20, 0x6a, 0x18, 0xb2, 0x12, 0x3f,
	0x0d, 0x3c, 0x5d, 0x18, 0x32, 0x59, 0x3f, 0xcd, 0x7d, 0x97, 0xa6, 0x69, 0x7c, 0x71, 0x61, 0xb4,
	0xb7, 0x46, 0x4a, 0xc3, 0xc6, 0xaf, 0x08, 0x85, 0x23, 0x64, 0x6d, 0x39, 0x0d, 0x11, 0xac, 0xe6,
	0x67, 0xf1, 0xfa, 0xd9, 0x5f, 0xf6, 0xb9, 0x65, 0xed, 0xc5, 0xce, 0x40, 0x13, 0xdc, 0xab, 0x97,
	0xef, 0xe1, 0xe4, 0x4f, 0xba, 0xff, 0xd5, 0x92, 0xec, 0x14, 0x92, 0xfd, 0xfb, 0x10, 0x80, 0xe0,
	0xfa, 0xdd, 0xd5, 0xf9, 0x6b, 0x7a, 0x7c, 0x2f, 0xfb, 0x32, 0x83, 0xa3, 0x1d, 0x3d, 0x48, 0xd1,
	0x23, 0x79, 0x04, 0xa0, 0x79, 0xaf, 0xa5, 0x2a, 0x35, 0x47, 0x9f, 0x20, 0x72, 0xc8, 0x0d, 0xc7,
	0xb1, 0xc7, 0x9e, 0x36, 0xf9, 0xc6, 0x30, 0x0b, 0x1a, 0x3a, 0xc0, 0xf4, 0xf8, 0x09, 0xc4, 0xae,
	0x13, 0x65, 0xcf, 0x3a, 0xee, 0xf3, 0x80, 0x83, 0xae, 0x0d, 0x42, 0x4e, 0x21, 0x42, 0xd1, 0x71,
	0x8d, 0xac, 0x1b, 0xec, 0x0b, 0xcd, 0xe9, 0x0e, 0x20, 0x05, 0x84, 0x77, 0xb2, 0x66, 0xf6, 0x2e,
	0xe3, 0x0b, 0xc5, 0xeb, 0x07, 0xb9, 0xff, 0x5d, 0x6f, 0xb8, 0x7c, 0xeb, 0x29, 0xfa, 0x4b, 0x74,
	0x7e, 0xf8, 0xf5, 0xc7, 0xe3, 0xd9, 0x37, 0x33, 0xbe, 0x9b, 0x51, 0x05, 0xf6, 0x0b, 0xbe, 0xfc,
	0x19, 0x00, 0x00, 0xff, 0xff, 0x49, 0xe7, 0xf0, 0xd4, 0x21, 0x03, 0x00, 0x00,
}
