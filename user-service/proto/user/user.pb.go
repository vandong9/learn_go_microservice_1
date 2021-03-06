// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Erros                []*Error `protobuf:"bytes,3,rep,name=erros,proto3" json:"erros,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErros() []*Error {
	if m != nil {
		return m.Erros
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Erros                []*Error `protobuf:"bytes,3,rep,name=erros,proto3" json:"erros,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErros() []*Error {
	if m != nil {
		return m.Erros
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterType((*Error)(nil), "user.Error")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x6d, 0xf3, 0xd1, 0x8f, 0x1b, 0xda, 0xc5, 0xe5, 0x3d, 0x18, 0xba, 0x78, 0xa4, 0x79, 0xf0,
	0x78, 0x22, 0x54, 0xa8, 0x6b, 0x17, 0x45, 0xa4, 0x74, 0x1b, 0xb5, 0xfb, 0x31, 0xb9, 0x60, 0x30,
	0xcd, 0xc4, 0x99, 0x69, 0xc5, 0x1f, 0xea, 0xff, 0x91, 0xb9, 0xd3, 0x4a, 0x2d, 0x82, 0x6e, 0x92,
	0x7b, 0xce, 0x3d, 0x5c, 0xce, 0x39, 0x0c, 0xfc, 0x6e, 0xb5, 0xb2, 0xea, 0x62, 0x6b, 0x48, 0xf3,
	0x67, 0xc6, 0x18, 0x23, 0x37, 0x67, 0x3b, 0x88, 0xee, 0x0d, 0x69, 0x1c, 0x43, 0xb0, 0x2a, 0x45,
	0x37, 0xed, 0xfe, 0x1f, 0xe6, 0xc1, 0xaa, 0x44, 0x84, 0xa8, 0x91, 0x1b, 0x12, 0x01, 0x33, 0x3c,
	0xa3, 0x80, 0x7e, 0xa1, 0x36, 0xad, 0x6c, 0x5e, 0x45, 0xc8, 0xf4, 0x01, 0xe2, 0x2f, 0x88, 0x69,
	0x23, 0xab, 0x5a, 0x44, 0xcc, 0x7b, 0x80, 0x13, 0x18, 0xb4, 0xd2, 0x98, 0x17, 0xa5, 0x4b, 0x11,
	0xf3, 0xe2, 0x03, 0x67, 0x43, 0xe8, 0xe7, 0xf4, 0xbc, 0x25, 0x63, 0x33, 0x05, 0x83, 0x9c, 0x4c,
	0xab, 0x1a, 0x43, 0xf8, 0x07, 0xd8, 0x16, 0x1b, 0x49, 0xe6, 0x30, 0x63, 0xbf, 0xce, 0x60, 0xce,
	0x3c, 0xa6, 0x10, 0xbb, 0xbf, 0x11, 0x41, 0x1a, 0x9e, 0x08, 0xfc, 0x02, 0xa7, 0x10, 0x93, 0xd6,
	0xca, 0x88, 0x90, 0x15, 0x89, 0x57, 0xdc, 0x68, 0xad, 0x74, 0xee, 0x37, 0xd9, 0x1a, 0xe2, 0x3b,
	0xf5, 0x44, 0x8d, 0xb3, 0x6d, 0xdd, 0xb0, 0xcf, 0xed, 0x81, 0x63, 0x77, 0xb2, 0xae, 0x4a, 0xce,
	0x3e, 0xc8, 0x3d, 0xf8, 0xc9, 0xdd, 0x2b, 0x88, 0x19, 0xbb, 0xf2, 0x0a, 0x55, 0x12, 0x9f, 0x8d,
	0x73, 0x9e, 0x31, 0x85, 0xa4, 0x24, 0x53, 0xe8, 0xaa, 0xb5, 0x95, 0x6a, 0xf6, 0xbd, 0x1e, 0x53,
	0xf3, 0xb7, 0x2e, 0x24, 0x2e, 0xc9, 0x2d, 0xe9, 0x5d, 0x55, 0x10, 0xfe, 0x83, 0xde, 0xb5, 0x26,
	0x69, 0x09, 0x8f, 0x62, 0x4e, 0xc6, 0x7e, 0x3e, 0x34, 0x96, 0x75, 0xf0, 0x2f, 0x84, 0x4b, 0xb2,
	0xdf, 0x88, 0xce, 0xa0, 0xb7, 0x24, 0xbb, 0xa8, 0x6b, 0x1c, 0x1d, 0x76, 0xdc, 0xfe, 0x17, 0xd2,
	0x29, 0x44, 0x8b, 0xad, 0x7d, 0xfc, 0x74, 0x70, 0x1f, 0x97, 0x6b, 0xcb, 0x3a, 0x78, 0x0e, 0xa3,
	0xb5, 0x6b, 0x45, 0x5a, 0xf2, 0x4d, 0x1e, 0xef, 0x4f, 0xc4, 0x0f, 0x3d, 0x7e, 0x6f, 0x97, 0xef,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xad, 0x09, 0x6f, 0x88, 0x02, 0x00, 0x00,
}
