// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: keto/authz/node/v1/node_service.proto

package node

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_keto_authz_node_v1_node_service_proto protoreflect.FileDescriptor

var file_keto_authz_node_v1_node_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x32, 0x0d, 0x0a, 0x0b, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x8b, 0x01, 0x0a, 0x19, 0x73,
	0x68, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b,
	0x65, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x6e, 0x6f, 0x64, 0x65, 0xaa, 0x02, 0x16, 0x4f, 0x72, 0x79, 0x2e, 0x4b, 0x65, 0x74,
	0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x16, 0x4f, 0x72, 0x79, 0x5c, 0x4b, 0x65, 0x74, 0x6f, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a,
	0x5c, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_keto_authz_node_v1_node_service_proto_goTypes = []interface{}{}
var file_keto_authz_node_v1_node_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_keto_authz_node_v1_node_service_proto_init() }
func file_keto_authz_node_v1_node_service_proto_init() {
	if File_keto_authz_node_v1_node_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_keto_authz_node_v1_node_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keto_authz_node_v1_node_service_proto_goTypes,
		DependencyIndexes: file_keto_authz_node_v1_node_service_proto_depIdxs,
	}.Build()
	File_keto_authz_node_v1_node_service_proto = out.File
	file_keto_authz_node_v1_node_service_proto_rawDesc = nil
	file_keto_authz_node_v1_node_service_proto_goTypes = nil
	file_keto_authz_node_v1_node_service_proto_depIdxs = nil
}
