// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/user/proto/user.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// 用户性别
type UserSex int32

const (
	UserSex_DEFAULT UserSex = 0
	UserSex_MEN     UserSex = 1
	UserSex_WOMEN   UserSex = 2
)

var UserSex_name = map[int32]string{
	0: "DEFAULT",
	1: "MEN",
	2: "WOMEN",
}
var UserSex_value = map[string]int32{
	"DEFAULT": 0,
	"MEN":     1,
	"WOMEN":   2,
}

func (x UserSex) String() string {
	return proto.EnumName(UserSex_name, int32(x))
}
func (UserSex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_d4aa2a9d2faea84e, []int{0}
}

// 用户基本信息
type UserBase struct {
	// @inject_tag: db:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id"`
	// @inject_tag: db:"user_name" valid:"required~用户名称必须存在"
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty" db:"user_name" valid:"required~用户名称必须存在"`
	// @inject_tag: db:"iphone" valid:"required~手机号必须存在"
	Iphone string `protobuf:"bytes,3,opt,name=iphone,proto3" json:"iphone,omitempty" db:"iphone" valid:"required~手机号必须存在"`
	// @inject_tag: db:"password"
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty" db:"password"`
	// @inject_tag: db:"sex"
	Sex                  UserSex  `protobuf:"varint,5,opt,name=sex,proto3,enum=user.UserSex" json:"sex,omitempty" db:"sex"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserBase) Reset()         { *m = UserBase{} }
func (m *UserBase) String() string { return proto.CompactTextString(m) }
func (*UserBase) ProtoMessage()    {}
func (*UserBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d4aa2a9d2faea84e, []int{0}
}
func (m *UserBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBase.Unmarshal(m, b)
}
func (m *UserBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBase.Marshal(b, m, deterministic)
}
func (dst *UserBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBase.Merge(dst, src)
}
func (m *UserBase) XXX_Size() int {
	return xxx_messageInfo_UserBase.Size(m)
}
func (m *UserBase) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBase.DiscardUnknown(m)
}

var xxx_messageInfo_UserBase proto.InternalMessageInfo

func (m *UserBase) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserBase) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserBase) GetIphone() string {
	if m != nil {
		return m.Iphone
	}
	return ""
}

func (m *UserBase) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserBase) GetSex() UserSex {
	if m != nil {
		return m.Sex
	}
	return UserSex_DEFAULT
}

// 用户Id
type UserId struct {
	// 用户id
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d4aa2a9d2faea84e, []int{1}
}
func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (dst *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(dst, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 用户全部选项
type UserAllOption struct {
	// 用户性别
	Sex UserSex `protobuf:"varint,1,opt,name=sex,proto3,enum=user.UserSex" json:"sex,omitempty"`
	// 页数
	Page *Page `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	// 用户名称
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAllOption) Reset()         { *m = UserAllOption{} }
func (m *UserAllOption) String() string { return proto.CompactTextString(m) }
func (*UserAllOption) ProtoMessage()    {}
func (*UserAllOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d4aa2a9d2faea84e, []int{2}
}
func (m *UserAllOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAllOption.Unmarshal(m, b)
}
func (m *UserAllOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAllOption.Marshal(b, m, deterministic)
}
func (dst *UserAllOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAllOption.Merge(dst, src)
}
func (m *UserAllOption) XXX_Size() int {
	return xxx_messageInfo_UserAllOption.Size(m)
}
func (m *UserAllOption) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAllOption.DiscardUnknown(m)
}

var xxx_messageInfo_UserAllOption proto.InternalMessageInfo

func (m *UserAllOption) GetSex() UserSex {
	if m != nil {
		return m.Sex
	}
	return UserSex_DEFAULT
}

