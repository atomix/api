// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/storage/primitive.proto

package storage

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Response type
type ResponseType int32

const (
	ResponseType_RESPONSE     ResponseType = 0
	ResponseType_OPEN_STREAM  ResponseType = 1
	ResponseType_CLOSE_STREAM ResponseType = 2
)

var ResponseType_name = map[int32]string{
	0: "RESPONSE",
	1: "OPEN_STREAM",
	2: "CLOSE_STREAM",
}

var ResponseType_value = map[string]int32{
	"RESPONSE":     0,
	"OPEN_STREAM":  1,
	"CLOSE_STREAM": 2,
}

func (x ResponseType) String() string {
	return proto.EnumName(ResponseType_name, int32(x))
}

func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f706c35595eab36b, []int{0}
}

// Response code
type ResponseCode int32

const (
	ResponseCode_OK             ResponseCode = 0
	ResponseCode_ERROR          ResponseCode = 1
	ResponseCode_NOT_LEADER     ResponseCode = 2
	ResponseCode_UNKNOWN        ResponseCode = 3
	ResponseCode_CANCELED       ResponseCode = 4
	ResponseCode_NOT_FOUND      ResponseCode = 5
	ResponseCode_ALREADY_EXISTS ResponseCode = 6
	ResponseCode_UNAUTHORIZED   ResponseCode = 7
	ResponseCode_FORBIDDEN      ResponseCode = 8
	ResponseCode_CONFLICT       ResponseCode = 9
	ResponseCode_INVALID        ResponseCode = 10
	ResponseCode_UNAVAILABLE    ResponseCode = 11
	ResponseCode_NOT_SUPPORTED  ResponseCode = 12
	ResponseCode_TIMEOUT        ResponseCode = 13
	ResponseCode_INTERNAL       ResponseCode = 14
)

var ResponseCode_name = map[int32]string{
	0:  "OK",
	1:  "ERROR",
	2:  "NOT_LEADER",
	3:  "UNKNOWN",
	4:  "CANCELED",
	5:  "NOT_FOUND",
	6:  "ALREADY_EXISTS",
	7:  "UNAUTHORIZED",
	8:  "FORBIDDEN",
	9:  "CONFLICT",
	10: "INVALID",
	11: "UNAVAILABLE",
	12: "NOT_SUPPORTED",
	13: "TIMEOUT",
	14: "INTERNAL",
}

var ResponseCode_value = map[string]int32{
	"OK":             0,
	"ERROR":          1,
	"NOT_LEADER":     2,
	"UNKNOWN":        3,
	"CANCELED":       4,
	"NOT_FOUND":      5,
	"ALREADY_EXISTS": 6,
	"UNAUTHORIZED":   7,
	"FORBIDDEN":      8,
	"CONFLICT":       9,
	"INVALID":        10,
	"UNAVAILABLE":    11,
	"NOT_SUPPORTED":  12,
	"TIMEOUT":        13,
	"INTERNAL":       14,
}

func (x ResponseCode) String() string {
	return proto.EnumName(ResponseCode_name, int32(x))
}

func (ResponseCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f706c35595eab36b, []int{1}
}

// Namespaced primitive identifier
type PrimitiveId struct {
	Type      string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *PrimitiveId) Reset()         { *m = PrimitiveId{} }
func (m *PrimitiveId) String() string { return proto.CompactTextString(m) }
func (*PrimitiveId) ProtoMessage()    {}
func (*PrimitiveId) Descriptor() ([]byte, []int) {
	return fileDescriptor_f706c35595eab36b, []int{0}
}
func (m *PrimitiveId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrimitiveId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrimitiveId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrimitiveId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveId.Merge(m, src)
}
func (m *PrimitiveId) XXX_Size() int {
	return m.Size()
}
func (m *PrimitiveId) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveId.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveId proto.InternalMessageInfo

func (m *PrimitiveId) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PrimitiveId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PrimitiveId) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Request header
type RequestHeader struct {
	Primitive PrimitiveId `protobuf:"bytes,1,opt,name=primitive,proto3" json:"primitive"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_f706c35595eab36b, []int{1}
}
func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return m.Size()
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetPrimitive() PrimitiveId {
	if m != nil {
		return m.Primitive
	}
	return PrimitiveId{}
}

// Response header
type ResponseHeader struct {
	Type ResponseType `protobuf:"varint,1,opt,name=type,proto3,enum=atomix.storage.ResponseType" json:"type,omitempty"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_f706c35595eab36b, []int{2}
}
func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return m.Size()
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetType() ResponseType {
	if m != nil {
		return m.Type
	}
	return ResponseType_RESPONSE
}

