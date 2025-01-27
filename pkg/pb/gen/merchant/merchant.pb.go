// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: pkg/pb/protos/merchant/merchant.proto

package merchant

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

type Status int32

const (
	Status_INACTIVE    Status = 0 // 停用
	Status_ACTIVE      Status = 1 // 啟用
	Status_Maintenance Status = 2 // 維護
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "INACTIVE",
		1: "ACTIVE",
		2: "Maintenance",
	}
	Status_value = map[string]int32{
		"INACTIVE":    0,
		"ACTIVE":      1,
		"Maintenance": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_pb_protos_merchant_merchant_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_pkg_pb_protos_merchant_merchant_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_pb_protos_merchant_merchant_proto_rawDescGZIP(), []int{0}
}

type ClientType int32

const (
	ClientType_FRONT ClientType = 0 // 前台
	ClientType_BACK  ClientType = 1 // 後台
)

// Enum value maps for ClientType.
var (
	ClientType_name = map[int32]string{
		0: "FRONT",
		1: "BACK",
	}
	ClientType_value = map[string]int32{
		"FRONT": 0,
		"BACK":  1,
	}
)

func (x ClientType) Enum() *ClientType {
	p := new(ClientType)
	*p = x
	return p
}

func (x ClientType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_pb_protos_merchant_merchant_proto_enumTypes[1].Descriptor()
}

func (ClientType) Type() protoreflect.EnumType {
	return &file_pkg_pb_protos_merchant_merchant_proto_enumTypes[1]
}

func (x ClientType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientType.Descriptor instead.
func (ClientType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_pb_protos_merchant_merchant_proto_rawDescGZIP(), []int{1}
}

type CreateMerchantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantName string   `protobuf:"bytes,1,opt,name=merchant_name,json=merchantName,proto3" json:"merchant_name,omitempty"`
	FrontDomain  string   `protobuf:"bytes,2,opt,name=front_domain,json=frontDomain,proto3" json:"front_domain,omitempty"`
	BackDomain   string   `protobuf:"bytes,3,opt,name=back_domain,json=backDomain,proto3" json:"back_domain,omitempty"`
	Currencies   []string `protobuf:"bytes,4,rep,name=currencies,proto3" json:"currencies,omitempty"`
	Status       Status   `protobuf:"varint,5,opt,name=status,proto3,enum=merchant.Status" json:"status,omitempty"`
}

func (x *CreateMerchantRequest) Reset() {
	*x = CreateMerchantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMerchantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMerchantRequest) ProtoMessage() {}

func (x *CreateMerchantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMerchantRequest.ProtoReflect.Descriptor instead.
func (*CreateMerchantRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protos_merchant_merchant_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMerchantRequest) GetMerchantName() string {
	if x != nil {
		return x.MerchantName
	}
	return ""
}

func (x *CreateMerchantRequest) GetFrontDomain() string {
	if x != nil {
		return x.FrontDomain
	}
	return ""
}

func (x *CreateMerchantRequest) GetBackDomain() string {
	if x != nil {
		return x.BackDomain
	}
	return ""
}

func (x *CreateMerchantRequest) GetCurrencies() []string {
	if x != nil {
		return x.Currencies
	}
	return nil
}

func (x *CreateMerchantRequest) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_INACTIVE
}

type MerchantInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantId   int64    `protobuf:"varint,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	MerchantName string   `protobuf:"bytes,2,opt,name=merchant_name,json=merchantName,proto3" json:"merchant_name,omitempty"`
	FrontDomain  string   `protobuf:"bytes,3,opt,name=front_domain,json=frontDomain,proto3" json:"front_domain,omitempty"`
	FrontSecret  string   `protobuf:"bytes,4,opt,name=front_secret,json=frontSecret,proto3" json:"front_secret,omitempty"`
	BackDomain   string   `protobuf:"bytes,5,opt,name=back_domain,json=backDomain,proto3" json:"back_domain,omitempty"`
	BackSecret   string   `protobuf:"bytes,6,opt,name=back_secret,json=backSecret,proto3" json:"back_secret,omitempty"`
	Currencies   []string `protobuf:"bytes,7,rep,name=currencies,proto3" json:"currencies,omitempty"`
	Status       Status   `protobuf:"varint,8,opt,name=status,proto3,enum=merchant.Status" json:"status,omitempty"`
}

func (x *MerchantInfo) Reset() {
	*x = MerchantInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantInfo) ProtoMessage() {}

func (x *MerchantInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantInfo.ProtoReflect.Descriptor instead.
func (*MerchantInfo) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protos_merchant_merchant_proto_rawDescGZIP(), []int{1}
}

func (x *MerchantInfo) GetMerchantId() int64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

func (x *MerchantInfo) GetMerchantName() string {
	if x != nil {
		return x.MerchantName
	}
	return ""
}

func (x *MerchantInfo) GetFrontDomain() string {
	if x != nil {
		return x.FrontDomain
	}
	return ""
}

func (x *MerchantInfo) GetFrontSecret() string {
	if x != nil {
		return x.FrontSecret
	}
	return ""
}

func (x *MerchantInfo) GetBackDomain() string {
	if x != nil {
		return x.BackDomain
	}
	return ""
}

func (x *MerchantInfo) GetBackSecret() string {
	if x != nil {
		return x.BackSecret
	}
	return ""
}

func (x *MerchantInfo) GetCurrencies() []string {
	if x != nil {
		return x.Currencies
	}
	return nil
}

func (x *MerchantInfo) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_INACTIVE
}

type UpdateMerchantStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantId int64  `protobuf:"varint,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	Status     Status `protobuf:"varint,2,opt,name=status,proto3,enum=merchant.Status" json:"status,omitempty"`
}

func (x *UpdateMerchantStatusRequest) Reset() {
	*x = UpdateMerchantStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMerchantStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMerchantStatusRequest) ProtoMessage() {}

func (x *UpdateMerchantStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMerchantStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateMerchantStatusRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protos_merchant_merchant_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMerchantStatusRequest) GetMerchantId() int64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

func (x *UpdateMerchantStatusRequest) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_INACTIVE
}

type UpdateMerchantStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=merchant.Status" json:"status,omitempty"`
}

func (x *UpdateMerchantStatusResponse) Reset() {
	*x = UpdateMerchantStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMerchantStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMerchantStatusResponse) ProtoMessage() {}

func (x *UpdateMerchantStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMerchantStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateMerchantStatusResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protos_merchant_merchant_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateMerchantStatusResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_INACTIVE
}

type GetMerchantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantId int64 `protobuf:"varint,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
}

func (x *GetMerchantRequest) Reset() {
	*x = GetMerchantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMerchantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMerchantRequest) ProtoMessage() {}

func (x *GetMerchantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMerchantRequest.ProtoReflect.Descriptor instead.
func (*GetMerchantRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protos_merchant_merchant_proto_rawDescGZIP(), []int{4}
}

func (x *GetMerchantRequest) GetMerchantId() int64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

type ValidClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     int64      `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string     `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	ClientType   ClientType `protobuf:"varint,3,opt,name=client_type,json=clientType,proto3,enum=merchant.ClientType" json:"client_type,omitempty"`
}

func (x *ValidClientRequest) Reset() {
	*x = ValidClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidClientRequest) ProtoMessage() {}

func (x *ValidClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidClientRequest.ProtoReflect.Descriptor instead.
func (*ValidClientRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protos_merchant_merchant_proto_rawDescGZIP(), []int{5}
}

func (x *ValidClientRequest) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *ValidClientRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *ValidClientRequest) GetClientType() ClientType {
	if x != nil {
		return x.ClientType
	}
	return ClientType_FRONT
}

type ValidClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid bool `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
}

func (x *ValidClientResponse) Reset() {
	*x = ValidClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidClientResponse) ProtoMessage() {}

func (x *ValidClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protos_merchant_merchant_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidClientResponse.ProtoReflect.Descriptor instead.
func (*ValidClientResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protos_merchant_merchant_proto_rawDescGZIP(), []int{6}
}

func (x *ValidClientResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

var File_pkg_pb_protos_merchant_merchant_proto protoreflect.FileDescriptor

var file_pkg_pb_protos_merchant_merchant_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x22, 0xca, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa6,
	0x02, 0x0a, 0x0c, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x61, 0x63, 0x6b, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x62, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x28, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x68, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x48, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x0b, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x30, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x2a, 0x33, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c,
	0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x02, 0x2a, 0x21, 0x0a, 0x0a, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x52, 0x4f, 0x4e, 0x54,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x32, 0xd3, 0x02, 0x0a,
	0x0f, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x49, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x65, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x25, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x12, 0x1c, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4a, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_protos_merchant_merchant_proto_rawDescOnce sync.Once
	file_pkg_pb_protos_merchant_merchant_proto_rawDescData = file_pkg_pb_protos_merchant_merchant_proto_rawDesc
)

func file_pkg_pb_protos_merchant_merchant_proto_rawDescGZIP() []byte {
	file_pkg_pb_protos_merchant_merchant_proto_rawDescOnce.Do(func() {
		file_pkg_pb_protos_merchant_merchant_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_protos_merchant_merchant_proto_rawDescData)
	})
	return file_pkg_pb_protos_merchant_merchant_proto_rawDescData
}

