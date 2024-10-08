// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: protos/messages.proto

package api

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreateTime string `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Value      []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	FileName   string `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileId     string `protobuf:"bytes,6,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_protos_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_protos_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Message) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Message) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Message) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Message) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type ReadUserMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Asc    bool   `protobuf:"varint,4,opt,name=asc,proto3" json:"asc,omitempty"`
}

func (x *ReadUserMessagesRequest) Reset() {
	*x = ReadUserMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadUserMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadUserMessagesRequest) ProtoMessage() {}

func (x *ReadUserMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadUserMessagesRequest.ProtoReflect.Descriptor instead.
func (*ReadUserMessagesRequest) Descriptor() ([]byte, []int) {
	return file_protos_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ReadUserMessagesRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReadUserMessagesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ReadUserMessagesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ReadUserMessagesRequest) GetAsc() bool {
	if x != nil {
		return x.Asc
	}
	return false
}

type ReadUserMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages   []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	IsLastPage bool       `protobuf:"varint,2,opt,name=is_last_page,json=isLastPage,proto3" json:"is_last_page,omitempty"`
}

func (x *ReadUserMessagesResponse) Reset() {
	*x = ReadUserMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadUserMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadUserMessagesResponse) ProtoMessage() {}

func (x *ReadUserMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadUserMessagesResponse.ProtoReflect.Descriptor instead.
func (*ReadUserMessagesResponse) Descriptor() ([]byte, []int) {
	return file_protos_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ReadUserMessagesResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ReadUserMessagesResponse) GetIsLastPage() bool {
	if x != nil {
		return x.IsLastPage
	}
	return false
}

type SaveMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SaveMessageRequest) Reset() {
	*x = SaveMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveMessageRequest) ProtoMessage() {}

func (x *SaveMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveMessageRequest.ProtoReflect.Descriptor instead.
func (*SaveMessageRequest) Descriptor() ([]byte, []int) {
	return file_protos_messages_proto_rawDescGZIP(), []int{3}
}

func (x *SaveMessageRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type SaveMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SaveMessageResponse) Reset() {
	*x = SaveMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveMessageResponse) ProtoMessage() {}

func (x *SaveMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveMessageResponse.ProtoReflect.Descriptor instead.
func (*SaveMessageResponse) Descriptor() ([]byte, []int) {
	return file_protos_messages_proto_rawDescGZIP(), []int{4}
}

func (x *SaveMessageResponse) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type GetServersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetServersRequest) Reset() {
	*x = GetServersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServersRequest) ProtoMessage() {}

func (x *GetServersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServersRequest.ProtoReflect.Descriptor instead.
func (*GetServersRequest) Descriptor() ([]byte, []int) {
	return file_protos_messages_proto_rawDescGZIP(), []int{5}
}

type GetServersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []*Server `protobuf:"bytes,3,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *GetServersResponse) Reset() {
	*x = GetServersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServersResponse) ProtoMessage() {}

func (x *GetServersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServersResponse.ProtoReflect.Descriptor instead.
func (*GetServersResponse) Descriptor() ([]byte, []int) {
	return file_protos_messages_proto_rawDescGZIP(), []int{6}
}

func (x *GetServersResponse) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RaftAddr string `protobuf:"bytes,2,opt,name=raft_addr,json=raftAddr,proto3" json:"raft_addr,omitempty"`
	IsLeader bool   `protobuf:"varint,3,opt,name=is_leader,json=isLeader,proto3" json:"is_leader,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_protos_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_protos_messages_proto_rawDescGZIP(), []int{7}
}

func (x *Server) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Server) GetRaftAddr() string {
	if x != nil {
		return x.RaftAddr
	}
	return ""
}

func (x *Server) GetIsLeader() bool {
	if x != nil {
		return x.IsLeader
	}
	return false
}

var File_protos_messages_proto protoreflect.FileDescriptor

var file_protos_messages_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x17, 0x52, 0x65, 0x61, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x73, 0x63, 0x22, 0x6e, 0x0a, 0x18, 0x52, 0x65,
	0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x44, 0x0a, 0x12, 0x53, 0x61,
	0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x45, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x22, 0x52, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x61, 0x66, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x61, 0x66, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x32, 0x92, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x64, 0x38, 0x37, 0x38, 0x2f, 0x67,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_messages_proto_rawDescOnce sync.Once
	file_protos_messages_proto_rawDescData = file_protos_messages_proto_rawDesc
)

func file_protos_messages_proto_rawDescGZIP() []byte {
	file_protos_messages_proto_rawDescOnce.Do(func() {
		file_protos_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_messages_proto_rawDescData)
	})
	return file_protos_messages_proto_rawDescData
}

var file_protos_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_messages_proto_goTypes = []interface{}{
	(*Message)(nil),                  // 0: messages.v1.Message
	(*ReadUserMessagesRequest)(nil),  // 1: messages.v1.ReadUserMessagesRequest
	(*ReadUserMessagesResponse)(nil), // 2: messages.v1.ReadUserMessagesResponse
	(*SaveMessageRequest)(nil),       // 3: messages.v1.SaveMessageRequest
	(*SaveMessageResponse)(nil),      // 4: messages.v1.SaveMessageResponse
	(*GetServersRequest)(nil),        // 5: messages.v1.GetServersRequest
	(*GetServersResponse)(nil),       // 6: messages.v1.GetServersResponse
	(*Server)(nil),                   // 7: messages.v1.Server
}
var file_protos_messages_proto_depIdxs = []int32{
	0, // 0: messages.v1.ReadUserMessagesResponse.messages:type_name -> messages.v1.Message
	0, // 1: messages.v1.SaveMessageRequest.message:type_name -> messages.v1.Message
	0, // 2: messages.v1.SaveMessageResponse.message:type_name -> messages.v1.Message
	7, // 3: messages.v1.GetServersResponse.servers:type_name -> messages.v1.Server
	5, // 4: messages.v1.Messages.GetServers:input_type -> messages.v1.GetServersRequest
	3, // 5: messages.v1.Messages.SaveMessage:input_type -> messages.v1.SaveMessageRequest
	1, // 6: messages.v1.Messages.ReadUserMessages:input_type -> messages.v1.ReadUserMessagesRequest
	6, // 7: messages.v1.Messages.GetServers:output_type -> messages.v1.GetServersResponse
	4, // 8: messages.v1.Messages.SaveMessage:output_type -> messages.v1.SaveMessageResponse
	2, // 9: messages.v1.Messages.ReadUserMessages:output_type -> messages.v1.ReadUserMessagesResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_messages_proto_init() }
func file_protos_messages_proto_init() {
	if File_protos_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_protos_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadUserMessagesRequest); i {
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
		file_protos_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadUserMessagesResponse); i {
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
		file_protos_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveMessageRequest); i {
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
		file_protos_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveMessageResponse); i {
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
		file_protos_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServersRequest); i {
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
		file_protos_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServersResponse); i {
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
		file_protos_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
			RawDescriptor: file_protos_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_messages_proto_goTypes,
		DependencyIndexes: file_protos_messages_proto_depIdxs,
		MessageInfos:      file_protos_messages_proto_msgTypes,
	}.Build()
	File_protos_messages_proto = out.File
	file_protos_messages_proto_rawDesc = nil
	file_protos_messages_proto_goTypes = nil
	file_protos_messages_proto_depIdxs = nil
}
