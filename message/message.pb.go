// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ConnectRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	X                    int64    `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int64    `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ConnectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConnectRequest) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *ConnectRequest) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type DisconnectRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectRequest) Reset()         { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()    {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *DisconnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectRequest.Unmarshal(m, b)
}
func (m *DisconnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectRequest.Marshal(b, m, deterministic)
}
func (m *DisconnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectRequest.Merge(m, src)
}
func (m *DisconnectRequest) XXX_Size() int {
	return xxx_messageInfo_DisconnectRequest.Size(m)
}
func (m *DisconnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectRequest proto.InternalMessageInfo

func (m *DisconnectRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReceiveMessageRequest struct {
	FromId               string   `protobuf:"bytes,1,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveMessageRequest) Reset()         { *m = ReceiveMessageRequest{} }
func (m *ReceiveMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ReceiveMessageRequest) ProtoMessage()    {}
func (*ReceiveMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *ReceiveMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveMessageRequest.Unmarshal(m, b)
}
func (m *ReceiveMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveMessageRequest.Marshal(b, m, deterministic)
}
func (m *ReceiveMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveMessageRequest.Merge(m, src)
}
func (m *ReceiveMessageRequest) XXX_Size() int {
	return xxx_messageInfo_ReceiveMessageRequest.Size(m)
}
func (m *ReceiveMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveMessageRequest proto.InternalMessageInfo

func (m *ReceiveMessageRequest) GetFromId() string {
	if m != nil {
		return m.FromId
	}
	return ""
}

type ReceiveMessageResponse struct {
	FromId               string   `protobuf:"bytes,1,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Params               []string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveMessageResponse) Reset()         { *m = ReceiveMessageResponse{} }
func (m *ReceiveMessageResponse) String() string { return proto.CompactTextString(m) }
func (*ReceiveMessageResponse) ProtoMessage()    {}
func (*ReceiveMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *ReceiveMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveMessageResponse.Unmarshal(m, b)
}
func (m *ReceiveMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveMessageResponse.Marshal(b, m, deterministic)
}
func (m *ReceiveMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveMessageResponse.Merge(m, src)
}
func (m *ReceiveMessageResponse) XXX_Size() int {
	return xxx_messageInfo_ReceiveMessageResponse.Size(m)
}
func (m *ReceiveMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveMessageResponse proto.InternalMessageInfo

func (m *ReceiveMessageResponse) GetFromId() string {
	if m != nil {
		return m.FromId
	}
	return ""
}

func (m *ReceiveMessageResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ReceiveMessageResponse) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

type SendMessageRequest struct {
	FromId               string   `protobuf:"bytes,1,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Params               []string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (m *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(m, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetFromId() string {
	if m != nil {
		return m.FromId
	}
	return ""
}

func (m *SendMessageRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SendMessageRequest) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "message.Empty")
	proto.RegisterType((*ConnectRequest)(nil), "message.ConnectRequest")
	proto.RegisterType((*DisconnectRequest)(nil), "message.DisconnectRequest")
	proto.RegisterType((*ReceiveMessageRequest)(nil), "message.ReceiveMessageRequest")
	proto.RegisterType((*ReceiveMessageResponse)(nil), "message.ReceiveMessageResponse")
	proto.RegisterType((*SendMessageRequest)(nil), "message.SendMessageRequest")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x65, 0x93, 0xda, 0xd0, 0x41, 0x17, 0x3a, 0x60, 0x1b, 0x2a, 0x68, 0x59, 0x2f, 0x3d, 0x95,
	0x52, 0x2f, 0x5e, 0x04, 0x3f, 0x0f, 0x1e, 0x44, 0x49, 0x4f, 0x0a, 0x22, 0x31, 0x19, 0x35, 0xd0,
	0x7c, 0x98, 0x5d, 0x4b, 0xf3, 0x3f, 0xfd, 0x41, 0x92, 0xa4, 0x59, 0x5b, 0x63, 0xc5, 0xde, 0xf6,
	0xed, 0xbc, 0x37, 0x3c, 0xde, 0x1b, 0xd8, 0x09, 0x49, 0x4a, 0xf7, 0x95, 0x86, 0x49, 0x1a, 0xab,
	0x18, 0xad, 0x05, 0x14, 0x16, 0x6c, 0x5d, 0x85, 0x89, 0xca, 0xc4, 0x1d, 0xf0, 0x8b, 0x38, 0x8a,
	0xc8, 0x53, 0x0e, 0xbd, 0x7f, 0x90, 0x54, 0xc8, 0xc1, 0x08, 0x7c, 0x9b, 0xf5, 0xd9, 0xa0, 0xe5,
	0x18, 0x81, 0x8f, 0x08, 0x8d, 0xc8, 0x0d, 0xc9, 0x36, 0x8a, 0x9f, 0xe2, 0x8d, 0xdb, 0xc0, 0xe6,
	0xb6, 0xd9, 0x67, 0x03, 0xd3, 0x61, 0xf3, 0x1c, 0x65, 0x76, 0xa3, 0x44, 0x99, 0x38, 0x84, 0xf6,
	0x65, 0x20, 0xbd, 0x3f, 0x97, 0x8a, 0x11, 0xec, 0x3a, 0xe4, 0x51, 0x30, 0xa3, 0x9b, 0xd2, 0x51,
	0x45, 0xec, 0x82, 0xf5, 0x92, 0xc6, 0xe1, 0x93, 0x66, 0x37, 0x73, 0x78, 0xed, 0x8b, 0x47, 0xe8,
	0xfc, 0x54, 0xc8, 0x24, 0x8e, 0x24, 0xad, 0x95, 0xe4, 0xce, 0x55, 0x96, 0x68, 0xe7, 0xf9, 0x1b,
	0x3b, 0xd0, 0x4c, 0xdc, 0xd4, 0x0d, 0xa5, 0x6d, 0xf6, 0xcd, 0x9c, 0x5b, 0x22, 0x71, 0x0f, 0x38,
	0xa1, 0xc8, 0xff, 0xa7, 0x9b, 0x4d, 0x56, 0x8f, 0x3f, 0x0d, 0xe0, 0x8b, 0xbd, 0x13, 0x4a, 0x67,
	0x81, 0x47, 0x38, 0x06, 0x6b, 0x91, 0x3a, 0x76, 0x87, 0x55, 0x45, 0xab, 0x3d, 0xf4, 0xb8, 0x1e,
	0x14, 0x4d, 0xe1, 0x31, 0xc0, 0x77, 0xae, 0xd8, 0xd3, 0xd3, 0x5a, 0xd8, 0x35, 0xe5, 0x04, 0xf8,
	0x6a, 0x74, 0xb8, 0xaf, 0x19, 0xbf, 0xb6, 0xd0, 0x3b, 0x58, 0x3b, 0x2f, 0x33, 0x1f, 0x31, 0x3c,
	0x85, 0xf6, 0x52, 0x60, 0xb7, 0xea, 0x8d, 0x52, 0x89, 0x7b, 0x5a, 0x57, 0x0f, 0xb3, 0x66, 0xeb,
	0x04, 0xf8, 0x12, 0xeb, 0x6c, 0x3a, 0xdd, 0x48, 0x7e, 0xde, 0x7a, 0xa8, 0xae, 0xf9, 0xb9, 0x59,
	0x5c, 0xf7, 0xd1, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xe4, 0xdc, 0xa9, 0xee, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*Empty, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*Empty, error)
	ReceiveMessage(ctx context.Context, in *ReceiveMessageRequest, opts ...grpc.CallOption) (MessageService_ReceiveMessageClient, error)
	SendMessageOthers(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Empty, error)
	SendMessageAll(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Empty, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/message.MessageService/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/message.MessageService/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) ReceiveMessage(ctx context.Context, in *ReceiveMessageRequest, opts ...grpc.CallOption) (MessageService_ReceiveMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageService_serviceDesc.Streams[0], "/message.MessageService/ReceiveMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceReceiveMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageService_ReceiveMessageClient interface {
	Recv() (*ReceiveMessageResponse, error)
	grpc.ClientStream
}

type messageServiceReceiveMessageClient struct {
	grpc.ClientStream
}

func (x *messageServiceReceiveMessageClient) Recv() (*ReceiveMessageResponse, error) {
	m := new(ReceiveMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageServiceClient) SendMessageOthers(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendMessageOthers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendMessageAll(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendMessageAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	Connect(context.Context, *ConnectRequest) (*Empty, error)
	Disconnect(context.Context, *DisconnectRequest) (*Empty, error)
	ReceiveMessage(*ReceiveMessageRequest, MessageService_ReceiveMessageServer) error
	SendMessageOthers(context.Context, *SendMessageRequest) (*Empty, error)
	SendMessageAll(context.Context, *SendMessageRequest) (*Empty, error)
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_ReceiveMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageServiceServer).ReceiveMessage(m, &messageServiceReceiveMessageServer{stream})
}

type MessageService_ReceiveMessageServer interface {
	Send(*ReceiveMessageResponse) error
	grpc.ServerStream
}

type messageServiceReceiveMessageServer struct {
	grpc.ServerStream
}

func (x *messageServiceReceiveMessageServer) Send(m *ReceiveMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MessageService_SendMessageOthers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendMessageOthers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendMessageOthers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendMessageOthers(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendMessageAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendMessageAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendMessageAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendMessageAll(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _MessageService_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _MessageService_Disconnect_Handler,
		},
		{
			MethodName: "SendMessageOthers",
			Handler:    _MessageService_SendMessageOthers_Handler,
		},
		{
			MethodName: "SendMessageAll",
			Handler:    _MessageService_SendMessageAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveMessage",
			Handler:       _MessageService_ReceiveMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "message.proto",
}
