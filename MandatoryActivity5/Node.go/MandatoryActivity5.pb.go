// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: MandatoryActivity5.proto

package Node_go

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

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bidder string `protobuf:"bytes,1,opt,name=bidder,proto3" json:"bidder,omitempty"`
	Amount int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	mi := &file_MandatoryActivity5_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MandatoryActivity5_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_MandatoryActivity5_proto_rawDescGZIP(), []int{0}
}

func (x *BidRequest) GetBidder() string {
	if x != nil {
		return x.Bidder
	}
	return ""
}

func (x *BidRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	mi := &file_MandatoryActivity5_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MandatoryActivity5_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_MandatoryActivity5_proto_rawDescGZIP(), []int{1}
}

func (x *BidResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResultRequest) Reset() {
	*x = ResultRequest{}
	mi := &file_MandatoryActivity5_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRequest) ProtoMessage() {}

func (x *ResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MandatoryActivity5_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRequest.ProtoReflect.Descriptor instead.
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return file_MandatoryActivity5_proto_rawDescGZIP(), []int{2}
}

func (x *ResultRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Highestbid string `protobuf:"bytes,1,opt,name=highestbid,proto3" json:"highestbid,omitempty"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	mi := &file_MandatoryActivity5_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MandatoryActivity5_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_MandatoryActivity5_proto_rawDescGZIP(), []int{3}
}

func (x *ResultResponse) GetHighestbid() string {
	if x != nil {
		return x.Highestbid
	}
	return ""
}

var File_MandatoryActivity5_proto protoreflect.FileDescriptor

var file_MandatoryActivity5_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x4d, 0x61, 0x6e, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x35, 0x22, 0x3c,
	0x0a, 0x0a, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x69,
	0x64, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x0b,
	0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x30, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x62, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x62,
	0x69, 0x64, 0x32, 0xa2, 0x01, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46,
	0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x1e, 0x2e, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x35, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x35, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x21, 0x2e, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x4d, 0x61, 0x6e, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x35, 0x2f, 0x4e, 0x6f,
	0x64, 0x65, 0x2e, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MandatoryActivity5_proto_rawDescOnce sync.Once
	file_MandatoryActivity5_proto_rawDescData = file_MandatoryActivity5_proto_rawDesc
)

func file_MandatoryActivity5_proto_rawDescGZIP() []byte {
	file_MandatoryActivity5_proto_rawDescOnce.Do(func() {
		file_MandatoryActivity5_proto_rawDescData = protoimpl.X.CompressGZIP(file_MandatoryActivity5_proto_rawDescData)
	})
	return file_MandatoryActivity5_proto_rawDescData
}

var file_MandatoryActivity5_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_MandatoryActivity5_proto_goTypes = []any{
	(*BidRequest)(nil),     // 0: MandatoryActivity5.BidRequest
	(*BidResponse)(nil),    // 1: MandatoryActivity5.BidResponse
	(*ResultRequest)(nil),  // 2: MandatoryActivity5.ResultRequest
	(*ResultResponse)(nil), // 3: MandatoryActivity5.ResultResponse
}
var file_MandatoryActivity5_proto_depIdxs = []int32{
	0, // 0: MandatoryActivity5.Auction.Bid:input_type -> MandatoryActivity5.BidRequest
	2, // 1: MandatoryActivity5.Auction.Result:input_type -> MandatoryActivity5.ResultRequest
	1, // 2: MandatoryActivity5.Auction.Bid:output_type -> MandatoryActivity5.BidResponse
	3, // 3: MandatoryActivity5.Auction.Result:output_type -> MandatoryActivity5.ResultResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MandatoryActivity5_proto_init() }
func file_MandatoryActivity5_proto_init() {
	if File_MandatoryActivity5_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MandatoryActivity5_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_MandatoryActivity5_proto_goTypes,
		DependencyIndexes: file_MandatoryActivity5_proto_depIdxs,
		MessageInfos:      file_MandatoryActivity5_proto_msgTypes,
	}.Build()
	File_MandatoryActivity5_proto = out.File
	file_MandatoryActivity5_proto_rawDesc = nil
	file_MandatoryActivity5_proto_goTypes = nil
	file_MandatoryActivity5_proto_depIdxs = nil
}