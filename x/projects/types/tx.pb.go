// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgAddProjectKeys struct {
	Creator     string       `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Project     string       `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	ProjectKeys []ProjectKey `protobuf:"bytes,3,rep,name=project_keys,json=projectKeys,proto3" json:"project_keys"`
}

func (m *MsgAddProjectKeys) Reset()         { *m = MsgAddProjectKeys{} }
func (m *MsgAddProjectKeys) String() string { return proto.CompactTextString(m) }
func (*MsgAddProjectKeys) ProtoMessage()    {}
func (*MsgAddProjectKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5dcbe7dfba713c0, []int{0}
}
func (m *MsgAddProjectKeys) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddProjectKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddProjectKeys.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddProjectKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddProjectKeys.Merge(m, src)
}
func (m *MsgAddProjectKeys) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddProjectKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddProjectKeys.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddProjectKeys proto.InternalMessageInfo

func (m *MsgAddProjectKeys) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAddProjectKeys) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *MsgAddProjectKeys) GetProjectKeys() []ProjectKey {
	if m != nil {
		return m.ProjectKeys
	}
	return nil
}

type MsgAddProjectKeysResponse struct {
}

func (m *MsgAddProjectKeysResponse) Reset()         { *m = MsgAddProjectKeysResponse{} }
func (m *MsgAddProjectKeysResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddProjectKeysResponse) ProtoMessage()    {}
func (*MsgAddProjectKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5dcbe7dfba713c0, []int{1}
}
func (m *MsgAddProjectKeysResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddProjectKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddProjectKeysResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddProjectKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddProjectKeysResponse.Merge(m, src)
}
func (m *MsgAddProjectKeysResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddProjectKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddProjectKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddProjectKeysResponse proto.InternalMessageInfo

type MsgSetAdminPolicy struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	Policy  Policy `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy"`
}

func (m *MsgSetAdminPolicy) Reset()         { *m = MsgSetAdminPolicy{} }
func (m *MsgSetAdminPolicy) String() string { return proto.CompactTextString(m) }
func (*MsgSetAdminPolicy) ProtoMessage()    {}
func (*MsgSetAdminPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5dcbe7dfba713c0, []int{2}
}
func (m *MsgSetAdminPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetAdminPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetAdminPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetAdminPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetAdminPolicy.Merge(m, src)
}
func (m *MsgSetAdminPolicy) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetAdminPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetAdminPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetAdminPolicy proto.InternalMessageInfo

func (m *MsgSetAdminPolicy) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSetAdminPolicy) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *MsgSetAdminPolicy) GetPolicy() Policy {
	if m != nil {
		return m.Policy
	}
	return Policy{}
}

type MsgSetAdminPolicyResponse struct {
}

func (m *MsgSetAdminPolicyResponse) Reset()         { *m = MsgSetAdminPolicyResponse{} }
func (m *MsgSetAdminPolicyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetAdminPolicyResponse) ProtoMessage()    {}
func (*MsgSetAdminPolicyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5dcbe7dfba713c0, []int{3}
}
func (m *MsgSetAdminPolicyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetAdminPolicyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetAdminPolicyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetAdminPolicyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetAdminPolicyResponse.Merge(m, src)
}
func (m *MsgSetAdminPolicyResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetAdminPolicyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetAdminPolicyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetAdminPolicyResponse proto.InternalMessageInfo

type MsgSetSubscriptionPolicy struct {
	Creator  string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Projects []string `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	Policy   Policy   `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy"`
}

func (m *MsgSetSubscriptionPolicy) Reset()         { *m = MsgSetSubscriptionPolicy{} }
func (m *MsgSetSubscriptionPolicy) String() string { return proto.CompactTextString(m) }
func (*MsgSetSubscriptionPolicy) ProtoMessage()    {}
func (*MsgSetSubscriptionPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5dcbe7dfba713c0, []int{4}
}
func (m *MsgSetSubscriptionPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetSubscriptionPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetSubscriptionPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetSubscriptionPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetSubscriptionPolicy.Merge(m, src)
}
func (m *MsgSetSubscriptionPolicy) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetSubscriptionPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetSubscriptionPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetSubscriptionPolicy proto.InternalMessageInfo

func (m *MsgSetSubscriptionPolicy) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSetSubscriptionPolicy) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

func (m *MsgSetSubscriptionPolicy) GetPolicy() Policy {
	if m != nil {
		return m.Policy
	}
	return Policy{}
}

type MsgSetSubscriptionPolicyResponse struct {
}

