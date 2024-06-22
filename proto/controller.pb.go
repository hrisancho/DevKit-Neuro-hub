// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.6
// source: controller.proto

package controller

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

type CRCType int32

const (
	CRCType_nothizng CRCType = 0
	CRCType_crc16    CRCType = 1
	CRCType_crc32    CRCType = 2
)

// Enum value maps for CRCType.
var (
	CRCType_name = map[int32]string{
		0: "nothizng",
		1: "crc16",
		2: "crc32",
	}
	CRCType_value = map[string]int32{
		"nothizng": 0,
		"crc16":    1,
		"crc32":    2,
	}
)

func (x CRCType) Enum() *CRCType {
	p := new(CRCType)
	*p = x
	return p
}

func (x CRCType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CRCType) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_proto_enumTypes[0].Descriptor()
}

func (CRCType) Type() protoreflect.EnumType {
	return &file_controller_proto_enumTypes[0]
}

func (x CRCType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CRCType.Descriptor instead.
func (CRCType) EnumDescriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0}
}

type ChannelsDataSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel1 float64 `protobuf:"fixed64,1,opt,name=channel1,proto3" json:"channel1,omitempty"`
	Channel2 float64 `protobuf:"fixed64,2,opt,name=channel2,proto3" json:"channel2,omitempty"`
	Channel3 float64 `protobuf:"fixed64,3,opt,name=channel3,proto3" json:"channel3,omitempty"`
	Channel4 float64 `protobuf:"fixed64,4,opt,name=channel4,proto3" json:"channel4,omitempty"`
	Channel5 float64 `protobuf:"fixed64,5,opt,name=channel5,proto3" json:"channel5,omitempty"`
	Channel6 float64 `protobuf:"fixed64,6,opt,name=channel6,proto3" json:"channel6,omitempty"`
	Channel7 float64 `protobuf:"fixed64,7,opt,name=channel7,proto3" json:"channel7,omitempty"`
	Channel8 float64 `protobuf:"fixed64,8,opt,name=channel8,proto3" json:"channel8,omitempty"`
	Id       uint64  `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChannelsDataSet) Reset() {
	*x = ChannelsDataSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelsDataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelsDataSet) ProtoMessage() {}

func (x *ChannelsDataSet) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelsDataSet.ProtoReflect.Descriptor instead.
func (*ChannelsDataSet) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0}
}

func (x *ChannelsDataSet) GetChannel1() float64 {
	if x != nil {
		return x.Channel1
	}
	return 0
}

func (x *ChannelsDataSet) GetChannel2() float64 {
	if x != nil {
		return x.Channel2
	}
	return 0
}

func (x *ChannelsDataSet) GetChannel3() float64 {
	if x != nil {
		return x.Channel3
	}
	return 0
}

func (x *ChannelsDataSet) GetChannel4() float64 {
	if x != nil {
		return x.Channel4
	}
	return 0
}

func (x *ChannelsDataSet) GetChannel5() float64 {
	if x != nil {
		return x.Channel5
	}
	return 0
}

func (x *ChannelsDataSet) GetChannel6() float64 {
	if x != nil {
		return x.Channel6
	}
	return 0
}

func (x *ChannelsDataSet) GetChannel7() float64 {
	if x != nil {
		return x.Channel7
	}
	return 0
}

func (x *ChannelsDataSet) GetChannel8() float64 {
	if x != nil {
		return x.Channel8
	}
	return 0
}

func (x *ChannelsDataSet) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RawDataPack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataMes []*ChannelsDataSet `protobuf:"bytes,1,rep,name=data_mes,json=dataMes,proto3" json:"data_mes,omitempty"`
}

func (x *RawDataPack) Reset() {
	*x = RawDataPack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawDataPack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawDataPack) ProtoMessage() {}

func (x *RawDataPack) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawDataPack.ProtoReflect.Descriptor instead.
func (*RawDataPack) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{1}
}

func (x *RawDataPack) GetDataMes() []*ChannelsDataSet {
	if x != nil {
		return x.DataMes
	}
	return nil
}

type EmgMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics *ChannelsDataSet `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *EmgMetric) Reset() {
	*x = EmgMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmgMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmgMetric) ProtoMessage() {}