// Response status
type ResponseStatus struct {
	Code    ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=atomix.storage.ResponseCode" json:"code,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *ResponseStatus) Reset()         { *m = ResponseStatus{} }
func (m *ResponseStatus) String() string { return proto.CompactTextString(m) }
func (*ResponseStatus) ProtoMessage()    {}
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f706c35595eab36b, []int{3}
}
func (m *ResponseStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseStatus.Merge(m, src)
}
func (m *ResponseStatus) XXX_Size() int {
	return m.Size()
}
func (m *ResponseStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseStatus proto.InternalMessageInfo

func (m *ResponseStatus) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_OK
}

func (m *ResponseStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("atomix.storage.ResponseType", ResponseType_name, ResponseType_value)
	proto.RegisterEnum("atomix.storage.ResponseCode", ResponseCode_name, ResponseCode_value)
	proto.RegisterType((*PrimitiveId)(nil), "atomix.storage.PrimitiveId")
	proto.RegisterType((*RequestHeader)(nil), "atomix.storage.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "atomix.storage.ResponseHeader")
	proto.RegisterType((*ResponseStatus)(nil), "atomix.storage.ResponseStatus")
}

func init() { proto.RegisterFile("atomix/storage/primitive.proto", fileDescriptor_f706c35595eab36b) }

var fileDescriptor_f706c35595eab36b = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0xed, 0x34, 0x4d, 0xea, 0x9b, 0xc4, 0xdf, 0x7c, 0x23, 0x16, 0x11, 0x54, 0x06, 0x65,
	0x85, 0xba, 0x48, 0x50, 0x79, 0x80, 0x6a, 0x62, 0x4f, 0xd4, 0x51, 0xdd, 0x19, 0x6b, 0x3c, 0x2e,
	0x7f, 0x84, 0x14, 0x99, 0x66, 0x14, 0x65, 0x91, 0xda, 0xc4, 0x2e, 0xa2, 0x6f, 0xc1, 0x63, 0x75,
	0xd9, 0x25, 0x2b, 0x84, 0x92, 0x1d, 0x4f, 0x81, 0xc6, 0xa9, 0x49, 0x60, 0xc3, 0xca, 0xd7, 0xf7,
	0xde, 0xf3, 0x3b, 0xe7, 0x5a, 0x06, 0x2f, 0x2d, 0xb3, 0xe5, 0xe2, 0xcb, 0xa8, 0x28, 0xb3, 0x55,
	0x3a, 0xd7, 0xa3, 0x7c, 0xb5, 0x58, 0x2e, 0xca, 0xc5, 0x67, 0x3d, 0xcc, 0x57, 0x59, 0x99, 0x61,
	0x77, 0x3b, 0x1f, 0x3e, 0xce, 0x9f, 0x3e, 0x99, 0x67, 0xf3, 0xac, 0x1a, 0x8d, 0x4c, 0xb5, 0xdd,
	0x1a, 0xc4, 0xd0, 0x89, 0x6a, 0x21, 0x9b, 0x61, 0x0c, 0xcd, 0xf2, 0x2e, 0xd7, 0x7d, 0xfb, 0x85,
	0xfd, 0xd2, 0x91, 0x55, 0x6d, 0x7a, 0x37, 0xe9, 0x52, 0xf7, 0x1b, 0xdb, 0x9e, 0xa9, 0xf1, 0x31,
	0x38, 0xe6, 0x59, 0xe4, 0xe9, 0xb5, 0xee, 0x1f, 0x54, 0x83, 0x5d, 0x63, 0x10, 0x41, 0x4f, 0xea,
	0x4f, 0xb7, 0xba, 0x28, 0xcf, 0x75, 0x3a, 0xd3, 0x2b, 0x7c, 0x06, 0xce, 0xef, 0x78, 0x15, 0xbb,
	0x73, 0xfa, 0x6c, 0xf8, 0x67, 0xbe, 0xe1, 0x5e, 0x8c, 0x71, 0xf3, 0xfe, 0xfb, 0x73, 0x4b, 0xee,
	0x34, 0x83, 0x31, 0xb8, 0x52, 0x17, 0x79, 0x76, 0x53, 0xe8, 0x47, 0xe4, 0xab, 0xbd, 0xa4, 0xee,
	0xe9, 0xf1, 0xdf, 0xb4, 0x7a, 0x5b, 0xdd, 0xe5, 0x7a, 0x7b, 0xc7, 0xe0, 0xc3, 0x8e, 0x11, 0x97,
	0x69, 0x79, 0x5b, 0x18, 0xc6, 0x75, 0x36, 0xfb, 0x27, 0xc3, 0xcf, 0x66, 0x5a, 0x56, 0x9b, 0xb8,
	0x0f, 0xed, 0xa5, 0x2e, 0x8a, 0x74, 0x5e, 0x7f, 0x8e, 0xfa, 0xf5, 0xe4, 0x0c, 0xba, 0xfb, 0x9e,
	0xb8, 0x0b, 0x47, 0x92, 0xc6, 0x91, 0xe0, 0x31, 0x45, 0x16, 0xfe, 0x0f, 0x3a, 0x22, 0xa2, 0x7c,
	0x1a, 0x2b, 0x49, 0xc9, 0x25, 0xb2, 0x31, 0x82, 0xae, 0x1f, 0x8a, 0x98, 0xd6, 0x9d, 0xc6, 0xc9,
	0x4f, 0x7b, 0x47, 0x30, 0x8e, 0xb8, 0x05, 0x0d, 0x71, 0x81, 0x2c, 0xec, 0xc0, 0x21, 0x95, 0x52,
	0x48, 0x64, 0x63, 0x17, 0x80, 0x0b, 0x35, 0x0d, 0x29, 0x09, 0xa8, 0x44, 0x0d, 0xdc, 0x81, 0x76,
	0xc2, 0x2f, 0xb8, 0x78, 0xc3, 0xd1, 0x81, 0x71, 0xf4, 0x09, 0xf7, 0x69, 0x48, 0x03, 0xd4, 0xc4,
	0x3d, 0x70, 0xcc, 0xea, 0x44, 0x24, 0x3c, 0x40, 0x87, 0x18, 0x83, 0x4b, 0x42, 0x49, 0x49, 0xf0,
	0x6e, 0x4a, 0xdf, 0xb2, 0x58, 0xc5, 0xa8, 0x65, 0x32, 0x24, 0x9c, 0x24, 0xea, 0x5c, 0x48, 0xf6,
	0x9e, 0x06, 0xa8, 0x6d, 0x44, 0x13, 0x21, 0xc7, 0x2c, 0x08, 0x28, 0x47, 0x47, 0x15, 0x51, 0xf0,
	0x49, 0xc8, 0x7c, 0x85, 0x1c, 0x63, 0xc6, 0xf8, 0x15, 0x09, 0x59, 0x80, 0xc0, 0x1c, 0x94, 0x70,
	0x72, 0x45, 0x58, 0x48, 0xc6, 0x21, 0x45, 0x1d, 0xfc, 0x3f, 0xf4, 0x8c, 0x5f, 0x9c, 0x44, 0x91,
	0x90, 0x8a, 0x06, 0xa8, 0x6b, 0x04, 0x8a, 0x5d, 0x52, 0x91, 0x28, 0xd4, 0x33, 0x2c, 0xc6, 0x15,
	0x95, 0x9c, 0x84, 0xc8, 0x1d, 0xf7, 0xef, 0xd7, 0x9e, 0xfd, 0xb0, 0xf6, 0xec, 0x1f, 0x6b, 0xcf,
	0xfe, 0xba, 0xf1, 0xac, 0x87, 0x8d, 0x67, 0x7d, 0xdb, 0x78, 0xd6, 0xc7, 0x56, 0xf5, 0x5f, 0xbe,
	0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x43, 0xd9, 0x14, 0xb9, 0xdf, 0x02, 0x00, 0x00,
}

