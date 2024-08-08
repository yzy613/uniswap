// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: router/v1/router.proto

package v1

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

type ExactInputSingleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenIn          string `protobuf:"bytes,1,opt,name=tokenIn,proto3" json:"tokenIn,omitempty"`
	TokenOut         string `protobuf:"bytes,2,opt,name=tokenOut,proto3" json:"tokenOut,omitempty"`
	Fee              uint32 `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Recipient        string `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Deadline         string `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	AmountIn         string `protobuf:"bytes,6,opt,name=amountIn,proto3" json:"amountIn,omitempty"`
	AmountOutMinimum string `protobuf:"bytes,7,opt,name=amountOutMinimum,proto3" json:"amountOutMinimum,omitempty"`
	PriceLimit       string `protobuf:"bytes,8,opt,name=priceLimit,proto3" json:"priceLimit,omitempty"`
}

func (x *ExactInputSingleRequest) Reset() {
	*x = ExactInputSingleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_router_v1_router_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExactInputSingleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExactInputSingleRequest) ProtoMessage() {}

func (x *ExactInputSingleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_router_v1_router_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExactInputSingleRequest.ProtoReflect.Descriptor instead.
func (*ExactInputSingleRequest) Descriptor() ([]byte, []int) {
	return file_router_v1_router_proto_rawDescGZIP(), []int{0}
}

func (x *ExactInputSingleRequest) GetTokenIn() string {
	if x != nil {
		return x.TokenIn
	}
	return ""
}

func (x *ExactInputSingleRequest) GetTokenOut() string {
	if x != nil {
		return x.TokenOut
	}
	return ""
}

func (x *ExactInputSingleRequest) GetFee() uint32 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *ExactInputSingleRequest) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *ExactInputSingleRequest) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *ExactInputSingleRequest) GetAmountIn() string {
	if x != nil {
		return x.AmountIn
	}
	return ""
}

func (x *ExactInputSingleRequest) GetAmountOutMinimum() string {
	if x != nil {
		return x.AmountOutMinimum
	}
	return ""
}

func (x *ExactInputSingleRequest) GetPriceLimit() string {
	if x != nil {
		return x.PriceLimit
	}
	return ""
}

type ExactInputSingleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmountOut string `protobuf:"bytes,1,opt,name=amountOut,proto3" json:"amountOut,omitempty"`
}

func (x *ExactInputSingleReply) Reset() {
	*x = ExactInputSingleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_router_v1_router_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExactInputSingleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExactInputSingleReply) ProtoMessage() {}

func (x *ExactInputSingleReply) ProtoReflect() protoreflect.Message {
	mi := &file_router_v1_router_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExactInputSingleReply.ProtoReflect.Descriptor instead.
func (*ExactInputSingleReply) Descriptor() ([]byte, []int) {
	return file_router_v1_router_proto_rawDescGZIP(), []int{1}
}

func (x *ExactInputSingleReply) GetAmountOut() string {
	if x != nil {
		return x.AmountOut
	}
	return ""
}

type ExactOutputSingleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenIn         string `protobuf:"bytes,1,opt,name=tokenIn,proto3" json:"tokenIn,omitempty"`
	TokenOut        string `protobuf:"bytes,2,opt,name=tokenOut,proto3" json:"tokenOut,omitempty"`
	Fee             uint32 `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Recipient       string `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Deadline        string `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	AmountOut       string `protobuf:"bytes,6,opt,name=amountOut,proto3" json:"amountOut,omitempty"`
	AmountInMaximum string `protobuf:"bytes,7,opt,name=amountInMaximum,proto3" json:"amountInMaximum,omitempty"`
	PriceLimit      string `protobuf:"bytes,8,opt,name=priceLimit,proto3" json:"priceLimit,omitempty"`
}

func (x *ExactOutputSingleRequest) Reset() {
	*x = ExactOutputSingleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_router_v1_router_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExactOutputSingleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExactOutputSingleRequest) ProtoMessage() {}

func (x *ExactOutputSingleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_router_v1_router_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExactOutputSingleRequest.ProtoReflect.Descriptor instead.
func (*ExactOutputSingleRequest) Descriptor() ([]byte, []int) {
	return file_router_v1_router_proto_rawDescGZIP(), []int{2}
}

func (x *ExactOutputSingleRequest) GetTokenIn() string {
	if x != nil {
		return x.TokenIn
	}
	return ""
}

func (x *ExactOutputSingleRequest) GetTokenOut() string {
	if x != nil {
		return x.TokenOut
	}
	return ""
}

func (x *ExactOutputSingleRequest) GetFee() uint32 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *ExactOutputSingleRequest) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *ExactOutputSingleRequest) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *ExactOutputSingleRequest) GetAmountOut() string {
	if x != nil {
		return x.AmountOut
	}
	return ""
}

