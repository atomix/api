// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/controller/database.proto

package controller

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Database identifier
type DatabaseId struct {
	// name is the name of the database
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// namespace is the namespace to which the database belongs
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabaseId) Reset()         { *m = DatabaseId{} }
func (m *DatabaseId) String() string { return proto.CompactTextString(m) }
func (*DatabaseId) ProtoMessage()    {}
func (*DatabaseId) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b4500d882e40b81, []int{0}
}
func (m *DatabaseId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseId.Unmarshal(m, b)
}
func (m *DatabaseId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseId.Marshal(b, m, deterministic)
}
func (m *DatabaseId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseId.Merge(m, src)
}
func (m *DatabaseId) XXX_Size() int {
	return xxx_messageInfo_DatabaseId.Size(m)
}
func (m *DatabaseId) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseId.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseId proto.InternalMessageInfo

func (m *DatabaseId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatabaseId) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Database configuration
type Database struct {
	// id is the database identifier
	ID *DatabaseId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// clusters is a list of clusters in the database
	Clusters []*Cluster `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// partitions is a list of partitions in the cluster
	Partitions           []*Partition `protobuf:"bytes,3,rep,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Database) Reset()         { *m = Database{} }
func (m *Database) String() string { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()    {}
func (*Database) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b4500d882e40b81, []int{1}
}
func (m *Database) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Database.Unmarshal(m, b)
}
func (m *Database) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Database.Marshal(b, m, deterministic)
}
func (m *Database) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Database.Merge(m, src)
}
func (m *Database) XXX_Size() int {
	return xxx_messageInfo_Database.Size(m)
}
func (m *Database) XXX_DiscardUnknown() {
	xxx_messageInfo_Database.DiscardUnknown(m)
}

var xxx_messageInfo_Database proto.InternalMessageInfo

func (m *Database) GetID() *DatabaseId {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Database) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *Database) GetPartitions() []*Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

// Cluster identifier
type ClusterId struct {
	// id is the unique numeric identifier for the cluster
	ID int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// database_id is the identifier for the database to which the cluster belonds
	DatabaseID           *DatabaseId `protobuf:"bytes,2,opt,name=database_id,json=databaseId,proto3" json:"database_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClusterId) Reset()         { *m = ClusterId{} }
func (m *ClusterId) String() string { return proto.CompactTextString(m) }
func (*ClusterId) ProtoMessage()    {}
func (*ClusterId) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b4500d882e40b81, []int{2}
}
func (m *ClusterId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterId.Unmarshal(m, b)
}
func (m *ClusterId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterId.Marshal(b, m, deterministic)
}
func (m *ClusterId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterId.Merge(m, src)
}
func (m *ClusterId) XXX_Size() int {
	return xxx_messageInfo_ClusterId.Size(m)
}
func (m *ClusterId) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterId.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterId proto.InternalMessageInfo

func (m *ClusterId) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ClusterId) GetDatabaseID() *DatabaseId {
	if m != nil {
		return m.DatabaseID
	}
	return nil
}

// Cluster configuration
type Cluster struct {
	// id is the cluster identifier
	ID *ClusterId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// partitions is a list of partitions in the cluster
	Partitions           []*Partition `protobuf:"bytes,2,rep,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b4500d882e40b81, []int{3}
}
func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetID() *ClusterId {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Cluster) GetPartitions() []*Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

