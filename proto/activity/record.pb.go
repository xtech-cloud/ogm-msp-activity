// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/activity/record.proto

package activity

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

// 频道订阅的请求
type ChannelSubRequest struct {
	Notification         string   `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelSubRequest) Reset()         { *m = ChannelSubRequest{} }
func (m *ChannelSubRequest) String() string { return proto.CompactTextString(m) }
func (*ChannelSubRequest) ProtoMessage()    {}
func (*ChannelSubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ca17cf1a8e4e2c, []int{0}
}

func (m *ChannelSubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelSubRequest.Unmarshal(m, b)
}
func (m *ChannelSubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelSubRequest.Marshal(b, m, deterministic)
}
func (m *ChannelSubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelSubRequest.Merge(m, src)
}
func (m *ChannelSubRequest) XXX_Size() int {
	return xxx_messageInfo_ChannelSubRequest.Size(m)
}
func (m *ChannelSubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelSubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelSubRequest proto.InternalMessageInfo

func (m *ChannelSubRequest) GetNotification() string {
	if m != nil {
		return m.Notification
	}
	return ""
}

// 频道取消订阅的请求
type ChannelUnsubRequest struct {
	Notification         string   `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelUnsubRequest) Reset()         { *m = ChannelUnsubRequest{} }
func (m *ChannelUnsubRequest) String() string { return proto.CompactTextString(m) }
func (*ChannelUnsubRequest) ProtoMessage()    {}
func (*ChannelUnsubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ca17cf1a8e4e2c, []int{1}
}

func (m *ChannelUnsubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelUnsubRequest.Unmarshal(m, b)
}
func (m *ChannelUnsubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelUnsubRequest.Marshal(b, m, deterministic)
}
func (m *ChannelUnsubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelUnsubRequest.Merge(m, src)
}
func (m *ChannelUnsubRequest) XXX_Size() int {
	return xxx_messageInfo_ChannelUnsubRequest.Size(m)
}
func (m *ChannelUnsubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelUnsubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelUnsubRequest proto.InternalMessageInfo

func (m *ChannelUnsubRequest) GetNotification() string {
	if m != nil {
		return m.Notification
	}
	return ""
}

// 获取频道列表的请求
type ChannelFetchRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelFetchRequest) Reset()         { *m = ChannelFetchRequest{} }
func (m *ChannelFetchRequest) String() string { return proto.CompactTextString(m) }
func (*ChannelFetchRequest) ProtoMessage()    {}
func (*ChannelFetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ca17cf1a8e4e2c, []int{2}
}

