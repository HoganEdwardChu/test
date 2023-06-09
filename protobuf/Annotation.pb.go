// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/Annotation.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PIntStringValue struct {
	IntValue             int32                   `protobuf:"varint,1,opt,name=intValue,proto3" json:"intValue,omitempty"`
	StringValue          *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=stringValue,proto3" json:"stringValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PIntStringValue) Reset()         { *m = PIntStringValue{} }
func (m *PIntStringValue) String() string { return proto.CompactTextString(m) }
func (*PIntStringValue) ProtoMessage()    {}
func (*PIntStringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe607a737cbfdf8, []int{0}
}

func (m *PIntStringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PIntStringValue.Unmarshal(m, b)
}
func (m *PIntStringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PIntStringValue.Marshal(b, m, deterministic)
}
func (m *PIntStringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PIntStringValue.Merge(m, src)
}
func (m *PIntStringValue) XXX_Size() int {
	return xxx_messageInfo_PIntStringValue.Size(m)
}
func (m *PIntStringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PIntStringValue.DiscardUnknown(m)
}

var xxx_messageInfo_PIntStringValue proto.InternalMessageInfo

func (m *PIntStringValue) GetIntValue() int32 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *PIntStringValue) GetStringValue() *wrapperspb.StringValue {
	if m != nil {
		return m.StringValue
	}
	return nil
}

type PIntStringStringValue struct {
	IntValue             int32                   `protobuf:"varint,1,opt,name=intValue,proto3" json:"intValue,omitempty"`
	StringValue1         *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=stringValue1,proto3" json:"stringValue1,omitempty"`
	StringValue2         *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=stringValue2,proto3" json:"stringValue2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PIntStringStringValue) Reset()         { *m = PIntStringStringValue{} }
func (m *PIntStringStringValue) String() string { return proto.CompactTextString(m) }
func (*PIntStringStringValue) ProtoMessage()    {}
func (*PIntStringStringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe607a737cbfdf8, []int{1}
}

func (m *PIntStringStringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PIntStringStringValue.Unmarshal(m, b)
}
func (m *PIntStringStringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PIntStringStringValue.Marshal(b, m, deterministic)
}
func (m *PIntStringStringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PIntStringStringValue.Merge(m, src)
}
func (m *PIntStringStringValue) XXX_Size() int {
	return xxx_messageInfo_PIntStringStringValue.Size(m)
}
func (m *PIntStringStringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PIntStringStringValue.DiscardUnknown(m)
}

var xxx_messageInfo_PIntStringStringValue proto.InternalMessageInfo

func (m *PIntStringStringValue) GetIntValue() int32 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *PIntStringStringValue) GetStringValue1() *wrapperspb.StringValue {
	if m != nil {
		return m.StringValue1
	}
	return nil
}

func (m *PIntStringStringValue) GetStringValue2() *wrapperspb.StringValue {
	if m != nil {
		return m.StringValue2
	}
	return nil
}

type PLongIntIntByteByteStringValue struct {
	LongValue            int64                   `protobuf:"varint,1,opt,name=longValue,proto3" json:"longValue,omitempty"`
	IntValue1            int32                   `protobuf:"varint,2,opt,name=intValue1,proto3" json:"intValue1,omitempty"`
	IntValue2            int32                   `protobuf:"varint,3,opt,name=intValue2,proto3" json:"intValue2,omitempty"`
	ByteValue1           int32                   `protobuf:"zigzag32,4,opt,name=byteValue1,proto3" json:"byteValue1,omitempty"`
	ByteValue2           int32                   `protobuf:"zigzag32,5,opt,name=byteValue2,proto3" json:"byteValue2,omitempty"`
	StringValue          *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=stringValue,proto3" json:"stringValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PLongIntIntByteByteStringValue) Reset()         { *m = PLongIntIntByteByteStringValue{} }
func (m *PLongIntIntByteByteStringValue) String() string { return proto.CompactTextString(m) }
func (*PLongIntIntByteByteStringValue) ProtoMessage()    {}
func (*PLongIntIntByteByteStringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe607a737cbfdf8, []int{2}
}

func (m *PLongIntIntByteByteStringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PLongIntIntByteByteStringValue.Unmarshal(m, b)
}
func (m *PLongIntIntByteByteStringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PLongIntIntByteByteStringValue.Marshal(b, m, deterministic)
}
func (m *PLongIntIntByteByteStringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PLongIntIntByteByteStringValue.Merge(m, src)
}
func (m *PLongIntIntByteByteStringValue) XXX_Size() int {
	return xxx_messageInfo_PLongIntIntByteByteStringValue.Size(m)
}
func (m *PLongIntIntByteByteStringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PLongIntIntByteByteStringValue.DiscardUnknown(m)
}

