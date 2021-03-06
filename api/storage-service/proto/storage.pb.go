// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/storage-service/proto/storage.proto

package storage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StatusCode int32

const (
	StatusCode_Unknown StatusCode = 0
	StatusCode_Ok      StatusCode = 1
	StatusCode_Failed  StatusCode = 2
)

var StatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}
var StatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_storage_0ff114ee40ef9e85, []int{0}
}

type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	TotalSize            int64    `protobuf:"varint,2,opt,name=TotalSize,proto3" json:"TotalSize,omitempty"`
	Received             int64    `protobuf:"varint,3,opt,name=Received,proto3" json:"Received,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_0ff114ee40ef9e85, []int{0}
}
func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (dst *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(dst, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Chunk) GetTotalSize() int64 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *Chunk) GetReceived() int64 {
	if m != nil {
		return m.Received
	}
	return 0
}

type UploadStatus struct {
	Message              string     `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	File                 string     `protobuf:"bytes,2,opt,name=File,proto3" json:"File,omitempty"`
	Size                 int64      `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	Code                 StatusCode `protobuf:"varint,4,opt,name=Code,proto3,enum=storage.StatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_0ff114ee40ef9e85, []int{1}
}
func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (dst *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(dst, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *UploadStatus) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *UploadStatus) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_Unknown
}

func init() {
	proto.RegisterType((*Chunk)(nil), "storage.Chunk")
	proto.RegisterType((*UploadStatus)(nil), "storage.UploadStatus")
	proto.RegisterEnum("storage.StatusCode", StatusCode_name, StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (StorageService_UploadClient, error)
}

type storageServiceClient struct {
	cc *grpc.ClientConn
}

func NewStorageServiceClient(cc *grpc.ClientConn) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (StorageService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StorageService_serviceDesc.Streams[0], "/storage.StorageService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageServiceUploadClient{stream}
	return x, nil
}

type StorageService_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type storageServiceUploadClient struct {
	grpc.ClientStream
}

func (x *storageServiceUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageServiceUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServiceServer is the server API for StorageService service.
type StorageServiceServer interface {
	Upload(StorageService_UploadServer) error
}

func RegisterStorageServiceServer(s *grpc.Server, srv StorageServiceServer) {
	s.RegisterService(&_StorageService_serviceDesc, srv)
}

func _StorageService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServiceServer).Upload(&storageServiceUploadServer{stream})
}

type StorageService_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type storageServiceUploadServer struct {
	grpc.ServerStream
}

func (x *storageServiceUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageServiceUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StorageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _StorageService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/storage-service/proto/storage.proto",
}

func init() {
	proto.RegisterFile("api/storage-service/proto/storage.proto", fileDescriptor_storage_0ff114ee40ef9e85)
}

var fileDescriptor_storage_0ff114ee40ef9e85 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xbb, 0x6d, 0x4d, 0xcd, 0x58, 0x42, 0x18, 0x11, 0x42, 0xf1, 0x10, 0x72, 0x69, 0x10,
	0xda, 0x42, 0xfb, 0x13, 0x82, 0xbd, 0x89, 0xb0, 0xb1, 0x27, 0x4f, 0x6b, 0x33, 0xd4, 0x25, 0x61,
	0x37, 0x24, 0xdb, 0x8a, 0xfe, 0x7a, 0xc9, 0xae, 0x49, 0xbd, 0xcd, 0xf7, 0x66, 0x79, 0x6f, 0xe7,
	0xc1, 0x52, 0xd4, 0x72, 0xd3, 0x1a, 0xdd, 0x88, 0x13, 0xad, 0x5a, 0x6a, 0x2e, 0xf2, 0x48, 0x9b,
	0xba, 0xd1, 0x46, 0xf7, 0xea, 0xda, 0x12, 0xce, 0xfe, 0x30, 0x79, 0x87, 0x9b, 0xec, 0xf3, 0xac,
	0x4a, 0x8c, 0x60, 0x96, 0x69, 0x65, 0x48, 0x99, 0x88, 0xc5, 0x2c, 0x9d, 0xf3, 0x1e, 0xf1, 0x11,
	0xfc, 0x37, 0x6d, 0x44, 0x95, 0xcb, 0x1f, 0x8a, 0xc6, 0x31, 0x4b, 0x27, 0xfc, 0x2a, 0xe0, 0x02,
	0x6e, 0x39, 0x1d, 0x49, 0x5e, 0xa8, 0x88, 0x26, 0x76, 0x39, 0x70, 0xf2, 0x0d, 0xf3, 0x43, 0x5d,
	0x69, 0x51, 0xe4, 0x46, 0x98, 0x73, 0xdb, 0x65, 0xbc, 0x50, 0xdb, 0x8a, 0x13, 0xd9, 0x0c, 0x9f,
	0xf7, 0x88, 0x08, 0xd3, 0xbd, 0xac, 0x9c, 0xbd, 0xcf, 0xed, 0xdc, 0x69, 0x36, 0xd2, 0xb9, 0xda,
	0x19, 0x97, 0x30, 0xcd, 0x74, 0x41, 0xd1, 0x34, 0x66, 0x69, 0xb0, 0xbd, 0x5f, 0xf7, 0x57, 0xb9,
	0x80, 0x6e, 0xc5, 0xed, 0x83, 0xa7, 0x15, 0xc0, 0x55, 0xc3, 0x3b, 0x98, 0x1d, 0x54, 0xa9, 0xf4,
	0x97, 0x0a, 0x47, 0xe8, 0xc1, 0xf8, 0xb5, 0x0c, 0x19, 0x02, 0x78, 0x7b, 0x21, 0x2b, 0x2a, 0xc2,
	0xf1, 0xf6, 0x19, 0x82, 0xdc, 0x59, 0xe5, 0xae, 0x35, 0xdc, 0x81, 0xe7, 0xfe, 0x8e, 0xc1, 0x90,
	0x62, 0x9b, 0x5a, 0x3c, 0x0c, 0xfc, 0xff, 0xb8, 0x64, 0x94, 0xb2, 0x0f, 0xcf, 0xb6, 0xbb, 0xfb,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xc1, 0x20, 0xec, 0x88, 0x01, 0x00, 0x00,
}
