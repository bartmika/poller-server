// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.9.1
// source: proto/telemetry.proto

package tpoller_server

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TelemetryLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TelemetryLabel) Reset() {
	*x = TelemetryLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_telemetry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryLabel) ProtoMessage() {}

func (x *TelemetryLabel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_telemetry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryLabel.ProtoReflect.Descriptor instead.
func (*TelemetryLabel) Descriptor() ([]byte, []int) {
	return file_proto_telemetry_proto_rawDescGZIP(), []int{0}
}

func (x *TelemetryLabel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TelemetryLabel) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TelemetryDatum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric    string               `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	Labels    []*TelemetryLabel    `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	Value     float64              `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TelemetryDatum) Reset() {
	*x = TelemetryDatum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_telemetry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryDatum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryDatum) ProtoMessage() {}

func (x *TelemetryDatum) ProtoReflect() protoreflect.Message {
	mi := &file_proto_telemetry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryDatum.ProtoReflect.Descriptor instead.
func (*TelemetryDatum) Descriptor() ([]byte, []int) {
	return file_proto_telemetry_proto_rawDescGZIP(), []int{1}
}

func (x *TelemetryDatum) GetMetric() string {
	if x != nil {
		return x.Metric
	}
	return ""
}

func (x *TelemetryDatum) GetLabels() []*TelemetryLabel {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TelemetryDatum) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TelemetryDatum) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_proto_telemetry_proto protoreflect.FileDescriptor

var file_proto_telemetry_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0e,
	0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x12, 0x2d, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x32, 0x4f, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12,
	0x42, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x6c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x61, 0x72, 0x74, 0x6d, 0x69, 0x6b, 0x61, 0x2f, 0x74, 0x70, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_telemetry_proto_rawDescOnce sync.Once
	file_proto_telemetry_proto_rawDescData = file_proto_telemetry_proto_rawDesc
)

func file_proto_telemetry_proto_rawDescGZIP() []byte {
	file_proto_telemetry_proto_rawDescOnce.Do(func() {
		file_proto_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_telemetry_proto_rawDescData)
	})
	return file_proto_telemetry_proto_rawDescData
}

var file_proto_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_telemetry_proto_goTypes = []interface{}{
	(*TelemetryLabel)(nil),      // 0: proto.TelemetryLabel
	(*TelemetryDatum)(nil),      // 1: proto.TelemetryDatum
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_proto_telemetry_proto_depIdxs = []int32{
	0, // 0: proto.TelemetryDatum.labels:type_name -> proto.TelemetryLabel
	2, // 1: proto.TelemetryDatum.timestamp:type_name -> google.protobuf.Timestamp
	3, // 2: proto.Telemetry.PollTelemeter:input_type -> google.protobuf.Empty
	1, // 3: proto.Telemetry.PollTelemeter:output_type -> proto.TelemetryDatum
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_telemetry_proto_init() }
func file_proto_telemetry_proto_init() {
	if File_proto_telemetry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_telemetry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryLabel); i {
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
		file_proto_telemetry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryDatum); i {
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
			RawDescriptor: file_proto_telemetry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_telemetry_proto_goTypes,
		DependencyIndexes: file_proto_telemetry_proto_depIdxs,
		MessageInfos:      file_proto_telemetry_proto_msgTypes,
	}.Build()
	File_proto_telemetry_proto = out.File
	file_proto_telemetry_proto_rawDesc = nil
	file_proto_telemetry_proto_goTypes = nil
	file_proto_telemetry_proto_depIdxs = nil
}