// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ShareCreateRequest struct {
	Volume               string   `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareCreateRequest) Reset()         { *m = ShareCreateRequest{} }
func (m *ShareCreateRequest) String() string { return proto.CompactTextString(m) }
func (*ShareCreateRequest) ProtoMessage()    {}
func (*ShareCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *ShareCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareCreateRequest.Unmarshal(m, b)
}
func (m *ShareCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareCreateRequest.Marshal(b, m, deterministic)
}
func (m *ShareCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareCreateRequest.Merge(m, src)
}
func (m *ShareCreateRequest) XXX_Size() int {
	return xxx_messageInfo_ShareCreateRequest.Size(m)
}
func (m *ShareCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShareCreateRequest proto.InternalMessageInfo

func (m *ShareCreateRequest) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

type ShareDeleteRequest struct {
	Volume               string   `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareDeleteRequest) Reset()         { *m = ShareDeleteRequest{} }
func (m *ShareDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*ShareDeleteRequest) ProtoMessage()    {}
func (*ShareDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *ShareDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareDeleteRequest.Unmarshal(m, b)
}
func (m *ShareDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareDeleteRequest.Marshal(b, m, deterministic)
}
func (m *ShareDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareDeleteRequest.Merge(m, src)
}
func (m *ShareDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_ShareDeleteRequest.Size(m)
}
func (m *ShareDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShareDeleteRequest proto.InternalMessageInfo

func (m *ShareDeleteRequest) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

type ShareGetRequest struct {
	Volume               string   `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareGetRequest) Reset()         { *m = ShareGetRequest{} }
func (m *ShareGetRequest) String() string { return proto.CompactTextString(m) }
func (*ShareGetRequest) ProtoMessage()    {}
func (*ShareGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *ShareGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareGetRequest.Unmarshal(m, b)
}
func (m *ShareGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareGetRequest.Marshal(b, m, deterministic)
}
func (m *ShareGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareGetRequest.Merge(m, src)
}
func (m *ShareGetRequest) XXX_Size() int {
	return xxx_messageInfo_ShareGetRequest.Size(m)
}
func (m *ShareGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShareGetRequest proto.InternalMessageInfo

func (m *ShareGetRequest) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

type ShareWatchRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareWatchRequest) Reset()         { *m = ShareWatchRequest{} }
func (m *ShareWatchRequest) String() string { return proto.CompactTextString(m) }
func (*ShareWatchRequest) ProtoMessage()    {}
func (*ShareWatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *ShareWatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareWatchRequest.Unmarshal(m, b)
}
func (m *ShareWatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareWatchRequest.Marshal(b, m, deterministic)
}
func (m *ShareWatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareWatchRequest.Merge(m, src)
}
func (m *ShareWatchRequest) XXX_Size() int {
	return xxx_messageInfo_ShareWatchRequest.Size(m)
}
func (m *ShareWatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareWatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShareWatchRequest proto.InternalMessageInfo

type ShareResponse struct {
	Volume               string   `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	ExportId             string   `protobuf:"bytes,2,opt,name=export_id,json=exportId,proto3" json:"export_id,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Error                string   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareResponse) Reset()         { *m = ShareResponse{} }
func (m *ShareResponse) String() string { return proto.CompactTextString(m) }
func (*ShareResponse) ProtoMessage()    {}
func (*ShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{4}
}

func (m *ShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareResponse.Unmarshal(m, b)
}
func (m *ShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareResponse.Marshal(b, m, deterministic)
}
func (m *ShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareResponse.Merge(m, src)
}
func (m *ShareResponse) XXX_Size() int {
	return xxx_messageInfo_ShareResponse.Size(m)
}
func (m *ShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShareResponse proto.InternalMessageInfo

func (m *ShareResponse) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *ShareResponse) GetExportId() string {
	if m != nil {
		return m.ExportId
	}
	return ""
}

func (m *ShareResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ShareResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ShareListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareListRequest) Reset()         { *m = ShareListRequest{} }
func (m *ShareListRequest) String() string { return proto.CompactTextString(m) }
func (*ShareListRequest) ProtoMessage()    {}
func (*ShareListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{5}
}

func (m *ShareListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareListRequest.Unmarshal(m, b)
}
func (m *ShareListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareListRequest.Marshal(b, m, deterministic)
}
func (m *ShareListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareListRequest.Merge(m, src)
}
func (m *ShareListRequest) XXX_Size() int {
	return xxx_messageInfo_ShareListRequest.Size(m)
}
func (m *ShareListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShareListRequest proto.InternalMessageInfo

type ShareListResponse struct {
	Shares               map[string]*ShareResponse `protobuf:"bytes,1,rep,name=shares,proto3" json:"shares,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ShareListResponse) Reset()         { *m = ShareListResponse{} }
func (m *ShareListResponse) String() string { return proto.CompactTextString(m) }
func (*ShareListResponse) ProtoMessage()    {}
func (*ShareListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{6}
}

func (m *ShareListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareListResponse.Unmarshal(m, b)
}
func (m *ShareListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareListResponse.Marshal(b, m, deterministic)
}
func (m *ShareListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareListResponse.Merge(m, src)
}
func (m *ShareListResponse) XXX_Size() int {
	return xxx_messageInfo_ShareListResponse.Size(m)
}
func (m *ShareListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShareListResponse proto.InternalMessageInfo

func (m *ShareListResponse) GetShares() map[string]*ShareResponse {
	if m != nil {
		return m.Shares
	}
	return nil
}

type LogWatchRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogWatchRequest) Reset()         { *m = LogWatchRequest{} }
func (m *LogWatchRequest) String() string { return proto.CompactTextString(m) }
func (*LogWatchRequest) ProtoMessage()    {}
func (*LogWatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{7}
}

func (m *LogWatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogWatchRequest.Unmarshal(m, b)
}
func (m *LogWatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogWatchRequest.Marshal(b, m, deterministic)
}
func (m *LogWatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogWatchRequest.Merge(m, src)
}
func (m *LogWatchRequest) XXX_Size() int {
	return xxx_messageInfo_LogWatchRequest.Size(m)
}
func (m *LogWatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogWatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogWatchRequest proto.InternalMessageInfo

type LogResponse struct {
	Line                 string   `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogResponse) Reset()         { *m = LogResponse{} }
func (m *LogResponse) String() string { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()    {}
func (*LogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{8}
}

func (m *LogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogResponse.Unmarshal(m, b)
}
func (m *LogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogResponse.Marshal(b, m, deterministic)
}
func (m *LogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResponse.Merge(m, src)
}
func (m *LogResponse) XXX_Size() int {
	return xxx_messageInfo_LogResponse.Size(m)
}
func (m *LogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogResponse proto.InternalMessageInfo

func (m *LogResponse) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{9}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	GitCommit            string   `protobuf:"bytes,2,opt,name=gitCommit,proto3" json:"gitCommit,omitempty"`
	BuildDate            string   `protobuf:"bytes,3,opt,name=buildDate,proto3" json:"buildDate,omitempty"`
	ApiVersion           int64    `protobuf:"varint,4,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	ApiMinVersion        int64    `protobuf:"varint,5,opt,name=apiMinVersion,proto3" json:"apiMinVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{10}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionResponse) GetGitCommit() string {
	if m != nil {
		return m.GitCommit
	}
	return ""
}

func (m *VersionResponse) GetBuildDate() string {
	if m != nil {
		return m.BuildDate
	}
	return ""
}

func (m *VersionResponse) GetApiVersion() int64 {
	if m != nil {
		return m.ApiVersion
	}
	return 0
}

func (m *VersionResponse) GetApiMinVersion() int64 {
	if m != nil {
		return m.ApiMinVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*ShareCreateRequest)(nil), "ShareCreateRequest")
	proto.RegisterType((*ShareDeleteRequest)(nil), "ShareDeleteRequest")
	proto.RegisterType((*ShareGetRequest)(nil), "ShareGetRequest")
	proto.RegisterType((*ShareWatchRequest)(nil), "ShareWatchRequest")
	proto.RegisterType((*ShareResponse)(nil), "ShareResponse")
	proto.RegisterType((*ShareListRequest)(nil), "ShareListRequest")
	proto.RegisterType((*ShareListResponse)(nil), "ShareListResponse")
	proto.RegisterMapType((map[string]*ShareResponse)(nil), "ShareListResponse.SharesEntry")
	proto.RegisterType((*LogWatchRequest)(nil), "LogWatchRequest")
	proto.RegisterType((*LogResponse)(nil), "LogResponse")
	proto.RegisterType((*VersionRequest)(nil), "VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "VersionResponse")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0x13, 0xea, 0x09, 0x6d, 0x9c, 0x09, 0x42, 0x96, 0x41, 0x55, 0xb1, 0x7a, 0x28,
	0x12, 0x5a, 0x50, 0x40, 0x08, 0x71, 0x6d, 0x11, 0xaa, 0x94, 0x5e, 0x5c, 0x09, 0x8e, 0xc8, 0x4d,
	0x46, 0xe9, 0x0a, 0xc7, 0x6b, 0x76, 0x37, 0x11, 0xfd, 0x11, 0xfe, 0x82, 0x5f, 0x44, 0xc8, 0xbb,
	0x6b, 0x27, 0x76, 0x8a, 0x72, 0xdb, 0x79, 0xf3, 0x9e, 0x67, 0xf6, 0xf9, 0xd9, 0xe0, 0xcb, 0x62,
	0xc6, 0x0a, 0x29, 0xb4, 0x88, 0x5f, 0x03, 0xde, 0xdc, 0xa5, 0x92, 0x2e, 0x24, 0xa5, 0x9a, 0x12,
	0xfa, 0xb9, 0x22, 0xa5, 0xf1, 0x19, 0xf4, 0xd7, 0x22, 0x5b, 0x2d, 0x29, 0xf4, 0x4e, 0xbd, 0x73,
	0x3f, 0x71, 0x55, 0xcd, 0xbe, 0xa4, 0x8c, 0xf6, 0xb3, 0x5f, 0xc1, 0xd0, 0xb0, 0xbf, 0x90, 0xde,
	0x47, 0x1d, 0xc3, 0xc8, 0x50, 0xbf, 0xa5, 0x7a, 0x76, 0xe7, 0xc8, 0x71, 0x01, 0x47, 0x06, 0x4c,
	0x48, 0x15, 0x22, 0x57, 0xf4, 0x3f, 0x35, 0x3e, 0x07, 0x9f, 0x7e, 0x15, 0x42, 0xea, 0xef, 0x7c,
	0x1e, 0x76, 0x4c, 0xeb, 0xd0, 0x02, 0x57, 0x73, 0x7c, 0x0a, 0x3d, 0xa5, 0x53, 0x4d, 0x61, 0xd7,
	0x34, 0x6c, 0x51, 0xa2, 0x24, 0xa5, 0x90, 0xe1, 0x81, 0x45, 0x4d, 0x11, 0x23, 0x04, 0x66, 0xe2,
	0x94, 0xab, 0x6a, 0xe5, 0xf8, 0xb7, 0xe7, 0x76, 0xb3, 0xa0, 0x5b, 0xe5, 0x03, 0xf4, 0x55, 0x09,
	0xaa, 0xd0, 0x3b, 0xed, 0x9e, 0x0f, 0x26, 0x27, 0x6c, 0x87, 0x63, 0x11, 0xf5, 0x39, 0xd7, 0xf2,
	0x3e, 0x71, 0xec, 0xe8, 0x0a, 0x06, 0x5b, 0x30, 0x06, 0xd0, 0xfd, 0x41, 0xf7, 0xee, 0x3a, 0xe5,
	0x11, 0xcf, 0xa0, 0xb7, 0x4e, 0xb3, 0x15, 0x99, 0x7b, 0x0c, 0x26, 0xc7, 0xac, 0x61, 0x41, 0x62,
	0x9b, 0x9f, 0x3a, 0x1f, 0xbd, 0x78, 0x04, 0xc3, 0xa9, 0x58, 0x34, 0x1c, 0x7b, 0x09, 0x83, 0xa9,
	0x58, 0xd4, 0x4b, 0x22, 0x1c, 0x64, 0x3c, 0x27, 0x67, 0x89, 0x39, 0xc7, 0x01, 0x1c, 0x7f, 0x25,
	0xa9, 0xb8, 0xc8, 0x2b, 0xd1, 0x1f, 0x0f, 0x86, 0x35, 0xe4, 0x94, 0x21, 0x3c, 0x5e, 0x5b, 0xc8,
	0xed, 0x56, 0x95, 0xf8, 0x02, 0xfc, 0x05, 0xd7, 0x17, 0x62, 0xb9, 0xe4, 0xda, 0x3d, 0x78, 0x03,
	0x94, 0xdd, 0xdb, 0x15, 0xcf, 0xe6, 0x97, 0x1b, 0xc3, 0x37, 0x00, 0x9e, 0x00, 0xa4, 0x05, 0x77,
	0xb3, 0x8c, 0xf3, 0xdd, 0x64, 0x0b, 0xc1, 0x33, 0x38, 0x4a, 0x0b, 0x7e, 0xcd, 0xf3, 0x8a, 0xd2,
	0x33, 0x94, 0x26, 0x38, 0xf9, 0xdb, 0x81, 0xb1, 0x31, 0xe5, 0x3a, 0xcd, 0xd3, 0x05, 0xc9, 0x1b,
	0x92, 0x6b, 0x3e, 0x23, 0x7c, 0xef, 0xac, 0xb5, 0x51, 0xc6, 0x31, 0xdb, 0x0d, 0x76, 0xd4, 0xb2,
	0x33, 0x7e, 0x54, 0xab, 0x6c, 0xa4, 0x2b, 0x55, 0x23, 0xe0, 0x0f, 0xa8, 0x18, 0x1c, 0x56, 0xd1,
	0xc6, 0x80, 0xb5, 0x52, 0xfe, 0xe0, 0x14, 0xbf, 0xce, 0x07, 0x8e, 0x58, 0x3b, 0x64, 0x11, 0xee,
	0xc6, 0xc7, 0xa8, 0x60, 0xf3, 0x55, 0xa0, 0xe3, 0x6c, 0xbf, 0xf0, 0xdd, 0x49, 0x6f, 0xbd, 0x72,
	0xb7, 0x2a, 0x17, 0x18, 0xb0, 0x56, 0x44, 0xa2, 0x27, 0x6c, 0x2b, 0x21, 0x86, 0xff, 0x06, 0xc0,
	0x59, 0x5b, 0xde, 0x66, 0xc8, 0x9a, 0xf1, 0x88, 0x02, 0xd6, 0x0a, 0xc7, 0x6d, 0xdf, 0xfc, 0x3a,
	0xde, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x06, 0x7a, 0xd8, 0xc4, 0x47, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShareManagerServiceClient is the client API for ShareManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShareManagerServiceClient interface {
	ShareCreate(ctx context.Context, in *ShareCreateRequest, opts ...grpc.CallOption) (*ShareResponse, error)
	ShareDelete(ctx context.Context, in *ShareDeleteRequest, opts ...grpc.CallOption) (*ShareResponse, error)
	ShareGet(ctx context.Context, in *ShareGetRequest, opts ...grpc.CallOption) (*ShareResponse, error)
	ShareList(ctx context.Context, in *ShareListRequest, opts ...grpc.CallOption) (*ShareListResponse, error)
	ShareWatch(ctx context.Context, in *ShareWatchRequest, opts ...grpc.CallOption) (ShareManagerService_ShareWatchClient, error)
	LogWatch(ctx context.Context, in *LogWatchRequest, opts ...grpc.CallOption) (ShareManagerService_LogWatchClient, error)
	VersionGet(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type shareManagerServiceClient struct {
	cc *grpc.ClientConn
}

func NewShareManagerServiceClient(cc *grpc.ClientConn) ShareManagerServiceClient {
	return &shareManagerServiceClient{cc}
}

func (c *shareManagerServiceClient) ShareCreate(ctx context.Context, in *ShareCreateRequest, opts ...grpc.CallOption) (*ShareResponse, error) {
	out := new(ShareResponse)
	err := c.cc.Invoke(ctx, "/ShareManagerService/ShareCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shareManagerServiceClient) ShareDelete(ctx context.Context, in *ShareDeleteRequest, opts ...grpc.CallOption) (*ShareResponse, error) {
	out := new(ShareResponse)
	err := c.cc.Invoke(ctx, "/ShareManagerService/ShareDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shareManagerServiceClient) ShareGet(ctx context.Context, in *ShareGetRequest, opts ...grpc.CallOption) (*ShareResponse, error) {
	out := new(ShareResponse)
	err := c.cc.Invoke(ctx, "/ShareManagerService/ShareGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shareManagerServiceClient) ShareList(ctx context.Context, in *ShareListRequest, opts ...grpc.CallOption) (*ShareListResponse, error) {
	out := new(ShareListResponse)
	err := c.cc.Invoke(ctx, "/ShareManagerService/ShareList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shareManagerServiceClient) ShareWatch(ctx context.Context, in *ShareWatchRequest, opts ...grpc.CallOption) (ShareManagerService_ShareWatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ShareManagerService_serviceDesc.Streams[0], "/ShareManagerService/ShareWatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &shareManagerServiceShareWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShareManagerService_ShareWatchClient interface {
	Recv() (*ShareResponse, error)
	grpc.ClientStream
}

type shareManagerServiceShareWatchClient struct {
	grpc.ClientStream
}

func (x *shareManagerServiceShareWatchClient) Recv() (*ShareResponse, error) {
	m := new(ShareResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shareManagerServiceClient) LogWatch(ctx context.Context, in *LogWatchRequest, opts ...grpc.CallOption) (ShareManagerService_LogWatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ShareManagerService_serviceDesc.Streams[1], "/ShareManagerService/LogWatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &shareManagerServiceLogWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShareManagerService_LogWatchClient interface {
	Recv() (*LogResponse, error)
	grpc.ClientStream
}

type shareManagerServiceLogWatchClient struct {
	grpc.ClientStream
}

func (x *shareManagerServiceLogWatchClient) Recv() (*LogResponse, error) {
	m := new(LogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shareManagerServiceClient) VersionGet(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/ShareManagerService/VersionGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShareManagerServiceServer is the server API for ShareManagerService service.
type ShareManagerServiceServer interface {
	ShareCreate(context.Context, *ShareCreateRequest) (*ShareResponse, error)
	ShareDelete(context.Context, *ShareDeleteRequest) (*ShareResponse, error)
	ShareGet(context.Context, *ShareGetRequest) (*ShareResponse, error)
	ShareList(context.Context, *ShareListRequest) (*ShareListResponse, error)
	ShareWatch(*ShareWatchRequest, ShareManagerService_ShareWatchServer) error
	LogWatch(*LogWatchRequest, ShareManagerService_LogWatchServer) error
	VersionGet(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedShareManagerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShareManagerServiceServer struct {
}

func (*UnimplementedShareManagerServiceServer) ShareCreate(ctx context.Context, req *ShareCreateRequest) (*ShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareCreate not implemented")
}
func (*UnimplementedShareManagerServiceServer) ShareDelete(ctx context.Context, req *ShareDeleteRequest) (*ShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareDelete not implemented")
}
func (*UnimplementedShareManagerServiceServer) ShareGet(ctx context.Context, req *ShareGetRequest) (*ShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareGet not implemented")
}
func (*UnimplementedShareManagerServiceServer) ShareList(ctx context.Context, req *ShareListRequest) (*ShareListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareList not implemented")
}
func (*UnimplementedShareManagerServiceServer) ShareWatch(req *ShareWatchRequest, srv ShareManagerService_ShareWatchServer) error {
	return status.Errorf(codes.Unimplemented, "method ShareWatch not implemented")
}
func (*UnimplementedShareManagerServiceServer) LogWatch(req *LogWatchRequest, srv ShareManagerService_LogWatchServer) error {
	return status.Errorf(codes.Unimplemented, "method LogWatch not implemented")
}
func (*UnimplementedShareManagerServiceServer) VersionGet(ctx context.Context, req *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VersionGet not implemented")
}

func RegisterShareManagerServiceServer(s *grpc.Server, srv ShareManagerServiceServer) {
	s.RegisterService(&_ShareManagerService_serviceDesc, srv)
}

func _ShareManagerService_ShareCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareManagerServiceServer).ShareCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShareManagerService/ShareCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareManagerServiceServer).ShareCreate(ctx, req.(*ShareCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShareManagerService_ShareDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareManagerServiceServer).ShareDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShareManagerService/ShareDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareManagerServiceServer).ShareDelete(ctx, req.(*ShareDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShareManagerService_ShareGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareManagerServiceServer).ShareGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShareManagerService/ShareGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareManagerServiceServer).ShareGet(ctx, req.(*ShareGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShareManagerService_ShareList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareManagerServiceServer).ShareList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShareManagerService/ShareList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareManagerServiceServer).ShareList(ctx, req.(*ShareListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShareManagerService_ShareWatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ShareWatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShareManagerServiceServer).ShareWatch(m, &shareManagerServiceShareWatchServer{stream})
}

type ShareManagerService_ShareWatchServer interface {
	Send(*ShareResponse) error
	grpc.ServerStream
}

type shareManagerServiceShareWatchServer struct {
	grpc.ServerStream
}

func (x *shareManagerServiceShareWatchServer) Send(m *ShareResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ShareManagerService_LogWatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogWatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShareManagerServiceServer).LogWatch(m, &shareManagerServiceLogWatchServer{stream})
}

type ShareManagerService_LogWatchServer interface {
	Send(*LogResponse) error
	grpc.ServerStream
}

type shareManagerServiceLogWatchServer struct {
	grpc.ServerStream
}

func (x *shareManagerServiceLogWatchServer) Send(m *LogResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ShareManagerService_VersionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareManagerServiceServer).VersionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShareManagerService/VersionGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareManagerServiceServer).VersionGet(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShareManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShareManagerService",
	HandlerType: (*ShareManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShareCreate",
			Handler:    _ShareManagerService_ShareCreate_Handler,
		},
		{
			MethodName: "ShareDelete",
			Handler:    _ShareManagerService_ShareDelete_Handler,
		},
		{
			MethodName: "ShareGet",
			Handler:    _ShareManagerService_ShareGet_Handler,
		},
		{
			MethodName: "ShareList",
			Handler:    _ShareManagerService_ShareList_Handler,
		},
		{
			MethodName: "VersionGet",
			Handler:    _ShareManagerService_VersionGet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ShareWatch",
			Handler:       _ShareManagerService_ShareWatch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LogWatch",
			Handler:       _ShareManagerService_LogWatch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}