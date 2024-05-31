// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: pkg/pb/chat/chat.proto

package chat

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

type GetFriendChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	OffSet   string `protobuf:"bytes,2,opt,name=OffSet,proto3" json:"OffSet,omitempty"`
	Limit    string `protobuf:"bytes,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	FriendID string `protobuf:"bytes,4,opt,name=FriendID,proto3" json:"FriendID,omitempty"`
}

func (x *GetFriendChatRequest) Reset() {
	*x = GetFriendChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendChatRequest) ProtoMessage() {}

func (x *GetFriendChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendChatRequest.ProtoReflect.Descriptor instead.
func (*GetFriendChatRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *GetFriendChatRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetFriendChatRequest) GetOffSet() string {
	if x != nil {
		return x.OffSet
	}
	return ""
}

func (x *GetFriendChatRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *GetFriendChatRequest) GetFriendID() string {
	if x != nil {
		return x.FriendID
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageID   string `protobuf:"bytes,1,opt,name=MessageID,proto3" json:"MessageID,omitempty"`
	SenderId    string `protobuf:"bytes,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RecipientId string `protobuf:"bytes,3,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	Content     string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp   string `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_chat_chat_proto_msgTypes[1]
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
	return file_pkg_pb_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetMessageID() string {
	if x != nil {
		return x.MessageID
	}
	return ""
}

func (x *Message) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Message) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type GetFriendChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendChat []*Message `protobuf:"bytes,1,rep,name=FriendChat,proto3" json:"FriendChat,omitempty"`
}

func (x *GetFriendChatResponse) Reset() {
	*x = GetFriendChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendChatResponse) ProtoMessage() {}

func (x *GetFriendChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendChatResponse.ProtoReflect.Descriptor instead.
func (*GetFriendChatResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *GetFriendChatResponse) GetFriendChat() []*Message {
	if x != nil {
		return x.FriendChat
	}
	return nil
}

var File_pkg_pb_chat_chat_proto protoreflect.FileDescriptor

var file_pkg_pb_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x78,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x53, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x46, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x68,
	0x61, 0x74, 0x32, 0x59, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_chat_chat_proto_rawDescOnce sync.Once
	file_pkg_pb_chat_chat_proto_rawDescData = file_pkg_pb_chat_chat_proto_rawDesc
)

func file_pkg_pb_chat_chat_proto_rawDescGZIP() []byte {
	file_pkg_pb_chat_chat_proto_rawDescOnce.Do(func() {
		file_pkg_pb_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_chat_chat_proto_rawDescData)
	})
	return file_pkg_pb_chat_chat_proto_rawDescData
}

var file_pkg_pb_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_pb_chat_chat_proto_goTypes = []interface{}{
	(*GetFriendChatRequest)(nil),  // 0: chat.GetFriendChatRequest
	(*Message)(nil),               // 1: chat.Message
	(*GetFriendChatResponse)(nil), // 2: chat.GetFriendChatResponse
}
var file_pkg_pb_chat_chat_proto_depIdxs = []int32{
	1, // 0: chat.GetFriendChatResponse.FriendChat:type_name -> chat.Message
	0, // 1: chat.ChatService.GetFriendChat:input_type -> chat.GetFriendChatRequest
	2, // 2: chat.ChatService.GetFriendChat:output_type -> chat.GetFriendChatResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_pb_chat_chat_proto_init() }
func file_pkg_pb_chat_chat_proto_init() {
	if File_pkg_pb_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendChatRequest); i {
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
		file_pkg_pb_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_pb_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendChatResponse); i {
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
			RawDescriptor: file_pkg_pb_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_chat_chat_proto_goTypes,
		DependencyIndexes: file_pkg_pb_chat_chat_proto_depIdxs,
		MessageInfos:      file_pkg_pb_chat_chat_proto_msgTypes,
	}.Build()
	File_pkg_pb_chat_chat_proto = out.File
	file_pkg_pb_chat_chat_proto_rawDesc = nil
	file_pkg_pb_chat_chat_proto_goTypes = nil
	file_pkg_pb_chat_chat_proto_depIdxs = nil
}
