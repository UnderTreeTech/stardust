// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stardust.proto

package stardust

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IdReq struct {
	BizType              int64    `protobuf:"varint,1,opt,name=biz_type,json=bizType,proto3" json:"biz_type" form:"biz_type" validate:"required,gte=0,lt=32"`
	Len                  int64    `protobuf:"varint,2,opt,name=len,proto3" json:"len" form:"len" validate:"required,gte=0,lte=1000"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdReq) Reset()         { *m = IdReq{} }
func (m *IdReq) String() string { return proto.CompactTextString(m) }
func (*IdReq) ProtoMessage()    {}
func (*IdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d205f85c3caec6, []int{0}
}
func (m *IdReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdReq.Merge(m, src)
}
func (m *IdReq) XXX_Size() int {
	return m.Size()
}
func (m *IdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IdReq.DiscardUnknown(m)
}

var xxx_messageInfo_IdReq proto.InternalMessageInfo

type IdsReply struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdsReply) Reset()         { *m = IdsReply{} }
func (m *IdsReply) String() string { return proto.CompactTextString(m) }
func (*IdsReply) ProtoMessage()    {}
func (*IdsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d205f85c3caec6, []int{1}
}
func (m *IdsReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdsReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdsReply.Merge(m, src)
}
func (m *IdsReply) XXX_Size() int {
	return m.Size()
}
func (m *IdsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IdsReply.DiscardUnknown(m)
}

var xxx_messageInfo_IdsReply proto.InternalMessageInfo

type IdReply struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdReply) Reset()         { *m = IdReply{} }
func (m *IdReply) String() string { return proto.CompactTextString(m) }
func (*IdReply) ProtoMessage()    {}
func (*IdReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d205f85c3caec6, []int{2}
}
func (m *IdReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdReply.Merge(m, src)
}
func (m *IdReply) XXX_Size() int {
	return m.Size()
}
func (m *IdReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IdReply.DiscardUnknown(m)
}

var xxx_messageInfo_IdReply proto.InternalMessageInfo

type ParseReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id" form:"id" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseReq) Reset()         { *m = ParseReq{} }
func (m *ParseReq) String() string { return proto.CompactTextString(m) }
func (*ParseReq) ProtoMessage()    {}
func (*ParseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d205f85c3caec6, []int{3}
}
func (m *ParseReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParseReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseReq.Merge(m, src)
}
func (m *ParseReq) XXX_Size() int {
	return m.Size()
}
func (m *ParseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseReq.DiscardUnknown(m)
}

var xxx_messageInfo_ParseReq proto.InternalMessageInfo

type ParseReply struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time"`
	Idc                  int64    `protobuf:"varint,2,opt,name=idc,proto3" json:"idc"`
	Worker               int64    `protobuf:"varint,3,opt,name=worker,proto3" json:"worker"`
	BizType              int64    `protobuf:"varint,4,opt,name=biz_type,json=bizType,proto3" json:"biz_type"`
	Sequence             int64    `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseReply) Reset()         { *m = ParseReply{} }
func (m *ParseReply) String() string { return proto.CompactTextString(m) }
func (*ParseReply) ProtoMessage()    {}
func (*ParseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d205f85c3caec6, []int{4}
}
func (m *ParseReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParseReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseReply.Merge(m, src)
}
func (m *ParseReply) XXX_Size() int {
	return m.Size()
}
func (m *ParseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseReply.DiscardUnknown(m)
}

var xxx_messageInfo_ParseReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IdReq)(nil), "service.stardust.v1.IdReq")
	proto.RegisterType((*IdsReply)(nil), "service.stardust.v1.IdsReply")
	proto.RegisterType((*IdReply)(nil), "service.stardust.v1.IdReply")
	proto.RegisterType((*ParseReq)(nil), "service.stardust.v1.ParseReq")
	proto.RegisterType((*ParseReply)(nil), "service.stardust.v1.ParseReply")
}

func init() { proto.RegisterFile("stardust.proto", fileDescriptor_71d205f85c3caec6) }

var fileDescriptor_71d205f85c3caec6 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0xc7, 0xeb, 0xb8, 0x4d, 0xac, 0xa1, 0x62, 0x61, 0x24, 0x64, 0xa2, 0xd6, 0x0e, 0x23, 0xa1,
	0x66, 0xd1, 0x3a, 0x69, 0xbb, 0x00, 0x55, 0x8a, 0x90, 0x02, 0x52, 0x65, 0x56, 0xc8, 0xc0, 0xa6,
	0x1b, 0x64, 0x7b, 0x5e, 0xc3, 0x08, 0x27, 0x76, 0x66, 0xc6, 0x41, 0xe9, 0x49, 0xb8, 0x05, 0x17,
	0xe0, 0x00, 0x5d, 0x72, 0x02, 0x0b, 0xc2, 0x0a, 0x2f, 0x73, 0x02, 0x34, 0x63, 0x3b, 0xcd, 0x22,
	0xc9, 0x66, 0x3c, 0xef, 0xbd, 0xbf, 0x7f, 0xf3, 0x3e, 0x66, 0xd0, 0x63, 0x2e, 0x02, 0x46, 0x32,
	0x2e, 0xdc, 0x94, 0x25, 0x22, 0x31, 0x9f, 0x70, 0x60, 0x33, 0x1a, 0x81, 0xbb, 0xf2, 0xcf, 0xce,
	0xdb, 0x67, 0x23, 0x2a, 0xbe, 0x64, 0xa1, 0x1b, 0x25, 0xe3, 0xde, 0x28, 0x19, 0x25, 0x3d, 0xa5,
	0x0d, 0xb3, 0x5b, 0x65, 0x29, 0x43, 0xed, 0x4a, 0x06, 0xfe, 0xa1, 0xa1, 0x03, 0x8f, 0xf8, 0x30,
	0x35, 0x6f, 0x90, 0x11, 0xd2, 0xbb, 0xcf, 0x62, 0x9e, 0x82, 0xa5, 0x75, 0xb4, 0xae, 0x3e, 0x7c,
	0x5d, 0xe4, 0xce, 0xca, 0xb7, 0xcc, 0x9d, 0xde, 0x6d, 0xc2, 0xc6, 0x57, 0xb8, 0xf6, 0xe0, 0xce,
	0x2c, 0x88, 0x29, 0x09, 0x04, 0x5c, 0x61, 0x06, 0xd3, 0x8c, 0x32, 0x20, 0xa7, 0x23, 0x01, 0x83,
	0xfe, 0x69, 0x2c, 0x06, 0x97, 0x17, 0xd8, 0x6f, 0x85, 0xf4, 0xee, 0xe3, 0x3c, 0x05, 0xf3, 0x1d,
	0xd2, 0x63, 0x98, 0x58, 0x0d, 0x85, 0x7d, 0x55, 0xe4, 0x8e, 0x34, 0x97, 0xb9, 0x73, 0x56, 0x12,
	0x63, 0x98, 0xec, 0x84, 0xc1, 0xe0, 0xbc, 0xdf, 0xef, 0x63, 0x5f, 0xfe, 0x85, 0x5f, 0x20, 0xc3,
	0x23, 0xdc, 0x87, 0x34, 0x9e, 0x9b, 0xcf, 0x90, 0x4e, 0x09, 0xb7, 0xb4, 0x8e, 0xde, 0xd5, 0x87,
	0x2d, 0xc9, 0xa5, 0x84, 0xfb, 0x72, 0xc1, 0xcf, 0x51, 0x4b, 0xd6, 0x25, 0x55, 0x4f, 0x51, 0x83,
	0x92, 0xaa, 0xa6, 0x66, 0x91, 0x3b, 0x0d, 0x4a, 0xfc, 0x06, 0x25, 0xf8, 0x0d, 0x32, 0xde, 0x07,
	0x8c, 0x83, 0xac, 0xfe, 0xe5, 0x9a, 0xe6, 0xa4, 0xd4, 0x2c, 0x73, 0xe7, 0xb8, 0xcc, 0x8f, 0x92,
	0x4d, 0xe9, 0x61, 0x05, 0xf9, 0xa9, 0x21, 0x54, 0x51, 0xe4, 0x59, 0x47, 0x68, 0x5f, 0xd0, 0x71,
	0xdd, 0x41, 0xa3, 0xc8, 0x1d, 0x65, 0xfb, 0x6a, 0x2d, 0xf3, 0x8d, 0xaa, 0x3e, 0x54, 0xf9, 0x46,
	0x32, 0xdf, 0xc8, 0xc4, 0xa8, 0xf9, 0x2d, 0x61, 0x5f, 0x81, 0x59, 0xba, 0x8a, 0xa2, 0x22, 0x77,
	0x2a, 0x8f, 0x5f, 0x7d, 0xcd, 0x93, 0xb5, 0x11, 0xed, 0x2b, 0xd5, 0xe1, 0xfa, 0x88, 0x1e, 0xfa,
	0xdd, 0x45, 0x06, 0x87, 0x69, 0x06, 0x93, 0x08, 0xac, 0x83, 0x07, 0x61, 0xed, 0xf3, 0x57, 0xbb,
	0x8b, 0x7f, 0x1a, 0x32, 0x3e, 0x88, 0x80, 0xbd, 0xcd, 0xb8, 0x30, 0x3d, 0x74, 0x78, 0x0d, 0xe2,
	0xd3, 0x84, 0x4e, 0x33, 0xf0, 0x08, 0x37, 0xdb, 0xee, 0x86, 0x1b, 0xe6, 0xaa, 0xeb, 0xd2, 0x3e,
	0xde, 0x12, 0xab, 0x26, 0x73, 0x8d, 0x1e, 0xad, 0xa1, 0x76, 0x92, 0x8e, 0xb6, 0xc6, 0x24, 0xc8,
	0x43, 0x2d, 0xd5, 0x5e, 0x8f, 0x98, 0x9b, 0x8f, 0xac, 0x47, 0xd8, 0x76, 0x76, 0x85, 0xd3, 0x78,
	0x3e, 0xb4, 0xef, 0xff, 0xd8, 0x7b, 0xf7, 0x0b, 0x5b, 0xfb, 0xb5, 0xb0, 0xb5, 0xdf, 0x0b, 0x5b,
	0xfb, 0xfe, 0xd7, 0xde, 0xbb, 0x31, 0x6a, 0x75, 0xd8, 0x54, 0x4f, 0xe2, 0xf2, 0x7f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc2, 0xa2, 0x9d, 0x4d, 0x68, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StarDustClient is the client API for StarDust service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StarDustClient interface {
	GetUniqueIds(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*IdsReply, error)
	GetUniqueId(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*IdReply, error)
	ParseId(ctx context.Context, in *ParseReq, opts ...grpc.CallOption) (*ParseReply, error)
}

type starDustClient struct {
	cc *grpc.ClientConn
}

func NewStarDustClient(cc *grpc.ClientConn) StarDustClient {
	return &starDustClient{cc}
}

func (c *starDustClient) GetUniqueIds(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*IdsReply, error) {
	out := new(IdsReply)
	err := c.cc.Invoke(ctx, "/service.stardust.v1.StarDust/GetUniqueIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starDustClient) GetUniqueId(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*IdReply, error) {
	out := new(IdReply)
	err := c.cc.Invoke(ctx, "/service.stardust.v1.StarDust/GetUniqueId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starDustClient) ParseId(ctx context.Context, in *ParseReq, opts ...grpc.CallOption) (*ParseReply, error) {
	out := new(ParseReply)
	err := c.cc.Invoke(ctx, "/service.stardust.v1.StarDust/ParseId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StarDustServer is the server API for StarDust service.
type StarDustServer interface {
	GetUniqueIds(context.Context, *IdReq) (*IdsReply, error)
	GetUniqueId(context.Context, *IdReq) (*IdReply, error)
	ParseId(context.Context, *ParseReq) (*ParseReply, error)
}

// UnimplementedStarDustServer can be embedded to have forward compatible implementations.
type UnimplementedStarDustServer struct {
}

func (*UnimplementedStarDustServer) GetUniqueIds(ctx context.Context, req *IdReq) (*IdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUniqueIds not implemented")
}
func (*UnimplementedStarDustServer) GetUniqueId(ctx context.Context, req *IdReq) (*IdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUniqueId not implemented")
}
func (*UnimplementedStarDustServer) ParseId(ctx context.Context, req *ParseReq) (*ParseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseId not implemented")
}

func RegisterStarDustServer(s *grpc.Server, srv StarDustServer) {
	s.RegisterService(&_StarDust_serviceDesc, srv)
}

func _StarDust_GetUniqueIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarDustServer).GetUniqueIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.stardust.v1.StarDust/GetUniqueIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarDustServer).GetUniqueIds(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarDust_GetUniqueId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarDustServer).GetUniqueId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.stardust.v1.StarDust/GetUniqueId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarDustServer).GetUniqueId(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarDust_ParseId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarDustServer).ParseId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.stardust.v1.StarDust/ParseId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarDustServer).ParseId(ctx, req.(*ParseReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _StarDust_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.stardust.v1.StarDust",
	HandlerType: (*StarDustServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUniqueIds",
			Handler:    _StarDust_GetUniqueIds_Handler,
		},
		{
			MethodName: "GetUniqueId",
			Handler:    _StarDust_GetUniqueId_Handler,
		},
		{
			MethodName: "ParseId",
			Handler:    _StarDust_ParseId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stardust.proto",
}

func (m *IdReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Len != 0 {
		i = encodeVarintStardust(dAtA, i, uint64(m.Len))
		i--
		dAtA[i] = 0x10
	}
	if m.BizType != 0 {
		i = encodeVarintStardust(dAtA, i, uint64(m.BizType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *IdsReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdsReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdsReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Ids) > 0 {
		dAtA2 := make([]byte, len(m.Ids)*10)
		var j1 int
		for _, num1 := range m.Ids {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintStardust(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IdReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintStardust(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ParseReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParseReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParseReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintStardust(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ParseReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParseReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParseReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Sequence != 0 {
		i = encodeVarintStardust(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x28
	}
	if m.BizType != 0 {
		i = encodeVarintStardust(dAtA, i, uint64(m.BizType))
		i--
		dAtA[i] = 0x20
	}
	if m.Worker != 0 {
		i = encodeVarintStardust(dAtA, i, uint64(m.Worker))
		i--
		dAtA[i] = 0x18
	}
	if m.Idc != 0 {
		i = encodeVarintStardust(dAtA, i, uint64(m.Idc))
		i--
		dAtA[i] = 0x10
	}
	if m.Time != 0 {
		i = encodeVarintStardust(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStardust(dAtA []byte, offset int, v uint64) int {
	offset -= sovStardust(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IdReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BizType != 0 {
		n += 1 + sovStardust(uint64(m.BizType))
	}
	if m.Len != 0 {
		n += 1 + sovStardust(uint64(m.Len))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IdsReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ids) > 0 {
		l = 0
		for _, e := range m.Ids {
			l += sovStardust(uint64(e))
		}
		n += 1 + sovStardust(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IdReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovStardust(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ParseReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovStardust(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ParseReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Time != 0 {
		n += 1 + sovStardust(uint64(m.Time))
	}
	if m.Idc != 0 {
		n += 1 + sovStardust(uint64(m.Idc))
	}
	if m.Worker != 0 {
		n += 1 + sovStardust(uint64(m.Worker))
	}
	if m.BizType != 0 {
		n += 1 + sovStardust(uint64(m.BizType))
	}
	if m.Sequence != 0 {
		n += 1 + sovStardust(uint64(m.Sequence))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStardust(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStardust(x uint64) (n int) {
	return sovStardust(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IdReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStardust
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
			return fmt.Errorf("proto: IdReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BizType", wireType)
			}
			m.BizType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStardust
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BizType |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Len", wireType)
			}
			m.Len = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStardust
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Len |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStardust(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStardust
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStardust
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IdsReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStardust
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
			return fmt.Errorf("proto: IdsReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdsReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStardust
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ids = append(m.Ids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStardust
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthStardust
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthStardust
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Ids) == 0 {
					m.Ids = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStardust
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ids = append(m.Ids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStardust(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStardust
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStardust
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IdReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStardust
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
			return fmt.Errorf("proto: IdReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStardust
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStardust(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStardust
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStardust
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParseReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStardust
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
			return fmt.Errorf("proto: ParseReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParseReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStardust
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStardust(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStardust
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStardust
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParseReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStardust
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
			return fmt.Errorf("proto: ParseReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParseReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStardust
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Idc", wireType)
			}
			m.Idc = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStardust
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Idc |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Worker", wireType)
			}
			m.Worker = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStardust
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Worker |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BizType", wireType)
			}
			m.BizType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStardust
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BizType |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStardust
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStardust(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStardust
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStardust
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStardust(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStardust
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
					return 0, ErrIntOverflowStardust
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
					return 0, ErrIntOverflowStardust
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
				return 0, ErrInvalidLengthStardust
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStardust
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStardust
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStardust        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStardust          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStardust = fmt.Errorf("proto: unexpected end of group")
)