// Cluster configuration
type ClusterConfig struct {
	// members is a list of cluster members
	Members []*MemberConfig `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	// partitions is a list of partitions owned by the cluster
	Partitions           []*PartitionId `protobuf:"bytes,2,rep,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ClusterConfig) Reset()         { *m = ClusterConfig{} }
func (m *ClusterConfig) String() string { return proto.CompactTextString(m) }
func (*ClusterConfig) ProtoMessage()    {}
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b4500d882e40b81, []int{4}
}
func (m *ClusterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterConfig.Unmarshal(m, b)
}
func (m *ClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterConfig.Marshal(b, m, deterministic)
}
func (m *ClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterConfig.Merge(m, src)
}
func (m *ClusterConfig) XXX_Size() int {
	return xxx_messageInfo_ClusterConfig.Size(m)
}
func (m *ClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterConfig proto.InternalMessageInfo

func (m *ClusterConfig) GetMembers() []*MemberConfig {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *ClusterConfig) GetPartitions() []*PartitionId {
	if m != nil {
		return m.Partitions
	}
	return nil
}

// Member configuration
type MemberConfig struct {
	// id is the unique member identifier
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// host is the member host
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// api_port is the port to use for the client API
	APIPort int32 `protobuf:"varint,3,opt,name=api_port,json=apiPort,proto3" json:"apiPort"`
	// protocol_port is the port to use for intra-cluster communication
	ProtocolPort         int32    `protobuf:"varint,4,opt,name=protocol_port,json=protocolPort,proto3" json:"protocolPort"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberConfig) Reset()         { *m = MemberConfig{} }
func (m *MemberConfig) String() string { return proto.CompactTextString(m) }
func (*MemberConfig) ProtoMessage()    {}
func (*MemberConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b4500d882e40b81, []int{5}
}
func (m *MemberConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberConfig.Unmarshal(m, b)
}
func (m *MemberConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberConfig.Marshal(b, m, deterministic)
}
func (m *MemberConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberConfig.Merge(m, src)
}
func (m *MemberConfig) XXX_Size() int {
	return xxx_messageInfo_MemberConfig.Size(m)
}
func (m *MemberConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MemberConfig proto.InternalMessageInfo

func (m *MemberConfig) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *MemberConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *MemberConfig) GetAPIPort() int32 {
	if m != nil {
		return m.APIPort
	}
	return 0
}

func (m *MemberConfig) GetProtocolPort() int32 {
	if m != nil {
		return m.ProtocolPort
	}
	return 0
}

func init() {
	proto.RegisterType((*DatabaseId)(nil), "atomix.controller.DatabaseId")
	proto.RegisterType((*Database)(nil), "atomix.controller.Database")
	proto.RegisterType((*ClusterId)(nil), "atomix.controller.ClusterId")
	proto.RegisterType((*Cluster)(nil), "atomix.controller.Cluster")
	proto.RegisterType((*ClusterConfig)(nil), "atomix.controller.ClusterConfig")
	proto.RegisterType((*MemberConfig)(nil), "atomix.controller.MemberConfig")
}

func init() { proto.RegisterFile("atomix/controller/database.proto", fileDescriptor_6b4500d882e40b81) }

var fileDescriptor_6b4500d882e40b81 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0xce, 0x93, 0x40,
	0x14, 0x0d, 0xb4, 0x7e, 0x94, 0xdb, 0xef, 0x33, 0x3a, 0x31, 0x5f, 0xb0, 0xa9, 0x52, 0x59, 0x75,
	0x05, 0xb1, 0x5a, 0x13, 0x13, 0xd3, 0x44, 0xdb, 0x0d, 0x0b, 0x4d, 0xc3, 0x0b, 0x34, 0x03, 0x43,
	0x71, 0x12, 0x60, 0x08, 0x4c, 0x13, 0x5d, 0xf8, 0x02, 0x3e, 0x87, 0xef, 0xe0, 0xdb, 0xb0, 0xe8,
	0xb2, 0x4f, 0x61, 0x18, 0xfe, 0x2d, 0x6a, 0x5c, 0x71, 0x98, 0x39, 0xe7, 0xdc, 0x7b, 0xee, 0x1d,
	0x58, 0x60, 0xce, 0x22, 0xfa, 0xc5, 0xf2, 0x58, 0xcc, 0x53, 0x16, 0x86, 0x7e, 0x6a, 0x11, 0xcc,
	0xb1, 0x8b, 0x33, 0xdf, 0x4c, 0x52, 0xc6, 0x19, 0x7a, 0x5c, 0x32, 0xcc, 0x96, 0x31, 0x7b, 0x1a,
	0x30, 0x16, 0x84, 0xbe, 0x25, 0x08, 0xee, 0xe9, 0x68, 0xe1, 0xf8, 0x6b, 0xc9, 0x9e, 0x3d, 0x09,
	0x58, 0xc0, 0x04, 0xb4, 0x0a, 0x54, 0x9d, 0xbe, 0xb8, 0xae, 0x92, 0xe0, 0x94, 0x53, 0x4e, 0x59,
	0x5c, 0x52, 0x8c, 0x0d, 0xc0, 0xae, 0x2a, 0x6c, 0x13, 0x84, 0x60, 0x1c, 0xe3, 0xc8, 0xd7, 0xa4,
	0x85, 0xb4, 0x54, 0x1d, 0x81, 0xd1, 0x1c, 0xd4, 0xe2, 0x9b, 0x25, 0xd8, 0xf3, 0x35, 0x59, 0x5c,
	0xb4, 0x07, 0xc6, 0x4f, 0x09, 0x26, 0xb5, 0x01, 0x5a, 0x83, 0x4c, 0x89, 0x10, 0x4f, 0x57, 0xcf,
	0xcc, 0xab, 0x00, 0x66, 0x5b, 0xe9, 0xc3, 0xcd, 0x39, 0xd7, 0x65, 0x7b, 0xe7, 0xc8, 0x94, 0xa0,
	0x37, 0x30, 0xf1, 0xc2, 0x53, 0xc6, 0xfd, 0x34, 0xd3, 0xe4, 0xc5, 0x68, 0x39, 0x5d, 0xcd, 0x06,
	0xc4, 0xdb, 0x92, 0xe2, 0x34, 0x5c, 0xf4, 0x0e, 0xa0, 0x89, 0x93, 0x69, 0x23, 0xa1, 0x9c, 0x0f,
	0x28, 0xf7, 0x35, 0xc9, 0xe9, 0xf0, 0x8d, 0x0c, 0xd4, 0xca, 0xd2, 0x26, 0xe8, 0xbe, 0xe9, 0xfc,
	0x41, 0xaf, 0xb5, 0x4f, 0x30, 0xad, 0xf7, 0x72, 0xa0, 0x44, 0xc4, 0xff, 0x67, 0xb4, 0x87, 0xe7,
	0x5c, 0x6f, 0x87, 0xba, 0x73, 0x80, 0x34, 0x77, 0xc6, 0x37, 0x50, 0xaa, 0xa2, 0xe8, 0x75, 0x67,
	0x58, 0xf3, 0x3f, 0xe7, 0xfd, 0x6d, 0x56, 0xfd, 0xcc, 0xf2, 0x7f, 0x66, 0xfe, 0x2e, 0xc1, 0x5d,
	0xe5, 0xbb, 0x65, 0xf1, 0x91, 0x06, 0xe8, 0x2d, 0x28, 0x91, 0x1f, 0xb9, 0xc5, 0xe8, 0x25, 0x61,
	0xa6, 0x0f, 0x98, 0x7d, 0x14, 0x8c, 0x52, 0xe1, 0xd4, 0x7c, 0xb4, 0x19, 0x68, 0xe5, 0xf9, 0xdf,
	0x5a, 0xb1, 0x49, 0xaf, 0x99, 0x1f, 0x12, 0xdc, 0x76, 0x9d, 0x3b, 0x4b, 0x50, 0x7b, 0x99, 0x11,
	0x8c, 0x3f, 0xb3, 0x8c, 0x57, 0x8f, 0x4f, 0x60, 0xf4, 0x12, 0x26, 0x38, 0xa1, 0x87, 0x84, 0xa5,
	0x5c, 0x1b, 0x89, 0xb5, 0xdd, 0x9f, 0x73, 0x5d, 0x79, 0xbf, 0xb7, 0xf7, 0x2c, 0xe5, 0x97, 0x5c,
	0x57, 0x70, 0x42, 0x0b, 0xe8, 0xd4, 0x00, 0xad, 0xe1, 0x4e, 0xbc, 0x79, 0x8f, 0x85, 0xa5, 0x6e,
	0x2c, 0x74, 0x8f, 0x2e, 0xb9, 0x7e, 0x5b, 0x5f, 0x08, 0x45, 0xef, 0xcf, 0xbd, 0x11, 0x7f, 0xaf,
	0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x76, 0xfe, 0x7f, 0xb3, 0x03, 0x00, 0x00,
}
