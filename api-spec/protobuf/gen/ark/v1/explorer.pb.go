// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: ark/v1/explorer.proto

package arkv1

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

type GetRoundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
}

func (x *GetRoundRequest) Reset() {
	*x = GetRoundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ark_v1_explorer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoundRequest) ProtoMessage() {}

func (x *GetRoundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ark_v1_explorer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoundRequest.ProtoReflect.Descriptor instead.
func (*GetRoundRequest) Descriptor() ([]byte, []int) {
	return file_ark_v1_explorer_proto_rawDescGZIP(), []int{0}
}

func (x *GetRoundRequest) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

type GetRoundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Round *Round `protobuf:"bytes,1,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *GetRoundResponse) Reset() {
	*x = GetRoundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ark_v1_explorer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoundResponse) ProtoMessage() {}

func (x *GetRoundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ark_v1_explorer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoundResponse.ProtoReflect.Descriptor instead.
func (*GetRoundResponse) Descriptor() ([]byte, []int) {
	return file_ark_v1_explorer_proto_rawDescGZIP(), []int{1}
}

func (x *GetRoundResponse) GetRound() *Round {
	if x != nil {
		return x.Round
	}
	return nil
}

type GetRoundByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRoundByIdRequest) Reset() {
	*x = GetRoundByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ark_v1_explorer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoundByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoundByIdRequest) ProtoMessage() {}

func (x *GetRoundByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ark_v1_explorer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoundByIdRequest.ProtoReflect.Descriptor instead.
func (*GetRoundByIdRequest) Descriptor() ([]byte, []int) {
	return file_ark_v1_explorer_proto_rawDescGZIP(), []int{2}
}

func (x *GetRoundByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRoundByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Round *Round `protobuf:"bytes,1,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *GetRoundByIdResponse) Reset() {
	*x = GetRoundByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ark_v1_explorer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoundByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoundByIdResponse) ProtoMessage() {}

func (x *GetRoundByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ark_v1_explorer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoundByIdResponse.ProtoReflect.Descriptor instead.
func (*GetRoundByIdResponse) Descriptor() ([]byte, []int) {
	return file_ark_v1_explorer_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoundByIdResponse) GetRound() *Round {
	if x != nil {
		return x.Round
	}
	return nil
}

type ListVtxosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ListVtxosRequest) Reset() {
	*x = ListVtxosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ark_v1_explorer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVtxosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVtxosRequest) ProtoMessage() {}

