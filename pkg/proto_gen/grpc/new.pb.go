// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: new.proto

package grpc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_INACTIVE Status = 0
	Status_ACTIVE   Status = 1
	Status_DRAFT    Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "INACTIVE",
		1: "ACTIVE",
		2: "DRAFT",
	}
	Status_value = map[string]int32{
		"INACTIVE": 0,
		"ACTIVE":   1,
		"DRAFT":    2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_new_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_new_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{0}
}

type Sort int32

const (
	Sort_DESC Sort = 0
	Sort_ASC  Sort = 1
)

// Enum value maps for Sort.
var (
	Sort_name = map[int32]string{
		0: "DESC",
		1: "ASC",
	}
	Sort_value = map[string]int32{
		"DESC": 0,
		"ASC":  1,
	}
)

func (x Sort) Enum() *Sort {
	p := new(Sort)
	*p = x
	return p
}

func (x Sort) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sort) Descriptor() protoreflect.EnumDescriptor {
	return file_new_proto_enumTypes[1].Descriptor()
}

func (Sort) Type() protoreflect.EnumType {
	return &file_new_proto_enumTypes[1]
}

func (x Sort) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sort.Descriptor instead.
func (Sort) EnumDescriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{1}
}

type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Media []*MediaItem `protobuf:"bytes,1,rep,name=Media,proto3" json:"Media,omitempty"`
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetMedia() []*MediaItem {
	if x != nil {
		return x.Media
	}
	return nil
}

type New struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *UUID                  `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title        string                 `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Text         string                 `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`
	ActivePeriod string                 `protobuf:"bytes,4,opt,name=ActivePeriod,proto3" json:"ActivePeriod,omitempty"`
	Status       *Status                `protobuf:"varint,5,opt,name=Status,proto3,enum=grpc.Status,oneof" json:"Status,omitempty"`
	Media        *Media                 `protobuf:"bytes,6,opt,name=Media,proto3,oneof" json:"Media,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=CreatedAt,proto3,oneof" json:"CreatedAt,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=UpdatedAt,proto3,oneof" json:"UpdatedAt,omitempty"`
	Tags         *string                `protobuf:"bytes,9,opt,name=Tags,proto3,oneof" json:"Tags,omitempty"`
}

func (x *New) Reset() {
	*x = New{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *New) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*New) ProtoMessage() {}

func (x *New) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use New.ProtoReflect.Descriptor instead.
func (*New) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{1}
}

func (x *New) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *New) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *New) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *New) GetActivePeriod() string {
	if x != nil {
		return x.ActivePeriod
	}
	return ""
}

func (x *New) GetStatus() Status {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return Status_INACTIVE
}

func (x *New) GetMedia() *Media {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *New) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *New) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *New) GetTags() string {
	if x != nil && x.Tags != nil {
		return *x.Tags
	}
	return ""
}

type NewsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	News []*New `protobuf:"bytes,1,rep,name=News,proto3" json:"News,omitempty"`
}

func (x *NewsList) Reset() {
	*x = NewsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsList) ProtoMessage() {}

func (x *NewsList) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsList.ProtoReflect.Descriptor instead.
func (*NewsList) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{2}
}

func (x *NewsList) GetNews() []*New {
	if x != nil {
		return x.News
	}
	return nil
}

type ListStatusFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status []Status `protobuf:"varint,1,rep,packed,name=Status,proto3,enum=grpc.Status" json:"Status,omitempty"`
}

func (x *ListStatusFilter) Reset() {
	*x = ListStatusFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStatusFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStatusFilter) ProtoMessage() {}

func (x *ListStatusFilter) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStatusFilter.ProtoReflect.Descriptor instead.
func (*ListStatusFilter) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{3}
}

func (x *ListStatusFilter) GetStatus() []Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sort   *Sort   `protobuf:"varint,1,opt,name=Sort,proto3,enum=grpc.Sort,oneof" json:"Sort,omitempty"`
	Author *UUID   `protobuf:"bytes,2,opt,name=Author,proto3,oneof" json:"Author,omitempty"`
	Status *Status `protobuf:"varint,3,opt,name=Status,proto3,enum=grpc.Status,oneof" json:"Status,omitempty"`
	Query  *string `protobuf:"bytes,4,opt,name=Query,proto3,oneof" json:"Query,omitempty"`
	Page   *int32  `protobuf:"varint,5,opt,name=Page,proto3,oneof" json:"Page,omitempty"`
	Token  string  `protobuf:"bytes,6,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{4}
}

func (x *ListRequest) GetSort() Sort {
	if x != nil && x.Sort != nil {
		return *x.Sort
	}
	return Sort_DESC
}

func (x *ListRequest) GetAuthor() *UUID {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *ListRequest) GetStatus() Status {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return Status_INACTIVE
}

func (x *ListRequest) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *ListRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// TODO: How to use uuid?
type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID Id = 1;
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{5}
}

func (x *InfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string                `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Text  string                `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
	Tags  *string               `protobuf:"bytes,5,opt,name=Tags,proto3,oneof" json:"Tags,omitempty"`
	Media []*MediaUploadRequest `protobuf:"bytes,6,rep,name=Media,proto3" json:"Media,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{6}
}

