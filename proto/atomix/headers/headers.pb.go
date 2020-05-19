// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/headers/headers.proto

package headers

import (
	fmt "fmt"
	primitive "github.com/atomix/api/proto/atomix/primitive"
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
	return fileDescriptor_95de8fb3e0e30ffd, []int{0}
}

type ResponseStatus int32

const (
	ResponseStatus_OK         ResponseStatus = 0
	ResponseStatus_ERROR      ResponseStatus = 1
	ResponseStatus_NOT_LEADER ResponseStatus = 2
)

var ResponseStatus_name = map[int32]string{
	0: "OK",
	1: "ERROR",
	2: "NOT_LEADER",
}

var ResponseStatus_value = map[string]int32{
	"OK":         0,
	"ERROR":      1,
	"NOT_LEADER": 2,
}

func (x ResponseStatus) String() string {
	return proto.EnumName(ResponseStatus_name, int32(x))
}

func (ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_95de8fb3e0e30ffd, []int{1}
}

type RequestHeader struct {
	Partition int32                 `protobuf:"varint,6,opt,name=partition,proto3" json:"partition,omitempty"`
	Primitive primitive.PrimitiveId `protobuf:"bytes,1,opt,name=primitive,proto3" json:"primitive"`
	SessionID uint64                `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	RequestID uint64                `protobuf:"varint,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Index     uint64                `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	Streams   []StreamHeader        `protobuf:"bytes,5,rep,name=streams,proto3" json:"streams"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_95de8fb3e0e30ffd, []int{0}
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

func (m *RequestHeader) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *RequestHeader) GetPrimitive() primitive.PrimitiveId {
	if m != nil {
		return m.Primitive
	}
	return primitive.PrimitiveId{}
}

func (m *RequestHeader) GetSessionID() uint64 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

func (m *RequestHeader) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func (m *RequestHeader) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RequestHeader) GetStreams() []StreamHeader {
	if m != nil {
		return m.Streams
	}
	return nil
}

type ResponseHeader struct {
	SessionID  uint64         `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	StreamID   uint64         `protobuf:"varint,2,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	ResponseID uint64         `protobuf:"varint,3,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
	Index      uint64         `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	Leader     string         `protobuf:"bytes,5,opt,name=leader,proto3" json:"leader,omitempty"`
	Status     ResponseStatus `protobuf:"varint,6,opt,name=status,proto3,enum=atomix.headers.ResponseStatus" json:"status,omitempty"`
	Type       ResponseType   `protobuf:"varint,7,opt,name=type,proto3,enum=atomix.headers.ResponseType" json:"type,omitempty"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_95de8fb3e0e30ffd, []int{1}
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

func (m *ResponseHeader) GetSessionID() uint64 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

func (m *ResponseHeader) GetStreamID() uint64 {
	if m != nil {
		return m.StreamID
	}
	return 0
}

func (m *ResponseHeader) GetResponseID() uint64 {
	if m != nil {
		return m.ResponseID
	}
	return 0
}

func (m *ResponseHeader) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ResponseHeader) GetLeader() string {
	if m != nil {
		return m.Leader
	}
	return ""
}

func (m *ResponseHeader) GetStatus() ResponseStatus {
	if m != nil {
		return m.Status
	}
	return ResponseStatus_OK
}

func (m *ResponseHeader) GetType() ResponseType {
	if m != nil {
		return m.Type
	}
	return ResponseType_RESPONSE
}

type StreamHeader struct {
	StreamID   uint64 `protobuf:"varint,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	ResponseID uint64 `protobuf:"varint,2,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
}

func (m *StreamHeader) Reset()         { *m = StreamHeader{} }
func (m *StreamHeader) String() string { return proto.CompactTextString(m) }
func (*StreamHeader) ProtoMessage()    {}
func (*StreamHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_95de8fb3e0e30ffd, []int{2}
}
func (m *StreamHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamHeader.Merge(m, src)
}
func (m *StreamHeader) XXX_Size() int {
	return m.Size()
}
func (m *StreamHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamHeader.DiscardUnknown(m)
}