var xxx_messageInfo_PLongIntIntByteByteStringValue proto.InternalMessageInfo

func (m *PLongIntIntByteByteStringValue) GetLongValue() int64 {
	if m != nil {
		return m.LongValue
	}
	return 0
}

func (m *PLongIntIntByteByteStringValue) GetIntValue1() int32 {
	if m != nil {
		return m.IntValue1
	}
	return 0
}

func (m *PLongIntIntByteByteStringValue) GetIntValue2() int32 {
	if m != nil {
		return m.IntValue2
	}
	return 0
}

func (m *PLongIntIntByteByteStringValue) GetByteValue1() int32 {
	if m != nil {
		return m.ByteValue1
	}
	return 0
}

func (m *PLongIntIntByteByteStringValue) GetByteValue2() int32 {
	if m != nil {
		return m.ByteValue2
	}
	return 0
}

func (m *PLongIntIntByteByteStringValue) GetStringValue() *wrapperspb.StringValue {
	if m != nil {
		return m.StringValue
	}
	return nil
}

type PIntBooleanIntBooleanValue struct {
	IntValue1            int32    `protobuf:"varint,1,opt,name=intValue1,proto3" json:"intValue1,omitempty"`
	BoolValue1           bool     `protobuf:"varint,2,opt,name=boolValue1,proto3" json:"boolValue1,omitempty"`
	IntValue2            int32    `protobuf:"varint,3,opt,name=intValue2,proto3" json:"intValue2,omitempty"`
	BoolValue2           bool     `protobuf:"varint,4,opt,name=boolValue2,proto3" json:"boolValue2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PIntBooleanIntBooleanValue) Reset()         { *m = PIntBooleanIntBooleanValue{} }
func (m *PIntBooleanIntBooleanValue) String() string { return proto.CompactTextString(m) }
func (*PIntBooleanIntBooleanValue) ProtoMessage()    {}
func (*PIntBooleanIntBooleanValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe607a737cbfdf8, []int{3}
}

func (m *PIntBooleanIntBooleanValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PIntBooleanIntBooleanValue.Unmarshal(m, b)
}
func (m *PIntBooleanIntBooleanValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PIntBooleanIntBooleanValue.Marshal(b, m, deterministic)
}
func (m *PIntBooleanIntBooleanValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PIntBooleanIntBooleanValue.Merge(m, src)
}
func (m *PIntBooleanIntBooleanValue) XXX_Size() int {
	return xxx_messageInfo_PIntBooleanIntBooleanValue.Size(m)
}
func (m *PIntBooleanIntBooleanValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PIntBooleanIntBooleanValue.DiscardUnknown(m)
}

var xxx_messageInfo_PIntBooleanIntBooleanValue proto.InternalMessageInfo

func (m *PIntBooleanIntBooleanValue) GetIntValue1() int32 {
	if m != nil {
		return m.IntValue1
	}
	return 0
}

func (m *PIntBooleanIntBooleanValue) GetBoolValue1() bool {
	if m != nil {
		return m.BoolValue1
	}
	return false
}

func (m *PIntBooleanIntBooleanValue) GetIntValue2() int32 {
	if m != nil {
		return m.IntValue2
	}
	return 0
}

func (m *PIntBooleanIntBooleanValue) GetBoolValue2() bool {
	if m != nil {
		return m.BoolValue2
	}
	return false
}

type PStringStringValue struct {
	StringValue1         *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=stringValue1,proto3" json:"stringValue1,omitempty"`
	StringValue2         *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=stringValue2,proto3" json:"stringValue2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PStringStringValue) Reset()         { *m = PStringStringValue{} }
func (m *PStringStringValue) String() string { return proto.CompactTextString(m) }
func (*PStringStringValue) ProtoMessage()    {}
func (*PStringStringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe607a737cbfdf8, []int{4}
}

func (m *PStringStringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PStringStringValue.Unmarshal(m, b)
}
func (m *PStringStringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PStringStringValue.Marshal(b, m, deterministic)
}
func (m *PStringStringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PStringStringValue.Merge(m, src)
}
func (m *PStringStringValue) XXX_Size() int {
	return xxx_messageInfo_PStringStringValue.Size(m)
}
func (m *PStringStringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PStringStringValue.DiscardUnknown(m)
}

