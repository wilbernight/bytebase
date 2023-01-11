// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: v1/deployment.proto

package v1

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

type DeploymentType int32

const (
	DeploymentType_DEPLOYMENT_TYPE_UNSPECIFIED DeploymentType = 0
	DeploymentType_DATABASE_CREATE             DeploymentType = 1
	DeploymentType_DATABASE_DDL                DeploymentType = 2
	DeploymentType_DATABASE_DDL_GHOST          DeploymentType = 3
	DeploymentType_DATABASE_DML                DeploymentType = 4
	DeploymentType_DATABASE_RESTORE_PITR       DeploymentType = 5
	DeploymentType_DATABASE_DML_ROLLBACK       DeploymentType = 6
)

// Enum value maps for DeploymentType.
var (
	DeploymentType_name = map[int32]string{
		0: "DEPLOYMENT_TYPE_UNSPECIFIED",
		1: "DATABASE_CREATE",
		2: "DATABASE_DDL",
		3: "DATABASE_DDL_GHOST",
		4: "DATABASE_DML",
		5: "DATABASE_RESTORE_PITR",
		6: "DATABASE_DML_ROLLBACK",
	}
	DeploymentType_value = map[string]int32{
		"DEPLOYMENT_TYPE_UNSPECIFIED": 0,
		"DATABASE_CREATE":             1,
		"DATABASE_DDL":                2,
		"DATABASE_DDL_GHOST":          3,
		"DATABASE_DML":                4,
		"DATABASE_RESTORE_PITR":       5,
		"DATABASE_DML_ROLLBACK":       6,
	}
)

func (x DeploymentType) Enum() *DeploymentType {
	p := new(DeploymentType)
	*p = x
	return p
}

func (x DeploymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_deployment_proto_enumTypes[0].Descriptor()
}

func (DeploymentType) Type() protoreflect.EnumType {
	return &file_v1_deployment_proto_enumTypes[0]
}

func (x DeploymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentType.Descriptor instead.
func (DeploymentType) EnumDescriptor() ([]byte, []int) {
	return file_v1_deployment_proto_rawDescGZIP(), []int{0}
}

var File_v1_deployment_proto protoreflect.FileDescriptor

var file_v1_deployment_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2a, 0xb8, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41,
	0x53, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x44,
	0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x44, 0x44, 0x4c, 0x10, 0x02, 0x12, 0x16, 0x0a,
	0x12, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x44, 0x44, 0x4c, 0x5f, 0x47, 0x48,
	0x4f, 0x53, 0x54, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53,
	0x45, 0x5f, 0x44, 0x4d, 0x4c, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x54, 0x41, 0x42,
	0x41, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x50, 0x49, 0x54, 0x52,
	0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x44,
	0x4d, 0x4c, 0x5f, 0x52, 0x4f, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x06, 0x42, 0x11, 0x5a,
	0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_deployment_proto_rawDescOnce sync.Once
	file_v1_deployment_proto_rawDescData = file_v1_deployment_proto_rawDesc
)

func file_v1_deployment_proto_rawDescGZIP() []byte {
	file_v1_deployment_proto_rawDescOnce.Do(func() {
		file_v1_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_deployment_proto_rawDescData)
	})
	return file_v1_deployment_proto_rawDescData
}

var file_v1_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_deployment_proto_goTypes = []interface{}{
	(DeploymentType)(0), // 0: bytebase.v1.DeploymentType
}
var file_v1_deployment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_deployment_proto_init() }
func file_v1_deployment_proto_init() {
	if File_v1_deployment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_deployment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_deployment_proto_goTypes,
		DependencyIndexes: file_v1_deployment_proto_depIdxs,
		EnumInfos:         file_v1_deployment_proto_enumTypes,
	}.Build()
	File_v1_deployment_proto = out.File
	file_v1_deployment_proto_rawDesc = nil
	file_v1_deployment_proto_goTypes = nil
	file_v1_deployment_proto_depIdxs = nil
}