var file_pkg_pb_protos_merchant_merchant_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_pb_protos_merchant_merchant_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_pb_protos_merchant_merchant_proto_goTypes = []any{
	(Status)(0),                          // 0: merchant.Status
	(ClientType)(0),                      // 1: merchant.ClientType
	(*CreateMerchantRequest)(nil),        // 2: merchant.CreateMerchantRequest
	(*MerchantInfo)(nil),                 // 3: merchant.MerchantInfo
	(*UpdateMerchantStatusRequest)(nil),  // 4: merchant.UpdateMerchantStatusRequest
	(*UpdateMerchantStatusResponse)(nil), // 5: merchant.UpdateMerchantStatusResponse
	(*GetMerchantRequest)(nil),           // 6: merchant.GetMerchantRequest
	(*ValidClientRequest)(nil),           // 7: merchant.ValidClientRequest
	(*ValidClientResponse)(nil),          // 8: merchant.ValidClientResponse
}
var file_pkg_pb_protos_merchant_merchant_proto_depIdxs = []int32{
	0, // 0: merchant.CreateMerchantRequest.status:type_name -> merchant.Status
	0, // 1: merchant.MerchantInfo.status:type_name -> merchant.Status
	0, // 2: merchant.UpdateMerchantStatusRequest.status:type_name -> merchant.Status
	0, // 3: merchant.UpdateMerchantStatusResponse.status:type_name -> merchant.Status
	1, // 4: merchant.ValidClientRequest.client_type:type_name -> merchant.ClientType
	2, // 5: merchant.MerchantService.CreateMerchant:input_type -> merchant.CreateMerchantRequest
	4, // 6: merchant.MerchantService.UpdateMerchantStatus:input_type -> merchant.UpdateMerchantStatusRequest
	6, // 7: merchant.MerchantService.GetMechant:input_type -> merchant.GetMerchantRequest
	7, // 8: merchant.MerchantService.ValidClient:input_type -> merchant.ValidClientRequest
	3, // 9: merchant.MerchantService.CreateMerchant:output_type -> merchant.MerchantInfo
	5, // 10: merchant.MerchantService.UpdateMerchantStatus:output_type -> merchant.UpdateMerchantStatusResponse
	3, // 11: merchant.MerchantService.GetMechant:output_type -> merchant.MerchantInfo
	8, // 12: merchant.MerchantService.ValidClient:output_type -> merchant.ValidClientResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_pb_protos_merchant_merchant_proto_init() }
func file_pkg_pb_protos_merchant_merchant_proto_init() {
	if File_pkg_pb_protos_merchant_merchant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_protos_merchant_merchant_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateMerchantRequest); i {
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
		file_pkg_pb_protos_merchant_merchant_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MerchantInfo); i {
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
		file_pkg_pb_protos_merchant_merchant_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateMerchantStatusRequest); i {
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
		file_pkg_pb_protos_merchant_merchant_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateMerchantStatusResponse); i {
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
		file_pkg_pb_protos_merchant_merchant_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetMerchantRequest); i {
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
		file_pkg_pb_protos_merchant_merchant_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ValidClientRequest); i {
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
		file_pkg_pb_protos_merchant_merchant_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ValidClientResponse); i {
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
			RawDescriptor: file_pkg_pb_protos_merchant_merchant_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_protos_merchant_merchant_proto_goTypes,
		DependencyIndexes: file_pkg_pb_protos_merchant_merchant_proto_depIdxs,
		EnumInfos:         file_pkg_pb_protos_merchant_merchant_proto_enumTypes,
		MessageInfos:      file_pkg_pb_protos_merchant_merchant_proto_msgTypes,
	}.Build()
	File_pkg_pb_protos_merchant_merchant_proto = out.File
	file_pkg_pb_protos_merchant_merchant_proto_rawDesc = nil
	file_pkg_pb_protos_merchant_merchant_proto_goTypes = nil
	file_pkg_pb_protos_merchant_merchant_proto_depIdxs = nil
}