func (m *MsgSetSubscriptionPolicyResponse) Reset()         { *m = MsgSetSubscriptionPolicyResponse{} }
func (m *MsgSetSubscriptionPolicyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetSubscriptionPolicyResponse) ProtoMessage()    {}
func (*MsgSetSubscriptionPolicyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5dcbe7dfba713c0, []int{5}
}
func (m *MsgSetSubscriptionPolicyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetSubscriptionPolicyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetSubscriptionPolicyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetSubscriptionPolicyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetSubscriptionPolicyResponse.Merge(m, src)
}
func (m *MsgSetSubscriptionPolicyResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetSubscriptionPolicyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetSubscriptionPolicyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetSubscriptionPolicyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddProjectKeys)(nil), "lavanet.lava.projects.MsgAddProjectKeys")
	proto.RegisterType((*MsgAddProjectKeysResponse)(nil), "lavanet.lava.projects.MsgAddProjectKeysResponse")
	proto.RegisterType((*MsgSetAdminPolicy)(nil), "lavanet.lava.projects.MsgSetAdminPolicy")
	proto.RegisterType((*MsgSetAdminPolicyResponse)(nil), "lavanet.lava.projects.MsgSetAdminPolicyResponse")
	proto.RegisterType((*MsgSetSubscriptionPolicy)(nil), "lavanet.lava.projects.MsgSetSubscriptionPolicy")
	proto.RegisterType((*MsgSetSubscriptionPolicyResponse)(nil), "lavanet.lava.projects.MsgSetSubscriptionPolicyResponse")
}

func init() { proto.RegisterFile("projects/tx.proto", fileDescriptor_b5dcbe7dfba713c0) }

