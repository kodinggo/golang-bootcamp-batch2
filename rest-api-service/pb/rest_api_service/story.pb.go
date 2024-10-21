// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: pb/rest_api_service/story.proto

package rest_api_service

import (
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
		mi := &file_pb_rest_api_service_story_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Story) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Story) ProtoMessage() {}

func (x *Story) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_story_proto_msgTypes[0]
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
	return file_pb_rest_api_service_story_proto_rawDescGZIP(), []int{0}
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
		mi := &file_pb_rest_api_service_story_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stories) ProtoMessage() {}

func (x *Stories) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rest_api_service_story_proto_msgTypes[1]
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
	return file_pb_rest_api_service_story_proto_rawDescGZIP(), []int{1}
}

func (x *Stories) GetStories() []*Story {
	if x != nil {
		return x.Stories
	}
	return nil
}

var File_pb_rest_api_service_story_proto protoreflect.FileDescriptor

var file_pb_rest_api_service_story_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x3f, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x34, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_rest_api_service_story_proto_rawDescOnce sync.Once
	file_pb_rest_api_service_story_proto_rawDescData = file_pb_rest_api_service_story_proto_rawDesc
)

func file_pb_rest_api_service_story_proto_rawDescGZIP() []byte {
	file_pb_rest_api_service_story_proto_rawDescOnce.Do(func() {
		file_pb_rest_api_service_story_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_rest_api_service_story_proto_rawDescData)
	})
	return file_pb_rest_api_service_story_proto_rawDescData
}

var file_pb_rest_api_service_story_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_rest_api_service_story_proto_goTypes = []interface{}{
	(*Story)(nil),                 // 0: pb.rest_api_service.Story
	(*Stories)(nil),               // 1: pb.rest_api_service.Stories
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_pb_rest_api_service_story_proto_depIdxs = []int32{
	2, // 0: pb.rest_api_service.Story.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: pb.rest_api_service.Story.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: pb.rest_api_service.Stories.stories:type_name -> pb.rest_api_service.Story
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_rest_api_service_story_proto_init() }
func file_pb_rest_api_service_story_proto_init() {
	if File_pb_rest_api_service_story_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_rest_api_service_story_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_rest_api_service_story_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_rest_api_service_story_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_rest_api_service_story_proto_goTypes,
		DependencyIndexes: file_pb_rest_api_service_story_proto_depIdxs,
		MessageInfos:      file_pb_rest_api_service_story_proto_msgTypes,
	}.Build()
	File_pb_rest_api_service_story_proto = out.File
	file_pb_rest_api_service_story_proto_rawDesc = nil
	file_pb_rest_api_service_story_proto_goTypes = nil
	file_pb_rest_api_service_story_proto_depIdxs = nil
}
