// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: pb/rest_api_service/service.proto

package rest_api_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type User_Role int32

const (
	User_AUTHOR   User_Role = 0
	User_ADMIN    User_Role = 1
	User_REVIEWER User_Role = 2
)

// Enum value maps for User_Role.
var (
	User_Role_name = map[int32]string{
		0: "AUTHOR",
		1: "ADMIN",
		2: "REVIEWER",
	}
	User_Role_value = map[string]int32{
		"AUTHOR":   0,
		"ADMIN":    1,
		"REVIEWER": 2,
	}
)

func (x User_Role) Enum() *User_Role {
	p := new(User_Role)
	*p = x
	return p
}

func (x User_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_rest_api_service_service_proto_enumTypes[0].Descriptor()
}

func (User_Role) Type() protoreflect.EnumType {
	return &file_pb_rest_api_service_service_proto_enumTypes[0]
}

func (x User_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_Role.Descriptor instead.
func (User_Role) EnumDescriptor() ([]byte, []int) {
	return file_pb_rest_api_service_service_proto_rawDescGZIP(), []int{0, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email     string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password  string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Role      User_Role              `protobuf:"varint,5,opt,name=role,proto3,enum=pb.rest_api_service.User_Role" json:"role,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rest_api_service_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_pb_rest_api_service_service_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRole() User_Role {
	if x != nil {
		return x.Role
	}
	return User_AUTHOR
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Story struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content   string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId  int64                  `protobuf:"varint,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Story) Reset() {
	*x = Story{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rest_api_service_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Story) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Story) ProtoMessage() {}

func (x *Story) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Story.ProtoReflect.Descriptor instead.
func (*Story) Descriptor() ([]byte, []int) {
	return file_pb_rest_api_service_service_proto_rawDescGZIP(), []int{1}
}

func (x *Story) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Story) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Story) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Story) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *Story) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Story) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Stories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stories []*Story `protobuf:"bytes,1,rep,name=stories,proto3" json:"stories,omitempty"`
}

func (x *Stories) Reset() {
	*x = Stories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rest_api_service_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stories) ProtoMessage() {}

func (x *Stories) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stories.ProtoReflect.Descriptor instead.
func (*Stories) Descriptor() ([]byte, []int) {
	return file_pb_rest_api_service_service_proto_rawDescGZIP(), []int{2}
}

func (x *Stories) GetStories() []*Story {
	if x != nil {
		return x.Stories
	}
	return nil
}

type StoryFindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortBy  string `protobuf:"bytes,1,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	Keyword string `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Cursor  string `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit   int64  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *StoryFindAllRequest) Reset() {
	*x = StoryFindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rest_api_service_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoryFindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoryFindAllRequest) ProtoMessage() {}

func (x *StoryFindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoryFindAllRequest.ProtoReflect.Descriptor instead.
func (*StoryFindAllRequest) Descriptor() ([]byte, []int) {
	return file_pb_rest_api_service_service_proto_rawDescGZIP(), []int{3}
}

func (x *StoryFindAllRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *StoryFindAllRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *StoryFindAllRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *StoryFindAllRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type StoryFindByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StoryFindByIDRequest) Reset() {
	*x = StoryFindByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rest_api_service_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoryFindByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoryFindByIDRequest) ProtoMessage() {}

func (x *StoryFindByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoryFindByIDRequest.ProtoReflect.Descriptor instead.
func (*StoryFindByIDRequest) Descriptor() ([]byte, []int) {
	return file_pb_rest_api_service_service_proto_rawDescGZIP(), []int{4}
}

func (x *StoryFindByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type StoryCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId int64  `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *StoryCreateRequest) Reset() {
	*x = StoryCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rest_api_service_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoryCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoryCreateRequest) ProtoMessage() {}