func (x *CreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateRequest) GetTags() string {
	if x != nil && x.Tags != nil {
		return *x.Tags
	}
	return ""
}

func (x *CreateRequest) GetMedia() []*MediaUploadRequest {
	if x != nil {
		return x.Media
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{7}
}

func (x *CreateResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type String struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *String) Reset() {
	*x = String{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String) ProtoMessage() {}

func (x *String) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use String.ProtoReflect.Descriptor instead.
func (*String) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{8}
}

func (x *String) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_new_proto protoreflect.FileDescriptor

var file_new_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x05, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x12, 0x25, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x22, 0x93, 0x03, 0x0a, 0x03, 0x4e, 0x65,
	0x77, 0x12, 0x1a, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x48, 0x01, 0x52, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x88, 0x01, 0x01, 0x12, 0x3d,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x03, 0x52, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04,
	0x54, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x04, 0x54, 0x61,
	0x67, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x61, 0x67, 0x73, 0x22,
	0x29, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x4e,
	0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4e, 0x65, 0x77, 0x52, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x22, 0x38, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x82, 0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x48, 0x00,
	0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x06, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x55, 0x49, 0x44, 0x48, 0x01, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x48, 0x02, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x53, 0x6f, 0x72, 0x74,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x50, 0x61, 0x67, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a,
	0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x54, 0x61, 0x67, 0x73, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x18, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x2a, 0x2d, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x02, 0x2a, 0x19, 0x0a, 0x04, 0x53, 0x6f, 0x72,
	0x74, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41,
	0x53, 0x43, 0x10, 0x01, 0x32, 0x96, 0x01, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x3b, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4e, 0x65, 0x77, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x12, 0x51, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x28, 0x01, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_new_proto_rawDescOnce sync.Once
	file_new_proto_rawDescData = file_new_proto_rawDesc
)

func file_new_proto_rawDescGZIP() []byte {
	file_new_proto_rawDescOnce.Do(func() {
		file_new_proto_rawDescData = protoimpl.X.CompressGZIP(file_new_proto_rawDescData)
	})
	return file_new_proto_rawDescData
}

var file_new_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_new_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_new_proto_goTypes = []interface{}{
	(Status)(0),                   // 0: grpc.Status
	(Sort)(0),                     // 1: grpc.Sort
	(*Media)(nil),                 // 2: grpc.Media
	(*New)(nil),                   // 3: grpc.New
	(*NewsList)(nil),              // 4: grpc.NewsList
	(*ListStatusFilter)(nil),      // 5: grpc.ListStatusFilter
	(*ListRequest)(nil),           // 6: grpc.ListRequest
	(*InfoRequest)(nil),           // 7: grpc.InfoRequest
	(*CreateRequest)(nil),         // 8: grpc.CreateRequest
	(*CreateResponse)(nil),        // 9: grpc.CreateResponse
	(*String)(nil),                // 10: grpc.String
	(*MediaItem)(nil),             // 11: grpc.MediaItem
	(*UUID)(nil),                  // 12: grpc.UUID
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
	(*MediaUploadRequest)(nil),    // 14: grpc.MediaUploadRequest
}
var file_new_proto_depIdxs = []int32{
	11, // 0: grpc.Media.Media:type_name -> grpc.MediaItem
	12, // 1: grpc.New.Id:type_name -> grpc.UUID
	0,  // 2: grpc.New.Status:type_name -> grpc.Status
	2,  // 3: grpc.New.Media:type_name -> grpc.Media
	13, // 4: grpc.New.CreatedAt:type_name -> google.protobuf.Timestamp
	13, // 5: grpc.New.UpdatedAt:type_name -> google.protobuf.Timestamp
	3,  // 6: grpc.NewsList.News:type_name -> grpc.New
	0,  // 7: grpc.ListStatusFilter.Status:type_name -> grpc.Status
	1,  // 8: grpc.ListRequest.Sort:type_name -> grpc.Sort
	12, // 9: grpc.ListRequest.Author:type_name -> grpc.UUID
	0,  // 10: grpc.ListRequest.Status:type_name -> grpc.Status
	14, // 11: grpc.CreateRequest.Media:type_name -> grpc.MediaUploadRequest
	6,  // 12: grpc.News.List:input_type -> grpc.ListRequest
	8,  // 13: grpc.News.Create:input_type -> grpc.CreateRequest
	4,  // 14: grpc.News.List:output_type -> grpc.NewsList
	9,  // 15: grpc.News.Create:output_type -> grpc.CreateResponse
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_new_proto_init() }
func file_new_proto_init() {
	if File_new_proto != nil {
		return
	}
	file_uuid_proto_init()
	file_media_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_new_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_new_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*New); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_new_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_new_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStatusFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_new_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_new_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_new_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_new_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_new_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*String); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_new_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_new_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_new_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_new_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_new_proto_goTypes,
		DependencyIndexes: file_new_proto_depIdxs,
		EnumInfos:         file_new_proto_enumTypes,
		MessageInfos:      file_new_proto_msgTypes,
	}.Build()
	File_new_proto = out.File
	file_new_proto_rawDesc = nil
	file_new_proto_goTypes = nil
	file_new_proto_depIdxs = nil
}