func (m *ChannelFetchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelFetchRequest.Unmarshal(m, b)
}
func (m *ChannelFetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelFetchRequest.Marshal(b, m, deterministic)
}
func (m *ChannelFetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelFetchRequest.Merge(m, src)
}
func (m *ChannelFetchRequest) XXX_Size() int {
	return xxx_messageInfo_ChannelFetchRequest.Size(m)
}
func (m *ChannelFetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelFetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelFetchRequest proto.InternalMessageInfo

func (m *ChannelFetchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ChannelFetchRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// 获取频道列表的回复
type ChannelFetchResponse struct {
	Status               *Status          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Channel              []*ChannelEntity `protobuf:"bytes,2,rep,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ChannelFetchResponse) Reset()         { *m = ChannelFetchResponse{} }
func (m *ChannelFetchResponse) String() string { return proto.CompactTextString(m) }
func (*ChannelFetchResponse) ProtoMessage()    {}
func (*ChannelFetchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ca17cf1a8e4e2c, []int{3}
}

func (m *ChannelFetchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelFetchResponse.Unmarshal(m, b)
}
func (m *ChannelFetchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelFetchResponse.Marshal(b, m, deterministic)
}
func (m *ChannelFetchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelFetchResponse.Merge(m, src)
}
func (m *ChannelFetchResponse) XXX_Size() int {
	return xxx_messageInfo_ChannelFetchResponse.Size(m)
}
func (m *ChannelFetchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelFetchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelFetchResponse proto.InternalMessageInfo

func (m *ChannelFetchResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ChannelFetchResponse) GetChannel() []*ChannelEntity {
	if m != nil {
		return m.Channel
	}
	return nil
}

func init() {
	proto.RegisterType((*ChannelSubRequest)(nil), "activity.ChannelSubRequest")
	proto.RegisterType((*ChannelUnsubRequest)(nil), "activity.ChannelUnsubRequest")
	proto.RegisterType((*ChannelFetchRequest)(nil), "activity.ChannelFetchRequest")
	proto.RegisterType((*ChannelFetchResponse)(nil), "activity.ChannelFetchResponse")
}

func init() {
	proto.RegisterFile("proto/activity/record.proto", fileDescriptor_d3ca17cf1a8e4e2c)
}

var fileDescriptor_d3ca17cf1a8e4e2c = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0xed, 0xc6, 0x3a, 0xf7, 0xd4, 0x83, 0xc6, 0xa1, 0xa5, 0x43, 0x19, 0x3d, 0xf5, 0xd4,
	0x61, 0x3d, 0x88, 0x57, 0x8b, 0x2f, 0xe7, 0x16, 0x3f, 0x40, 0x1a, 0x53, 0x1a, 0x1c, 0xc9, 0x6c,
	0x9e, 0x0a, 0xfb, 0xb8, 0x7e, 0x13, 0x21, 0x4b, 0x5c, 0xe7, 0x54, 0xf0, 0xf8, 0xfc, 0x5f, 0x7e,
	0xb4, 0x7f, 0x02, 0xb3, 0x55, 0xab, 0x50, 0x2d, 0x28, 0x43, 0xf1, 0x2e, 0x70, 0xbd, 0x68, 0x39,
	0x53, 0xed, 0x4b, 0x6a, 0x54, 0x72, 0xe8, 0xe4, 0xe8, 0x7b, 0x4c, 0x37, 0xb4, 0xe5, 0x36, 0x16,
	0xdf, 0xc0, 0x49, 0xde, 0x50, 0x29, 0xf9, 0xb2, 0xec, 0xaa, 0x82, 0xbf, 0x75, 0x5c, 0x23, 0x89,
	0xe1, 0x48, 0x2a, 0x14, 0xb5, 0x60, 0x14, 0x85, 0x92, 0xa1, 0x37, 0xf7, 0x92, 0x49, 0xb1, 0xa3,
	0xc5, 0xb7, 0x70, 0x6a, 0x8b, 0xcf, 0x52, 0xff, 0xaf, 0x9a, 0x7f, 0x55, 0x1f, 0x38, 0xb2, 0xc6,
	0x55, 0xcf, 0xc0, 0x57, 0x75, 0xad, 0x39, 0x9a, 0xd2, 0xb0, 0xb0, 0x17, 0x99, 0xc2, 0x88, 0xa9,
	0x4e, 0x62, 0x38, 0x30, 0xf2, 0xe6, 0x88, 0x35, 0x4c, 0x77, 0x21, 0x7a, 0xa5, 0xa4, 0xe6, 0x24,
	0x01, 0x5f, 0x23, 0xc5, 0x4e, 0x1b, 0x4a, 0x90, 0x1d, 0xa7, 0xee, 0xc7, 0xd3, 0xd2, 0xe8, 0x85,
	0xf5, 0xc9, 0x15, 0x8c, 0xd9, 0x86, 0x10, 0x0e, 0xe6, 0xc3, 0x24, 0xc8, 0xce, 0xb7, 0x51, 0x8b,
	0xbe, 0x97, 0x28, 0x70, 0x5d, 0xb8, 0x5c, 0xf6, 0xe1, 0xc1, 0xd8, 0x5a, 0x24, 0x87, 0x49, 0xd9,
	0x55, 0x9a, 0xb5, 0xa2, 0xe2, 0x64, 0xb6, 0x57, 0xdd, 0xce, 0x19, 0xf5, 0xb8, 0x77, 0x4b, 0x2a,
	0x5f, 0xdd, 0xb7, 0xc6, 0x07, 0xe4, 0x11, 0x02, 0x33, 0x9f, 0xc5, 0x5c, 0xec, 0x61, 0xfa, 0xe3,
	0xfe, 0x05, 0x7a, 0x82, 0x91, 0xd9, 0xe1, 0x07, 0x44, 0x7f, 0xe4, 0xe8, 0xf2, 0x37, 0xdb, 0x91,
	0x2a, 0xdf, 0x3c, 0x8c, 0xeb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x2a, 0x00, 0x14, 0x5e,
	0x02, 0x00, 0x00,
}