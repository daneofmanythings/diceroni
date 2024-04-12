// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: internal/grpc/proto/roller.proto

package roller

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_roller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_roller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_roller_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_roller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_roller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_roller_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type RollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiceString string `protobuf:"bytes,1,opt,name=dice_string,json=diceString,proto3" json:"dice_string,omitempty"`
	CallerId   string `protobuf:"bytes,2,opt,name=caller_id,json=callerId,proto3" json:"caller_id,omitempty"`
}

func (x *RollRequest) Reset() {
	*x = RollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_roller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollRequest) ProtoMessage() {}

func (x *RollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_roller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollRequest.ProtoReflect.Descriptor instead.
func (*RollRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_roller_proto_rawDescGZIP(), []int{2}
}

func (x *RollRequest) GetDiceString() string {
	if x != nil {
		return x.DiceString
	}
	return ""
}

func (x *RollRequest) GetCallerId() string {
	if x != nil {
		return x.CallerId
	}
	return ""
}

type DiceRollMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseLiteral string   `protobuf:"bytes,1,opt,name=response_literal,json=responseLiteral,proto3" json:"response_literal,omitempty"`
	Tags            []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	RawRolls        []uint32 `protobuf:"varint,3,rep,packed,name=raw_rolls,json=rawRolls,proto3" json:"raw_rolls,omitempty"`
	FinalRolls      []uint32 `protobuf:"varint,4,rep,packed,name=final_rolls,json=finalRolls,proto3" json:"final_rolls,omitempty"`
	Value           int64    `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DiceRollMetadata) Reset() {
	*x = DiceRollMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_roller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiceRollMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiceRollMetadata) ProtoMessage() {}

func (x *DiceRollMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_roller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiceRollMetadata.ProtoReflect.Descriptor instead.
func (*DiceRollMetadata) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_roller_proto_rawDescGZIP(), []int{3}
}

func (x *DiceRollMetadata) GetResponseLiteral() string {
	if x != nil {
		return x.ResponseLiteral
	}
	return ""
}

func (x *DiceRollMetadata) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *DiceRollMetadata) GetRawRolls() []uint32 {
	if x != nil {
		return x.RawRolls
	}
	return nil
}

func (x *DiceRollMetadata) GetFinalRolls() []uint32 {
	if x != nil {
		return x.FinalRolls
	}
	return nil
}

func (x *DiceRollMetadata) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type RollData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestLiteral string              `protobuf:"bytes,1,opt,name=request_literal,json=requestLiteral,proto3" json:"request_literal,omitempty"`
	Value          int64               `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Metadata       []*DiceRollMetadata `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *RollData) Reset() {
	*x = RollData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_roller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollData) ProtoMessage() {}

func (x *RollData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_roller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollData.ProtoReflect.Descriptor instead.
func (*RollData) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_roller_proto_rawDescGZIP(), []int{4}
}

func (x *RollData) GetRequestLiteral() string {
	if x != nil {
		return x.RequestLiteral
	}
	return ""
}

func (x *RollData) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *RollData) GetMetadata() []*DiceRollMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type MyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details []*any1.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *MyStatus) Reset() {
	*x = MyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_roller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyStatus) ProtoMessage() {}

func (x *MyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_roller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyStatus.ProtoReflect.Descriptor instead.
func (*MyStatus) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_roller_proto_rawDescGZIP(), []int{5}
}

func (x *MyStatus) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MyStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MyStatus) GetDetails() []*any1.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

type RollResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*RollResponse_Data
	//	*RollResponse_Status
	Message isRollResponse_Message `protobuf_oneof:"message"`
}

func (x *RollResponse) Reset() {
	*x = RollResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_proto_roller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollResponse) ProtoMessage() {}

func (x *RollResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_roller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollResponse.ProtoReflect.Descriptor instead.
func (*RollResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_roller_proto_rawDescGZIP(), []int{6}
}

func (m *RollResponse) GetMessage() isRollResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *RollResponse) GetData() *RollData {
	if x, ok := x.GetMessage().(*RollResponse_Data); ok {
		return x.Data
	}
	return nil
}

func (x *RollResponse) GetStatus() *MyStatus {
	if x, ok := x.GetMessage().(*RollResponse_Status); ok {
		return x.Status
	}
	return nil
}

type isRollResponse_Message interface {
	isRollResponse_Message()
}

type RollResponse_Data struct {
	Data *RollData `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type RollResponse_Status struct {
	Status *MyStatus `protobuf:"bytes,2,opt,name=status,proto3,oneof"`
}

func (*RollResponse_Data) isRollResponse_Message() {}

func (*RollResponse_Status) isRollResponse_Message() {}

var File_internal_grpc_proto_roller_proto protoreflect.FileDescriptor

var file_internal_grpc_proto_roller_proto_rawDesc = []byte{
	0x0a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x4b, 0x0a, 0x0b,
	0x52, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x10, 0x44, 0x69,
	0x63, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29,
	0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x61, 0x77, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x08, 0x72, 0x61, 0x77, 0x52, 0x6f, 0x6c, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x83, 0x01, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27,
	0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x69, 0x63,
	0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x08, 0x4d, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0x75, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x6f, 0x6c,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x09, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6e, 0x65, 0x6f, 0x66, 0x6d, 0x61, 0x6e,
	0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_grpc_proto_roller_proto_rawDescOnce sync.Once
	file_internal_grpc_proto_roller_proto_rawDescData = file_internal_grpc_proto_roller_proto_rawDesc
)

func file_internal_grpc_proto_roller_proto_rawDescGZIP() []byte {
	file_internal_grpc_proto_roller_proto_rawDescOnce.Do(func() {
		file_internal_grpc_proto_roller_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpc_proto_roller_proto_rawDescData)
	})
	return file_internal_grpc_proto_roller_proto_rawDescData
}