var xxx_messageInfo_StreamHeader proto.InternalMessageInfo

func (m *StreamHeader) GetStreamID() uint64 {
	if m != nil {
		return m.StreamID
	}
	return 0
}

func (m *StreamHeader) GetResponseID() uint64 {
	if m != nil {
		return m.ResponseID
	}
	return 0
}

func init() {
	proto.RegisterEnum("atomix.headers.ResponseType", ResponseType_name, ResponseType_value)
	proto.RegisterEnum("atomix.headers.ResponseStatus", ResponseStatus_name, ResponseStatus_value)
	proto.RegisterType((*RequestHeader)(nil), "atomix.headers.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "atomix.headers.ResponseHeader")
	proto.RegisterType((*StreamHeader)(nil), "atomix.headers.StreamHeader")
}

func init() { proto.RegisterFile("atomix/headers/headers.proto", fileDescriptor_95de8fb3e0e30ffd) }

var fileDescriptor_95de8fb3e0e30ffd = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0xb5, 0x8a, 0xed, 0x44, 0x63, 0x47, 0x15, 0x4b, 0x28, 0x22, 0xb8, 0xb2, 0xf0, 0x49,
	0x0d, 0x45, 0x2e, 0x0e, 0xf4, 0x54, 0x28, 0x76, 0x2d, 0xa8, 0x68, 0x6a, 0x99, 0x95, 0xef, 0x46,
	0x45, 0x4b, 0xaa, 0x52, 0x5b, 0xaa, 0x76, 0x53, 0x92, 0xb7, 0xe8, 0x23, 0xf4, 0x19, 0xfa, 0x14,
	0x39, 0xe6, 0xd8, 0x93, 0x29, 0xf2, 0x8b, 0x14, 0x69, 0x57, 0x55, 0x5c, 0x08, 0xa1, 0x27, 0xcd,
	0xce, 0x7e, 0x33, 0xf3, 0xcf, 0xbf, 0x08, 0xfa, 0x11, 0x4f, 0xd7, 0xc9, 0xf5, 0xe8, 0x13, 0x8d,
	0x62, 0x9a, 0xb3, 0xfa, 0xeb, 0x66, 0x79, 0xca, 0x53, 0xac, 0x8b, 0x5b, 0x57, 0x66, 0x4f, 0x6d,
	0x49, 0x67, 0x79, 0xb2, 0x4e, 0x78, 0xf2, 0x8d, 0x36, 0x91, 0xa8, 0x38, 0x3d, 0xb9, 0x4c, 0x2f,
	0xd3, 0x2a, 0x1c, 0x95, 0x91, 0xc8, 0x0e, 0x7f, 0xa8, 0x70, 0x4c, 0xe8, 0xd7, 0x2b, 0xca, 0xf8,
	0xbb, 0xaa, 0x15, 0xee, 0x83, 0x96, 0x45, 0x39, 0x4f, 0x78, 0x92, 0x6e, 0xcc, 0x8e, 0x8d, 0x9c,
	0x36, 0x69, 0x12, 0x78, 0x02, 0xda, 0xdf, 0xc6, 0x26, 0xb2, 0x91, 0xd3, 0x1d, 0x3f, 0x73, 0xa5,
	0x96, 0x66, 0xe2, 0xa2, 0x8e, 0xfc, 0x78, 0xda, 0xba, 0xdd, 0x0e, 0x14, 0xd2, 0x54, 0xe1, 0x17,
	0x00, 0x8c, 0x32, 0x96, 0xa4, 0x9b, 0x55, 0x12, 0x9b, 0xaa, 0x8d, 0x9c, 0xd6, 0xf4, 0xb8, 0xd8,
	0x0e, 0xb4, 0x50, 0x64, 0xfd, 0x19, 0xd1, 0x24, 0xe0, 0xc7, 0x25, 0x9d, 0x0b, 0x7d, 0x25, 0x7d,
	0xd0, 0xd0, 0x52, 0x75, 0x49, 0x4b, 0xc0, 0x8f, 0xf1, 0x09, 0xb4, 0x93, 0x4d, 0x4c, 0xaf, 0xcd,
	0x56, 0x09, 0x12, 0x71, 0xc0, 0xaf, 0xe1, 0x90, 0xf1, 0x9c, 0x46, 0x6b, 0x66, 0xb6, 0xed, 0x03,
	0xa7, 0x3b, 0xee, 0xbb, 0xfb, 0xf6, 0xb9, 0x61, 0x75, 0x2d, 0x1c, 0x90, 0x8a, 0xeb, 0x92, 0xe1,
	0x4f, 0x15, 0x74, 0x42, 0x59, 0x96, 0x6e, 0x18, 0x95, 0x1e, 0xed, 0xaf, 0x80, 0x1e, 0x59, 0xe1,
	0x39, 0x68, 0xa2, 0x57, 0xb3, 0x6f, 0xaf, 0xd8, 0x0e, 0x8e, 0xc4, 0x50, 0x7f, 0x46, 0x8e, 0xc4,
	0xb5, 0x1f, 0xe3, 0x11, 0x74, 0x73, 0x39, 0xaa, 0x59, 0x57, 0x2f, 0xb6, 0x03, 0xa8, 0x15, 0xf8,
	0x33, 0x02, 0x35, 0xf2, 0xe0, 0xc2, 0x4f, 0xa1, 0xf3, 0xa5, 0x52, 0x6a, 0xb6, 0x6d, 0xe4, 0x68,
	0x44, 0x9e, 0xf0, 0x2b, 0xe8, 0x30, 0x1e, 0xf1, 0x2b, 0x56, 0x3d, 0xac, 0x3e, 0xb6, 0xfe, 0xf5,
	0xa1, 0x9e, 0x12, 0x56, 0x14, 0x91, 0x34, 0x7e, 0x09, 0x2d, 0x7e, 0x93, 0x51, 0xf3, 0xb0, 0xaa,
	0xea, 0x3f, 0x54, 0xb5, 0xbc, 0xc9, 0x28, 0xa9, 0xc8, 0xe1, 0x67, 0xe8, 0xdd, 0xf7, 0x74, 0xdf,
	0x03, 0xf4, 0x3f, 0x1e, 0xa8, 0x8f, 0x79, 0x70, 0xf6, 0x06, 0x7a, 0xf7, 0x15, 0xe0, 0x1e, 0x1c,
	0x11, 0x2f, 0x5c, 0x04, 0xf3, 0xd0, 0x33, 0x14, 0xfc, 0x04, 0xba, 0xc1, 0xc2, 0x9b, 0xaf, 0xc2,
	0x25, 0xf1, 0x26, 0x1f, 0x0c, 0x84, 0x0d, 0xe8, 0xbd, 0xbd, 0x08, 0x42, 0xaf, 0xce, 0xa8, 0x67,
	0xe7, 0xcd, 0x03, 0x8b, 0xc5, 0x71, 0x07, 0xd4, 0xe0, 0xbd, 0xa1, 0x60, 0x0d, 0xda, 0x1e, 0x21,
	0x01, 0x31, 0x10, 0xd6, 0x01, 0xe6, 0xc1, 0x72, 0x75, 0xe1, 0x4d, 0x66, 0x1e, 0x31, 0xd4, 0xa9,
	0x79, 0x5b, 0x58, 0xe8, 0xae, 0xb0, 0xd0, 0xef, 0xc2, 0x42, 0xdf, 0x77, 0x96, 0x72, 0xb7, 0xb3,
	0x94, 0x5f, 0x3b, 0x4b, 0xf9, 0xd8, 0xa9, 0x7e, 0xad, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xb7, 0x01, 0x2b, 0xeb, 0xc2, 0x03, 0x00, 0x00,
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
	if m.Partition != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.Partition))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Streams) > 0 {
		for iNdEx := len(m.Streams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Streams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHeaders(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Index != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x20
	}
	if m.RequestID != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.RequestID))
		i--
		dAtA[i] = 0x18
	}
	if m.SessionID != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.SessionID))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Primitive.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintHeaders(dAtA, i, uint64(size))
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
		i = encodeVarintHeaders(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x38
	}
	if m.Status != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Leader) > 0 {
		i -= len(m.Leader)
		copy(dAtA[i:], m.Leader)
		i = encodeVarintHeaders(dAtA, i, uint64(len(m.Leader)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Index != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x20
	}
	if m.ResponseID != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.ResponseID))
		i--
		dAtA[i] = 0x18
	}
	if m.StreamID != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.StreamID))
		i--
		dAtA[i] = 0x10
	}
	if m.SessionID != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.SessionID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StreamHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ResponseID != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.ResponseID))
		i--
		dAtA[i] = 0x10
	}
	if m.StreamID != 0 {
		i = encodeVarintHeaders(dAtA, i, uint64(m.StreamID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHeaders(dAtA []byte, offset int, v uint64) int {
	offset -= sovHeaders(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Primitive.Size()
	n += 1 + l + sovHeaders(uint64(l))
	if m.SessionID != 0 {
		n += 1 + sovHeaders(uint64(m.SessionID))
	}
	if m.RequestID != 0 {
		n += 1 + sovHeaders(uint64(m.RequestID))
	}
	if m.Index != 0 {
		n += 1 + sovHeaders(uint64(m.Index))
	}
	if len(m.Streams) > 0 {
		for _, e := range m.Streams {
			l = e.Size()
			n += 1 + l + sovHeaders(uint64(l))
		}
	}
	if m.Partition != 0 {
		n += 1 + sovHeaders(uint64(m.Partition))
	}
	return n
}

func (m *ResponseHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SessionID != 0 {
		n += 1 + sovHeaders(uint64(m.SessionID))
	}
	if m.StreamID != 0 {
		n += 1 + sovHeaders(uint64(m.StreamID))
	}
	if m.ResponseID != 0 {
		n += 1 + sovHeaders(uint64(m.ResponseID))
	}
	if m.Index != 0 {
		n += 1 + sovHeaders(uint64(m.Index))
	}
	l = len(m.Leader)
	if l > 0 {
		n += 1 + l + sovHeaders(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovHeaders(uint64(m.Status))
	}
	if m.Type != 0 {
		n += 1 + sovHeaders(uint64(m.Type))
	}
	return n
}

func (m *StreamHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StreamID != 0 {
		n += 1 + sovHeaders(uint64(m.StreamID))
	}
	if m.ResponseID != 0 {
		n += 1 + sovHeaders(uint64(m.ResponseID))
	}
	return n
}

func sovHeaders(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHeaders(x uint64) (n int) {
	return sovHeaders(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaders
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
					return ErrIntOverflowHeaders
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
				return ErrInvalidLengthHeaders
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHeaders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Primitive.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			m.SessionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			m.RequestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Streams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
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
				return ErrInvalidLengthHeaders
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHeaders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Streams = append(m.Streams, StreamHeader{})
			if err := m.Streams[len(m.Streams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Partition", wireType)
			}
			m.Partition = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Partition |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeaders(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeaders
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHeaders
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
				return ErrIntOverflowHeaders
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
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			m.SessionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamID", wireType)
			}
			m.StreamID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseID", wireType)
			}
			m.ResponseID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResponseID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
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
				return ErrInvalidLengthHeaders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHeaders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Leader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ResponseStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
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
			skippy, err := skipHeaders(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeaders
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHeaders
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
func (m *StreamHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaders
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
			return fmt.Errorf("proto: StreamHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamID", wireType)
			}
			m.StreamID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseID", wireType)
			}
			m.ResponseID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResponseID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeaders(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeaders
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHeaders
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
func skipHeaders(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeaders
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
					return 0, ErrIntOverflowHeaders
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
					return 0, ErrIntOverflowHeaders
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
				return 0, ErrInvalidLengthHeaders
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHeaders
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHeaders
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHeaders        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeaders          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHeaders = fmt.Errorf("proto: unexpected end of group")
)