func (x *ExactOutputSingleRequest) GetAmountInMaximum() string {
	if x != nil {
		return x.AmountInMaximum
	}
	return ""
}

func (x *ExactOutputSingleRequest) GetPriceLimit() string {
	if x != nil {
		return x.PriceLimit
	}
	return ""
}

type ExactOutputSingleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmountIn string `protobuf:"bytes,1,opt,name=amountIn,proto3" json:"amountIn,omitempty"`
}

func (x *ExactOutputSingleReply) Reset() {
	*x = ExactOutputSingleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_router_v1_router_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExactOutputSingleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExactOutputSingleReply) ProtoMessage() {}

func (x *ExactOutputSingleReply) ProtoReflect() protoreflect.Message {
	mi := &file_router_v1_router_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExactOutputSingleReply.ProtoReflect.Descriptor instead.
func (*ExactOutputSingleReply) Descriptor() ([]byte, []int) {
	return file_router_v1_router_proto_rawDescGZIP(), []int{3}
}

func (x *ExactOutputSingleReply) GetAmountIn() string {
	if x != nil {
		return x.AmountIn
	}
	return ""
}

var File_router_v1_router_proto protoreflect.FileDescriptor

var file_router_v1_router_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x17, 0x45, 0x78, 0x61, 0x63, 0x74, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x12,
	0x2a, 0x0a, 0x10, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x4d, 0x69, 0x6e, 0x69,
	0x6d, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x4f, 0x75, 0x74, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x35, 0x0a, 0x15, 0x45,
	0x78, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f,
	0x75, 0x74, 0x22, 0x84, 0x02, 0x0a, 0x18, 0x45, 0x78, 0x61, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x4f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x12,
	0x28, 0x0a, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x4d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x34, 0x0a, 0x16, 0x45, 0x78, 0x61,
	0x63, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x32,
	0x9c, 0x02, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x65,
	0x78, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x2f, 0x65, 0x78, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x11, 0x65, 0x78, 0x61, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x63, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x61, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x61,
	0x63, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x42, 0x2b,
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x18, 0x75, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_router_v1_router_proto_rawDescOnce sync.Once
	file_router_v1_router_proto_rawDescData = file_router_v1_router_proto_rawDesc
)

func file_router_v1_router_proto_rawDescGZIP() []byte {
	file_router_v1_router_proto_rawDescOnce.Do(func() {
		file_router_v1_router_proto_rawDescData = protoimpl.X.CompressGZIP(file_router_v1_router_proto_rawDescData)
	})
	return file_router_v1_router_proto_rawDescData
}

var file_router_v1_router_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_router_v1_router_proto_goTypes = []any{
	(*ExactInputSingleRequest)(nil),  // 0: api.router.v1.ExactInputSingleRequest
	(*ExactInputSingleReply)(nil),    // 1: api.router.v1.ExactInputSingleReply
	(*ExactOutputSingleRequest)(nil), // 2: api.router.v1.ExactOutputSingleRequest
	(*ExactOutputSingleReply)(nil),   // 3: api.router.v1.ExactOutputSingleReply
}
var file_router_v1_router_proto_depIdxs = []int32{
	0, // 0: api.router.v1.Router.exactInputSingle:input_type -> api.router.v1.ExactInputSingleRequest
	2, // 1: api.router.v1.Router.exactOutputSingle:input_type -> api.router.v1.ExactOutputSingleRequest
	1, // 2: api.router.v1.Router.exactInputSingle:output_type -> api.router.v1.ExactInputSingleReply
	3, // 3: api.router.v1.Router.exactOutputSingle:output_type -> api.router.v1.ExactOutputSingleReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_router_v1_router_proto_init() }
func file_router_v1_router_proto_init() {
	if File_router_v1_router_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_router_v1_router_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExactInputSingleRequest); i {
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
		file_router_v1_router_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ExactInputSingleReply); i {
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
		file_router_v1_router_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ExactOutputSingleRequest); i {
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
		file_router_v1_router_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ExactOutputSingleReply); i {
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
			RawDescriptor: file_router_v1_router_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_router_v1_router_proto_goTypes,
		DependencyIndexes: file_router_v1_router_proto_depIdxs,
		MessageInfos:      file_router_v1_router_proto_msgTypes,
	}.Build()
	File_router_v1_router_proto = out.File
	file_router_v1_router_proto_rawDesc = nil
	file_router_v1_router_proto_goTypes = nil
	file_router_v1_router_proto_depIdxs = nil
}
