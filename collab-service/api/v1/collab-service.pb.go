// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: v1/collab-service.proto

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

type CreateEditorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateEditorRequest) Reset() {
	*x = CreateEditorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_collab_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEditorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEditorRequest) ProtoMessage() {}

func (x *CreateEditorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_collab_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEditorRequest.ProtoReflect.Descriptor instead.
func (*CreateEditorRequest) Descriptor() ([]byte, []int) {
	return file_v1_collab_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEditorRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateEditorRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CreateEditorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateEditorResponse) Reset() {
	*x = CreateEditorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_collab_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEditorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEditorResponse) ProtoMessage() {}

func (x *CreateEditorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_collab_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEditorResponse.ProtoReflect.Descriptor instead.
func (*CreateEditorResponse) Descriptor() ([]byte, []int) {
	return file_v1_collab_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEditorResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetEditorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetEditorRequest) Reset() {
	*x = GetEditorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_collab_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEditorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEditorRequest) ProtoMessage() {}

func (x *GetEditorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_collab_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEditorRequest.ProtoReflect.Descriptor instead.
func (*GetEditorRequest) Descriptor() ([]byte, []int) {
	return file_v1_collab_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetEditorRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetEditorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetEditorResponse) Reset() {
	*x = GetEditorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_collab_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEditorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEditorResponse) ProtoMessage() {}

func (x *GetEditorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_collab_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEditorResponse.ProtoReflect.Descriptor instead.
func (*GetEditorResponse) Descriptor() ([]byte, []int) {
	return file_v1_collab_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetEditorResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type UpdateEditorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *UpdateEditorRequest) Reset() {
	*x = UpdateEditorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_collab_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEditorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEditorRequest) ProtoMessage() {}

func (x *UpdateEditorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_collab_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEditorRequest.ProtoReflect.Descriptor instead.
func (*UpdateEditorRequest) Descriptor() ([]byte, []int) {
	return file_v1_collab_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEditorRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *UpdateEditorRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type UpdateEditorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *UpdateEditorResponse) Reset() {
	*x = UpdateEditorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_collab_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEditorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEditorResponse) ProtoMessage() {}

func (x *UpdateEditorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_collab_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEditorResponse.ProtoReflect.Descriptor instead.
func (*UpdateEditorResponse) Descriptor() ([]byte, []int) {
	return file_v1_collab_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEditorResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type DeleteEditorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *DeleteEditorRequest) Reset() {
	*x = DeleteEditorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_collab_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEditorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEditorRequest) ProtoMessage() {}

func (x *DeleteEditorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_collab_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEditorRequest.ProtoReflect.Descriptor instead.
func (*DeleteEditorRequest) Descriptor() ([]byte, []int) {
	return file_v1_collab_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteEditorRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type DeleteEditorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *DeleteEditorResponse) Reset() {
	*x = DeleteEditorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_collab_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEditorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEditorResponse) ProtoMessage() {}

func (x *DeleteEditorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_collab_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEditorResponse.ProtoReflect.Descriptor instead.
func (*DeleteEditorResponse) Descriptor() ([]byte, []int) {
	return file_v1_collab_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEditorResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

var File_v1_collab_service_proto protoreflect.FileDescriptor

var file_v1_collab_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x76, 0x31, 0x22, 0x48, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x2a, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x31, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x48, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x35, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x22, 0x34, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x32, 0xc1,
	0x02, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x12,
	0x1d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64,
	0x69, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x12, 0x1d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_collab_service_proto_rawDescOnce sync.Once
	file_v1_collab_service_proto_rawDescData = file_v1_collab_service_proto_rawDesc
)

func file_v1_collab_service_proto_rawDescGZIP() []byte {
	file_v1_collab_service_proto_rawDescOnce.Do(func() {
		file_v1_collab_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_collab_service_proto_rawDescData)
	})
	return file_v1_collab_service_proto_rawDescData
}

var file_v1_collab_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_collab_service_proto_goTypes = []interface{}{
	(*CreateEditorRequest)(nil),  // 0: token.v1.CreateEditorRequest
	(*CreateEditorResponse)(nil), // 1: token.v1.CreateEditorResponse
	(*GetEditorRequest)(nil),     // 2: token.v1.GetEditorRequest
	(*GetEditorResponse)(nil),    // 3: token.v1.GetEditorResponse
	(*UpdateEditorRequest)(nil),  // 4: token.v1.UpdateEditorRequest
	(*UpdateEditorResponse)(nil), // 5: token.v1.UpdateEditorResponse
	(*DeleteEditorRequest)(nil),  // 6: token.v1.DeleteEditorRequest
	(*DeleteEditorResponse)(nil), // 7: token.v1.DeleteEditorResponse
}
var file_v1_collab_service_proto_depIdxs = []int32{
	0, // 0: token.v1.TokenService.CreateEditor:input_type -> token.v1.CreateEditorRequest
	2, // 1: token.v1.TokenService.GetEditor:input_type -> token.v1.GetEditorRequest
	4, // 2: token.v1.TokenService.UpdateEditor:input_type -> token.v1.UpdateEditorRequest
	6, // 3: token.v1.TokenService.DeleteEditor:input_type -> token.v1.DeleteEditorRequest
	1, // 4: token.v1.TokenService.CreateEditor:output_type -> token.v1.CreateEditorResponse
	3, // 5: token.v1.TokenService.GetEditor:output_type -> token.v1.GetEditorResponse
	5, // 6: token.v1.TokenService.UpdateEditor:output_type -> token.v1.UpdateEditorResponse
	7, // 7: token.v1.TokenService.DeleteEditor:output_type -> token.v1.DeleteEditorResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_collab_service_proto_init() }
func file_v1_collab_service_proto_init() {
	if File_v1_collab_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_collab_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEditorRequest); i {
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
		file_v1_collab_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEditorResponse); i {
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
		file_v1_collab_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEditorRequest); i {
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
		file_v1_collab_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEditorResponse); i {
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
		file_v1_collab_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEditorRequest); i {
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
		file_v1_collab_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEditorResponse); i {
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
		file_v1_collab_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEditorRequest); i {
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
		file_v1_collab_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEditorResponse); i {
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
			RawDescriptor: file_v1_collab_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_collab_service_proto_goTypes,
		DependencyIndexes: file_v1_collab_service_proto_depIdxs,
		MessageInfos:      file_v1_collab_service_proto_msgTypes,
	}.Build()
	File_v1_collab_service_proto = out.File
	file_v1_collab_service_proto_rawDesc = nil
	file_v1_collab_service_proto_goTypes = nil
	file_v1_collab_service_proto_depIdxs = nil
}
