// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: echo-grpc.proto

package grpc

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action       string           `protobuf:"bytes,1,opt,name=Action,proto3" json:"Action,omitempty"`
	Msg          string           `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	MapI64String map[int64]string `protobuf:"bytes,3,rep,name=MapI64String,proto3" json:"MapI64String,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_echo_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_echo_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Request) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Request) GetMapI64String() map[int64]string {
	if x != nil {
		return x.MapI64String
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action       string           `protobuf:"bytes,1,opt,name=Action,proto3" json:"Action,omitempty"`
	Msg          string           `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	MapI64String map[int64]string `protobuf:"bytes,3,rep,name=MapI64String,proto3" json:"MapI64String,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_echo_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_echo_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Response) GetMapI64String() map[int64]string {
	if x != nil {
		return x.MapI64String
	}
	return nil
}

var File_echo_grpc_proto protoreflect.FileDescriptor

var file_echo_grpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x63, 0x68, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0xbd, 0x01, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73,
	0x67, 0x12, 0x47, 0x0a, 0x0c, 0x4d, 0x61, 0x70, 0x49, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x49, 0x36,
	0x34, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x4d, 0x61,
	0x70, 0x49, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x3f, 0x0a, 0x11, 0x4d, 0x61,
	0x70, 0x49, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbf, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x48, 0x0a, 0x0c, 0x4d, 0x61, 0x70, 0x49, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x61, 0x70,
	0x49, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x4d, 0x61, 0x70, 0x49, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x3f, 0x0a, 0x11,
	0x4d, 0x61, 0x70, 0x49, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x37, 0x0a,
	0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x2f, 0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x12, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x65, 0x67, 0x6f, 0x2f, 0x6b,
	0x69, 0x74, 0x65, 0x78, 0x2d, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_echo_grpc_proto_rawDescOnce sync.Once
	file_echo_grpc_proto_rawDescData = file_echo_grpc_proto_rawDesc
)

func file_echo_grpc_proto_rawDescGZIP() []byte {
	file_echo_grpc_proto_rawDescOnce.Do(func() {
		file_echo_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_grpc_proto_rawDescData)
	})
	return file_echo_grpc_proto_rawDescData
}

var file_echo_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_echo_grpc_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: protobuf.Request
	(*Response)(nil), // 1: protobuf.Response
	nil,              // 2: protobuf.Request.MapI64StringEntry
	nil,              // 3: protobuf.Response.MapI64StringEntry
}
var file_echo_grpc_proto_depIdxs = []int32{
	2, // 0: protobuf.Request.MapI64String:type_name -> protobuf.Request.MapI64StringEntry
	3, // 1: protobuf.Response.MapI64String:type_name -> protobuf.Response.MapI64StringEntry
	0, // 2: protobuf.Echo.echo:input_type -> protobuf.Request
	1, // 3: protobuf.Echo.echo:output_type -> protobuf.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_echo_grpc_proto_init() }
func file_echo_grpc_proto_init() {
	if File_echo_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_echo_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_echo_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_grpc_proto_goTypes,
		DependencyIndexes: file_echo_grpc_proto_depIdxs,
		MessageInfos:      file_echo_grpc_proto_msgTypes,
	}.Build()
	File_echo_grpc_proto = out.File
	file_echo_grpc_proto_rawDesc = nil
	file_echo_grpc_proto_goTypes = nil
	file_echo_grpc_proto_depIdxs = nil
}