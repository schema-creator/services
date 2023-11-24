// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: project-service/v1/project-service.proto

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

type CreateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateProjectRequest) Reset() {
	*x = CreateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_service_v1_project_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRequest) ProtoMessage() {}

func (x *CreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_service_v1_project_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_project_service_v1_project_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProjectRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	LastImage  string `protobuf:"bytes,3,opt,name=last_image,json=lastImage,proto3" json:"last_image,omitempty"`
	IsPersonal bool   `protobuf:"varint,4,opt,name=is_personal,json=isPersonal,proto3" json:"is_personal,omitempty"`
}

func (x *CreateProjectResponse) Reset() {
	*x = CreateProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_service_v1_project_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectResponse) ProtoMessage() {}

func (x *CreateProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_service_v1_project_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectResponse.ProtoReflect.Descriptor instead.
func (*CreateProjectResponse) Descriptor() ([]byte, []int) {
	return file_project_service_v1_project_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProjectResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProjectResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateProjectResponse) GetLastImage() string {
	if x != nil {
		return x.LastImage
	}
	return ""
}

func (x *CreateProjectResponse) GetIsPersonal() bool {
	if x != nil {
		return x.IsPersonal
	}
	return false
}

var File_project_service_v1_project_service_proto protoreflect.FileDescriptor

var file_project_service_v1_project_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0x2c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x7d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x32, 0x60, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_service_v1_project_service_proto_rawDescOnce sync.Once
	file_project_service_v1_project_service_proto_rawDescData = file_project_service_v1_project_service_proto_rawDesc
)

func file_project_service_v1_project_service_proto_rawDescGZIP() []byte {
	file_project_service_v1_project_service_proto_rawDescOnce.Do(func() {
		file_project_service_v1_project_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_service_v1_project_service_proto_rawDescData)
	})
	return file_project_service_v1_project_service_proto_rawDescData
}

var file_project_service_v1_project_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_project_service_v1_project_service_proto_goTypes = []interface{}{
	(*CreateProjectRequest)(nil),  // 0: project.CreateProjectRequest
	(*CreateProjectResponse)(nil), // 1: project.CreateProjectResponse
}
var file_project_service_v1_project_service_proto_depIdxs = []int32{
	0, // 0: project.ProjectService.CreateProject:input_type -> project.CreateProjectRequest
	1, // 1: project.ProjectService.CreateProject:output_type -> project.CreateProjectResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_service_v1_project_service_proto_init() }
func file_project_service_v1_project_service_proto_init() {
	if File_project_service_v1_project_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_service_v1_project_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectRequest); i {
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
		file_project_service_v1_project_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectResponse); i {
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
			RawDescriptor: file_project_service_v1_project_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_service_v1_project_service_proto_goTypes,
		DependencyIndexes: file_project_service_v1_project_service_proto_depIdxs,
		MessageInfos:      file_project_service_v1_project_service_proto_msgTypes,
	}.Build()
	File_project_service_v1_project_service_proto = out.File
	file_project_service_v1_project_service_proto_rawDesc = nil
	file_project_service_v1_project_service_proto_goTypes = nil
	file_project_service_v1_project_service_proto_depIdxs = nil
}