var xxx_messageInfo_PStringStringValue proto.InternalMessageInfo

func (m *PStringStringValue) GetStringValue1() *wrapperspb.StringValue {
	if m != nil {
		return m.StringValue1
	}
	return nil
}

func (m *PStringStringValue) GetStringValue2() *wrapperspb.StringValue {
	if m != nil {
		return m.StringValue2
	}
	return nil
}

type PAnnotationValue struct {
	// Types that are valid to be assigned to Field:
	//	*PAnnotationValue_StringValue
	//	*PAnnotationValue_BoolValue
	//	*PAnnotationValue_IntValue
	//	*PAnnotationValue_LongValue
	//	*PAnnotationValue_ShortValue
	//	*PAnnotationValue_DoubleValue
	//	*PAnnotationValue_BinaryValue
	//	*PAnnotationValue_ByteValue
	//	*PAnnotationValue_IntStringValue
	//	*PAnnotationValue_StringStringValue
	//	*PAnnotationValue_IntStringStringValue
	//	*PAnnotationValue_LongIntIntByteByteStringValue
	//	*PAnnotationValue_IntBooleanIntBooleanValue
	Field                isPAnnotationValue_Field `protobuf_oneof:"field"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PAnnotationValue) Reset()         { *m = PAnnotationValue{} }
func (m *PAnnotationValue) String() string { return proto.CompactTextString(m) }
func (*PAnnotationValue) ProtoMessage()    {}
func (*PAnnotationValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe607a737cbfdf8, []int{5}
}

func (m *PAnnotationValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PAnnotationValue.Unmarshal(m, b)
}
func (m *PAnnotationValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PAnnotationValue.Marshal(b, m, deterministic)
}
func (m *PAnnotationValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PAnnotationValue.Merge(m, src)
}
func (m *PAnnotationValue) XXX_Size() int {
	return xxx_messageInfo_PAnnotationValue.Size(m)
}
func (m *PAnnotationValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PAnnotationValue.DiscardUnknown(m)
}

var xxx_messageInfo_PAnnotationValue proto.InternalMessageInfo

type isPAnnotationValue_Field interface {
	isPAnnotationValue_Field()
}

type PAnnotationValue_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=stringValue,proto3,oneof"`
}

type PAnnotationValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=boolValue,proto3,oneof"`
}

type PAnnotationValue_IntValue struct {
	IntValue int32 `protobuf:"varint,3,opt,name=intValue,proto3,oneof"`
}

type PAnnotationValue_LongValue struct {
	LongValue int64 `protobuf:"varint,4,opt,name=longValue,proto3,oneof"`
}

type PAnnotationValue_ShortValue struct {
	ShortValue int32 `protobuf:"zigzag32,5,opt,name=shortValue,proto3,oneof"`
}

type PAnnotationValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,6,opt,name=doubleValue,proto3,oneof"`
}

type PAnnotationValue_BinaryValue struct {
	BinaryValue []byte `protobuf:"bytes,7,opt,name=binaryValue,proto3,oneof"`
}

type PAnnotationValue_ByteValue struct {
	ByteValue int32 `protobuf:"zigzag32,8,opt,name=byteValue,proto3,oneof"`
}

type PAnnotationValue_IntStringValue struct {
	IntStringValue *PIntStringValue `protobuf:"bytes,9,opt,name=intStringValue,proto3,oneof"`
}

type PAnnotationValue_StringStringValue struct {
	StringStringValue *PStringStringValue `protobuf:"bytes,10,opt,name=stringStringValue,proto3,oneof"`
}

type PAnnotationValue_IntStringStringValue struct {
	IntStringStringValue *PIntStringStringValue `protobuf:"bytes,11,opt,name=intStringStringValue,proto3,oneof"`
}

type PAnnotationValue_LongIntIntByteByteStringValue struct {
	LongIntIntByteByteStringValue *PLongIntIntByteByteStringValue `protobuf:"bytes,12,opt,name=longIntIntByteByteStringValue,proto3,oneof"`
}

type PAnnotationValue_IntBooleanIntBooleanValue struct {
	IntBooleanIntBooleanValue *PIntBooleanIntBooleanValue `protobuf:"bytes,13,opt,name=intBooleanIntBooleanValue,proto3,oneof"`
}

func (*PAnnotationValue_StringValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_BoolValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_IntValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_LongValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_ShortValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_DoubleValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_BinaryValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_ByteValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_IntStringValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_StringStringValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_IntStringStringValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_LongIntIntByteByteStringValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_IntBooleanIntBooleanValue) isPAnnotationValue_Field() {}

