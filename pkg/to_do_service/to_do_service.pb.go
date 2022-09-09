// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/to_do_service/to_do_service.proto

package proto_to_do_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ToDo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ToDo) Reset() {
	*x = ToDo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo) ProtoMessage() {}

func (x *ToDo) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo.ProtoReflect.Descriptor instead.
func (*ToDo) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{0}
}

func (x *ToDo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ToDo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ToDo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CreateToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToDo *ToDo `protobuf:"bytes,1,opt,name=to_do,json=toDo,proto3" json:"to_do,omitempty"`
}

func (x *CreateToDoRequest) Reset() {
	*x = CreateToDoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateToDoRequest) ProtoMessage() {}

func (x *CreateToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateToDoRequest.ProtoReflect.Descriptor instead.
func (*CreateToDoRequest) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateToDoRequest) GetToDo() *ToDo {
	if x != nil {
		return x.ToDo
	}
	return nil
}

type CreateToDoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateToDoResponse) Reset() {
	*x = CreateToDoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateToDoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateToDoResponse) ProtoMessage() {}

func (x *CreateToDoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateToDoResponse.ProtoReflect.Descriptor instead.
func (*CreateToDoResponse) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateToDoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllToDoRequest) Reset() {
	*x = GetAllToDoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllToDoRequest) ProtoMessage() {}

func (x *GetAllToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllToDoRequest.ProtoReflect.Descriptor instead.
func (*GetAllToDoRequest) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{3}
}

type GetAllToDoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo []*ToDo `protobuf:"bytes,1,rep,name=todo,proto3" json:"todo,omitempty"`
}

func (x *GetAllToDoResponse) Reset() {
	*x = GetAllToDoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllToDoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllToDoResponse) ProtoMessage() {}

func (x *GetAllToDoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllToDoResponse.ProtoReflect.Descriptor instead.
func (*GetAllToDoResponse) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllToDoResponse) GetTodo() []*ToDo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type UpdateFieldToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text  string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *UpdateFieldToDoRequest) Reset() {
	*x = UpdateFieldToDoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFieldToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFieldToDoRequest) ProtoMessage() {}

func (x *UpdateFieldToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFieldToDoRequest.ProtoReflect.Descriptor instead.
func (*UpdateFieldToDoRequest) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateFieldToDoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFieldToDoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateFieldToDoRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type UpdateFieldToDoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateFieldToDoResponse) Reset() {
	*x = UpdateFieldToDoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFieldToDoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFieldToDoResponse) ProtoMessage() {}

func (x *UpdateFieldToDoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFieldToDoResponse.ProtoReflect.Descriptor instead.
func (*UpdateFieldToDoResponse) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFieldToDoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text  string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *UpdateToDoRequest) Reset() {
	*x = UpdateToDoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateToDoRequest) ProtoMessage() {}

func (x *UpdateToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateToDoRequest.ProtoReflect.Descriptor instead.
func (*UpdateToDoRequest) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateToDoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateToDoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateToDoRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type UpdateToDoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateToDoResponse) Reset() {
	*x = UpdateToDoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateToDoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateToDoResponse) ProtoMessage() {}

func (x *UpdateToDoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateToDoResponse.ProtoReflect.Descriptor instead.
func (*UpdateToDoResponse) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateToDoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteToDoRequest) Reset() {
	*x = DeleteToDoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteToDoRequest) ProtoMessage() {}

func (x *DeleteToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteToDoRequest.ProtoReflect.Descriptor instead.
func (*DeleteToDoRequest) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteToDoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteToDoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteToDoResponse) Reset() {
	*x = DeleteToDoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteToDoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteToDoResponse) ProtoMessage() {}

func (x *DeleteToDoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteToDoResponse.ProtoReflect.Descriptor instead.
func (*DeleteToDoResponse) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteToDoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetToDoByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetToDoByIdRequest) Reset() {
	*x = GetToDoByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetToDoByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetToDoByIdRequest) ProtoMessage() {}

