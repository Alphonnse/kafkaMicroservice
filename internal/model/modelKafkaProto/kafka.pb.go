// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: kafka.proto

package modelProto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Termometer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic       string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Temperature float32                `protobuf:"fixed32,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Humidity    float32                `protobuf:"fixed32,3,opt,name=humidity,proto3" json:"humidity,omitempty"`
	Location    string                 `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Timestamp   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Termometer) Reset() {
	*x = Termometer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Termometer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Termometer) ProtoMessage() {}

func (x *Termometer) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Termometer.ProtoReflect.Descriptor instead.
func (*Termometer) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{0}
}

func (x *Termometer) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Termometer) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *Termometer) GetHumidity() float32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *Termometer) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Termometer) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_kafka_proto protoreflect.FileDescriptor

var file_kafka_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x0a, 0x54,
	0x65, 0x72, 0x6d, 0x6f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x6c, 0x70, 0x68, 0x6f, 0x6e, 0x6e, 0x73, 0x65, 0x2f, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kafka_proto_rawDescOnce sync.Once
	file_kafka_proto_rawDescData = file_kafka_proto_rawDesc
)

func file_kafka_proto_rawDescGZIP() []byte {
	file_kafka_proto_rawDescOnce.Do(func() {
		file_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(file_kafka_proto_rawDescData)
	})
	return file_kafka_proto_rawDescData
}

var file_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kafka_proto_goTypes = []interface{}{
	(*Termometer)(nil),            // 0: kafkaModel.Termometer
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_kafka_proto_depIdxs = []int32{
	1, // 0: kafkaModel.Termometer.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kafka_proto_init() }
func file_kafka_proto_init() {
	if File_kafka_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kafka_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Termometer); i {
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
			RawDescriptor: file_kafka_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kafka_proto_goTypes,
		DependencyIndexes: file_kafka_proto_depIdxs,
		MessageInfos:      file_kafka_proto_msgTypes,
	}.Build()
	File_kafka_proto = out.File
	file_kafka_proto_rawDesc = nil
	file_kafka_proto_goTypes = nil
	file_kafka_proto_depIdxs = nil
}