func (m *UserAllOption) GetPage() *Page {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *UserAllOption) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// 用户全部
type UserAll struct {
	// 用户信息
	All []*UserBase `protobuf:"bytes,1,rep,name=all,proto3" json:"all,omitempty"`
	// 页数
	Page                 *Page    `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAll) Reset()         { *m = UserAll{} }
func (m *UserAll) String() string { return proto.CompactTextString(m) }
func (*UserAll) ProtoMessage()    {}
func (*UserAll) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d4aa2a9d2faea84e, []int{3}
}
func (m *UserAll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAll.Unmarshal(m, b)
}
func (m *UserAll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAll.Marshal(b, m, deterministic)
}
func (dst *UserAll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAll.Merge(dst, src)
}
func (m *UserAll) XXX_Size() int {
	return xxx_messageInfo_UserAll.Size(m)
}
func (m *UserAll) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAll.DiscardUnknown(m)
}

var xxx_messageInfo_UserAll proto.InternalMessageInfo

func (m *UserAll) GetAll() []*UserBase {
	if m != nil {
		return m.All
	}
	return nil
}

func (m *UserAll) GetPage() *Page {
	if m != nil {
		return m.Page
	}
	return nil
}

// 空消息
type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d4aa2a9d2faea84e, []int{4}
}
func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (dst *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(dst, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

// 分页
type Page struct {
	// 页
	PageIndex int64 `protobuf:"varint,1,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	// 每页大小
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// 总页数
	PageTotal int64 `protobuf:"varint,3,opt,name=page_total,json=pageTotal,proto3" json:"page_total,omitempty"`
	// 条数
	Count int64 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	// 总条数
	Total                int64    `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_d4aa2a9d2faea84e, []int{5}
}
func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (dst *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(dst, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *Page) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Page) GetPageTotal() int64 {
	if m != nil {
		return m.PageTotal
	}
	return 0
}

func (m *Page) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Page) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*UserBase)(nil), "user.UserBase")
	proto.RegisterType((*UserId)(nil), "user.UserId")
	proto.RegisterType((*UserAllOption)(nil), "user.UserAllOption")
	proto.RegisterType((*UserAll)(nil), "user.UserAll")
	proto.RegisterType((*Null)(nil), "user.Null")
	proto.RegisterType((*Page)(nil), "user.Page")
	proto.RegisterEnum("user.UserSex", UserSex_name, UserSex_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// 用户添加
	UserAdd(ctx context.Context, in *UserBase, opts ...grpc.CallOption) (*UserBase, error)
	// 用户更新
	UserUpdate(ctx context.Context, in *UserBase, opts ...grpc.CallOption) (*UserBase, error)
	// 用户删除
	UserDelete(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserId, error)
	// 用户查询一个
	UserQueryOne(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserBase, error)
	// 用户查询全部
	UserQueryAll(ctx context.Context, in *UserAllOption, opts ...grpc.CallOption) (*UserAll, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserAdd(ctx context.Context, in *UserBase, opts ...grpc.CallOption) (*UserBase, error) {
	out := new(UserBase)
	err := c.cc.Invoke(ctx, "/user.UserService/UserAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserUpdate(ctx context.Context, in *UserBase, opts ...grpc.CallOption) (*UserBase, error) {
	out := new(UserBase)
	err := c.cc.Invoke(ctx, "/user.UserService/UserUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserDelete(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := c.cc.Invoke(ctx, "/user.UserService/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserQueryOne(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserBase, error) {
	out := new(UserBase)
	err := c.cc.Invoke(ctx, "/user.UserService/UserQueryOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserQueryAll(ctx context.Context, in *UserAllOption, opts ...grpc.CallOption) (*UserAll, error) {
	out := new(UserAll)
	err := c.cc.Invoke(ctx, "/user.UserService/UserQueryAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// 用户添加
	UserAdd(context.Context, *UserBase) (*UserBase, error)
	// 用户更新
	UserUpdate(context.Context, *UserBase) (*UserBase, error)
	// 用户删除
	UserDelete(context.Context, *UserId) (*UserId, error)
	// 用户查询一个
	UserQueryOne(context.Context, *UserId) (*UserBase, error)
	// 用户查询全部
	UserQueryAll(context.Context, *UserAllOption) (*UserAll, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserAdd(ctx, req.(*UserBase))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserUpdate(ctx, req.(*UserBase))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserDelete(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserQueryOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserQueryOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserQueryOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserQueryOne(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserQueryAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAllOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserQueryAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserQueryAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserQueryAll(ctx, req.(*UserAllOption))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserAdd",
			Handler:    _UserService_UserAdd_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _UserService_UserUpdate_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _UserService_UserDelete_Handler,
		},
		{
			MethodName: "UserQueryOne",
			Handler:    _UserService_UserQueryOne_Handler,
		},
		{
			MethodName: "UserQueryAll",
			Handler:    _UserService_UserQueryAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "srv/user/proto/user.proto",
}

func init() { proto.RegisterFile("srv/user/proto/user.proto", fileDescriptor_user_d4aa2a9d2faea84e) }

var fileDescriptor_user_d4aa2a9d2faea84e = []byte{
	// 784 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5f, 0x8b, 0xe3, 0x54,
	0x14, 0xdf, 0x9b, 0xb4, 0x9d, 0xf6, 0xf4, 0xcf, 0xd6, 0xbb, 0x0a, 0x35, 0x5a, 0xb7, 0xe4, 0x41,
	0x4b, 0xd8, 0x26, 0x9d, 0xae, 0xb0, 0xb4, 0x8a, 0x6c, 0xea, 0xac, 0x50, 0x46, 0xa7, 0x6b, 0x66,
	0x07, 0xc5, 0x97, 0xe5, 0xb6, 0xb9, 0xa4, 0x91, 0x34, 0x09, 0x49, 0xba, 0xdb, 0xae, 0x08, 0xfe,
	0x01, 0x59, 0x91, 0xc5, 0xa5, 0xfa, 0x22, 0x28, 0xe8, 0xd3, 0x0a, 0x3e, 0x28, 0xbe, 0x2c, 0xf8,
	0x25, 0x7c, 0x98, 0x17, 0x3f, 0x80, 0xed, 0xe0, 0xb7, 0x90, 0x7b, 0x6f, 0xdd, 0xcc, 0x74, 0x04,
	0x7d, 0xca, 0x39, 0xbf, 0x3f, 0xe7, 0xdc, 0x73, 0x92, 0x1b, 0x78, 0x36, 0x8e, 0xee, 0x18, 0xb3,
	0x98, 0x46, 0x46, 0x18, 0x05, 0x49, 0xc0, 0x43, 0x9d, 0x87, 0x38, 0xc3, 0x62, 0xe5, 0x79, 0x27,
	0x08, 0x1c, 0x8f, 0x1a, 0x24, 0x74, 0x0d, 0xe2, 0xfb, 0x41, 0x42, 0x12, 0x37, 0xf0, 0x63, 0xa1,
	0x51, 0xae, 0xf0, 0xc7, 0xb8, 0xe5, 0x50, 0xbf, 0x15, 0xdf, 0x25, 0x8e, 0x43, 0x23, 0x23, 0x08,
	0xb9, 0xe2, 0xbc, 0x5a, 0xfd, 0x03, 0x41, 0xfe, 0x28, 0xa6, 0x51, 0x9f, 0xc4, 0x14, 0xd7, 0x41,
	0x72, 0xed, 0x1a, 0x6a, 0xa0, 0xa6, 0xdc, 0x2f, 0x2f, 0x4d, 0xd0, 0xf2, 0x27, 0x5f, 0xff, 0xb8,
	0xbe, 0xff, 0xd8, 0xb5, 0x2d, 0xc9, 0xb5, 0xb1, 0x0e, 0x05, 0xd6, 0xff, 0xb6, 0x4f, 0xa6, 0xb4,
	0x26, 0x35, 0x50, 0xb3, 0xd0, 0x7f, 0x6a, 0x69, 0x56, 0xb4, 0x92, 0x50, 0xad, 0x1e, 0x3e, 0x38,
	0x79, 0xf4, 0x8b, 0x95, 0x67, 0x9a, 0x03, 0x32, 0xa5, 0xf8, 0x45, 0xc8, 0xb9, 0xe1, 0x24, 0xf0,
	0x69, 0x4d, 0xe6, 0xe2, 0xca, 0xd2, 0x2c, 0x6a, 0x85, 0xf5, 0xe7, 0x5f, 0xac, 0xbf, 0xfd, 0x6d,
	0xf5, 0xe5, 0x63, 0x6b, 0xc3, 0xe2, 0x97, 0x20, 0x1f, 0x92, 0x38, 0xbe, 0x1b, 0x44, 0x76, 0x2d,
	0xc3, 0x95, 0xc5, 0xa5, 0x99, 0xd7, 0x72, 0xab, 0x9f, 0x3f, 0x3b, 0xf9, 0xfe, 0x63, 0xeb, 0x09,
	0x89, 0xaf, 0x80, 0x1c, 0xd3, 0x79, 0x2d, 0xdb, 0x40, 0xcd, 0x4a, 0xa7, 0xac, 0xf3, 0xc5, 0xb0,
	0xc3, 0x1f, 0xd2, 0xf9, 0xc6, 0xb2, 0xfe, 0xe8, 0xd1, 0xea, 0xfe, 0x4f, 0x16, 0x93, 0xa9, 0x35,
	0xc8, 0x31, 0x72, 0x60, 0xe3, 0x4a, 0x3a, 0x17, 0x1b, 0x44, 0x9d, 0x42, 0x99, 0x31, 0xa6, 0xe7,
	0x0d, 0xf9, 0x62, 0xf0, 0x65, 0x51, 0x18, 0xfd, 0x4b, 0x61, 0x5e, 0x0b, 0xbf, 0x00, 0x99, 0x90,
	0x38, 0x62, 0xea, 0x62, 0x07, 0x84, 0xe2, 0x26, 0x71, 0xa8, 0xc5, 0x71, 0xfc, 0xdc, 0xe9, 0xd5,
	0xf0, 0x69, 0xd3, 0x3d, 0xa8, 0xfb, 0xb0, 0xb3, 0x69, 0x87, 0x1b, 0x20, 0x13, 0xcf, 0xab, 0xa1,
	0x86, 0xdc, 0x2c, 0x76, 0x2a, 0x69, 0x23, 0xb6, 0x7e, 0x8b, 0x51, 0xff, 0xd5, 0x49, 0xcd, 0x41,
	0xe6, 0x60, 0xe6, 0x79, 0xea, 0x03, 0x04, 0x19, 0x06, 0xe3, 0x3a, 0x00, 0x23, 0x6e, 0xbb, 0xbe,
	0xbd, 0x19, 0x41, 0xb6, 0x0a, 0x0c, 0x19, 0x30, 0x80, 0x9d, 0x8c, 0xd3, 0xb1, 0x7b, 0x4f, 0x14,
	0x95, 0xd9, 0x42, 0x1d, 0x7a, 0xe8, 0xde, 0x4b, 0xbd, 0x49, 0x90, 0x10, 0x8f, 0x9f, 0x7b, 0xe3,
	0xbd, 0xc5, 0x00, 0xfc, 0x34, 0x64, 0xc7, 0xc1, 0xcc, 0x4f, 0xf8, 0x5b, 0x91, 0x2d, 0x91, 0x30,
	0x54, 0xe8, 0xb3, 0x02, 0xe5, 0x89, 0xa6, 0x89, 0x21, 0x0f, 0xe9, 0x1c, 0x17, 0x61, 0x67, 0xef,
	0xc6, 0x1b, 0xe6, 0xd1, 0x9b, 0xb7, 0xaa, 0x17, 0xf0, 0x0e, 0xc8, 0x6f, 0xdd, 0x38, 0xa8, 0x22,
	0x5c, 0x80, 0xec, 0x3b, 0x43, 0x16, 0x4a, 0x9d, 0x63, 0x09, 0x8a, 0x42, 0x1c, 0xdd, 0x71, 0xc7,
	0x14, 0xbf, 0xb2, 0x59, 0x90, 0x6d, 0xe3, 0xad, 0x9d, 0x28, 0x5b, 0xb9, 0x5a, 0xfd, 0xe4, 0xf8,
	0xcf, 0xaf, 0x24, 0x50, 0xb3, 0xfc, 0x5a, 0xf4, 0x90, 0x86, 0x5f, 0x03, 0x60, 0xec, 0x51, 0x68,
	0x93, 0x84, 0xfe, 0x5f, 0xbf, 0x92, 0xfa, 0x5f, 0x15, 0xfe, 0x3d, 0xea, 0xd1, 0x84, 0xe2, 0x52,
	0xaa, 0x1f, 0xd8, 0xca, 0x99, 0x4c, 0xc5, 0xdc, 0x5b, 0xd2, 0x40, 0xdc, 0xce, 0x0f, 0x5c, 0xfb,
	0x43, 0xfc, 0x3a, 0x94, 0x18, 0xfb, 0xf6, 0x8c, 0x46, 0x8b, 0xa1, 0xbf, 0xed, 0xdf, 0xee, 0xfe,
	0x0c, 0xaf, 0x70, 0x11, 0x9f, 0xaa, 0x30, 0x42, 0x1a, 0xde, 0x3b, 0x55, 0x84, 0x7d, 0x25, 0x97,
	0x52, 0xdb, 0x93, 0x6f, 0x54, 0x29, 0x9f, 0x01, 0xff, 0x19, 0x04, 0x8b, 0x41, 0x46, 0x48, 0xeb,
	0x1f, 0xa3, 0xa5, 0xf9, 0x3b, 0xc2, 0x43, 0xc8, 0x89, 0xfb, 0xa8, 0x5e, 0x07, 0x98, 0xbb, 0x24,
	0x98, 0x52, 0xdf, 0xb9, 0xd6, 0xc5, 0xf5, 0x49, 0x92, 0x84, 0x71, 0xcf, 0x30, 0x1c, 0x37, 0x99,
	0xcc, 0x46, 0xfa, 0x38, 0x98, 0x1a, 0x29, 0xad, 0x5c, 0x62, 0x4f, 0x11, 0x5f, 0x5f, 0x50, 0x32,
	0xd1, 0x7d, 0x9a, 0x74, 0xe4, 0x5d, 0xbd, 0xad, 0x54, 0x76, 0x3b, 0xd7, 0xf4, 0xb6, 0xde, 0xd6,
	0x77, 0x7b, 0xdd, 0x6e, 0xb7, 0xab, 0x21, 0xd4, 0xa9, 0x92, 0x30, 0xf4, 0xdc, 0x31, 0xff, 0x93,
	0x18, 0xef, 0xc7, 0x81, 0xdf, 0x3b, 0x87, 0x58, 0x18, 0xe4, 0x97, 0xdb, 0x57, 0x71, 0x11, 0x0a,
	0xeb, 0x5f, 0x7f, 0x58, 0x7f, 0xf7, 0xe9, 0x5f, 0xdf, 0x3c, 0x7c, 0xef, 0x32, 0xd4, 0x01, 0xcc,
	0xd0, 0xdd, 0xa7, 0x0b, 0x73, 0x96, 0x4c, 0xf0, 0xc5, 0xbc, 0xa4, 0x14, 0xde, 0x6d, 0x99, 0x37,
	0x07, 0xad, 0x7d, 0xba, 0x68, 0x48, 0xa3, 0x2a, 0x54, 0xce, 0x08, 0x2e, 0x8c, 0x72, 0xfc, 0x3f,
	0x75, 0xf5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0x3c, 0x17, 0x33, 0x16, 0x05, 0x00, 0x00,
}