func (x *GetToDoByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetToDoByIdRequest.ProtoReflect.Descriptor instead.
func (*GetToDoByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{11}
}

func (x *GetToDoByIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetToDoByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToDo *ToDo `protobuf:"bytes,1,opt,name=to_do,json=toDo,proto3" json:"to_do,omitempty"`
}

func (x *GetToDoByIdResponse) Reset() {
	*x = GetToDoByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_to_do_service_to_do_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetToDoByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetToDoByIdResponse) ProtoMessage() {}

func (x *GetToDoByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_to_do_service_to_do_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetToDoByIdResponse.ProtoReflect.Descriptor instead.
func (*GetToDoByIdResponse) Descriptor() ([]byte, []int) {
	return file_api_to_do_service_to_do_service_proto_rawDescGZIP(), []int{12}
}

func (x *GetToDoByIdResponse) GetToDo() *ToDo {
	if x != nil {
		return x.ToDo
	}
	return nil
}

var File_api_to_do_service_to_do_service_proto protoreflect.FileDescriptor

var file_api_to_do_service_to_do_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f,
	0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x04, 0x54,
	0x6f, 0x44, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x44, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2f, 0x0a, 0x05, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x44,
	0x6f, 0x22, 0x2e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x44, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04,
	0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x52, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x6f, 0x44, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x33, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
	0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f,
	0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x44, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x46, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x44, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f, 0x64,
	0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x44,
	0x6f, 0x52, 0x04, 0x74, 0x6f, 0x44, 0x6f, 0x32, 0x85, 0x06, 0x0a, 0x0b, 0x54, 0x6f, 0x44, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f, 0x64,
	0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x74, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x27, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x44, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f,
	0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x12, 0x7e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x44, 0x6f,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f, 0x64, 0x6f,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x44, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x44, 0x6f, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x90, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x6f, 0x44, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f,
	0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a, 0x15,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x7b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f,
	0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x32, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x78, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x6f, 0x44, 0x6f, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42,
	0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65,
	0x6b, 0x69, 0x74, 0x76, 0x61, 0x6e, 0x64, 0x2f, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_to_do_service_to_do_service_proto_rawDescOnce sync.Once
	file_api_to_do_service_to_do_service_proto_rawDescData = file_api_to_do_service_to_do_service_proto_rawDesc
)

func file_api_to_do_service_to_do_service_proto_rawDescGZIP() []byte {
	file_api_to_do_service_to_do_service_proto_rawDescOnce.Do(func() {
		file_api_to_do_service_to_do_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_to_do_service_to_do_service_proto_rawDescData)
	})
	return file_api_to_do_service_to_do_service_proto_rawDescData
}

