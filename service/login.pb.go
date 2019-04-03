// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/login.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	model "github.com/imfuxiao/grpc_learn/model"
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

func init() { proto.RegisterFile("service/login.proto", fileDescriptor_fd8c1bb84b00e5b2) }

var fileDescriptor_fd8c1bb84b00e5b2 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0xcf, 0xc9, 0x4f, 0xcf, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x87, 0x0a, 0x4a, 0x09, 0xe6, 0xe6, 0xa7, 0xa4, 0xe6, 0xe8, 0x83, 0x49, 0x88, 0x9c, 0x91,
	0x2d, 0x17, 0x8f, 0x0f, 0x48, 0x69, 0x30, 0x44, 0x89, 0x90, 0x2e, 0x17, 0x2b, 0x98, 0x2f, 0x24,
	0xac, 0x07, 0x51, 0x06, 0xe6, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48, 0xf1, 0x43, 0x05,
	0x43, 0x8b, 0x53, 0x8b, 0x3c, 0xf3, 0xd2, 0xf2, 0x9d, 0x34, 0xa2, 0xd4, 0xd2, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x33, 0x73, 0xd3, 0x4a, 0x2b, 0x32, 0x13, 0xf3, 0xf5,
	0xd3, 0x8b, 0x0a, 0x92, 0xe3, 0x73, 0x52, 0x13, 0x8b, 0xf2, 0xf4, 0xa1, 0x76, 0x27, 0xb1, 0x81,
	0xed, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xcf, 0xfe, 0x97, 0xa2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginServiceClient interface {
	Login(ctx context.Context, in *model.LoginRequest, opts ...grpc.CallOption) (*model.UserInfo, error)
}

type loginServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoginServiceClient(cc *grpc.ClientConn) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *model.LoginRequest, opts ...grpc.CallOption) (*model.UserInfo, error) {
	out := new(model.UserInfo)
	err := c.cc.Invoke(ctx, "/service.LoginService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
type LoginServiceServer interface {
	Login(context.Context, *model.LoginRequest) (*model.UserInfo, error)
}

// UnimplementedLoginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (*UnimplementedLoginServiceServer) Login(ctx context.Context, req *model.LoginRequest) (*model.UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterLoginServiceServer(s *grpc.Server, srv LoginServiceServer) {
	s.RegisterService(&_LoginService_serviceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.LoginService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*model.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/login.proto",
}
