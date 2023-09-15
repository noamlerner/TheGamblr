// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: stage.proto

package proto

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

type Stage int32

const (
	Stage_PRE_FLOP       Stage = 0
	Stage_FLOP           Stage = 1
	Stage_TURN           Stage = 2
	Stage_RIVER          Stage = 3
	Stage_ROUND_COMPLETE Stage = 4
	Stage_GAME_COMPLETE  Stage = 5
)

// Enum value maps for Stage.
var (
	Stage_name = map[int32]string{
		0: "PRE_FLOP",
		1: "FLOP",
		2: "TURN",
		3: "RIVER",
		4: "ROUND_COMPLETE",
		5: "GAME_COMPLETE",
	}
	Stage_value = map[string]int32{
		"PRE_FLOP":       0,
		"FLOP":           1,
		"TURN":           2,
		"RIVER":          3,
		"ROUND_COMPLETE": 4,
		"GAME_COMPLETE":  5,
	}
)

func (x Stage) Enum() *Stage {
	p := new(Stage)
	*p = x
	return p
}

func (x Stage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Stage) Descriptor() protoreflect.EnumDescriptor {
	return file_stage_proto_enumTypes[0].Descriptor()
}

func (Stage) Type() protoreflect.EnumType {
	return &file_stage_proto_enumTypes[0]
}

func (x Stage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Stage.Descriptor instead.
func (Stage) EnumDescriptor() ([]byte, []int) {
	return file_stage_proto_rawDescGZIP(), []int{0}
}

var File_stage_proto protoreflect.FileDescriptor

var file_stage_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74,
	0x68, 0x65, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x72, 0x2a, 0x5b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x50, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x4c, 0x4f, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x55, 0x52,
	0x4e, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x49, 0x56, 0x45, 0x52, 0x10, 0x03, 0x12, 0x12,
	0x0a, 0x0e, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0x05, 0x42, 0x11, 0x5a, 0x0f, 0x74, 0x68, 0x65, 0x67, 0x61, 0x6d, 0x62,
	0x6c, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stage_proto_rawDescOnce sync.Once
	file_stage_proto_rawDescData = file_stage_proto_rawDesc
)

func file_stage_proto_rawDescGZIP() []byte {
	file_stage_proto_rawDescOnce.Do(func() {
		file_stage_proto_rawDescData = protoimpl.X.CompressGZIP(file_stage_proto_rawDescData)
	})
	return file_stage_proto_rawDescData
}

var file_stage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_stage_proto_goTypes = []interface{}{
	(Stage)(0), // 0: thegamblr.Stage
}
var file_stage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stage_proto_init() }
func file_stage_proto_init() {
	if File_stage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stage_proto_goTypes,
		DependencyIndexes: file_stage_proto_depIdxs,
		EnumInfos:         file_stage_proto_enumTypes,
	}.Build()
	File_stage_proto = out.File
	file_stage_proto_rawDesc = nil
	file_stage_proto_goTypes = nil
	file_stage_proto_depIdxs = nil
}
