// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: efg/backstage/custom_options.proto

package custom_options

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_efg_backstage_custom_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51001,
		Name:          "efg.backstage.custom_options.owner",
		Tag:           "bytes,51001,opt,name=owner",
		Filename:      "efg/backstage/custom_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51002,
		Name:          "efg.backstage.custom_options.system",
		Tag:           "bytes,51002,opt,name=system",
		Filename:      "efg/backstage/custom_options.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional string owner = 51001;
	E_Owner = &file_efg_backstage_custom_options_proto_extTypes[0]
	// optional string system = 51002;
	E_System = &file_efg_backstage_custom_options_proto_extTypes[1]
)

var File_efg_backstage_custom_options_proto protoreflect.FileDescriptor

var file_efg_backstage_custom_options_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x66, 0x67, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x65, 0x66, 0x67, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb9,
	0x8e, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x3a, 0x3c, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xba, 0x8e, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x88, 0x01, 0x01, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x63,
	0x65, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x62,
	0x61, 0x63, 0x6b, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_efg_backstage_custom_options_proto_goTypes = []interface{}{
	(*descriptorpb.ServiceOptions)(nil), // 0: google.protobuf.ServiceOptions
}
var file_efg_backstage_custom_options_proto_depIdxs = []int32{
	0, // 0: efg.backstage.custom_options.owner:extendee -> google.protobuf.ServiceOptions
	0, // 1: efg.backstage.custom_options.system:extendee -> google.protobuf.ServiceOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_efg_backstage_custom_options_proto_init() }
func file_efg_backstage_custom_options_proto_init() {
	if File_efg_backstage_custom_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_efg_backstage_custom_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_efg_backstage_custom_options_proto_goTypes,
		DependencyIndexes: file_efg_backstage_custom_options_proto_depIdxs,
		ExtensionInfos:    file_efg_backstage_custom_options_proto_extTypes,
	}.Build()
	File_efg_backstage_custom_options_proto = out.File
	file_efg_backstage_custom_options_proto_rawDesc = nil
	file_efg_backstage_custom_options_proto_goTypes = nil
	file_efg_backstage_custom_options_proto_depIdxs = nil
}