func (m *PrimitiveId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrimitiveId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrimitiveId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPrimitive(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPrimitive(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintPrimitive(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Primitive.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPrimitive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ResponseHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintPrimitive(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintPrimitive(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintPrimitive(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPrimitive(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrimitive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PrimitiveId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovPrimitive(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPrimitive(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPrimitive(uint64(l))
	}
	return n
}

func (m *RequestHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Primitive.Size()
	n += 1 + l + sovPrimitive(uint64(l))
	return n
}

func (m *ResponseHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovPrimitive(uint64(m.Type))
	}
	return n
}

func (m *ResponseStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovPrimitive(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovPrimitive(uint64(l))
	}
	return n
}

func sovPrimitive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrimitive(x uint64) (n int) {
	return sovPrimitive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PrimitiveId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrimitive
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PrimitiveId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrimitiveId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrimitive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrimitive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPrimitive
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
func (m *RequestHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrimitive
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RequestHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Primitive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Primitive.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrimitive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrimitive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPrimitive
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
func (m *ResponseHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrimitive
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResponseHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ResponseType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPrimitive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrimitive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPrimitive
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
func (m *ResponseStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrimitive
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResponseStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= ResponseCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrimitive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrimitive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPrimitive
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
func skipPrimitive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrimitive
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
					return 0, ErrIntOverflowPrimitive
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPrimitive
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
			if length < 0 {
				return 0, ErrInvalidLengthPrimitive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrimitive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrimitive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrimitive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrimitive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrimitive = fmt.Errorf("proto: unexpected end of group")
)