func (x *StoryCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoryCreateRequest.ProtoReflect.Descriptor instead.
func (*StoryCreateRequest) Descriptor() ([]byte, []int) {
	return file_pb_rest_api_service_service_proto_rawDescGZIP(), []int{5}
}

func (x *StoryCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *StoryCreateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *StoryCreateRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type StoryUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId int64  `protobuf:"varint,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *StoryUpdateRequest) Reset() {
	*x = StoryUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rest_api_service_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoryUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoryUpdateRequest) ProtoMessage() {}

func (x *StoryUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoryUpdateRequest.ProtoReflect.Descriptor instead.
func (*StoryUpdateRequest) Descriptor() ([]byte, []int) {
	return file_pb_rest_api_service_service_proto_rawDescGZIP(), []int{6}
}

func (x *StoryUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StoryUpdateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *StoryUpdateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *StoryUpdateRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type StoryDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StoryDeleteRequest) Reset() {
	*x = StoryDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rest_api_service_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoryDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoryDeleteRequest) ProtoMessage() {}

func (x *StoryDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoryDeleteRequest.ProtoReflect.Descriptor instead.
func (*StoryDeleteRequest) Descriptor() ([]byte, []int) {
	return file_pb_rest_api_service_service_proto_rawDescGZIP(), []int{7}
}

func (x *StoryDeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_pb_rest_api_service_service_proto protoreflect.FileDescriptor

var file_pb_rest_api_service_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x2b, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x55, 0x54, 0x48, 0x4f,
	0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x45, 0x52, 0x10, 0x02, 0x22, 0xda, 0x01, 0x0a,
	0x05, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3f, 0x0a, 0x07, 0x53, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x76, 0x0a, 0x13, 0x53, 0x74,
	0x6f, 0x72, 0x79, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x12, 0x53, 0x74,
	0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x71, 0x0a,
	0x12, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x22, 0x24, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x9d, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x08, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x29, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x4d, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x4d, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x49, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_rest_api_service_service_proto_rawDescOnce sync.Once
	file_pb_rest_api_service_service_proto_rawDescData = file_pb_rest_api_service_service_proto_rawDesc
)

func file_pb_rest_api_service_service_proto_rawDescGZIP() []byte {
	file_pb_rest_api_service_service_proto_rawDescOnce.Do(func() {
		file_pb_rest_api_service_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_rest_api_service_service_proto_rawDescData)
	})
	return file_pb_rest_api_service_service_proto_rawDescData
}

var file_pb_rest_api_service_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_rest_api_service_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pb_rest_api_service_service_proto_goTypes = []interface{}{
	(User_Role)(0),                // 0: pb.rest_api_service.User.Role
	(*User)(nil),                  // 1: pb.rest_api_service.User
	(*Story)(nil),                 // 2: pb.rest_api_service.Story
	(*Stories)(nil),               // 3: pb.rest_api_service.Stories
	(*StoryFindAllRequest)(nil),   // 4: pb.rest_api_service.StoryFindAllRequest
	(*StoryFindByIDRequest)(nil),  // 5: pb.rest_api_service.StoryFindByIDRequest
	(*StoryCreateRequest)(nil),    // 6: pb.rest_api_service.StoryCreateRequest
	(*StoryUpdateRequest)(nil),    // 7: pb.rest_api_service.StoryUpdateRequest
	(*StoryDeleteRequest)(nil),    // 8: pb.rest_api_service.StoryDeleteRequest
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_pb_rest_api_service_service_proto_depIdxs = []int32{
	0,  // 0: pb.rest_api_service.User.role:type_name -> pb.rest_api_service.User.Role
	9,  // 1: pb.rest_api_service.User.created_at:type_name -> google.protobuf.Timestamp
	9,  // 2: pb.rest_api_service.User.updated_at:type_name -> google.protobuf.Timestamp
	9,  // 3: pb.rest_api_service.Story.created_at:type_name -> google.protobuf.Timestamp
	9,  // 4: pb.rest_api_service.Story.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 5: pb.rest_api_service.Stories.stories:type_name -> pb.rest_api_service.Story
	4,  // 6: pb.rest_api_service.StoryService.FindAll:input_type -> pb.rest_api_service.StoryFindAllRequest
	5,  // 7: pb.rest_api_service.StoryService.FindByID:input_type -> pb.rest_api_service.StoryFindByIDRequest
	6,  // 8: pb.rest_api_service.StoryService.Create:input_type -> pb.rest_api_service.StoryCreateRequest
	7,  // 9: pb.rest_api_service.StoryService.Update:input_type -> pb.rest_api_service.StoryUpdateRequest
	8,  // 10: pb.rest_api_service.StoryService.Delete:input_type -> pb.rest_api_service.StoryDeleteRequest
	3,  // 11: pb.rest_api_service.StoryService.FindAll:output_type -> pb.rest_api_service.Stories
	2,  // 12: pb.rest_api_service.StoryService.FindByID:output_type -> pb.rest_api_service.Story
	2,  // 13: pb.rest_api_service.StoryService.Create:output_type -> pb.rest_api_service.Story
	2,  // 14: pb.rest_api_service.StoryService.Update:output_type -> pb.rest_api_service.Story
	10, // 15: pb.rest_api_service.StoryService.Delete:output_type -> google.protobuf.Empty
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pb_rest_api_service_service_proto_init() }
func file_pb_rest_api_service_service_proto_init() {
	if File_pb_rest_api_service_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_rest_api_service_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_pb_rest_api_service_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Story); i {
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
		file_pb_rest_api_service_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stories); i {
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
		file_pb_rest_api_service_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoryFindAllRequest); i {
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
		file_pb_rest_api_service_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoryFindByIDRequest); i {
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
		file_pb_rest_api_service_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoryCreateRequest); i {
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
		file_pb_rest_api_service_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoryUpdateRequest); i {
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
		file_pb_rest_api_service_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoryDeleteRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_rest_api_service_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_rest_api_service_service_proto_goTypes,
		DependencyIndexes: file_pb_rest_api_service_service_proto_depIdxs,
		EnumInfos:         file_pb_rest_api_service_service_proto_enumTypes,
		MessageInfos:      file_pb_rest_api_service_service_proto_msgTypes,
	}.Build()
	File_pb_rest_api_service_service_proto = out.File
	file_pb_rest_api_service_service_proto_rawDesc = nil
	file_pb_rest_api_service_service_proto_goTypes = nil
	file_pb_rest_api_service_service_proto_depIdxs = nil
}
