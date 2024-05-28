// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: sso/channel_management.proto

package sso

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

type CreateChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId   string `protobuf:"bytes,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
}

func (x *CreateChannelRequest) Reset() {
	*x = CreateChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelRequest) ProtoMessage() {}

func (x *CreateChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelRequest.ProtoReflect.Descriptor instead.
func (*CreateChannelRequest) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{0}
}

func (x *CreateChannelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateChannelRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateChannelRequest) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

type CreateChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *CreateChannelResponse) Reset() {
	*x = CreateChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelResponse) ProtoMessage() {}

func (x *CreateChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelResponse.ProtoReflect.Descriptor instead.
func (*CreateChannelResponse) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChannelResponse) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type DeleteChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *DeleteChannelRequest) Reset() {
	*x = DeleteChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChannelRequest) ProtoMessage() {}

func (x *DeleteChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChannelRequest.ProtoReflect.Descriptor instead.
func (*DeleteChannelRequest) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteChannelRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type DeleteChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteChannelResponse) Reset() {
	*x = DeleteChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChannelResponse) ProtoMessage() {}

func (x *DeleteChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChannelResponse.ProtoReflect.Descriptor instead.
func (*DeleteChannelResponse) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{3}
}

type UpdateChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId   string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateChannelRequest) Reset() {
	*x = UpdateChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChannelRequest) ProtoMessage() {}

func (x *UpdateChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChannelRequest.ProtoReflect.Descriptor instead.
func (*UpdateChannelRequest) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateChannelRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *UpdateChannelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateChannelRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateChannelResponse) Reset() {
	*x = UpdateChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChannelResponse) ProtoMessage() {}

func (x *UpdateChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChannelResponse.ProtoReflect.Descriptor instead.
func (*UpdateChannelResponse) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{5}
}

type JoinChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *JoinChannelRequest) Reset() {
	*x = JoinChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinChannelRequest) ProtoMessage() {}

func (x *JoinChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinChannelRequest.ProtoReflect.Descriptor instead.
func (*JoinChannelRequest) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{6}
}

func (x *JoinChannelRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *JoinChannelRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type JoinChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinChannelResponse) Reset() {
	*x = JoinChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinChannelResponse) ProtoMessage() {}

func (x *JoinChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinChannelResponse.ProtoReflect.Descriptor instead.
func (*JoinChannelResponse) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{7}
}

type KickUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *KickUserRequest) Reset() {
	*x = KickUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickUserRequest) ProtoMessage() {}

func (x *KickUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickUserRequest.ProtoReflect.Descriptor instead.
func (*KickUserRequest) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{8}
}

func (x *KickUserRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *KickUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type KickUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KickUserResponse) Reset() {
	*x = KickUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickUserResponse) ProtoMessage() {}

func (x *KickUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickUserResponse.ProtoReflect.Descriptor instead.
func (*KickUserResponse) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{9}
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	WithMedia bool   `protobuf:"varint,4,opt,name=with_media,json=withMedia,proto3" json:"with_media,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{10}
}

func (x *SendMessageRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SendMessageRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SendMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendMessageRequest) GetWithMedia() bool {
	if x != nil {
		return x.WithMedia
	}
	return false
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{11}
}

type SendMediaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Media     []byte `protobuf:"bytes,3,opt,name=media,proto3" json:"media,omitempty"`
	MediaType string `protobuf:"bytes,4,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
}

func (x *SendMediaRequest) Reset() {
	*x = SendMediaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMediaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMediaRequest) ProtoMessage() {}

func (x *SendMediaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMediaRequest.ProtoReflect.Descriptor instead.
func (*SendMediaRequest) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{12}
}

func (x *SendMediaRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SendMediaRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SendMediaRequest) GetMedia() []byte {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *SendMediaRequest) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

type SendMediaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMediaResponse) Reset() {
	*x = SendMediaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_channel_management_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMediaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMediaResponse) ProtoMessage() {}

func (x *SendMediaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_channel_management_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMediaResponse.ProtoReflect.Descriptor instead.
func (*SendMediaResponse) Descriptor() ([]byte, []int) {
	return file_sso_channel_management_proto_rawDescGZIP(), []int{13}
}

var File_sso_channel_management_proto protoreflect.FileDescriptor

var file_sso_channel_management_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x73, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x73, 0x73, 0x6f, 0x22, 0x6b, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x22, 0x36, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22,
	0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c,
	0x0a, 0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13,
	0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x49, 0x0a, 0x0f, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x12,
	0x0a, 0x10, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x77, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x7f, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf2, 0x03, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x48, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x19,
	0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x73, 0x6f, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x19, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x19, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x73, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x4a,
	0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x17, 0x2e, 0x73, 0x73, 0x6f,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x08, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x73, 0x73,
	0x6f, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x73, 0x6f, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x15, 0x2e, 0x73, 0x73,
	0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x73, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sso_channel_management_proto_rawDescOnce sync.Once
	file_sso_channel_management_proto_rawDescData = file_sso_channel_management_proto_rawDesc
)

func file_sso_channel_management_proto_rawDescGZIP() []byte {
	file_sso_channel_management_proto_rawDescOnce.Do(func() {
		file_sso_channel_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_channel_management_proto_rawDescData)
	})
	return file_sso_channel_management_proto_rawDescData
}

var file_sso_channel_management_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_sso_channel_management_proto_goTypes = []interface{}{
	(*CreateChannelRequest)(nil),  // 0: sso.CreateChannelRequest
	(*CreateChannelResponse)(nil), // 1: sso.CreateChannelResponse
	(*DeleteChannelRequest)(nil),  // 2: sso.DeleteChannelRequest
	(*DeleteChannelResponse)(nil), // 3: sso.DeleteChannelResponse
	(*UpdateChannelRequest)(nil),  // 4: sso.UpdateChannelRequest
	(*UpdateChannelResponse)(nil), // 5: sso.UpdateChannelResponse
	(*JoinChannelRequest)(nil),    // 6: sso.JoinChannelRequest
	(*JoinChannelResponse)(nil),   // 7: sso.JoinChannelResponse
	(*KickUserRequest)(nil),       // 8: sso.KickUserRequest
	(*KickUserResponse)(nil),      // 9: sso.KickUserResponse
	(*SendMessageRequest)(nil),    // 10: sso.SendMessageRequest
	(*SendMessageResponse)(nil),   // 11: sso.SendMessageResponse
	(*SendMediaRequest)(nil),      // 12: sso.SendMediaRequest
	(*SendMediaResponse)(nil),     // 13: sso.SendMediaResponse
}
var file_sso_channel_management_proto_depIdxs = []int32{
	0,  // 0: sso.ChannelManagement.CreateChannel:input_type -> sso.CreateChannelRequest
	2,  // 1: sso.ChannelManagement.DeleteChannel:input_type -> sso.DeleteChannelRequest
	4,  // 2: sso.ChannelManagement.UpdateChannel:input_type -> sso.UpdateChannelRequest
	6,  // 3: sso.ChannelManagement.JoinChannel:input_type -> sso.JoinChannelRequest
	8,  // 4: sso.ChannelManagement.KickUser:input_type -> sso.KickUserRequest
	10, // 5: sso.ChannelManagement.SendMessage:input_type -> sso.SendMessageRequest
	12, // 6: sso.ChannelManagement.SendMedia:input_type -> sso.SendMediaRequest
	1,  // 7: sso.ChannelManagement.CreateChannel:output_type -> sso.CreateChannelResponse
	3,  // 8: sso.ChannelManagement.DeleteChannel:output_type -> sso.DeleteChannelResponse
	5,  // 9: sso.ChannelManagement.UpdateChannel:output_type -> sso.UpdateChannelResponse
	7,  // 10: sso.ChannelManagement.JoinChannel:output_type -> sso.JoinChannelResponse
	9,  // 11: sso.ChannelManagement.KickUser:output_type -> sso.KickUserResponse
	11, // 12: sso.ChannelManagement.SendMessage:output_type -> sso.SendMessageResponse
	13, // 13: sso.ChannelManagement.SendMedia:output_type -> sso.SendMediaResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_sso_channel_management_proto_init() }
func file_sso_channel_management_proto_init() {
	if File_sso_channel_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sso_channel_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChannelRequest); i {
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
		file_sso_channel_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChannelResponse); i {
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
		file_sso_channel_management_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChannelRequest); i {
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
		file_sso_channel_management_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChannelResponse); i {
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
		file_sso_channel_management_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChannelRequest); i {
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
		file_sso_channel_management_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChannelResponse); i {
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
		file_sso_channel_management_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinChannelRequest); i {
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
		file_sso_channel_management_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinChannelResponse); i {
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
		file_sso_channel_management_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickUserRequest); i {
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
		file_sso_channel_management_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickUserResponse); i {
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
		file_sso_channel_management_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_sso_channel_management_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_sso_channel_management_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMediaRequest); i {
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
		file_sso_channel_management_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMediaResponse); i {
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
			RawDescriptor: file_sso_channel_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_channel_management_proto_goTypes,
		DependencyIndexes: file_sso_channel_management_proto_depIdxs,
		MessageInfos:      file_sso_channel_management_proto_msgTypes,
	}.Build()
	File_sso_channel_management_proto = out.File
	file_sso_channel_management_proto_rawDesc = nil
	file_sso_channel_management_proto_goTypes = nil
	file_sso_channel_management_proto_depIdxs = nil
}