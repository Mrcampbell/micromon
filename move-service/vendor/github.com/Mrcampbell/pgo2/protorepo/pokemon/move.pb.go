// Code generated by protoc-gen-go. DO NOT EDIT.
// source: move.proto

package pokemon

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("move.proto", fileDescriptor_a434d77a673f848c) }

var fileDescriptor_a434d77a673f848c = []byte{
	// 70 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcd, 0x2f, 0x4b,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0xc8, 0xcf, 0x4e, 0xcd, 0xcd, 0xcf, 0x33,
	0xe2, 0xe5, 0xe2, 0xf6, 0xcd, 0x2f, 0x4b, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x75, 0xe2,
	0x8c, 0x82, 0xc9, 0x24, 0xb1, 0x81, 0x55, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xbe,
	0x83, 0x99, 0x37, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MoveServiceClient is the client API for MoveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoveServiceClient interface {
}

type moveServiceClient struct {
	cc *grpc.ClientConn
}

func NewMoveServiceClient(cc *grpc.ClientConn) MoveServiceClient {
	return &moveServiceClient{cc}
}

// MoveServiceServer is the server API for MoveService service.
type MoveServiceServer interface {
}

// UnimplementedMoveServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMoveServiceServer struct {
}

func RegisterMoveServiceServer(s *grpc.Server, srv MoveServiceServer) {
	s.RegisterService(&_MoveService_serviceDesc, srv)
}

var _MoveService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.MoveService",
	HandlerType: (*MoveServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "move.proto",
}