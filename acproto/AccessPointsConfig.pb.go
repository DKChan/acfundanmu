// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.0
// source: AccessPointsConfig.proto

package acproto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AccessPointsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptimalAps            []*AccessPoint `protobuf:"bytes,1,rep,name=optimalAps,proto3" json:"optimalAps,omitempty"`
	BackupAps             []*AccessPoint `protobuf:"bytes,2,rep,name=backupAps,proto3" json:"backupAps,omitempty"`
	AvailablePorts        []uint32       `protobuf:"varint,3,rep,packed,name=availablePorts,proto3" json:"availablePorts,omitempty"`
	ForeceLastConnectedAp *AccessPoint   `protobuf:"bytes,4,opt,name=foreceLastConnectedAp,proto3" json:"foreceLastConnectedAp,omitempty"`
}

func (x *AccessPointsConfig) Reset() {
	*x = AccessPointsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AccessPointsConfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessPointsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessPointsConfig) ProtoMessage() {}

func (x *AccessPointsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_AccessPointsConfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessPointsConfig.ProtoReflect.Descriptor instead.
func (*AccessPointsConfig) Descriptor() ([]byte, []int) {
	return file_AccessPointsConfig_proto_rawDescGZIP(), []int{0}
}

func (x *AccessPointsConfig) GetOptimalAps() []*AccessPoint {
	if x != nil {
		return x.OptimalAps
	}
	return nil
}

func (x *AccessPointsConfig) GetBackupAps() []*AccessPoint {
	if x != nil {
		return x.BackupAps
	}
	return nil
}

func (x *AccessPointsConfig) GetAvailablePorts() []uint32 {
	if x != nil {
		return x.AvailablePorts
	}
	return nil
}

func (x *AccessPointsConfig) GetForeceLastConnectedAp() *AccessPoint {
	if x != nil {
		return x.ForeceLastConnectedAp
	}
	return nil
}

var File_AccessPointsConfig_proto protoreflect.FileDescriptor

var file_AccessPointsConfig_proto_rawDesc = []byte{
	0x0a, 0x18, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75,
	0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a, 0x11, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x01, 0x0a, 0x12, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x37, 0x0a, 0x0a, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x41, 0x70, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d,
	0x75, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x6f,
	0x70, 0x74, 0x69, 0x6d, 0x61, 0x6c, 0x41, 0x70, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x41, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x41,
	0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x70, 0x73,
	0x12, 0x26, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x4d, 0x0a, 0x15, 0x66, 0x6f, 0x72, 0x65,
	0x63, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44,
	0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x15, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AccessPointsConfig_proto_rawDescOnce sync.Once
	file_AccessPointsConfig_proto_rawDescData = file_AccessPointsConfig_proto_rawDesc
)

func file_AccessPointsConfig_proto_rawDescGZIP() []byte {
	file_AccessPointsConfig_proto_rawDescOnce.Do(func() {
		file_AccessPointsConfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_AccessPointsConfig_proto_rawDescData)
	})
	return file_AccessPointsConfig_proto_rawDescData
}

var file_AccessPointsConfig_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AccessPointsConfig_proto_goTypes = []interface{}{
	(*AccessPointsConfig)(nil), // 0: AcFunDanmu.AccessPointsConfig
	(*AccessPoint)(nil),        // 1: AcFunDanmu.AccessPoint
}
var file_AccessPointsConfig_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.AccessPointsConfig.optimalAps:type_name -> AcFunDanmu.AccessPoint
	1, // 1: AcFunDanmu.AccessPointsConfig.backupAps:type_name -> AcFunDanmu.AccessPoint
	1, // 2: AcFunDanmu.AccessPointsConfig.foreceLastConnectedAp:type_name -> AcFunDanmu.AccessPoint
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_AccessPointsConfig_proto_init() }
func file_AccessPointsConfig_proto_init() {
	if File_AccessPointsConfig_proto != nil {
		return
	}
	file_AccessPoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AccessPointsConfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessPointsConfig); i {
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
			RawDescriptor: file_AccessPointsConfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AccessPointsConfig_proto_goTypes,
		DependencyIndexes: file_AccessPointsConfig_proto_depIdxs,
		MessageInfos:      file_AccessPointsConfig_proto_msgTypes,
	}.Build()
	File_AccessPointsConfig_proto = out.File
	file_AccessPointsConfig_proto_rawDesc = nil
	file_AccessPointsConfig_proto_goTypes = nil
	file_AccessPointsConfig_proto_depIdxs = nil
}