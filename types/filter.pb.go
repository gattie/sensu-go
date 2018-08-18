// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filter.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// EventFilter is a filter specification.
type EventFilter struct {
	// Name is the unique identifier for a filter
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Action specifies to allow/deny events to continue through the pipeline
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	// Statements is an array of boolean expressions that are &&'d together
	// to determine if the event matches this filter.
	Statements []string `protobuf:"bytes,3,rep,name=statements" json:"statements"`
	// When indicates a TimeWindowWhen that a filter uses to filter by days & times
	When *TimeWindowWhen `protobuf:"bytes,6,opt,name=when" json:"when,omitempty"`
	// Namespace to which the filter belongs to
	Namespace string `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The type of the filter (govaluate or js)
	Type                 string   `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventFilter) Reset()         { *m = EventFilter{} }
func (m *EventFilter) String() string { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()    {}
func (*EventFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_a8bf061271ca41e8, []int{0}
}
func (m *EventFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EventFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFilter.Merge(dst, src)
}
func (m *EventFilter) XXX_Size() int {
	return m.Size()
}
func (m *EventFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventFilter proto.InternalMessageInfo

func (m *EventFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventFilter) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *EventFilter) GetStatements() []string {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *EventFilter) GetWhen() *TimeWindowWhen {
	if m != nil {
		return m.When
	}
	return nil
}

func (m *EventFilter) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *EventFilter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*EventFilter)(nil), "sensu.types.EventFilter")
}
func (this *EventFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EventFilter)
	if !ok {
		that2, ok := that.(EventFilter)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Action != that1.Action {
		return false
	}
	if len(this.Statements) != len(that1.Statements) {
		return false
	}
	for i := range this.Statements {
		if this.Statements[i] != that1.Statements[i] {
			return false
		}
	}
	if !this.When.Equal(that1.When) {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *EventFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFilter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFilter(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Action) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFilter(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
	}
	if len(m.Statements) > 0 {
		for _, s := range m.Statements {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.When != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintFilter(dAtA, i, uint64(m.When.Size()))
		n1, err := m.When.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Namespace) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintFilter(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintFilter(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintFilter(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedEventFilter(r randyFilter, easy bool) *EventFilter {
	this := &EventFilter{}
	this.Name = string(randStringFilter(r))
	this.Action = string(randStringFilter(r))
	v1 := r.Intn(10)
	this.Statements = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.Statements[i] = string(randStringFilter(r))
	}
	if r.Intn(10) != 0 {
		this.When = NewPopulatedTimeWindowWhen(r, easy)
	}
	this.Namespace = string(randStringFilter(r))
	this.Type = string(randStringFilter(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedFilter(r, 9)
	}
	return this
}

type randyFilter interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFilter(r randyFilter) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFilter(r randyFilter) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneFilter(r)
	}
	return string(tmps)
}
func randUnrecognizedFilter(r randyFilter, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldFilter(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldFilter(dAtA []byte, r randyFilter, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateFilter(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateFilter(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *EventFilter) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFilter(uint64(l))
	}
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovFilter(uint64(l))
	}
	if len(m.Statements) > 0 {
		for _, s := range m.Statements {
			l = len(s)
			n += 1 + l + sovFilter(uint64(l))
		}
	}
	if m.When != nil {
		l = m.When.Size()
		n += 1 + l + sovFilter(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovFilter(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovFilter(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFilter(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFilter(x uint64) (n int) {
	return sovFilter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statements", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statements = append(m.Statements, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field When", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.When == nil {
				m.When = &TimeWindowWhen{}
			}
			if err := m.When.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFilter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFilter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFilter
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFilter
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFilter
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipFilter(dAtA[start:])
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
	ErrInvalidLengthFilter = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFilter   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("filter.proto", fileDescriptor_filter_a8bf061271ca41e8) }

var fileDescriptor_filter_a8bf061271ca41e8 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcb, 0xcc, 0x29,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x4e, 0xcd, 0x2b, 0x2e, 0xd5,
	0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x49, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c,
	0x30, 0x0b, 0xa2, 0x57, 0x4a, 0xb0, 0x24, 0x33, 0x37, 0x35, 0xbe, 0x3c, 0x33, 0x2f, 0x25, 0xbf,
	0x1c, 0x22, 0xa4, 0x74, 0x86, 0x91, 0x8b, 0xdb, 0xb5, 0x2c, 0x35, 0xaf, 0xc4, 0x0d, 0x6c, 0x89,
	0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98,
	0x2d, 0x24, 0xc6, 0xc5, 0x96, 0x98, 0x5c, 0x92, 0x99, 0x9f, 0x27, 0xc1, 0x04, 0x16, 0x85, 0xf2,
	0x84, 0xf4, 0xb8, 0xb8, 0x8a, 0x4b, 0x12, 0x4b, 0x52, 0x73, 0x53, 0xf3, 0x4a, 0x8a, 0x25, 0x98,
	0x15, 0x98, 0x35, 0x38, 0x9d, 0xf8, 0x5e, 0xdd, 0x93, 0x47, 0x12, 0x0d, 0x42, 0x62, 0x0b, 0xe9,
	0x73, 0xb1, 0x94, 0x67, 0xa4, 0xe6, 0x49, 0xb0, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x49, 0xeb, 0x21,
	0xf9, 0x44, 0x2f, 0x24, 0x33, 0x37, 0x35, 0x1c, 0xec, 0xb0, 0xf0, 0x8c, 0xd4, 0xbc, 0x20, 0xb0,
	0x42, 0x21, 0x19, 0x2e, 0x4e, 0x90, 0x03, 0x8a, 0x0b, 0x12, 0x93, 0x53, 0x25, 0xd8, 0xc1, 0x76,
	0x23, 0x04, 0x40, 0x4e, 0x05, 0xe9, 0x95, 0xe0, 0x80, 0x38, 0x15, 0xc4, 0x76, 0x52, 0xfe, 0xf1,
	0x50, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x1d, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0xa2, 0x58, 0xc1, 0x76,
	0x25, 0xb1, 0x81, 0xbd, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x39, 0x61, 0x47, 0x59,
	0x01, 0x00, 0x00,
}
