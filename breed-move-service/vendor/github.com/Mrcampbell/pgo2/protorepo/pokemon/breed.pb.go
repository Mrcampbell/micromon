// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pokemon/breed.proto

package pokemon

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type BreedSummary struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type_1               Type     `protobuf:"varint,3,opt,name=type_1,json=type1,proto3,enum=pokemon.Type" json:"type_1,omitempty"`
	Type_2               Type     `protobuf:"varint,4,opt,name=type_2,json=type2,proto3,enum=pokemon.Type" json:"type_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BreedSummary) Reset()         { *m = BreedSummary{} }
func (m *BreedSummary) String() string { return proto.CompactTextString(m) }
func (*BreedSummary) ProtoMessage()    {}
func (*BreedSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{0}
}

func (m *BreedSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BreedSummary.Unmarshal(m, b)
}
func (m *BreedSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BreedSummary.Marshal(b, m, deterministic)
}
func (m *BreedSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BreedSummary.Merge(m, src)
}
func (m *BreedSummary) XXX_Size() int {
	return xxx_messageInfo_BreedSummary.Size(m)
}
func (m *BreedSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_BreedSummary.DiscardUnknown(m)
}

var xxx_messageInfo_BreedSummary proto.InternalMessageInfo

func (m *BreedSummary) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BreedSummary) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BreedSummary) GetType_1() Type {
	if m != nil {
		return m.Type_1
	}
	return Type_UNKNOWN
}

func (m *BreedSummary) GetType_2() Type {
	if m != nil {
		return m.Type_2
	}
	return Type_UNKNOWN
}

type BreedDetail struct {
	Summary              *BreedSummary `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	Stats                *StatGroup    `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	BreedMoves           []*BreedMove  `protobuf:"bytes,3,rep,name=breed_moves,json=breedMoves,proto3" json:"breed_moves,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BreedDetail) Reset()         { *m = BreedDetail{} }
func (m *BreedDetail) String() string { return proto.CompactTextString(m) }
func (*BreedDetail) ProtoMessage()    {}
func (*BreedDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{1}
}

func (m *BreedDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BreedDetail.Unmarshal(m, b)
}
func (m *BreedDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BreedDetail.Marshal(b, m, deterministic)
}
func (m *BreedDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BreedDetail.Merge(m, src)
}
func (m *BreedDetail) XXX_Size() int {
	return xxx_messageInfo_BreedDetail.Size(m)
}
func (m *BreedDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_BreedDetail.DiscardUnknown(m)
}

var xxx_messageInfo_BreedDetail proto.InternalMessageInfo

func (m *BreedDetail) GetSummary() *BreedSummary {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *BreedDetail) GetStats() *StatGroup {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *BreedDetail) GetBreedMoves() []*BreedMove {
	if m != nil {
		return m.BreedMoves
	}
	return nil
}

type GetBreedSummaryRequest struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VersionGroupId       VersionGroup `protobuf:"varint,2,opt,name=version_group_id,json=versionGroupId,proto3,enum=pokemon.VersionGroup" json:"version_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetBreedSummaryRequest) Reset()         { *m = GetBreedSummaryRequest{} }
func (m *GetBreedSummaryRequest) String() string { return proto.CompactTextString(m) }
func (*GetBreedSummaryRequest) ProtoMessage()    {}
func (*GetBreedSummaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{2}
}

func (m *GetBreedSummaryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreedSummaryRequest.Unmarshal(m, b)
}
func (m *GetBreedSummaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreedSummaryRequest.Marshal(b, m, deterministic)
}
func (m *GetBreedSummaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreedSummaryRequest.Merge(m, src)
}
func (m *GetBreedSummaryRequest) XXX_Size() int {
	return xxx_messageInfo_GetBreedSummaryRequest.Size(m)
}
func (m *GetBreedSummaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreedSummaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreedSummaryRequest proto.InternalMessageInfo

func (m *GetBreedSummaryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetBreedSummaryRequest) GetVersionGroupId() VersionGroup {
	if m != nil {
		return m.VersionGroupId
	}
	return VersionGroup_UNKNOWN_VERSION_GROUP
}

type GetBreedSummaryResponse struct {
	Summary              *BreedSummary `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetBreedSummaryResponse) Reset()         { *m = GetBreedSummaryResponse{} }