func (x *ListVtxosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ark_v1_explorer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVtxosRequest.ProtoReflect.Descriptor instead.
func (*ListVtxosRequest) Descriptor() ([]byte, []int) {
	return file_ark_v1_explorer_proto_rawDescGZIP(), []int{4}
}

func (x *ListVtxosRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ListVtxosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpendableVtxos []*Vtxo `protobuf:"bytes,1,rep,name=spendable_vtxos,json=spendableVtxos,proto3" json:"spendable_vtxos,omitempty"`
	SpentVtxos     []*Vtxo `protobuf:"bytes,2,rep,name=spent_vtxos,json=spentVtxos,proto3" json:"spent_vtxos,omitempty"`
}

func (x *ListVtxosResponse) Reset() {
	*x = ListVtxosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ark_v1_explorer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVtxosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVtxosResponse) ProtoMessage() {}

func (x *ListVtxosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ark_v1_explorer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVtxosResponse.ProtoReflect.Descriptor instead.
func (*ListVtxosResponse) Descriptor() ([]byte, []int) {
	return file_ark_v1_explorer_proto_rawDescGZIP(), []int{5}
}

func (x *ListVtxosResponse) GetSpendableVtxos() []*Vtxo {
	if x != nil {
		return x.SpendableVtxos
	}
	return nil
}

func (x *ListVtxosResponse) GetSpentVtxos() []*Vtxo {
	if x != nil {
		return x.SpentVtxos
	}
	return nil
}

var File_ark_v1_explorer_proto protoreflect.FileDescriptor

var file_ark_v1_explorer_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61,
	0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x72,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x05,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x74, 0x78,
	0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x79, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x74, 0x78, 0x6f, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0f, 0x73, 0x70, 0x65, 0x6e,
	0x64, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x74, 0x78, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x74, 0x78, 0x6f, 0x52,
	0x0e, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x74, 0x78, 0x6f, 0x73, 0x12,
	0x2d, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x74, 0x78, 0x6f, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x74,
	0x78, 0x6f, 0x52, 0x0a, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x56, 0x74, 0x78, 0x6f, 0x73, 0x32, 0xaf,
	0x02, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x57, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x17,
	0x2e, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x7b, 0x74, 0x78, 0x69, 0x64, 0x7d, 0x12, 0x64, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x61, 0x72,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x72, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x5d, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x74, 0x78, 0x6f, 0x73, 0x12, 0x18,
	0x2e, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x74, 0x78, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x72, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x74, 0x78, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x74, 0x78, 0x6f, 0x73, 0x2f, 0x7b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d,
	0x42, 0x93, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x42,
	0x0d, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6b,
	0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x61, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x72, 0x6b, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x41, 0x72, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x06, 0x41, 0x72, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x41, 0x72, 0x6b, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41,
	0x72, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ark_v1_explorer_proto_rawDescOnce sync.Once
	file_ark_v1_explorer_proto_rawDescData = file_ark_v1_explorer_proto_rawDesc
)

func file_ark_v1_explorer_proto_rawDescGZIP() []byte {
	file_ark_v1_explorer_proto_rawDescOnce.Do(func() {
		file_ark_v1_explorer_proto_rawDescData = protoimpl.X.CompressGZIP(file_ark_v1_explorer_proto_rawDescData)
	})
	return file_ark_v1_explorer_proto_rawDescData
}

var file_ark_v1_explorer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ark_v1_explorer_proto_goTypes = []interface{}{
	(*GetRoundRequest)(nil),      // 0: ark.v1.GetRoundRequest
	(*GetRoundResponse)(nil),     // 1: ark.v1.GetRoundResponse
	(*GetRoundByIdRequest)(nil),  // 2: ark.v1.GetRoundByIdRequest
	(*GetRoundByIdResponse)(nil), // 3: ark.v1.GetRoundByIdResponse
	(*ListVtxosRequest)(nil),     // 4: ark.v1.ListVtxosRequest
	(*ListVtxosResponse)(nil),    // 5: ark.v1.ListVtxosResponse
	(*Round)(nil),                // 6: ark.v1.Round
	(*Vtxo)(nil),                 // 7: ark.v1.Vtxo
}
var file_ark_v1_explorer_proto_depIdxs = []int32{
	6, // 0: ark.v1.GetRoundResponse.round:type_name -> ark.v1.Round
	6, // 1: ark.v1.GetRoundByIdResponse.round:type_name -> ark.v1.Round
	7, // 2: ark.v1.ListVtxosResponse.spendable_vtxos:type_name -> ark.v1.Vtxo
	7, // 3: ark.v1.ListVtxosResponse.spent_vtxos:type_name -> ark.v1.Vtxo
	0, // 4: ark.v1.ExplorerService.GetRound:input_type -> ark.v1.GetRoundRequest
	2, // 5: ark.v1.ExplorerService.GetRoundById:input_type -> ark.v1.GetRoundByIdRequest
	4, // 6: ark.v1.ExplorerService.ListVtxos:input_type -> ark.v1.ListVtxosRequest
	1, // 7: ark.v1.ExplorerService.GetRound:output_type -> ark.v1.GetRoundResponse
	3, // 8: ark.v1.ExplorerService.GetRoundById:output_type -> ark.v1.GetRoundByIdResponse
	5, // 9: ark.v1.ExplorerService.ListVtxos:output_type -> ark.v1.ListVtxosResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ark_v1_explorer_proto_init() }
func file_ark_v1_explorer_proto_init() {
	if File_ark_v1_explorer_proto != nil {
		return
	}
	file_ark_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ark_v1_explorer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoundRequest); i {
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
		file_ark_v1_explorer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoundResponse); i {
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
		file_ark_v1_explorer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoundByIdRequest); i {
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
		file_ark_v1_explorer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoundByIdResponse); i {
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
		file_ark_v1_explorer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVtxosRequest); i {
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
		file_ark_v1_explorer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVtxosResponse); i {
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
			RawDescriptor: file_ark_v1_explorer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ark_v1_explorer_proto_goTypes,
		DependencyIndexes: file_ark_v1_explorer_proto_depIdxs,
		MessageInfos:      file_ark_v1_explorer_proto_msgTypes,
	}.Build()
	File_ark_v1_explorer_proto = out.File
	file_ark_v1_explorer_proto_rawDesc = nil
	file_ark_v1_explorer_proto_goTypes = nil
	file_ark_v1_explorer_proto_depIdxs = nil
}
