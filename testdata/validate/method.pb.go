// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: testdata/validate/method.proto

package validate

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

var File_testdata_validate_method_proto protoreflect.FileDescriptor

var file_testdata_validate_method_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x83, 0x01, 0x0a, 0x0a, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x45, 0x78, 0x70, 0x72, 0x12, 0x75, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x3e, 0xd2, 0x49, 0x3b, 0x0a, 0x39, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5b, 0x22, 0x78, 0x2d, 0x69, 0x73, 0x2d, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x22, 0x5d, 0x20, 0x3d, 0x3d, 0x20, 0x22, 0x74, 0x72, 0x75, 0x65, 0x22, 0x32,
	0x82, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x71, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3a, 0xd2, 0x49, 0x37, 0x0a, 0x35, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x5b, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x48, 0x64, 0x72, 0x5d, 0x20, 0x3d, 0x3d, 0x20, 0x22, 0x74,
	0x72, 0x75, 0x65, 0x22, 0x32, 0xa4, 0x01, 0x0a, 0x12, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x03,
	0x52, 0x70, 0x63, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x56, 0xd2, 0x49, 0x53, 0x0a, 0x35, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5b, 0x69, 0x73, 0x41, 0x64,
	0x6d, 0x48, 0x64, 0x72, 0x5d, 0x20, 0x3d, 0x3d, 0x20, 0x22, 0x74, 0x72, 0x75, 0x65, 0x22, 0x1a,
	0x1a, 0x0a, 0x18, 0x12, 0x16, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x48, 0x64, 0x72, 0x12,
	0x0a, 0x78, 0x2d, 0x69, 0x73, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x2e, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x61, 0x6b, 0x78, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x65, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_testdata_validate_method_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil), // 0: google.protobuf.Empty
}
var file_testdata_validate_method_proto_depIdxs = []int32{
	0, // 0: testdata.validate.MethodExpr.Rpc:input_type -> google.protobuf.Empty
	0, // 1: testdata.validate.MethodOptions.Rpc:input_type -> google.protobuf.Empty
	0, // 2: testdata.validate.MethodLocalOptions.Rpc:input_type -> google.protobuf.Empty
	0, // 3: testdata.validate.MethodExpr.Rpc:output_type -> google.protobuf.Empty
	0, // 4: testdata.validate.MethodOptions.Rpc:output_type -> google.protobuf.Empty
	0, // 5: testdata.validate.MethodLocalOptions.Rpc:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testdata_validate_method_proto_init() }
func file_testdata_validate_method_proto_init() {
	if File_testdata_validate_method_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testdata_validate_method_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_testdata_validate_method_proto_goTypes,
		DependencyIndexes: file_testdata_validate_method_proto_depIdxs,
	}.Build()
	File_testdata_validate_method_proto = out.File
	file_testdata_validate_method_proto_rawDesc = nil
	file_testdata_validate_method_proto_goTypes = nil
	file_testdata_validate_method_proto_depIdxs = nil
}