func (m *GetBreedSummaryResponse) String() string { return proto.CompactTextString(m) }
func (*GetBreedSummaryResponse) ProtoMessage()    {}
func (*GetBreedSummaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{3}
}

func (m *GetBreedSummaryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreedSummaryResponse.Unmarshal(m, b)
}
func (m *GetBreedSummaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreedSummaryResponse.Marshal(b, m, deterministic)
}
func (m *GetBreedSummaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreedSummaryResponse.Merge(m, src)
}
func (m *GetBreedSummaryResponse) XXX_Size() int {
	return xxx_messageInfo_GetBreedSummaryResponse.Size(m)
}
func (m *GetBreedSummaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreedSummaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreedSummaryResponse proto.InternalMessageInfo

func (m *GetBreedSummaryResponse) GetSummary() *BreedSummary {
	if m != nil {
		return m.Summary
	}
	return nil
}

type GetBreedDetailRequest struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VersionGroupId       VersionGroup `protobuf:"varint,2,opt,name=version_group_id,json=versionGroupId,proto3,enum=pokemon.VersionGroup" json:"version_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetBreedDetailRequest) Reset()         { *m = GetBreedDetailRequest{} }
func (m *GetBreedDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetBreedDetailRequest) ProtoMessage()    {}
func (*GetBreedDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{4}
}

func (m *GetBreedDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreedDetailRequest.Unmarshal(m, b)
}
func (m *GetBreedDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreedDetailRequest.Marshal(b, m, deterministic)
}
func (m *GetBreedDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreedDetailRequest.Merge(m, src)
}
func (m *GetBreedDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetBreedDetailRequest.Size(m)
}
func (m *GetBreedDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreedDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreedDetailRequest proto.InternalMessageInfo

func (m *GetBreedDetailRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetBreedDetailRequest) GetVersionGroupId() VersionGroup {
	if m != nil {
		return m.VersionGroupId
	}
	return VersionGroup_UNKNOWN_VERSION_GROUP
}

type GetBreedDetailResponse struct {
	Detail               *BreedDetail `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetBreedDetailResponse) Reset()         { *m = GetBreedDetailResponse{} }
func (m *GetBreedDetailResponse) String() string { return proto.CompactTextString(m) }
func (*GetBreedDetailResponse) ProtoMessage()    {}
func (*GetBreedDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a06e1c6efe7453, []int{5}
}

