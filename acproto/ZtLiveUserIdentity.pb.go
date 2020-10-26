// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: ZtLiveUserIdentity.proto

package acproto

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

type ZtLiveUserIdentity_ManagerType int32

const (
	ZtLiveUserIdentity_UNKNOWN_MANAGER_TYPE ZtLiveUserIdentity_ManagerType = 0
	ZtLiveUserIdentity_NORMAL               ZtLiveUserIdentity_ManagerType = 1
)

// Enum value maps for ZtLiveUserIdentity_ManagerType.
var (
	ZtLiveUserIdentity_ManagerType_name = map[int32]string{
		0: "UNKNOWN_MANAGER_TYPE",
		1: "NORMAL",
	}
	ZtLiveUserIdentity_ManagerType_value = map[string]int32{
		"UNKNOWN_MANAGER_TYPE": 0,
		"NORMAL":               1,
	}
)

func (x ZtLiveUserIdentity_ManagerType) Enum() *ZtLiveUserIdentity_ManagerType {
	p := new(ZtLiveUserIdentity_ManagerType)
	*p = x
	return p
}

func (x ZtLiveUserIdentity_ManagerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZtLiveUserIdentity_ManagerType) Descriptor() protoreflect.EnumDescriptor {
	return file_ZtLiveUserIdentity_proto_enumTypes[0].Descriptor()
}

func (ZtLiveUserIdentity_ManagerType) Type() protoreflect.EnumType {
	return &file_ZtLiveUserIdentity_proto_enumTypes[0]
}

func (x ZtLiveUserIdentity_ManagerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZtLiveUserIdentity_ManagerType.Descriptor instead.
func (ZtLiveUserIdentity_ManagerType) EnumDescriptor() ([]byte, []int) {
	return file_ZtLiveUserIdentity_proto_rawDescGZIP(), []int{0, 0}
}

type ZtLiveUserIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerType ZtLiveUserIdentity_ManagerType `protobuf:"varint,1,opt,name=managerType,proto3,enum=AcFunDanmu.ZtLiveUserIdentity_ManagerType" json:"managerType,omitempty"`
}

func (x *ZtLiveUserIdentity) Reset() {
	*x = ZtLiveUserIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ZtLiveUserIdentity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZtLiveUserIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZtLiveUserIdentity) ProtoMessage() {}

func (x *ZtLiveUserIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_ZtLiveUserIdentity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZtLiveUserIdentity.ProtoReflect.Descriptor instead.
func (*ZtLiveUserIdentity) Descriptor() ([]byte, []int) {
	return file_ZtLiveUserIdentity_proto_rawDescGZIP(), []int{0}
}

func (x *ZtLiveUserIdentity) GetManagerType() ZtLiveUserIdentity_ManagerType {
	if x != nil {
		return x.ManagerType
	}
	return ZtLiveUserIdentity_UNKNOWN_MANAGER_TYPE
}

var File_ZtLiveUserIdentity_proto protoreflect.FileDescriptor

var file_ZtLiveUserIdentity_proto_rawDesc = []byte{
	0x0a, 0x18, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75,
	0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0x97, 0x01, 0x0a, 0x12, 0x5a, 0x74, 0x4c, 0x69, 0x76,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x4c, 0x0a,
	0x0b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e,
	0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x33, 0x0a, 0x0b, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ZtLiveUserIdentity_proto_rawDescOnce sync.Once
	file_ZtLiveUserIdentity_proto_rawDescData = file_ZtLiveUserIdentity_proto_rawDesc
)

func file_ZtLiveUserIdentity_proto_rawDescGZIP() []byte {
	file_ZtLiveUserIdentity_proto_rawDescOnce.Do(func() {
		file_ZtLiveUserIdentity_proto_rawDescData = protoimpl.X.CompressGZIP(file_ZtLiveUserIdentity_proto_rawDescData)
	})
	return file_ZtLiveUserIdentity_proto_rawDescData
}

var file_ZtLiveUserIdentity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ZtLiveUserIdentity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ZtLiveUserIdentity_proto_goTypes = []interface{}{
	(ZtLiveUserIdentity_ManagerType)(0), // 0: AcFunDanmu.ZtLiveUserIdentity.ManagerType
	(*ZtLiveUserIdentity)(nil),          // 1: AcFunDanmu.ZtLiveUserIdentity
}
var file_ZtLiveUserIdentity_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.ZtLiveUserIdentity.managerType:type_name -> AcFunDanmu.ZtLiveUserIdentity.ManagerType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ZtLiveUserIdentity_proto_init() }
func file_ZtLiveUserIdentity_proto_init() {
	if File_ZtLiveUserIdentity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ZtLiveUserIdentity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZtLiveUserIdentity); i {
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
			RawDescriptor: file_ZtLiveUserIdentity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ZtLiveUserIdentity_proto_goTypes,
		DependencyIndexes: file_ZtLiveUserIdentity_proto_depIdxs,
		EnumInfos:         file_ZtLiveUserIdentity_proto_enumTypes,
		MessageInfos:      file_ZtLiveUserIdentity_proto_msgTypes,
	}.Build()
	File_ZtLiveUserIdentity_proto = out.File
	file_ZtLiveUserIdentity_proto_rawDesc = nil
	file_ZtLiveUserIdentity_proto_goTypes = nil
	file_ZtLiveUserIdentity_proto_depIdxs = nil
}