var file_api_to_do_service_to_do_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_to_do_service_to_do_service_proto_goTypes = []interface{}{
	(*ToDo)(nil),                    // 0: api.to_do_service.v1.ToDo
	(*CreateToDoRequest)(nil),       // 1: api.to_do_service.v1.CreateToDoRequest
	(*CreateToDoResponse)(nil),      // 2: api.to_do_service.v1.CreateToDoResponse
	(*GetAllToDoRequest)(nil),       // 3: api.to_do_service.v1.GetAllToDoRequest
	(*GetAllToDoResponse)(nil),      // 4: api.to_do_service.v1.GetAllToDoResponse
	(*UpdateFieldToDoRequest)(nil),  // 5: api.to_do_service.v1.UpdateFieldToDoRequest
	(*UpdateFieldToDoResponse)(nil), // 6: api.to_do_service.v1.UpdateFieldToDoResponse
	(*UpdateToDoRequest)(nil),       // 7: api.to_do_service.v1.UpdateToDoRequest
	(*UpdateToDoResponse)(nil),      // 8: api.to_do_service.v1.UpdateToDoResponse
	(*DeleteToDoRequest)(nil),       // 9: api.to_do_service.v1.DeleteToDoRequest
	(*DeleteToDoResponse)(nil),      // 10: api.to_do_service.v1.DeleteToDoResponse
	(*GetToDoByIdRequest)(nil),      // 11: api.to_do_service.v1.GetToDoByIdRequest
	(*GetToDoByIdResponse)(nil),     // 12: api.to_do_service.v1.GetToDoByIdResponse
}
var file_api_to_do_service_to_do_service_proto_depIdxs = []int32{
	0,  // 0: api.to_do_service.v1.CreateToDoRequest.to_do:type_name -> api.to_do_service.v1.ToDo
	0,  // 1: api.to_do_service.v1.GetAllToDoResponse.todo:type_name -> api.to_do_service.v1.ToDo
	0,  // 2: api.to_do_service.v1.GetToDoByIdResponse.to_do:type_name -> api.to_do_service.v1.ToDo
	1,  // 3: api.to_do_service.v1.ToDoService.CreateToDo:input_type -> api.to_do_service.v1.CreateToDoRequest
	3,  // 4: api.to_do_service.v1.ToDoService.GetAllToDo:input_type -> api.to_do_service.v1.GetAllToDoRequest
	11, // 5: api.to_do_service.v1.ToDoService.GetToDoById:input_type -> api.to_do_service.v1.GetToDoByIdRequest
	5,  // 6: api.to_do_service.v1.ToDoService.UpdateFieldToDo:input_type -> api.to_do_service.v1.UpdateFieldToDoRequest
	7,  // 7: api.to_do_service.v1.ToDoService.UpdateToDo:input_type -> api.to_do_service.v1.UpdateToDoRequest
	9,  // 8: api.to_do_service.v1.ToDoService.DeleteToDo:input_type -> api.to_do_service.v1.DeleteToDoRequest
	2,  // 9: api.to_do_service.v1.ToDoService.CreateToDo:output_type -> api.to_do_service.v1.CreateToDoResponse
	4,  // 10: api.to_do_service.v1.ToDoService.GetAllToDo:output_type -> api.to_do_service.v1.GetAllToDoResponse
	12, // 11: api.to_do_service.v1.ToDoService.GetToDoById:output_type -> api.to_do_service.v1.GetToDoByIdResponse
	6,  // 12: api.to_do_service.v1.ToDoService.UpdateFieldToDo:output_type -> api.to_do_service.v1.UpdateFieldToDoResponse
	8,  // 13: api.to_do_service.v1.ToDoService.UpdateToDo:output_type -> api.to_do_service.v1.UpdateToDoResponse
	10, // 14: api.to_do_service.v1.ToDoService.DeleteToDo:output_type -> api.to_do_service.v1.DeleteToDoResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_to_do_service_to_do_service_proto_init() }
func file_api_to_do_service_to_do_service_proto_init() {
	if File_api_to_do_service_to_do_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_to_do_service_to_do_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDo); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateToDoRequest); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateToDoResponse); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllToDoRequest); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllToDoResponse); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFieldToDoRequest); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFieldToDoResponse); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateToDoRequest); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateToDoResponse); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteToDoRequest); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteToDoResponse); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetToDoByIdRequest); i {
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
		file_api_to_do_service_to_do_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetToDoByIdResponse); i {
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
			RawDescriptor: file_api_to_do_service_to_do_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_to_do_service_to_do_service_proto_goTypes,
		DependencyIndexes: file_api_to_do_service_to_do_service_proto_depIdxs,
		MessageInfos:      file_api_to_do_service_to_do_service_proto_msgTypes,
	}.Build()
	File_api_to_do_service_to_do_service_proto = out.File
	file_api_to_do_service_to_do_service_proto_rawDesc = nil
	file_api_to_do_service_to_do_service_proto_goTypes = nil
	file_api_to_do_service_to_do_service_proto_depIdxs = nil
}