func (m *GetBreedDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreedDetailResponse.Unmarshal(m, b)
}
func (m *GetBreedDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreedDetailResponse.Marshal(b, m, deterministic)
}
func (m *GetBreedDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreedDetailResponse.Merge(m, src)
}
func (m *GetBreedDetailResponse) XXX_Size() int {
	return xxx_messageInfo_GetBreedDetailResponse.Size(m)
}
func (m *GetBreedDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreedDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreedDetailResponse proto.InternalMessageInfo

func (m *GetBreedDetailResponse) GetDetail() *BreedDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func init() {
	proto.RegisterType((*BreedSummary)(nil), "pokemon.BreedSummary")
	proto.RegisterType((*BreedDetail)(nil), "pokemon.BreedDetail")
	proto.RegisterType((*GetBreedSummaryRequest)(nil), "pokemon.GetBreedSummaryRequest")
	proto.RegisterType((*GetBreedSummaryResponse)(nil), "pokemon.GetBreedSummaryResponse")
	proto.RegisterType((*GetBreedDetailRequest)(nil), "pokemon.GetBreedDetailRequest")
	proto.RegisterType((*GetBreedDetailResponse)(nil), "pokemon.GetBreedDetailResponse")
}

func init() { proto.RegisterFile("pokemon/breed.proto", fileDescriptor_b9a06e1c6efe7453) }

var fileDescriptor_b9a06e1c6efe7453 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0xa4, 0x4d, 0xd5, 0x09, 0x98, 0x6a, 0xa0, 0xa9, 0x65, 0xa1, 0x12, 0x59, 0x1c,
	0x72, 0x40, 0xb1, 0xe2, 0x3e, 0x00, 0x52, 0x85, 0xa8, 0x40, 0xe2, 0xe2, 0x22, 0x0e, 0x5c, 0xa2,
	0x0d, 0x1e, 0x85, 0x15, 0xb5, 0xd7, 0x78, 0x37, 0x96, 0x22, 0x04, 0x07, 0x5e, 0x81, 0x33, 0x4f,
	0xc5, 0x2b, 0x70, 0xe7, 0x15, 0x90, 0xf7, 0x8f, 0xb1, 0x1b, 0x72, 0xe0, 0xc0, 0xcd, 0x99, 0xf9,
	0x66, 0xbe, 0xdf, 0x7e, 0xce, 0x1a, 0xee, 0x97, 0xe2, 0x03, 0xe5, 0xa2, 0x88, 0x57, 0x15, 0x51,
	0x36, 0x2f, 0x2b, 0xa1, 0x04, 0x1e, 0xd9, 0x62, 0xf8, 0x70, 0x2d, 0xc4, 0xfa, 0x86, 0x62, 0x56,
	0xf2, 0x98, 0x15, 0x85, 0x50, 0x4c, 0x71, 0x51, 0x48, 0x23, 0x0b, 0xd1, 0xcd, 0xaa, 0x6d, 0x49,
	0xb6, 0x76, 0xe6, 0x6a, 0x52, 0x31, 0xb5, 0xae, 0xc4, 0xa6, 0xbc, 0xdd, 0xd0, 0x46, 0xb9, 0xa8,
	0xdd, 0x44, 0xe8, 0x1a, 0x35, 0x55, 0x92, 0x8b, 0xa2, 0x33, 0x14, 0x7d, 0x81, 0x3b, 0x97, 0x8d,
	0xfc, 0x7a, 0x93, 0xe7, 0xac, 0xda, 0xa2, 0x0f, 0x03, 0x9e, 0x05, 0xde, 0xd4, 0x9b, 0x1d, 0xa7,
	0x03, 0x9e, 0x21, 0xc2, 0x41, 0xc1, 0x72, 0x0a, 0x06, 0xba, 0xa2, 0x9f, 0xf1, 0x31, 0x8c, 0x1a,
	0x9e, 0xe5, 0x22, 0x18, 0x4e, 0xbd, 0x99, 0x9f, 0xdc, 0x9d, 0x5b, 0x83, 0xf9, 0xeb, 0x6d, 0x49,
	0xe9, 0x61, 0xd3, 0x5c, 0xb4, 0xaa, 0x24, 0x38, 0xd8, 0xab, 0x4a, 0xa2, 0xef, 0x1e, 0x8c, 0x35,
	0xc0, 0x33, 0x52, 0x8c, 0xdf, 0x60, 0x0c, 0x47, 0xd2, 0xa0, 0x68, 0x88, 0x71, 0x72, 0xda, 0x8e,
	0x75, 0x39, 0x53, 0xa7, 0xc2, 0x19, 0x1c, 0x36, 0x41, 0x48, 0x4d, 0x38, 0x4e, 0xb0, 0x95, 0x5f,
	0x2b, 0xa6, 0xae, 0x9a, 0x93, 0xa6, 0x46, 0x80, 0x17, 0x30, 0xd6, 0xc9, 0x2c, 0x9b, 0x68, 0x64,
	0x30, 0x9c, 0x0e, 0x7b, 0x7a, 0xbd, 0xfe, 0x95, 0xa8, 0x29, 0x85, 0x95, 0x7b, 0x94, 0x11, 0x87,
	0xc9, 0x15, 0xa9, 0x9e, 0x35, 0x7d, 0xdc, 0x90, 0x54, 0x3b, 0x49, 0x3d, 0x85, 0x13, 0x9b, 0xef,
	0x52, 0x07, 0xbc, 0xe4, 0x99, 0x66, 0xf2, 0x3b, 0x47, 0x78, 0x63, 0x04, 0x06, 0xcb, 0xaf, 0x3b,
	0xbf, 0x5e, 0x64, 0xd1, 0x4b, 0x38, 0xdb, 0xb1, 0x92, 0xa5, 0x28, 0x24, 0xfd, 0x73, 0x2a, 0xd1,
	0x7b, 0x38, 0x75, 0xbb, 0x4c, 0xb0, 0xff, 0x8d, 0xfa, 0xf9, 0x9f, 0x80, 0x9c, 0x93, 0x85, 0x7e,
	0x02, 0xa3, 0x4c, 0x57, 0x2c, 0xf3, 0x83, 0x3e, 0xb3, 0x55, 0x5b, 0x4d, 0xf2, 0xcb, 0x73, 0xff,
	0x44, 0xaa, 0x6a, 0xfe, 0x8e, 0x90, 0x83, 0xdf, 0x5f, 0x8c, 0xe7, 0xed, 0x82, 0xbf, 0x9e, 0x2d,
	0x7c, 0xb4, 0xb7, 0x6f, 0x88, 0xa2, 0xc9, 0xd7, 0x1f, 0x3f, 0xbf, 0x0d, 0x4e, 0xd0, 0x8f, 0xeb,
	0x85, 0xb9, 0x25, 0xf1, 0x27, 0x9e, 0x7d, 0x46, 0x05, 0xf7, 0x6e, 0x25, 0x8f, 0xbb, 0xbb, 0xfa,
	0xaf, 0x3f, 0x9c, 0xee, 0x17, 0x58, 0xb7, 0x73, 0xed, 0x16, 0xe0, 0xa4, 0xef, 0x16, 0xdb, 0x77,
	0x74, 0x79, 0xfc, 0xd6, 0x7d, 0x05, 0x56, 0x23, 0x7d, 0x19, 0x2f, 0x7e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xa1, 0x03, 0xb3, 0xa4, 0x2c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BreedServiceClient is the client API for BreedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BreedServiceClient interface {
	GetBreedDetail(ctx context.Context, in *GetBreedDetailRequest, opts ...grpc.CallOption) (*GetBreedDetailResponse, error)
	GetBreedSummary(ctx context.Context, in *GetBreedSummaryRequest, opts ...grpc.CallOption) (*GetBreedSummaryResponse, error)
}

type breedServiceClient struct {
	cc *grpc.ClientConn
}

func NewBreedServiceClient(cc *grpc.ClientConn) BreedServiceClient {
	return &breedServiceClient{cc}
}

func (c *breedServiceClient) GetBreedDetail(ctx context.Context, in *GetBreedDetailRequest, opts ...grpc.CallOption) (*GetBreedDetailResponse, error) {
	out := new(GetBreedDetailResponse)
	err := c.cc.Invoke(ctx, "/pokemon.BreedService/GetBreedDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *breedServiceClient) GetBreedSummary(ctx context.Context, in *GetBreedSummaryRequest, opts ...grpc.CallOption) (*GetBreedSummaryResponse, error) {
	out := new(GetBreedSummaryResponse)
	err := c.cc.Invoke(ctx, "/pokemon.BreedService/GetBreedSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BreedServiceServer is the server API for BreedService service.
type BreedServiceServer interface {
	GetBreedDetail(context.Context, *GetBreedDetailRequest) (*GetBreedDetailResponse, error)
	GetBreedSummary(context.Context, *GetBreedSummaryRequest) (*GetBreedSummaryResponse, error)
}

// UnimplementedBreedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBreedServiceServer struct {
}

func (*UnimplementedBreedServiceServer) GetBreedDetail(ctx context.Context, req *GetBreedDetailRequest) (*GetBreedDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBreedDetail not implemented")
}
func (*UnimplementedBreedServiceServer) GetBreedSummary(ctx context.Context, req *GetBreedSummaryRequest) (*GetBreedSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBreedSummary not implemented")
}

func RegisterBreedServiceServer(s *grpc.Server, srv BreedServiceServer) {
	s.RegisterService(&_BreedService_serviceDesc, srv)
}

func _BreedService_GetBreedDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBreedDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreedServiceServer).GetBreedDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.BreedService/GetBreedDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreedServiceServer).GetBreedDetail(ctx, req.(*GetBreedDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BreedService_GetBreedSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBreedSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreedServiceServer).GetBreedSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.BreedService/GetBreedSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreedServiceServer).GetBreedSummary(ctx, req.(*GetBreedSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BreedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.BreedService",
	HandlerType: (*BreedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBreedDetail",
			Handler:    _BreedService_GetBreedDetail_Handler,
		},
		{
			MethodName: "GetBreedSummary",
			Handler:    _BreedService_GetBreedSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon/breed.proto",
}
