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
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_pb_rest_api_service_service_proto protoreflect.FileDescriptor

var file_pb_rest_api_service_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x97, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x51, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x70, 0x62,
	0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x25, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x4d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x70,
	0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x4d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x70, 0x62,
	0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x47, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x62, 0x2f,
	0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pb_rest_api_service_service_proto_goTypes = []interface{}{
	(*StoryFindAllRequest)(nil), // 0: pb.rest_api_service.StoryFindAllRequest
	(*StoryByIDRequest)(nil),    // 1: pb.rest_api_service.StoryByIDRequest
	(*StoryCreateRequest)(nil),  // 2: pb.rest_api_service.StoryCreateRequest
	(*StoryUpdateRequest)(nil),  // 3: pb.rest_api_service.StoryUpdateRequest
	(*Stories)(nil),             // 4: pb.rest_api_service.Stories
	(*Story)(nil),               // 5: pb.rest_api_service.Story
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_pb_rest_api_service_service_proto_depIdxs = []int32{
	0, // 0: pb.rest_api_service.StoryService.FindAll:input_type -> pb.rest_api_service.StoryFindAllRequest
	1, // 1: pb.rest_api_service.StoryService.FindByID:input_type -> pb.rest_api_service.StoryByIDRequest
	2, // 2: pb.rest_api_service.StoryService.Create:input_type -> pb.rest_api_service.StoryCreateRequest
	3, // 3: pb.rest_api_service.StoryService.Update:input_type -> pb.rest_api_service.StoryUpdateRequest
	1, // 4: pb.rest_api_service.StoryService.Delete:input_type -> pb.rest_api_service.StoryByIDRequest
	4, // 5: pb.rest_api_service.StoryService.FindAll:output_type -> pb.rest_api_service.Stories
	5, // 6: pb.rest_api_service.StoryService.FindByID:output_type -> pb.rest_api_service.Story
	5, // 7: pb.rest_api_service.StoryService.Create:output_type -> pb.rest_api_service.Story
	5, // 8: pb.rest_api_service.StoryService.Update:output_type -> pb.rest_api_service.Story
	6, // 9: pb.rest_api_service.StoryService.Delete:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_rest_api_service_service_proto_init() }
func file_pb_rest_api_service_service_proto_init() {
	if File_pb_rest_api_service_service_proto != nil {
		return
	}
	file_pb_rest_api_service_user_proto_init()
	file_pb_rest_api_service_story_proto_init()
	file_pb_rest_api_service_story_request_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_rest_api_service_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_rest_api_service_service_proto_goTypes,
		DependencyIndexes: file_pb_rest_api_service_service_proto_depIdxs,
	}.Build()
	File_pb_rest_api_service_service_proto = out.File
	file_pb_rest_api_service_service_proto_rawDesc = nil
	file_pb_rest_api_service_service_proto_goTypes = nil
	file_pb_rest_api_service_service_proto_depIdxs = nil
}
