// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test-service.proto

package protobuffer_def // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

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

type CMD int32

const (
	CMD_TEST_1 CMD = 0
	CMD_TEST_2 CMD = 1
)

var CMD_name = map[int32]string{
	0: "TEST_1",
	1: "TEST_2",
}
var CMD_value = map[string]int32{
	"TEST_1": 0,
	"TEST_2": 1,
}

func (x CMD) String() string {
	return proto.EnumName(CMD_name, int32(x))
}
func (CMD) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_test_service_1d13ca228077dacc, []int{0}
}

type BaseRequest struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	C                    CMD      `protobuf:"varint,2,opt,name=c,proto3,enum=CMD" json:"c,omitempty"`
	Body                 *any.Any `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseRequest) Reset()         { *m = BaseRequest{} }
func (m *BaseRequest) String() string { return proto.CompactTextString(m) }
func (*BaseRequest) ProtoMessage()    {}
func (*BaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_service_1d13ca228077dacc, []int{0}
}
func (m *BaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseRequest.Unmarshal(m, b)
}
func (m *BaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseRequest.Marshal(b, m, deterministic)
}
func (dst *BaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRequest.Merge(dst, src)
}
func (m *BaseRequest) XXX_Size() int {
	return xxx_messageInfo_BaseRequest.Size(m)
}
func (m *BaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRequest proto.InternalMessageInfo

func (m *BaseRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *BaseRequest) GetC() CMD {
	if m != nil {
		return m.C
	}
	return CMD_TEST_1
}

func (m *BaseRequest) GetBody() *any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

type BaseResponse struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	C                    CMD      `protobuf:"varint,2,opt,name=c,proto3,enum=CMD" json:"c,omitempty"`
	Body                 *any.Any `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_service_1d13ca228077dacc, []int{1}
}
func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (dst *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(dst, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *BaseResponse) GetC() CMD {
	if m != nil {
		return m.C
	}
	return CMD_TEST_1
}

func (m *BaseResponse) GetBody() *any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*BaseRequest)(nil), "BaseRequest")
	proto.RegisterType((*BaseResponse)(nil), "BaseResponse")
	proto.RegisterEnum("CMD", CMD_name, CMD_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	BaseInterface(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) BaseInterface(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/TestService/BaseInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	BaseInterface(context.Context, *BaseRequest) (*BaseResponse, error)
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_BaseInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).BaseInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestService/BaseInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).BaseInterface(ctx, req.(*BaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BaseInterface",
			Handler:    _TestService_BaseInterface_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test-service.proto",
}

func init() { proto.RegisterFile("test-service.proto", fileDescriptor_test_service_1d13ca228077dacc) }

var fileDescriptor_test_service_1d13ca228077dacc = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92,
	0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x13, 0xf3, 0x2a,
	0x21, 0x52, 0x4a, 0x99, 0x5c, 0xdc, 0x4e, 0x89, 0xc5, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0x32, 0x5c, 0x9c, 0x45, 0x10, 0xa6, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x42, 0x40, 0x48, 0x88, 0x8b, 0x31, 0x59, 0x82, 0x49, 0x81, 0x51, 0x83, 0xcf, 0x88, 0x45,
	0xcf, 0xd9, 0xd7, 0x25, 0x88, 0x31, 0x59, 0x48, 0x83, 0x8b, 0x25, 0x29, 0x3f, 0xa5, 0x52, 0x82,
	0x59, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x44, 0x0f, 0x62, 0x95, 0x1e, 0xcc, 0x2a, 0x3d, 0xc7, 0xbc,
	0xca, 0x20, 0xb0, 0x0a, 0xa5, 0x2c, 0x2e, 0x1e, 0x88, 0x55, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9,
	0xb4, 0xb4, 0x4b, 0x4b, 0x96, 0x8b, 0xd9, 0xd9, 0xd7, 0x45, 0x88, 0x8b, 0x8b, 0x2d, 0xc4, 0x35,
	0x38, 0x24, 0xde, 0x50, 0x80, 0x01, 0xce, 0x36, 0x12, 0x60, 0x34, 0xb2, 0xe5, 0xe2, 0x0e, 0x49,
	0x2d, 0x2e, 0x09, 0x86, 0x84, 0x92, 0x90, 0x1e, 0x17, 0x2f, 0xc8, 0x65, 0x9e, 0x79, 0x25, 0xa9,
	0x45, 0x69, 0x89, 0xc9, 0xa9, 0x42, 0x3c, 0x7a, 0x48, 0x81, 0x22, 0xc5, 0xab, 0x87, 0xec, 0x6e,
	0x25, 0x06, 0x27, 0xe1, 0x28, 0x41, 0x3d, 0x6b, 0x98, 0xad, 0x69, 0xa9, 0x45, 0xf1, 0x29, 0xa9,
	0x69, 0x49, 0x6c, 0x60, 0x01, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x3d, 0x85, 0xfc,
	0x81, 0x01, 0x00, 0x00,
}
