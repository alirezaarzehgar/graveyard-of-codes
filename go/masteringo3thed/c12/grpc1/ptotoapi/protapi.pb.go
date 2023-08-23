// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: protapi.proto

package protoapi

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

type NumberArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A float32 `protobuf:"fixed32,1,opt,name=A,proto3" json:"A,omitempty"`
	B float32 `protobuf:"fixed32,2,opt,name=B,proto3" json:"B,omitempty"`
}

func (x *NumberArgs) Reset() {
	*x = NumberArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberArgs) ProtoMessage() {}

func (x *NumberArgs) ProtoReflect() protoreflect.Message {
	mi := &file_protapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberArgs.ProtoReflect.Descriptor instead.
func (*NumberArgs) Descriptor() ([]byte, []int) {
	return file_protapi_proto_rawDescGZIP(), []int{0}
}

func (x *NumberArgs) GetA() float32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *NumberArgs) GetB() float32 {
	if x != nil {
		return x.B
	}
	return 0
}

type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float32 `protobuf:"fixed32,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_protapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_protapi_proto_rawDescGZIP(), []int{1}
}

func (x *Number) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_protapi_proto protoreflect.FileDescriptor

var file_protapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x28, 0x0a, 0x0a, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x12, 0x0c, 0x0a,
	0x01, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x41, 0x12, 0x0c, 0x0a, 0x01, 0x42,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x42, 0x22, 0x1e, 0x0a, 0x06, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x7c, 0x0a, 0x04, 0x4d, 0x61, 0x74,
	0x68, 0x12, 0x1b, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0b, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x07, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x05, 0x4d, 0x69, 0x6e, 0x75, 0x73, 0x12, 0x0b, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x41, 0x72, 0x67, 0x73, 0x1a, 0x07, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x03, 0x44, 0x69, 0x76, 0x12, 0x0b, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x72, 0x67,
	0x73, 0x1a, 0x07, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x03, 0x4d, 0x75,
	0x6c, 0x12, 0x0b, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x07,
	0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protapi_proto_rawDescOnce sync.Once
	file_protapi_proto_rawDescData = file_protapi_proto_rawDesc
)

func file_protapi_proto_rawDescGZIP() []byte {
	file_protapi_proto_rawDescOnce.Do(func() {
		file_protapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_protapi_proto_rawDescData)
	})
	return file_protapi_proto_rawDescData
}

var file_protapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protapi_proto_goTypes = []interface{}{
	(*NumberArgs)(nil), // 0: NumberArgs
	(*Number)(nil),     // 1: Number
}
var file_protapi_proto_depIdxs = []int32{
	0, // 0: Math.Add:input_type -> NumberArgs
	0, // 1: Math.Minus:input_type -> NumberArgs
	0, // 2: Math.Div:input_type -> NumberArgs
	0, // 3: Math.Mul:input_type -> NumberArgs
	1, // 4: Math.Add:output_type -> Number
	1, // 5: Math.Minus:output_type -> Number
	1, // 6: Math.Div:output_type -> Number
	1, // 7: Math.Mul:output_type -> Number
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protapi_proto_init() }
func file_protapi_proto_init() {
	if File_protapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberArgs); i {
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
		file_protapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
			RawDescriptor: file_protapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protapi_proto_goTypes,
		DependencyIndexes: file_protapi_proto_depIdxs,
		MessageInfos:      file_protapi_proto_msgTypes,
	}.Build()
	File_protapi_proto = out.File
	file_protapi_proto_rawDesc = nil
	file_protapi_proto_goTypes = nil
	file_protapi_proto_depIdxs = nil
}
