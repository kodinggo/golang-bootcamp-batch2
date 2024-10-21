// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: pb/rest_api_service/story_request.proto

package rest_api_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

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
		mi := &file_pb_rest_api_service_story_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoryFindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoryFindAllRequest) ProtoMessage() {}

func (x *StoryFindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_story_request_proto_msgTypes[0]
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
	return file_pb_rest_api_service_story_request_proto_rawDescGZIP(), []int{0}
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

type StoryByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StoryByIDRequest) Reset() {
	*x = StoryByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rest_api_service_story_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoryByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoryByIDRequest) ProtoMessage() {}

func (x *StoryByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_story_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoryByIDRequest.ProtoReflect.Descriptor instead.
func (*StoryByIDRequest) Descriptor() ([]byte, []int) {
	return file_pb_rest_api_service_story_request_proto_rawDescGZIP(), []int{1}
}

func (x *StoryByIDRequest) GetId() int64 {
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
		mi := &file_pb_rest_api_service_story_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoryCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoryCreateRequest) ProtoMessage() {}

func (x *StoryCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_story_request_proto_msgTypes[2]
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
	return file_pb_rest_api_service_story_request_proto_rawDescGZIP(), []int{2}
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
		mi := &file_pb_rest_api_service_story_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoryUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoryUpdateRequest) ProtoMessage() {}

func (x *StoryUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_story_request_proto_msgTypes[3]
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
	return file_pb_rest_api_service_story_request_proto_rawDescGZIP(), []int{3}
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

var File_pb_rest_api_service_story_request_proto protoreflect.FileDescriptor

var file_pb_rest_api_service_story_request_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x62, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x76,
	0x0a, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x42,
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
	0x42, 0x15, 0x5a, 0x13, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_rest_api_service_story_request_proto_rawDescOnce sync.Once
	file_pb_rest_api_service_story_request_proto_rawDescData = file_pb_rest_api_service_story_request_proto_rawDesc
)

func file_pb_rest_api_service_story_request_proto_rawDescGZIP() []byte {
	file_pb_rest_api_service_story_request_proto_rawDescOnce.Do(func() {
		file_pb_rest_api_service_story_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_rest_api_service_story_request_proto_rawDescData)
	})
	return file_pb_rest_api_service_story_request_proto_rawDescData
}

var file_pb_rest_api_service_story_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_rest_api_service_story_request_proto_goTypes = []interface{}{
	(*StoryFindAllRequest)(nil), // 0: pb.rest_api_service.StoryFindAllRequest
	(*StoryByIDRequest)(nil),    // 1: pb.rest_api_service.StoryByIDRequest
	(*StoryCreateRequest)(nil),  // 2: pb.rest_api_service.StoryCreateRequest
	(*StoryUpdateRequest)(nil),  // 3: pb.rest_api_service.StoryUpdateRequest
}
var file_pb_rest_api_service_story_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_rest_api_service_story_request_proto_init() }
func file_pb_rest_api_service_story_request_proto_init() {
	if File_pb_rest_api_service_story_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_rest_api_service_story_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_rest_api_service_story_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoryByIDRequest); i {
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
		file_pb_rest_api_service_story_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_rest_api_service_story_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_rest_api_service_story_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_rest_api_service_story_request_proto_goTypes,
		DependencyIndexes: file_pb_rest_api_service_story_request_proto_depIdxs,
		MessageInfos:      file_pb_rest_api_service_story_request_proto_msgTypes,
	}.Build()
	File_pb_rest_api_service_story_request_proto = out.File
	file_pb_rest_api_service_story_request_proto_rawDesc = nil
	file_pb_rest_api_service_story_request_proto_goTypes = nil
	file_pb_rest_api_service_story_request_proto_depIdxs = nil
}