func (x *EmgMetric) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmgMetric.ProtoReflect.Descriptor instead.
func (*EmgMetric) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{2}
}

func (x *EmgMetric) GetMetrics() *ChannelsDataSet {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type MessageWitchCRC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  []byte  `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Kind CRCType `protobuf:"varint,2,opt,name=kind,proto3,enum=controller.CRCType" json:"kind,omitempty"`
	Crc  int64   `protobuf:"varint,3,opt,name=crc,proto3" json:"crc,omitempty"`
}

func (x *MessageWitchCRC) Reset() {
	*x = MessageWitchCRC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWitchCRC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWitchCRC) ProtoMessage() {}

func (x *MessageWitchCRC) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWitchCRC.ProtoReflect.Descriptor instead.
func (*MessageWitchCRC) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{3}
}

func (x *MessageWitchCRC) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *MessageWitchCRC) GetKind() CRCType {
	if x != nil {
		return x.Kind
	}
	return CRCType_nothizng
}

func (x *MessageWitchCRC) GetCrc() int64 {
	if x != nil {
		return x.Crc
	}
	return 0
}

var File_controller_proto protoreflect.FileDescriptor

var file_controller_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0x81,
	0x02, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x31, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x33, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x35, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x35, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x36, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x37, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x38, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x38, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x45, 0x0a, 0x0b, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x63,
	0x6b, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74,
	0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x09, 0x45, 0x6d, 0x67,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x35, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x65, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x5e, 0x0a,
	0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x63, 0x68, 0x43, 0x52, 0x43,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x27, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x52,
	0x43, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x72, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x72, 0x63, 0x2a, 0x2d, 0x0a,
	0x07, 0x43, 0x52, 0x43, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x68,
	0x69, 0x7a, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x63, 0x72, 0x63, 0x31, 0x36, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x63, 0x72, 0x63, 0x33, 0x32, 0x10, 0x02, 0x42, 0x2b, 0x48, 0x03,
	0x5a, 0x24, 0x44, 0x65, 0x76, 0x4b, 0x69, 0x74, 0x2d, 0x4e, 0x65, 0x75, 0x72, 0x6f, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0xf8, 0x01, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_controller_proto_rawDescOnce sync.Once
	file_controller_proto_rawDescData = file_controller_proto_rawDesc
)

func file_controller_proto_rawDescGZIP() []byte {
	file_controller_proto_rawDescOnce.Do(func() {
		file_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_proto_rawDescData)
	})
	return file_controller_proto_rawDescData
}

var file_controller_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_controller_proto_goTypes = []interface{}{
	(CRCType)(0),            // 0: controller.CRCType
	(*ChannelsDataSet)(nil), // 1: controller.ChannelsDataSet
	(*RawDataPack)(nil),     // 2: controller.RawDataPack
	(*EmgMetric)(nil),       // 3: controller.EmgMetric
	(*MessageWitchCRC)(nil), // 4: controller.MessageWitchCRC
}
var file_controller_proto_depIdxs = []int32{
	1, // 0: controller.RawDataPack.data_mes:type_name -> controller.ChannelsDataSet
	1, // 1: controller.EmgMetric.metrics:type_name -> controller.ChannelsDataSet
	0, // 2: controller.MessageWitchCRC.kind:type_name -> controller.CRCType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_controller_proto_init() }
func file_controller_proto_init() {
	if File_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelsDataSet); i {
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
		file_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawDataPack); i {
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
		file_controller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmgMetric); i {
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
		file_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWitchCRC); i {
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
			RawDescriptor: file_controller_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_proto_goTypes,
		DependencyIndexes: file_controller_proto_depIdxs,
		EnumInfos:         file_controller_proto_enumTypes,
		MessageInfos:      file_controller_proto_msgTypes,
	}.Build()
	File_controller_proto = out.File
	file_controller_proto_rawDesc = nil
	file_controller_proto_goTypes = nil
	file_controller_proto_depIdxs = nil
}