// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/event.proto

package event

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

type EventInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Senderid int64  `protobuf:"varint,1,opt,name=senderid,proto3" json:"senderid,omitempty"`
	Eventid  int64  `protobuf:"varint,2,opt,name=eventid,proto3" json:"eventid,omitempty"`
	Time     int64  `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *EventInfo) Reset() {
	*x = EventInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventInfo) ProtoMessage() {}

func (x *EventInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventInfo.ProtoReflect.Descriptor instead.
func (*EventInfo) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventInfo) GetSenderid() int64 {
	if x != nil {
		return x.Senderid
	}
	return 0
}

func (x *EventInfo) GetEventid() int64 {
	if x != nil {
		return x.Eventid
	}
	return 0
}

func (x *EventInfo) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *EventInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MakeEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Senderid int64  `protobuf:"varint,1,opt,name=senderid,proto3" json:"senderid,omitempty"`
	Time     int64  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MakeEventRequest) Reset() {
	*x = MakeEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeEventRequest) ProtoMessage() {}

func (x *MakeEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeEventRequest.ProtoReflect.Descriptor instead.
func (*MakeEventRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{1}
}

func (x *MakeEventRequest) GetSenderid() int64 {
	if x != nil {
		return x.Senderid
	}
	return 0
}

func (x *MakeEventRequest) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *MakeEventRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MakeEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eventid int64 `protobuf:"varint,1,opt,name=eventid,proto3" json:"eventid,omitempty"`
}

func (x *MakeEventResponse) Reset() {
	*x = MakeEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeEventResponse) ProtoMessage() {}

func (x *MakeEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeEventResponse.ProtoReflect.Descriptor instead.
func (*MakeEventResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{2}
}

func (x *MakeEventResponse) GetEventid() int64 {
	if x != nil {
		return x.Eventid
	}
	return 0
}

type GetEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Senderid int64 `protobuf:"varint,1,opt,name=senderid,proto3" json:"senderid,omitempty"`
	Eventid  int64 `protobuf:"varint,2,opt,name=eventid,proto3" json:"eventid,omitempty"`
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{3}
}

func (x *GetEventRequest) GetSenderid() int64 {
	if x != nil {
		return x.Senderid
	}
	return 0
}

func (x *GetEventRequest) GetEventid() int64 {
	if x != nil {
		return x.Eventid
	}
	return 0
}

type GetEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventInfo *EventInfo `protobuf:"bytes,1,opt,name=eventInfo,proto3" json:"eventInfo,omitempty"`
}

func (x *GetEventResponse) Reset() {
	*x = GetEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventResponse) ProtoMessage() {}

func (x *GetEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventResponse.ProtoReflect.Descriptor instead.
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{4}
}

func (x *GetEventResponse) GetEventInfo() *EventInfo {
	if x != nil {
		return x.EventInfo
	}
	return nil
}

type DeleteEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Senderid int64 `protobuf:"varint,1,opt,name=senderid,proto3" json:"senderid,omitempty"`
	Eventid  int64 `protobuf:"varint,2,opt,name=eventid,proto3" json:"eventid,omitempty"`
}

func (x *DeleteEventRequest) Reset() {
	*x = DeleteEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventRequest) ProtoMessage() {}

func (x *DeleteEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventRequest.ProtoReflect.Descriptor instead.
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteEventRequest) GetSenderid() int64 {
	if x != nil {
		return x.Senderid
	}
	return 0
}

func (x *DeleteEventRequest) GetEventid() int64 {
	if x != nil {
		return x.Eventid
	}
	return 0
}

type DeleteEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Senderid int64 `protobuf:"varint,1,opt,name=senderid,proto3" json:"senderid,omitempty"`
}

func (x *DeleteEventResponse) Reset() {
	*x = DeleteEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventResponse) ProtoMessage() {}

func (x *DeleteEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventResponse.ProtoReflect.Descriptor instead.
func (*DeleteEventResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteEventResponse) GetSenderid() int64 {
	if x != nil {
		return x.Senderid
	}
	return 0
}

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Senderid int64 `protobuf:"varint,1,opt,name=senderid,proto3" json:"senderid,omitempty"`
	Time     int64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{7}
}

func (x *GetEventsRequest) GetSenderid() int64 {
	if x != nil {
		return x.Senderid
	}
	return 0
}

func (x *GetEventsRequest) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type GetEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventInfo []*EventInfo `protobuf:"bytes,1,rep,name=eventInfo,proto3" json:"eventInfo,omitempty"`
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{8}
}

func (x *GetEventsResponse) GetEventInfo() []*EventInfo {
	if x != nil {
		return x.EventInfo
	}
	return nil
}

type GetEventsStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Senderid int64 `protobuf:"varint,1,opt,name=senderid,proto3" json:"senderid,omitempty"`
	Time     int64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GetEventsStreamRequest) Reset() {
	*x = GetEventsStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsStreamRequest) ProtoMessage() {}

func (x *GetEventsStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsStreamRequest.ProtoReflect.Descriptor instead.
func (*GetEventsStreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{9}
}

func (x *GetEventsStreamRequest) GetSenderid() int64 {
	if x != nil {
		return x.Senderid
	}
	return 0
}

func (x *GetEventsStreamRequest) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type GetEventsStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventInfo *EventInfo `protobuf:"bytes,1,opt,name=EventInfo,proto3" json:"EventInfo,omitempty"`
}

func (x *GetEventsStreamResponse) Reset() {
	*x = GetEventsStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsStreamResponse) ProtoMessage() {}

func (x *GetEventsStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsStreamResponse.ProtoReflect.Descriptor instead.
func (*GetEventsStreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{10}
}

func (x *GetEventsStreamResponse) GetEventInfo() *EventInfo {
	if x != nil {
		return x.EventInfo
	}
	return nil
}

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{11}
}

type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{12}
}

func (x *TestResponse) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

var File_proto_event_proto protoreflect.FileDescriptor

var file_proto_event_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x69,
	0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x10, 0x4d, 0x61, 0x6b,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x2d, 0x0a, 0x11, 0x4d, 0x61, 0x6b, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x64,
	0x22, 0x47, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x4a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x22, 0x31, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64,
	0x22, 0x42, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x48, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x32, 0xbf, 0x03,
	0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x09, 0x4d, 0x61, 0x6b, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x12, 0x4c, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x08, 0x5a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_event_proto_rawDescOnce sync.Once
	file_proto_event_proto_rawDescData = file_proto_event_proto_rawDesc
)

func file_proto_event_proto_rawDescGZIP() []byte {
	file_proto_event_proto_rawDescOnce.Do(func() {
		file_proto_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_event_proto_rawDescData)
	})
	return file_proto_event_proto_rawDescData
}

var file_proto_event_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_event_proto_goTypes = []interface{}{
	(*EventInfo)(nil),               // 0: event_api.EventInfo
	(*MakeEventRequest)(nil),        // 1: event_api.MakeEventRequest
	(*MakeEventResponse)(nil),       // 2: event_api.MakeEventResponse
	(*GetEventRequest)(nil),         // 3: event_api.GetEventRequest
	(*GetEventResponse)(nil),        // 4: event_api.GetEventResponse
	(*DeleteEventRequest)(nil),      // 5: event_api.DeleteEventRequest
	(*DeleteEventResponse)(nil),     // 6: event_api.DeleteEventResponse
	(*GetEventsRequest)(nil),        // 7: event_api.GetEventsRequest
	(*GetEventsResponse)(nil),       // 8: event_api.GetEventsResponse
	(*GetEventsStreamRequest)(nil),  // 9: event_api.GetEventsStreamRequest
	(*GetEventsStreamResponse)(nil), // 10: event_api.GetEventsStreamResponse
	(*TestRequest)(nil),             // 11: event_api.TestRequest
	(*TestResponse)(nil),            // 12: event_api.TestResponse
}
var file_proto_event_proto_depIdxs = []int32{
	0,  // 0: event_api.GetEventResponse.eventInfo:type_name -> event_api.EventInfo
	0,  // 1: event_api.GetEventsResponse.eventInfo:type_name -> event_api.EventInfo
	0,  // 2: event_api.GetEventsStreamResponse.EventInfo:type_name -> event_api.EventInfo
	11, // 3: event_api.Event.Test:input_type -> event_api.TestRequest
	1,  // 4: event_api.Event.MakeEvent:input_type -> event_api.MakeEventRequest
	3,  // 5: event_api.Event.GetEvent:input_type -> event_api.GetEventRequest
	7,  // 6: event_api.Event.GetEvents:input_type -> event_api.GetEventsRequest
	9,  // 7: event_api.Event.GetEventsStream:input_type -> event_api.GetEventsStreamRequest
	5,  // 8: event_api.Event.DeleteEvent:input_type -> event_api.DeleteEventRequest
	12, // 9: event_api.Event.Test:output_type -> event_api.TestResponse
	2,  // 10: event_api.Event.MakeEvent:output_type -> event_api.MakeEventResponse
	4,  // 11: event_api.Event.GetEvent:output_type -> event_api.GetEventResponse
	8,  // 12: event_api.Event.GetEvents:output_type -> event_api.GetEventsResponse
	10, // 13: event_api.Event.GetEventsStream:output_type -> event_api.GetEventsStreamResponse
	6,  // 14: event_api.Event.DeleteEvent:output_type -> event_api.DeleteEventResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_event_proto_init() }
func file_proto_event_proto_init() {
	if File_proto_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventInfo); i {
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
		file_proto_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeEventRequest); i {
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
		file_proto_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeEventResponse); i {
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
		file_proto_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventRequest); i {
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
		file_proto_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventResponse); i {
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
		file_proto_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventRequest); i {
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
		file_proto_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventResponse); i {
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
		file_proto_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsRequest); i {
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
		file_proto_event_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsResponse); i {
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
		file_proto_event_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsStreamRequest); i {
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
		file_proto_event_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsStreamResponse); i {
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
		file_proto_event_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequest); i {
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
		file_proto_event_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResponse); i {
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
			RawDescriptor: file_proto_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_event_proto_goTypes,
		DependencyIndexes: file_proto_event_proto_depIdxs,
		MessageInfos:      file_proto_event_proto_msgTypes,
	}.Build()
	File_proto_event_proto = out.File
	file_proto_event_proto_rawDesc = nil
	file_proto_event_proto_goTypes = nil
	file_proto_event_proto_depIdxs = nil
}
