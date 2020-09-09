// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: ZtLiveScMessage.proto

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

type ZtLiveScMessage_CompressionType int32

const (
	ZtLiveScMessage_UNKNOWN ZtLiveScMessage_CompressionType = 0
	ZtLiveScMessage_NONE    ZtLiveScMessage_CompressionType = 1
	ZtLiveScMessage_GZIP    ZtLiveScMessage_CompressionType = 2
)

// Enum value maps for ZtLiveScMessage_CompressionType.
var (
	ZtLiveScMessage_CompressionType_name = map[int32]string{
		0: "UNKNOWN",
		1: "NONE",
		2: "GZIP",
	}
	ZtLiveScMessage_CompressionType_value = map[string]int32{
		"UNKNOWN": 0,
		"NONE":    1,
		"GZIP":    2,
	}
)

func (x ZtLiveScMessage_CompressionType) Enum() *ZtLiveScMessage_CompressionType {
	p := new(ZtLiveScMessage_CompressionType)
	*p = x
	return p
}

func (x ZtLiveScMessage_CompressionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZtLiveScMessage_CompressionType) Descriptor() protoreflect.EnumDescriptor {
	return file_ZtLiveScMessage_proto_enumTypes[0].Descriptor()
}

func (ZtLiveScMessage_CompressionType) Type() protoreflect.EnumType {
	return &file_ZtLiveScMessage_proto_enumTypes[0]
}

func (x ZtLiveScMessage_CompressionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZtLiveScMessage_CompressionType.Descriptor instead.
func (ZtLiveScMessage_CompressionType) EnumDescriptor() ([]byte, []int) {
	return file_ZtLiveScMessage_proto_rawDescGZIP(), []int{0, 0}
}

type ZtLiveScMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType       string                          `protobuf:"bytes,1,opt,name=messageType,proto3" json:"messageType,omitempty"`
	CompressionType   ZtLiveScMessage_CompressionType `protobuf:"varint,2,opt,name=compressionType,proto3,enum=AcFunDanmu.ZtLiveScMessage_CompressionType" json:"compressionType,omitempty"`
	Payload           []byte                          `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	LiveId            string                          `protobuf:"bytes,4,opt,name=liveId,proto3" json:"liveId,omitempty"`
	Ticket            string                          `protobuf:"bytes,5,opt,name=ticket,proto3" json:"ticket,omitempty"`
	ServerTimestampMs int64                           `protobuf:"varint,6,opt,name=serverTimestampMs,proto3" json:"serverTimestampMs,omitempty"`
}

func (x *ZtLiveScMessage) Reset() {
	*x = ZtLiveScMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ZtLiveScMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZtLiveScMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZtLiveScMessage) ProtoMessage() {}

func (x *ZtLiveScMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ZtLiveScMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZtLiveScMessage.ProtoReflect.Descriptor instead.
func (*ZtLiveScMessage) Descriptor() ([]byte, []int) {
	return file_ZtLiveScMessage_proto_rawDescGZIP(), []int{0}
}

func (x *ZtLiveScMessage) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *ZtLiveScMessage) GetCompressionType() ZtLiveScMessage_CompressionType {
	if x != nil {
		return x.CompressionType
	}
	return ZtLiveScMessage_UNKNOWN
}

func (x *ZtLiveScMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ZtLiveScMessage) GetLiveId() string {
	if x != nil {
		return x.LiveId
	}
	return ""
}

func (x *ZtLiveScMessage) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

func (x *ZtLiveScMessage) GetServerTimestampMs() int64 {
	if x != nil {
		return x.ServerTimestampMs
	}
	return 0
}

var File_ZtLiveScMessage_proto protoreflect.FileDescriptor

var file_ZtLiveScMessage_proto_rawDesc = []byte{
	0x0a, 0x15, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61,
	0x6e, 0x6d, 0x75, 0x22, 0xb6, 0x02, 0x0a, 0x0f, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x63,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x55, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e,
	0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x69,
	0x76, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x76, 0x65,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x73, 0x22, 0x32, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x5a, 0x49, 0x50, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ZtLiveScMessage_proto_rawDescOnce sync.Once
	file_ZtLiveScMessage_proto_rawDescData = file_ZtLiveScMessage_proto_rawDesc
)

func file_ZtLiveScMessage_proto_rawDescGZIP() []byte {
	file_ZtLiveScMessage_proto_rawDescOnce.Do(func() {
		file_ZtLiveScMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_ZtLiveScMessage_proto_rawDescData)
	})
	return file_ZtLiveScMessage_proto_rawDescData
}

var file_ZtLiveScMessage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ZtLiveScMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ZtLiveScMessage_proto_goTypes = []interface{}{
	(ZtLiveScMessage_CompressionType)(0), // 0: AcFunDanmu.ZtLiveScMessage.CompressionType
	(*ZtLiveScMessage)(nil),              // 1: AcFunDanmu.ZtLiveScMessage
}
var file_ZtLiveScMessage_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.ZtLiveScMessage.compressionType:type_name -> AcFunDanmu.ZtLiveScMessage.CompressionType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ZtLiveScMessage_proto_init() }
func file_ZtLiveScMessage_proto_init() {
	if File_ZtLiveScMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ZtLiveScMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZtLiveScMessage); i {
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
			RawDescriptor: file_ZtLiveScMessage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ZtLiveScMessage_proto_goTypes,
		DependencyIndexes: file_ZtLiveScMessage_proto_depIdxs,
		EnumInfos:         file_ZtLiveScMessage_proto_enumTypes,
		MessageInfos:      file_ZtLiveScMessage_proto_msgTypes,
	}.Build()
	File_ZtLiveScMessage_proto = out.File
	file_ZtLiveScMessage_proto_rawDesc = nil
	file_ZtLiveScMessage_proto_goTypes = nil
	file_ZtLiveScMessage_proto_depIdxs = nil
}