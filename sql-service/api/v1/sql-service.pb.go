// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: v1/sql-service.proto

package v1

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

type ConvertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data        string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	ConvertType int32  `protobuf:"varint,2,opt,name=convert_type,json=convertType,proto3" json:"convert_type,omitempty"`
}

func (x *ConvertRequest) Reset() {
	*x = ConvertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sql_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertRequest) ProtoMessage() {}

func (x *ConvertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sql_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertRequest.ProtoReflect.Descriptor instead.
func (*ConvertRequest) Descriptor() ([]byte, []int) {
	return file_v1_sql_service_proto_rawDescGZIP(), []int{0}
}

func (x *ConvertRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *ConvertRequest) GetConvertType() int32 {
	if x != nil {
		return x.ConvertType
	}
	return 0
}

type ConvertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConvertData string `protobuf:"bytes,1,opt,name=convert_data,json=convertData,proto3" json:"convert_data,omitempty"`
}

func (x *ConvertResponse) Reset() {
	*x = ConvertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sql_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertResponse) ProtoMessage() {}

func (x *ConvertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sql_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertResponse.ProtoReflect.Descriptor instead.
func (*ConvertResponse) Descriptor() ([]byte, []int) {
	return file_v1_sql_service_proto_rawDescGZIP(), []int{1}
}

func (x *ConvertResponse) GetConvertData() string {
	if x != nil {
		return x.ConvertData
	}
	return ""
}

var File_v1_sql_service_proto protoreflect.FileDescriptor

var file_v1_sql_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x73, 0x71, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x47,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x34, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x32, 0x44, 0x0a,
	0x03, 0x53, 0x71, 0x6c, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x44,
	0x44, 0x4c, 0x12, 0x16, 0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x71, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x71, 0x6c, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_sql_service_proto_rawDescOnce sync.Once
	file_v1_sql_service_proto_rawDescData = file_v1_sql_service_proto_rawDesc
)

func file_v1_sql_service_proto_rawDescGZIP() []byte {
	file_v1_sql_service_proto_rawDescOnce.Do(func() {
		file_v1_sql_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_sql_service_proto_rawDescData)
	})
	return file_v1_sql_service_proto_rawDescData
}

var file_v1_sql_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_sql_service_proto_goTypes = []interface{}{
	(*ConvertRequest)(nil),  // 0: sql.v1.ConvertRequest
	(*ConvertResponse)(nil), // 1: sql.v1.ConvertResponse
}
var file_v1_sql_service_proto_depIdxs = []int32{
	0, // 0: sql.v1.Sql.ConvertDDL:input_type -> sql.v1.ConvertRequest
	1, // 1: sql.v1.Sql.ConvertDDL:output_type -> sql.v1.ConvertResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_sql_service_proto_init() }
func file_v1_sql_service_proto_init() {
	if File_v1_sql_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_sql_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertRequest); i {
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
		file_v1_sql_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertResponse); i {
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
			RawDescriptor: file_v1_sql_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_sql_service_proto_goTypes,
		DependencyIndexes: file_v1_sql_service_proto_depIdxs,
		MessageInfos:      file_v1_sql_service_proto_msgTypes,
	}.Build()
	File_v1_sql_service_proto = out.File
	file_v1_sql_service_proto_rawDesc = nil
	file_v1_sql_service_proto_goTypes = nil
	file_v1_sql_service_proto_depIdxs = nil
}