func (m *PAnnotationValue) GetField() isPAnnotationValue_Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *PAnnotationValue) GetStringValue() string {
	if x, ok := m.GetField().(*PAnnotationValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *PAnnotationValue) GetBoolValue() bool {
	if x, ok := m.GetField().(*PAnnotationValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *PAnnotationValue) GetIntValue() int32 {
	if x, ok := m.GetField().(*PAnnotationValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *PAnnotationValue) GetLongValue() int64 {
	if x, ok := m.GetField().(*PAnnotationValue_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (m *PAnnotationValue) GetShortValue() int32 {
	if x, ok := m.GetField().(*PAnnotationValue_ShortValue); ok {
		return x.ShortValue
	}
	return 0
}

func (m *PAnnotationValue) GetDoubleValue() float64 {
	if x, ok := m.GetField().(*PAnnotationValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *PAnnotationValue) GetBinaryValue() []byte {
	if x, ok := m.GetField().(*PAnnotationValue_BinaryValue); ok {
		return x.BinaryValue
	}
	return nil
}

func (m *PAnnotationValue) GetByteValue() int32 {
	if x, ok := m.GetField().(*PAnnotationValue_ByteValue); ok {
		return x.ByteValue
	}
	return 0
}

func (m *PAnnotationValue) GetIntStringValue() *PIntStringValue {
	if x, ok := m.GetField().(*PAnnotationValue_IntStringValue); ok {
		return x.IntStringValue
	}
	return nil
}

func (m *PAnnotationValue) GetStringStringValue() *PStringStringValue {
	if x, ok := m.GetField().(*PAnnotationValue_StringStringValue); ok {
		return x.StringStringValue
	}
	return nil
}

func (m *PAnnotationValue) GetIntStringStringValue() *PIntStringStringValue {
	if x, ok := m.GetField().(*PAnnotationValue_IntStringStringValue); ok {
		return x.IntStringStringValue
	}
	return nil
}

func (m *PAnnotationValue) GetLongIntIntByteByteStringValue() *PLongIntIntByteByteStringValue {
	if x, ok := m.GetField().(*PAnnotationValue_LongIntIntByteByteStringValue); ok {
		return x.LongIntIntByteByteStringValue
	}
	return nil
}

func (m *PAnnotationValue) GetIntBooleanIntBooleanValue() *PIntBooleanIntBooleanValue {
	if x, ok := m.GetField().(*PAnnotationValue_IntBooleanIntBooleanValue); ok {
		return x.IntBooleanIntBooleanValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PAnnotationValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PAnnotationValue_StringValue)(nil),
		(*PAnnotationValue_BoolValue)(nil),
		(*PAnnotationValue_IntValue)(nil),
		(*PAnnotationValue_LongValue)(nil),
		(*PAnnotationValue_ShortValue)(nil),
		(*PAnnotationValue_DoubleValue)(nil),
		(*PAnnotationValue_BinaryValue)(nil),
		(*PAnnotationValue_ByteValue)(nil),
		(*PAnnotationValue_IntStringValue)(nil),
		(*PAnnotationValue_StringStringValue)(nil),
		(*PAnnotationValue_IntStringStringValue)(nil),
		(*PAnnotationValue_LongIntIntByteByteStringValue)(nil),
		(*PAnnotationValue_IntBooleanIntBooleanValue)(nil),
	}
}

type PAnnotation struct {
	Key                  int32             `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *PAnnotationValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PAnnotation) Reset()         { *m = PAnnotation{} }
func (m *PAnnotation) String() string { return proto.CompactTextString(m) }
func (*PAnnotation) ProtoMessage()    {}
func (*PAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_abe607a737cbfdf8, []int{6}
}

func (m *PAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PAnnotation.Unmarshal(m, b)
}
func (m *PAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PAnnotation.Marshal(b, m, deterministic)
}
func (m *PAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PAnnotation.Merge(m, src)
}
func (m *PAnnotation) XXX_Size() int {
	return xxx_messageInfo_PAnnotation.Size(m)
}
func (m *PAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_PAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_PAnnotation proto.InternalMessageInfo

func (m *PAnnotation) GetKey() int32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *PAnnotation) GetValue() *PAnnotationValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*PIntStringValue)(nil), "v1.PIntStringValue")
	proto.RegisterType((*PIntStringStringValue)(nil), "v1.PIntStringStringValue")
	proto.RegisterType((*PLongIntIntByteByteStringValue)(nil), "v1.PLongIntIntByteByteStringValue")
	proto.RegisterType((*PIntBooleanIntBooleanValue)(nil), "v1.PIntBooleanIntBooleanValue")
	proto.RegisterType((*PStringStringValue)(nil), "v1.PStringStringValue")
	proto.RegisterType((*PAnnotationValue)(nil), "v1.PAnnotationValue")
	proto.RegisterType((*PAnnotation)(nil), "v1.PAnnotation")
}

func init() { proto.RegisterFile("v1/Annotation.proto", fileDescriptor_abe607a737cbfdf8) }

var fileDescriptor_abe607a737cbfdf8 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xf5, 0x36, 0xb8, 0x4d, 0x26, 0x81, 0xb6, 0xdb, 0x82, 0xdc, 0x28, 0x58, 0xc1, 0xa7, 0x88,
	0x83, 0xa3, 0x98, 0x23, 0x02, 0x41, 0x0e, 0xc8, 0x11, 0x48, 0x58, 0x8b, 0xc4, 0x81, 0x03, 0x92,
	0x93, 0xba, 0xc1, 0xe0, 0xee, 0x5a, 0xce, 0xc6, 0x28, 0x7f, 0xc2, 0x9f, 0xf0, 0x11, 0xfc, 0x11,
	0x27, 0xb4, 0xeb, 0xb8, 0xbb, 0x76, 0xd2, 0x28, 0x11, 0x87, 0x48, 0x9b, 0xd9, 0x99, 0x37, 0xf3,
	0x66, 0xdf, 0x8c, 0xe1, 0x22, 0x1f, 0x0d, 0xdf, 0x52, 0xca, 0x78, 0xc8, 0x63, 0x46, 0xdd, 0x34,
	0x63, 0x9c, 0xe1, 0xa3, 0x7c, 0xd4, 0xb5, 0xe7, 0x8c, 0xcd, 0x93, 0x68, 0x28, 0x2d, 0xd3, 0xe5,
	0xcd, 0xf0, 0x67, 0x16, 0xa6, 0x69, 0x94, 0x2d, 0x0a, 0x1f, 0xe7, 0x16, 0x4e, 0x83, 0x09, 0xe5,
	0x9f, 0x78, 0x16, 0xd3, 0xf9, 0xe7, 0x30, 0x59, 0x46, 0xb8, 0x0b, 0xcd, 0x98, 0x72, 0x79, 0xb6,
	0x50, 0x1f, 0x0d, 0x4c, 0x72, 0xf7, 0x1f, 0xbf, 0x86, 0xf6, 0x42, 0xb9, 0x5a, 0x47, 0x7d, 0x34,
	0x68, 0x7b, 0x3d, 0xb7, 0x48, 0xe2, 0x96, 0x49, 0x5c, 0x0d, 0x8e, 0xe8, 0x01, 0xce, 0x6f, 0x04,
	0x8f, 0x55, 0xbe, 0x7d, 0xb3, 0xbe, 0x81, 0x8e, 0x06, 0x32, 0xda, 0x2b, 0x6d, 0x25, 0xa2, 0x86,
	0xe0, 0x59, 0x8d, 0x03, 0x11, 0x3c, 0xe7, 0x2f, 0x02, 0x3b, 0xf8, 0xc0, 0xe8, 0x7c, 0x42, 0xf9,
	0x84, 0xf2, 0xf1, 0x8a, 0x47, 0xe2, 0xa7, 0x53, 0xe8, 0x41, 0x2b, 0x61, 0x65, 0x6b, 0x04, 0x87,
	0x06, 0x51, 0x06, 0x71, 0x5b, 0x12, 0x2a, 0x18, 0x98, 0x44, 0x19, 0xf4, 0xdb, 0xa2, 0x3a, 0xed,
	0xd6, 0xc3, 0x36, 0xc0, 0x74, 0xc5, 0xa3, 0x75, 0xf0, 0x83, 0x3e, 0x1a, 0x9c, 0x13, 0xcd, 0x52,
	0xb9, 0xf7, 0x2c, 0xb3, 0x76, 0xef, 0xd5, 0x9f, 0xed, 0xf8, 0xd0, 0x67, 0xfb, 0x85, 0xa0, 0x2b,
	0x9e, 0x6d, 0xcc, 0x58, 0x12, 0x85, 0x54, 0x9d, 0xb6, 0x50, 0x43, 0x75, 0x6a, 0xa2, 0x38, 0xc6,
	0x12, 0x8d, 0x79, 0x93, 0x68, 0x96, 0x3d, 0xa8, 0x97, 0xbe, 0x9e, 0xa4, 0xae, 0x47, 0x7b, 0xa2,
	0x34, 0x1c, 0x6c, 0xca, 0xa9, 0x2e, 0x19, 0xf4, 0xdf, 0x92, 0x39, 0x3a, 0x58, 0x32, 0x7f, 0x4c,
	0x38, 0x0b, 0xd4, 0x54, 0x16, 0x85, 0x39, 0xd5, 0xa7, 0x10, 0x75, 0xb5, 0x7c, 0xa3, 0xd2, 0x6e,
	0x6c, 0x43, 0xeb, 0x8e, 0x61, 0xd1, 0x30, 0xdf, 0x20, 0xca, 0x84, 0x7b, 0xda, 0xac, 0xc8, 0x86,
	0xf9, 0x86, 0x36, 0x2d, 0xb6, 0x2e, 0x43, 0xd1, 0xb0, 0x86, 0x88, 0x56, 0x42, 0xec, 0x03, 0x2c,
	0xbe, 0xb1, 0x6c, 0x1d, 0x2f, 0xc5, 0xe2, 0x1b, 0x44, 0xb3, 0x89, 0x1a, 0xaf, 0xd9, 0x72, 0x9a,
	0x44, 0x4a, 0x2e, 0x48, 0xd4, 0xa8, 0x19, 0x85, 0xcf, 0x34, 0xa6, 0x61, 0xb6, 0x2a, 0x7c, 0x4e,
	0xfa, 0x68, 0xd0, 0x11, 0x3e, 0x9a, 0x51, 0xf2, 0x28, 0x45, 0x68, 0x35, 0xd7, 0x89, 0x94, 0x09,
	0xbf, 0x82, 0x47, 0x71, 0x65, 0xf7, 0x58, 0x2d, 0xd9, 0xe4, 0x0b, 0x37, 0x1f, 0xb9, 0xb5, 0xb5,
	0xe4, 0x1b, 0xa4, 0xe6, 0x8c, 0xdf, 0xc1, 0xf9, 0xa2, 0xfe, 0xf0, 0x16, 0x48, 0x84, 0x27, 0x12,
	0x61, 0x43, 0x16, 0xbe, 0x41, 0x36, 0x43, 0xf0, 0x47, 0xb8, 0x8c, 0xb7, 0xac, 0x24, 0xab, 0x2d,
	0xa1, 0xae, 0xaa, 0xc5, 0x54, 0xd1, 0xb6, 0x06, 0xe2, 0xef, 0xf0, 0x34, 0xd9, 0xb5, 0x29, 0xac,
	0x8e, 0x44, 0x76, 0x24, 0xf2, 0xce, 0x9d, 0xe2, 0x1b, 0x64, 0x37, 0x14, 0xfe, 0x0a, 0x57, 0xf1,
	0x7d, 0x83, 0x69, 0x3d, 0x94, 0x79, 0xec, 0x92, 0xc1, 0x76, 0x2f, 0xdf, 0x20, 0xf7, 0x43, 0x8c,
	0x4f, 0xc0, 0xbc, 0x89, 0xa3, 0xe4, 0xda, 0x79, 0x0f, 0x6d, 0x4d, 0xcc, 0xf8, 0x0c, 0x1a, 0x3f,
	0xa2, 0xd5, 0x7a, 0xda, 0xc5, 0x11, 0x3f, 0x07, 0x33, 0xd7, 0xbe, 0x0a, 0x97, 0x32, 0x6b, 0x4d,
	0xfe, 0xa4, 0x70, 0x19, 0xbf, 0x84, 0x67, 0x33, 0x76, 0xeb, 0xd2, 0x30, 0x8f, 0xb2, 0x19, 0xcb,
	0x52, 0x37, 0x8d, 0x69, 0xca, 0x62, 0xca, 0xdd, 0x79, 0x96, 0xce, 0x5c, 0x9e, 0x85, 0xb3, 0x68,
	0x7c, 0xaa, 0x82, 0x03, 0x31, 0x6d, 0x01, 0xfa, 0xd2, 0x18, 0xe6, 0xa3, 0xe9, 0xb1, 0x9c, 0xbd,
	0x17, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x60, 0x71, 0xc8, 0xf5, 0x06, 0x00, 0x00,
}