var file_internal_grpc_proto_roller_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_grpc_proto_roller_proto_goTypes = []interface{}{
	(*PingRequest)(nil),      // 0: google.rpc.PingRequest
	(*PingResponse)(nil),     // 1: google.rpc.PingResponse
	(*RollRequest)(nil),      // 2: google.rpc.RollRequest
	(*DiceRollMetadata)(nil), // 3: google.rpc.DiceRollMetadata
	(*RollData)(nil),         // 4: google.rpc.RollData
	(*MyStatus)(nil),         // 5: google.rpc.MyStatus
	(*RollResponse)(nil),     // 6: google.rpc.RollResponse
	(*any1.Any)(nil),         // 7: google.protobuf.Any
}
var file_internal_grpc_proto_roller_proto_depIdxs = []int32{
	3, // 0: google.rpc.RollData.metadata:type_name -> google.rpc.DiceRollMetadata
	7, // 1: google.rpc.MyStatus.details:type_name -> google.protobuf.Any
	4, // 2: google.rpc.RollResponse.data:type_name -> google.rpc.RollData
	5, // 3: google.rpc.RollResponse.status:type_name -> google.rpc.MyStatus
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_grpc_proto_roller_proto_init() }
func file_internal_grpc_proto_roller_proto_init() {
	if File_internal_grpc_proto_roller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpc_proto_roller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_internal_grpc_proto_roller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_internal_grpc_proto_roller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollRequest); i {
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
		file_internal_grpc_proto_roller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiceRollMetadata); i {
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
		file_internal_grpc_proto_roller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollData); i {
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
		file_internal_grpc_proto_roller_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyStatus); i {
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
		file_internal_grpc_proto_roller_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollResponse); i {
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
	file_internal_grpc_proto_roller_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*RollResponse_Data)(nil),
		(*RollResponse_Status)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_grpc_proto_roller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_grpc_proto_roller_proto_goTypes,
		DependencyIndexes: file_internal_grpc_proto_roller_proto_depIdxs,
		MessageInfos:      file_internal_grpc_proto_roller_proto_msgTypes,
	}.Build()
	File_internal_grpc_proto_roller_proto = out.File
	file_internal_grpc_proto_roller_proto_rawDesc = nil
	file_internal_grpc_proto_roller_proto_goTypes = nil
	file_internal_grpc_proto_roller_proto_depIdxs = nil
}
