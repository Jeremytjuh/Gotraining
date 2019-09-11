// Code generated by protoc-gen-go. DO NOT EDIT.
// source: LoginTest.proto

package LoginTest

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d18a4d45edf1a1e, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// The response message containing the greetings
type LoginReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d18a4d45edf1a1e, []int{1}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "LoginTest.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "LoginTest.LoginReply")
}

func init() { proto.RegisterFile("LoginTest.proto", fileDescriptor_6d18a4d45edf1a1e) }

var fileDescriptor_6d18a4d45edf1a1e = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xdf, 0xca, 0xd3, 0x30,
	0x18, 0xc6, 0xe9, 0xc0, 0x3f, 0x5f, 0x10, 0x94, 0xc2, 0x70, 0x54, 0xc1, 0xd7, 0x1e, 0x88, 0x88,
	0x6d, 0x50, 0xcf, 0x76, 0x36, 0x61, 0x3b, 0x12, 0xc4, 0xb9, 0x1b, 0x48, 0xd3, 0x77, 0x49, 0xa0,
	0xcd, 0x1b, 0x93, 0x74, 0xb5, 0xa7, 0x5e, 0x82, 0x5e, 0x80, 0x17, 0xe2, 0x81, 0x17, 0xe1, 0x2d,
	0x78, 0x21, 0xb2, 0x6e, 0x9d, 0x1f, 0xec, 0x28, 0xef, 0x2f, 0xcf, 0xc3, 0x93, 0xe4, 0x09, 0x7b,
	0xf8, 0x81, 0x94, 0xb1, 0x3b, 0x0c, 0xb1, 0x74, 0x9e, 0x22, 0xa5, 0x37, 0x97, 0x8d, 0xec, 0xa9,
	0x22, 0x52, 0x0d, 0x72, 0xe1, 0x0c, 0x17, 0xd6, 0x52, 0x14, 0xd1, 0x90, 0x0d, 0x27, 0x63, 0xf6,
	0x7a, 0x5c, 0x64, 0xa1, 0xd0, 0x16, 0xa1, 0x17, 0x4a, 0xa1, 0xe7, 0xe4, 0x46, 0xc7, 0xb5, 0x3b,
	0xdf, 0xb0, 0x07, 0x63, 0xf0, 0x16, 0xbf, 0x74, 0x18, 0x62, 0x9a, 0xb1, 0xfb, 0x5d, 0x40, 0x6f,
	0x45, 0x8b, 0x8b, 0x04, 0x92, 0x97, 0x37, 0xdb, 0x0b, 0x1f, 0x35, 0x27, 0x42, 0xe8, 0xc9, 0xd7,
	0x8b, 0xd9, 0x49, 0x9b, 0x38, 0x7f, 0xc1, 0xd8, 0x39, 0xc7, 0x35, 0x43, 0xba, 0x60, 0xf7, 0x5a,
	0x0c, 0x41, 0xa8, 0x29, 0x64, 0xc2, 0xb7, 0x3f, 0x67, 0xec, 0xce, 0x68, 0x4c, 0x77, 0xd3, 0xf0,
	0xb8, 0xfc, 0xff, 0xd6, 0xdb, 0x77, 0xc9, 0xe6, 0xd7, 0x82, 0x6b, 0x86, 0xfc, 0xc9, 0xb7, 0x3f,
	0x7f, 0x7f, 0xcc, 0xe6, 0xf9, 0x23, 0x7e, 0x78, 0xc3, 0xf1, 0xab, 0x68, 0x5d, 0x83, 0x1c, 0xa5,
	0xa6, 0x65, 0xf2, 0x2a, 0xfb, 0x9d, 0x7c, 0x5f, 0xfd, 0x4a, 0xd2, 0xfd, 0xea, 0xbd, 0x89, 0x1f,
	0xf7, 0xeb, 0x03, 0xfa, 0x21, 0x6a, 0x63, 0xd5, 0x67, 0xf4, 0x07, 0x23, 0x11, 0x6a, 0x0c, 0xd2,
	0x9b, 0xb1, 0x0e, 0x28, 0x0a, 0xe8, 0xb5, 0x91, 0x1a, 0x82, 0xa6, 0xae, 0xa9, 0xc1, 0x52, 0x84,
	0x0a, 0xa1, 0x0b, 0x58, 0x83, 0xb1, 0xe0, 0x1a, 0x21, 0x11, 0x68, 0x0f, 0x51, 0x23, 0xd4, 0x24,
	0xbb, 0x16, 0xed, 0xa9, 0x3c, 0x90, 0xd4, 0x1e, 0xe1, 0x79, 0xf6, 0x89, 0x3d, 0xdb, 0x18, 0x5b,
	0x03, 0x75, 0x11, 0x5a, 0xf2, 0x08, 0xa2, 0x3a, 0x8e, 0x6b, 0xa9, 0xe9, 0x7c, 0x62, 0x5a, 0xea,
	0x18, 0x5d, 0x58, 0x72, 0xae, 0x4c, 0xd4, 0x5d, 0x55, 0x4a, 0x6a, 0xb9, 0xf2, 0x4e, 0x16, 0x28,
	0x29, 0x0c, 0x21, 0xe2, 0x19, 0x95, 0x88, 0xd8, 0x8b, 0xa1, 0xba, 0x3b, 0x7e, 0xcc, 0xbb, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x12, 0xf4, 0x63, 0xc1, 0x02, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginClient interface {
	// Sends a greeting
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type loginClient struct {
	cc *grpc.ClientConn
}

func NewLoginClient(cc *grpc.ClientConn) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/LoginTest.Login/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
type LoginServer interface {
	// Sends a greeting
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

// UnimplementedLoginServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (*UnimplementedLoginServer) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LoginTest.Login/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LoginTest.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Login_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "LoginTest.proto",
}
