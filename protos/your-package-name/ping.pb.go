// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: ping.proto

package your_package_name

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

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_ping_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_ping_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_ping_proto protoreflect.FileDescriptor

var file_ping_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x32, 0xe9, 0x01, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x32, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x11, 0x2e,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x1a, 0x11, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x50,
	0x6f, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x6f,
	0x6e, 0x67, 0x12, 0x11, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x11, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x28, 0x01, 0x12, 0x35, 0x0a, 0x09,
	0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x73, 0x12, 0x11, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x11, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x0a, 0x50, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x6f, 0x6e, 0x67,
	0x73, 0x12, 0x11, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x1a, 0x11, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x15, 0x5a,
	0x13, 0x2e, 0x2f, 0x79, 0x6f, 0x75, 0x72, 0x2d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2d,
	0x6e, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ping_proto_rawDescOnce sync.Once
	file_ping_proto_rawDescData = file_ping_proto_rawDesc
)

func file_ping_proto_rawDescGZIP() []byte {
	file_ping_proto_rawDescOnce.Do(func() {
		file_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_ping_proto_rawDescData)
	})
	return file_ping_proto_rawDescData
}

var file_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ping_proto_goTypes = []interface{}{
	(*Ping)(nil), // 0: healthCheck.Ping
	(*Pong)(nil), // 1: healthCheck.Pong
}
var file_ping_proto_depIdxs = []int32{
	0, // 0: healthCheck.HealthCheck.PingPong:input_type -> healthCheck.Ping
	0, // 1: healthCheck.HealthCheck.PingsPong:input_type -> healthCheck.Ping
	0, // 2: healthCheck.HealthCheck.PingPongs:input_type -> healthCheck.Ping
	0, // 3: healthCheck.HealthCheck.PingsPongs:input_type -> healthCheck.Ping
	1, // 4: healthCheck.HealthCheck.PingPong:output_type -> healthCheck.Pong
	1, // 5: healthCheck.HealthCheck.PingsPong:output_type -> healthCheck.Pong
	1, // 6: healthCheck.HealthCheck.PingPongs:output_type -> healthCheck.Pong
	1, // 7: healthCheck.HealthCheck.PingsPongs:output_type -> healthCheck.Pong
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ping_proto_init() }
func file_ping_proto_init() {
	if File_ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
			RawDescriptor: file_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ping_proto_goTypes,
		DependencyIndexes: file_ping_proto_depIdxs,
		MessageInfos:      file_ping_proto_msgTypes,
	}.Build()
	File_ping_proto = out.File
	file_ping_proto_rawDesc = nil
	file_ping_proto_goTypes = nil
	file_ping_proto_depIdxs = nil
}