var fileDescriptor_b5dcbe7dfba713c0 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0xcf,
	0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcd,
	0x49, 0x2c, 0x4b, 0xcc, 0x4b, 0x2d, 0xd1, 0x03, 0xd1, 0x7a, 0x30, 0x79, 0x29, 0x31, 0xb8, 0x4a,
	0x28, 0x03, 0xa2, 0x5c, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xcc, 0xd4, 0x07, 0xb1, 0x20, 0xa2,
	0x4a, 0x93, 0x19, 0xb9, 0x04, 0x7d, 0x8b, 0xd3, 0x1d, 0x53, 0x52, 0x02, 0x20, 0xaa, 0xbd, 0x53,
	0x2b, 0x8b, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x90, 0x0c, 0xd4, 0x58, 0x09, 0x26, 0x88, 0x0c, 0x94, 0x2b,
	0xe4, 0xc5, 0xc5, 0x03, 0x65, 0xc6, 0x67, 0xa7, 0x56, 0x16, 0x4b, 0x30, 0x2b, 0x30, 0x6b, 0x70,
	0x1b, 0x29, 0xea, 0x61, 0x75, 0xa5, 0x1e, 0xc2, 0x36, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82,
	0xb8, 0x0b, 0x10, 0xf6, 0x2b, 0x49, 0x73, 0x49, 0x62, 0x38, 0x2a, 0x28, 0xb5, 0xb8, 0x20, 0x3f,
	0xaf, 0x38, 0x55, 0xa9, 0x05, 0xe2, 0xe4, 0xe0, 0xd4, 0x12, 0xc7, 0x94, 0xdc, 0xcc, 0xbc, 0x80,
	0xfc, 0x9c, 0xcc, 0xe4, 0x4a, 0xb2, 0x9c, 0x6c, 0xcd, 0xc5, 0x56, 0x00, 0xd6, 0x2d, 0xc1, 0xac,
	0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0x8b, 0xcb, 0xb1, 0x60, 0x45, 0x50, 0x87, 0x42, 0xb5, 0x40, 0xdd,
	0x88, 0xea, 0x0a, 0xb8, 0x1b, 0x7b, 0x19, 0xb9, 0x24, 0x20, 0xb2, 0xc1, 0xa5, 0x49, 0xc5, 0xc9,
	0x45, 0x99, 0x05, 0x25, 0x99, 0xf9, 0x84, 0x9d, 0x2a, 0xc5, 0xc5, 0x01, 0xb3, 0x54, 0x82, 0x49,
	0x81, 0x59, 0x83, 0x33, 0x08, 0xce, 0xa7, 0xcc, 0xb1, 0x4a, 0x5c, 0x0a, 0xb8, 0x9c, 0x03, 0x73,
	0xb3, 0xd1, 0x53, 0x26, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x1c, 0x2e, 0x3e, 0xb4, 0xe4, 0xa0,
	0x81, 0xc3, 0x2a, 0x8c, 0x38, 0x92, 0x32, 0x20, 0x56, 0x25, 0xcc, 0x56, 0x90, 0x6d, 0x68, 0x31,
	0x89, 0xc7, 0x36, 0x54, 0x95, 0xf8, 0x6c, 0xc3, 0x1e, 0x2f, 0x42, 0x8d, 0x8c, 0x5c, 0xa2, 0xd8,
	0x23, 0x45, 0x1f, 0xaf, 0x59, 0x98, 0x1a, 0xa4, 0xcc, 0x49, 0xd4, 0x00, 0x73, 0x83, 0x93, 0xd3,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c,
	0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x69, 0xa4, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x0d, 0x07, 0xd3, 0xfa, 0x15, 0xfa, 0x88, 0xec, 0x5f,
	0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xce, 0xbd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b,
	0xf8, 0x81, 0x0b, 0x17, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	AddProjectKeys(ctx context.Context, in *MsgAddProjectKeys, opts ...grpc.CallOption) (*MsgAddProjectKeysResponse, error)
	SetAdminPolicy(ctx context.Context, in *MsgSetAdminPolicy, opts ...grpc.CallOption) (*MsgSetAdminPolicyResponse, error)
	SetSubscriptionPolicy(ctx context.Context, in *MsgSetSubscriptionPolicy, opts ...grpc.CallOption) (*MsgSetSubscriptionPolicyResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddProjectKeys(ctx context.Context, in *MsgAddProjectKeys, opts ...grpc.CallOption) (*MsgAddProjectKeysResponse, error) {
	out := new(MsgAddProjectKeysResponse)
	err := c.cc.Invoke(ctx, "/lavanet.lava.projects.Msg/AddProjectKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetAdminPolicy(ctx context.Context, in *MsgSetAdminPolicy, opts ...grpc.CallOption) (*MsgSetAdminPolicyResponse, error) {
	out := new(MsgSetAdminPolicyResponse)
	err := c.cc.Invoke(ctx, "/lavanet.lava.projects.Msg/SetAdminPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetSubscriptionPolicy(ctx context.Context, in *MsgSetSubscriptionPolicy, opts ...grpc.CallOption) (*MsgSetSubscriptionPolicyResponse, error) {
	out := new(MsgSetSubscriptionPolicyResponse)
	err := c.cc.Invoke(ctx, "/lavanet.lava.projects.Msg/SetSubscriptionPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AddProjectKeys(context.Context, *MsgAddProjectKeys) (*MsgAddProjectKeysResponse, error)
	SetAdminPolicy(context.Context, *MsgSetAdminPolicy) (*MsgSetAdminPolicyResponse, error)
	SetSubscriptionPolicy(context.Context, *MsgSetSubscriptionPolicy) (*MsgSetSubscriptionPolicyResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddProjectKeys(ctx context.Context, req *MsgAddProjectKeys) (*MsgAddProjectKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProjectKeys not implemented")
}
func (*UnimplementedMsgServer) SetAdminPolicy(ctx context.Context, req *MsgSetAdminPolicy) (*MsgSetAdminPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAdminPolicy not implemented")
}
func (*UnimplementedMsgServer) SetSubscriptionPolicy(ctx context.Context, req *MsgSetSubscriptionPolicy) (*MsgSetSubscriptionPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSubscriptionPolicy not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddProjectKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddProjectKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddProjectKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lavanet.lava.projects.Msg/AddProjectKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddProjectKeys(ctx, req.(*MsgAddProjectKeys))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetAdminPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetAdminPolicy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetAdminPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lavanet.lava.projects.Msg/SetAdminPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetAdminPolicy(ctx, req.(*MsgSetAdminPolicy))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetSubscriptionPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetSubscriptionPolicy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetSubscriptionPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lavanet.lava.projects.Msg/SetSubscriptionPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetSubscriptionPolicy(ctx, req.(*MsgSetSubscriptionPolicy))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lavanet.lava.projects.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProjectKeys",
			Handler:    _Msg_AddProjectKeys_Handler,
		},
		{
			MethodName: "SetAdminPolicy",
			Handler:    _Msg_SetAdminPolicy_Handler,
		},
		{
			MethodName: "SetSubscriptionPolicy",
			Handler:    _Msg_SetSubscriptionPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "projects/tx.proto",
}

func (m *MsgAddProjectKeys) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddProjectKeys) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddProjectKeys) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProjectKeys) > 0 {
		for iNdEx := len(m.ProjectKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProjectKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Project) > 0 {
		i -= len(m.Project)
		copy(dAtA[i:], m.Project)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Project)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddProjectKeysResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddProjectKeysResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddProjectKeysResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetAdminPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetAdminPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetAdminPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Policy.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Project) > 0 {
		i -= len(m.Project)
		copy(dAtA[i:], m.Project)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Project)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetAdminPolicyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetAdminPolicyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetAdminPolicyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetSubscriptionPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetSubscriptionPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetSubscriptionPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Policy.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Projects) > 0 {
		for iNdEx := len(m.Projects) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Projects[iNdEx])
			copy(dAtA[i:], m.Projects[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Projects[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetSubscriptionPolicyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetSubscriptionPolicyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetSubscriptionPolicyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddProjectKeys) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Project)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ProjectKeys) > 0 {
		for _, e := range m.ProjectKeys {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgAddProjectKeysResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetAdminPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Project)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Policy.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSetAdminPolicyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetSubscriptionPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Projects) > 0 {
		for _, s := range m.Projects {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = m.Policy.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSetSubscriptionPolicyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddProjectKeys) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddProjectKeys: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddProjectKeys: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Project", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Project = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectKeys = append(m.ProjectKeys, ProjectKey{})
			if err := m.ProjectKeys[len(m.ProjectKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAddProjectKeysResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddProjectKeysResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddProjectKeysResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetAdminPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetAdminPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetAdminPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Project", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Project = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetAdminPolicyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetAdminPolicyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetAdminPolicyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetSubscriptionPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetSubscriptionPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetSubscriptionPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Projects", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Projects = append(m.Projects, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetSubscriptionPolicyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetSubscriptionPolicyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetSubscriptionPolicyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)