// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: validate/validate.proto

package validate

import (
	options "github.com/Neakxs/protocel/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ValidateOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options                          *options.Options `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	RequiredSupportDisabled          bool             `protobuf:"varint,2,opt,name=required_support_disabled,json=requiredSupportDisabled,proto3" json:"required_support_disabled,omitempty"`
	ResourceReferenceSupportDisabled bool             `protobuf:"varint,3,opt,name=resource_reference_support_disabled,json=resourceReferenceSupportDisabled,proto3" json:"resource_reference_support_disabled,omitempty"`
}

func (x *ValidateOptions) Reset() {
	*x = ValidateOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validate_validate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateOptions) ProtoMessage() {}

func (x *ValidateOptions) ProtoReflect() protoreflect.Message {
	mi := &file_validate_validate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateOptions.ProtoReflect.Descriptor instead.
func (*ValidateOptions) Descriptor() ([]byte, []int) {
	return file_validate_validate_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateOptions) GetOptions() *options.Options {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *ValidateOptions) GetRequiredSupportDisabled() bool {
	if x != nil {
		return x.RequiredSupportDisabled
	}
	return false
}

func (x *ValidateOptions) GetResourceReferenceSupportDisabled() bool {
	if x != nil {
		return x.ResourceReferenceSupportDisabled
	}
	return false
}

type ValidateRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expr    string           `protobuf:"bytes,1,opt,name=expr,proto3" json:"expr,omitempty"`
	Exprs   []string         `protobuf:"bytes,2,rep,name=exprs,proto3" json:"exprs,omitempty"`
	Options *options.Options `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *ValidateRule) Reset() {
	*x = ValidateRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validate_validate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRule) ProtoMessage() {}

func (x *ValidateRule) ProtoReflect() protoreflect.Message {
	mi := &file_validate_validate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRule.ProtoReflect.Descriptor instead.
func (*ValidateRule) Descriptor() ([]byte, []int) {
	return file_validate_validate_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateRule) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *ValidateRule) GetExprs() []string {
	if x != nil {
		return x.Exprs
	}
	return nil
}

func (x *ValidateRule) GetOptions() *options.Options {
	if x != nil {
		return x.Options
	}
	return nil
}

var file_validate_validate_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*ValidateOptions)(nil),
		Field:         1178,
		Name:          "protocel.validate.file",
		Tag:           "bytes,1178,opt,name=file",
		Filename:      "validate/validate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*ValidateRule)(nil),
		Field:         1178,
		Name:          "protocel.validate.service",
		Tag:           "bytes,1178,opt,name=service",
		Filename:      "validate/validate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*ValidateRule)(nil),
		Field:         1178,
		Name:          "protocel.validate.method",
		Tag:           "bytes,1178,opt,name=method",
		Filename:      "validate/validate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*ValidateRule)(nil),
		Field:         1178,
		Name:          "protocel.validate.message",
		Tag:           "bytes,1178,opt,name=message",
		Filename:      "validate/validate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*ValidateRule)(nil),
		Field:         1178,
		Name:          "protocel.validate.field",
		Tag:           "bytes,1178,opt,name=field",
		Filename:      "validate/validate.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional protocel.validate.ValidateOptions file = 1178;
	E_File = &file_validate_validate_proto_extTypes[0]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional protocel.validate.ValidateRule service = 1178;
	E_Service = &file_validate_validate_proto_extTypes[1]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional protocel.validate.ValidateRule method = 1178;
	E_Method = &file_validate_validate_proto_extTypes[2]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional protocel.validate.ValidateRule message = 1178;
	E_Message = &file_validate_validate_proto_extTypes[3]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional protocel.validate.ValidateRule field = 1178;
	E_Field = &file_validate_validate_proto_extTypes[4]
)

var File_validate_validate_proto protoreflect.FileDescriptor

var file_validate_validate_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x65, 0x6c, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x65, 0x6c, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a,
	0x0a, 0x19, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x17, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x23, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x6d, 0x0a, 0x0c, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x70,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x70, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x78, 0x70, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78,
	0x70, 0x72, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x65, 0x6c, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x55, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9a,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x65, 0x6c,
	0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x3a,
	0x5b, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9a, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x65, 0x6c, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x58, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9a, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x65, 0x6c, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3a, 0x5b, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x9a, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x65, 0x6c, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x3a, 0x55, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9a, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x65, 0x6c, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x61, 0x6b, 0x78, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x65, 0x6c, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_validate_validate_proto_rawDescOnce sync.Once
	file_validate_validate_proto_rawDescData = file_validate_validate_proto_rawDesc
)

func file_validate_validate_proto_rawDescGZIP() []byte {
	file_validate_validate_proto_rawDescOnce.Do(func() {
		file_validate_validate_proto_rawDescData = protoimpl.X.CompressGZIP(file_validate_validate_proto_rawDescData)
	})
	return file_validate_validate_proto_rawDescData
}

var file_validate_validate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_validate_validate_proto_goTypes = []interface{}{
	(*ValidateOptions)(nil),             // 0: protocel.validate.ValidateOptions
	(*ValidateRule)(nil),                // 1: protocel.validate.ValidateRule
	(*options.Options)(nil),             // 2: protocel.options.Options
	(*descriptorpb.FileOptions)(nil),    // 3: google.protobuf.FileOptions
	(*descriptorpb.ServiceOptions)(nil), // 4: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),  // 5: google.protobuf.MethodOptions
	(*descriptorpb.MessageOptions)(nil), // 6: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 7: google.protobuf.FieldOptions
}
var file_validate_validate_proto_depIdxs = []int32{
	2,  // 0: protocel.validate.ValidateOptions.options:type_name -> protocel.options.Options
	2,  // 1: protocel.validate.ValidateRule.options:type_name -> protocel.options.Options
	3,  // 2: protocel.validate.file:extendee -> google.protobuf.FileOptions
	4,  // 3: protocel.validate.service:extendee -> google.protobuf.ServiceOptions
	5,  // 4: protocel.validate.method:extendee -> google.protobuf.MethodOptions
	6,  // 5: protocel.validate.message:extendee -> google.protobuf.MessageOptions
	7,  // 6: protocel.validate.field:extendee -> google.protobuf.FieldOptions
	0,  // 7: protocel.validate.file:type_name -> protocel.validate.ValidateOptions
	1,  // 8: protocel.validate.service:type_name -> protocel.validate.ValidateRule
	1,  // 9: protocel.validate.method:type_name -> protocel.validate.ValidateRule
	1,  // 10: protocel.validate.message:type_name -> protocel.validate.ValidateRule
	1,  // 11: protocel.validate.field:type_name -> protocel.validate.ValidateRule
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	7,  // [7:12] is the sub-list for extension type_name
	2,  // [2:7] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_validate_validate_proto_init() }
func file_validate_validate_proto_init() {
	if File_validate_validate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_validate_validate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateOptions); i {
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
		file_validate_validate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRule); i {
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
			RawDescriptor: file_validate_validate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_validate_validate_proto_goTypes,
		DependencyIndexes: file_validate_validate_proto_depIdxs,
		MessageInfos:      file_validate_validate_proto_msgTypes,
		ExtensionInfos:    file_validate_validate_proto_extTypes,
	}.Build()
	File_validate_validate_proto = out.File
	file_validate_validate_proto_rawDesc = nil
	file_validate_validate_proto_goTypes = nil
	file_validate_validate_proto_depIdxs = nil
}
