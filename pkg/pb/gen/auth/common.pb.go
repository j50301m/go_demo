// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: pkg/pb/protos/auth/common.proto

package auth

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protos_auth_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protos_auth_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protos_auth_common_proto_rawDescGZIP(), []int{0}
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId     int64   `protobuf:"varint,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	RoleName   string  `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	PermIds    []int64 `protobuf:"varint,3,rep,packed,name=perm_ids,json=permIds,proto3" json:"perm_ids,omitempty"`
	ClientType int32   `protobuf:"varint,4,opt,name=clientType,proto3" json:"clientType,omitempty"`             // 使用 pkg/enum/client_type 的id作為參數
	IsSystem   bool    `protobuf:"varint,5,opt,name=is_system,json=isSystem,proto3" json:"is_system,omitempty"` // 是否為系統預設角色  如果是則不可編輯和刪除
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_protos_auth_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_protos_auth_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_pkg_pb_protos_auth_common_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *Role) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *Role) GetPermIds() []int64 {
	if x != nil {
		return x.PermIds
	}
	return nil
}

func (x *Role) GetClientType() int32 {
	if x != nil {
		return x.ClientType
	}
	return 0
}

func (x *Role) GetIsSystem() bool {
	if x != nil {
		return x.IsSystem
	}
	return false
}

var File_pkg_pb_protos_auth_common_proto protoreflect.FileDescriptor

var file_pkg_pb_protos_auth_common_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x94, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x6d, 0x49, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_protos_auth_common_proto_rawDescOnce sync.Once
	file_pkg_pb_protos_auth_common_proto_rawDescData = file_pkg_pb_protos_auth_common_proto_rawDesc
)

func file_pkg_pb_protos_auth_common_proto_rawDescGZIP() []byte {
	file_pkg_pb_protos_auth_common_proto_rawDescOnce.Do(func() {
		file_pkg_pb_protos_auth_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_protos_auth_common_proto_rawDescData)
	})
	return file_pkg_pb_protos_auth_common_proto_rawDescData
}

var file_pkg_pb_protos_auth_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_pb_protos_auth_common_proto_goTypes = []any{
	(*Empty)(nil), // 0: auth.Empty
	(*Role)(nil),  // 1: auth.Role
}
var file_pkg_pb_protos_auth_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pb_protos_auth_common_proto_init() }
func file_pkg_pb_protos_auth_common_proto_init() {
	if File_pkg_pb_protos_auth_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_protos_auth_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_pkg_pb_protos_auth_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Role); i {
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
			RawDescriptor: file_pkg_pb_protos_auth_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_pb_protos_auth_common_proto_goTypes,
		DependencyIndexes: file_pkg_pb_protos_auth_common_proto_depIdxs,
		MessageInfos:      file_pkg_pb_protos_auth_common_proto_msgTypes,
	}.Build()
	File_pkg_pb_protos_auth_common_proto = out.File
	file_pkg_pb_protos_auth_common_proto_rawDesc = nil
	file_pkg_pb_protos_auth_common_proto_goTypes = nil
	file_pkg_pb_protos_auth_common_proto_depIdxs = nil
}
