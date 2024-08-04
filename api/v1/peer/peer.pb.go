// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Hubble

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: peer/peer.proto

package peer

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

// ChangeNotificationType defines the peer change notification type.
type ChangeNotificationType int32

const (
	ChangeNotificationType_UNKNOWN      ChangeNotificationType = 0
	ChangeNotificationType_PEER_ADDED   ChangeNotificationType = 1
	ChangeNotificationType_PEER_DELETED ChangeNotificationType = 2
	ChangeNotificationType_PEER_UPDATED ChangeNotificationType = 3
)

// Enum value maps for ChangeNotificationType.
var (
	ChangeNotificationType_name = map[int32]string{
		0: "UNKNOWN",
		1: "PEER_ADDED",
		2: "PEER_DELETED",
		3: "PEER_UPDATED",
	}
	ChangeNotificationType_value = map[string]int32{
		"UNKNOWN":      0,
		"PEER_ADDED":   1,
		"PEER_DELETED": 2,
		"PEER_UPDATED": 3,
	}
)

func (x ChangeNotificationType) Enum() *ChangeNotificationType {
	p := new(ChangeNotificationType)
	*p = x
	return p
}

func (x ChangeNotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeNotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_peer_peer_proto_enumTypes[0].Descriptor()
}

func (ChangeNotificationType) Type() protoreflect.EnumType {
	return &file_peer_peer_proto_enumTypes[0]
}

func (x ChangeNotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeNotificationType.Descriptor instead.
func (ChangeNotificationType) EnumDescriptor() ([]byte, []int) {
	return file_peer_peer_proto_rawDescGZIP(), []int{0}
}

type NotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyRequest) Reset() {
	*x = NotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRequest) ProtoMessage() {}

func (x *NotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peer_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRequest.ProtoReflect.Descriptor instead.
func (*NotifyRequest) Descriptor() ([]byte, []int) {
	return file_peer_peer_proto_rawDescGZIP(), []int{0}
}

// ChangeNotification indicates a change regarding a hubble peer.
type ChangeNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the peer, typically the hostname. The name includes
	// the cluster name if a value other than default has been specified.
	// This value can be used to uniquely identify the host.
	// When the cluster name is not the default, the cluster name is prepended
	// to the peer name and a forward slash is added.
	//
	// Examples:
	//   - runtime1
	//   - testcluster/runtime1
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Address is the address of the peer's gRPC service.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// ChangeNotificationType indicates the type of change, ie whether the peer
	// was added, deleted or updated.
	Type ChangeNotificationType `protobuf:"varint,3,opt,name=type,proto3,enum=peer.ChangeNotificationType" json:"type,omitempty"`
	// TLS provides information to connect to the Address with TLS enabled.
	// If not set, TLS shall be assumed to be disabled.
	Tls *TLS `protobuf:"bytes,4,opt,name=tls,proto3" json:"tls,omitempty"`
}

func (x *ChangeNotification) Reset() {
	*x = ChangeNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeNotification) ProtoMessage() {}

func (x *ChangeNotification) ProtoReflect() protoreflect.Message {
	mi := &file_peer_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeNotification.ProtoReflect.Descriptor instead.
func (*ChangeNotification) Descriptor() ([]byte, []int) {
	return file_peer_peer_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeNotification) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChangeNotification) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ChangeNotification) GetType() ChangeNotificationType {
	if x != nil {
		return x.Type
	}
	return ChangeNotificationType_UNKNOWN
}

func (x *ChangeNotification) GetTls() *TLS {
	if x != nil {
		return x.Tls
	}
	return nil
}

// TLS provides information to establish a TLS connection to the peer.
type TLS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ServerName is used to verify the hostname on the returned certificate.
	ServerName string `protobuf:"bytes,1,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
}

func (x *TLS) Reset() {
	*x = TLS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_peer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLS) ProtoMessage() {}

func (x *TLS) ProtoReflect() protoreflect.Message {
	mi := &file_peer_peer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLS.ProtoReflect.Descriptor instead.
func (*TLS) Descriptor() ([]byte, []int) {
	return file_peer_peer_proto_rawDescGZIP(), []int{2}
}

func (x *TLS) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

var File_peer_peer_proto protoreflect.FileDescriptor

var file_peer_peer_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x70, 0x65, 0x65, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x12, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x65,
	0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70,
	0x65, 0x65, 0x72, 0x2e, 0x54, 0x4c, 0x53, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x22, 0x26, 0x0a, 0x03,
	0x54, 0x4c, 0x53, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x59, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50,
	0x45, 0x45, 0x52, 0x5f, 0x41, 0x44, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50,
	0x45, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a,
	0x0c, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x32,
	0x43, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x13, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_peer_proto_rawDescOnce sync.Once
	file_peer_peer_proto_rawDescData = file_peer_peer_proto_rawDesc
)

func file_peer_peer_proto_rawDescGZIP() []byte {
	file_peer_peer_proto_rawDescOnce.Do(func() {
		file_peer_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_peer_proto_rawDescData)
	})
	return file_peer_peer_proto_rawDescData
}

var file_peer_peer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_peer_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_peer_peer_proto_goTypes = []any{
	(ChangeNotificationType)(0), // 0: peer.ChangeNotificationType
	(*NotifyRequest)(nil),       // 1: peer.NotifyRequest
	(*ChangeNotification)(nil),  // 2: peer.ChangeNotification
	(*TLS)(nil),                 // 3: peer.TLS
}
var file_peer_peer_proto_depIdxs = []int32{
	0, // 0: peer.ChangeNotification.type:type_name -> peer.ChangeNotificationType
	3, // 1: peer.ChangeNotification.tls:type_name -> peer.TLS
	1, // 2: peer.Peer.Notify:input_type -> peer.NotifyRequest
	2, // 3: peer.Peer.Notify:output_type -> peer.ChangeNotification
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_peer_peer_proto_init() }
func file_peer_peer_proto_init() {
	if File_peer_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peer_peer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NotifyRequest); i {
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
		file_peer_peer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ChangeNotification); i {
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
		file_peer_peer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TLS); i {
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
			RawDescriptor: file_peer_peer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_peer_peer_proto_goTypes,
		DependencyIndexes: file_peer_peer_proto_depIdxs,
		EnumInfos:         file_peer_peer_proto_enumTypes,
		MessageInfos:      file_peer_peer_proto_msgTypes,
	}.Build()
	File_peer_peer_proto = out.File
	file_peer_peer_proto_rawDesc = nil
	file_peer_peer_proto_goTypes = nil
	file_peer_peer_proto_depIdxs = nil
}
