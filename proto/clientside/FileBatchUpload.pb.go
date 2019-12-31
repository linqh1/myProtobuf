// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/clientside/FileBatchUpload.proto

package clientside

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the user's name.
type FileInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_98a64c3d74a8b42b, []int{0}
}

func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfo.Unmarshal(m, b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
}
func (m *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(m, src)
}
func (m *FileInfo) XXX_Size() int {
	return xxx_messageInfo_FileInfo.Size(m)
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type FileUploadResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileUploadResponse) Reset()         { *m = FileUploadResponse{} }
func (m *FileUploadResponse) String() string { return proto.CompactTextString(m) }
func (*FileUploadResponse) ProtoMessage()    {}
func (*FileUploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98a64c3d74a8b42b, []int{1}
}

func (m *FileUploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileUploadResponse.Unmarshal(m, b)
}
func (m *FileUploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileUploadResponse.Marshal(b, m, deterministic)
}
func (m *FileUploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileUploadResponse.Merge(m, src)
}
func (m *FileUploadResponse) XXX_Size() int {
	return xxx_messageInfo_FileUploadResponse.Size(m)
}
func (m *FileUploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileUploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileUploadResponse proto.InternalMessageInfo

func (m *FileUploadResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*FileInfo)(nil), "clientside.FileInfo")
	proto.RegisterType((*FileUploadResponse)(nil), "clientside.FileUploadResponse")
}

func init() {
	proto.RegisterFile("proto/clientside/FileBatchUpload.proto", fileDescriptor_98a64c3d74a8b42b)
}

var fileDescriptor_98a64c3d74a8b42b = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x29, 0xce, 0x4c, 0x49, 0xd5, 0x77, 0xcb, 0xcc,
	0x49, 0x75, 0x4a, 0x2c, 0x49, 0xce, 0x08, 0x2d, 0xc8, 0xc9, 0x4f, 0x4c, 0xd1, 0x03, 0x2b, 0x10,
	0xe2, 0x42, 0xa8, 0x50, 0x92, 0xe3, 0xe2, 0x00, 0x29, 0xf2, 0xcc, 0x4b, 0xcb, 0x17, 0x12, 0xe2,
	0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xf4,
	0xb8, 0x84, 0x40, 0xf2, 0x10, 0xfd, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12,
	0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0xc5, 0x30, 0xae, 0x51, 0x28, 0x17, 0x3f,
	0x9a, 0xa5, 0x42, 0x4e, 0x5c, 0x6c, 0x50, 0x96, 0x88, 0x1e, 0xc2, 0x66, 0x3d, 0x98, 0xb5, 0x52,
	0x72, 0xe8, 0xa2, 0xa8, 0x96, 0x29, 0x31, 0x68, 0x30, 0x26, 0xb1, 0x81, 0x5d, 0x6e, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x18, 0xdc, 0x49, 0x3e, 0xe3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileBatchUploadClient is the client API for FileBatchUpload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileBatchUploadClient interface {
	// Sends a greeting
	Upload(ctx context.Context, opts ...grpc.CallOption) (FileBatchUpload_UploadClient, error)
}

type fileBatchUploadClient struct {
	cc *grpc.ClientConn
}

func NewFileBatchUploadClient(cc *grpc.ClientConn) FileBatchUploadClient {
	return &fileBatchUploadClient{cc}
}

func (c *fileBatchUploadClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FileBatchUpload_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileBatchUpload_serviceDesc.Streams[0], "/clientside.FileBatchUpload/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileBatchUploadUploadClient{stream}
	return x, nil
}

type FileBatchUpload_UploadClient interface {
	Send(*FileInfo) error
	CloseAndRecv() (*FileUploadResponse, error)
	grpc.ClientStream
}

type fileBatchUploadUploadClient struct {
	grpc.ClientStream
}

func (x *fileBatchUploadUploadClient) Send(m *FileInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileBatchUploadUploadClient) CloseAndRecv() (*FileUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileBatchUploadServer is the server API for FileBatchUpload service.
type FileBatchUploadServer interface {
	// Sends a greeting
	Upload(FileBatchUpload_UploadServer) error
}

// UnimplementedFileBatchUploadServer can be embedded to have forward compatible implementations.
type UnimplementedFileBatchUploadServer struct {
}

func (*UnimplementedFileBatchUploadServer) Upload(srv FileBatchUpload_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterFileBatchUploadServer(s *grpc.Server, srv FileBatchUploadServer) {
	s.RegisterService(&_FileBatchUpload_serviceDesc, srv)
}

func _FileBatchUpload_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileBatchUploadServer).Upload(&fileBatchUploadUploadServer{stream})
}

type FileBatchUpload_UploadServer interface {
	SendAndClose(*FileUploadResponse) error
	Recv() (*FileInfo, error)
	grpc.ServerStream
}

type fileBatchUploadUploadServer struct {
	grpc.ServerStream
}

func (x *fileBatchUploadUploadServer) SendAndClose(m *FileUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileBatchUploadUploadServer) Recv() (*FileInfo, error) {
	m := new(FileInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileBatchUpload_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clientside.FileBatchUpload",
	HandlerType: (*FileBatchUploadServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _FileBatchUpload_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/clientside/FileBatchUpload.proto",
}
