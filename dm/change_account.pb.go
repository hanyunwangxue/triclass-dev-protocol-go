// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: dm/change_account.proto

package dm

import (
	proto "github.com/golang/protobuf/proto"
	protocol_go "github.com/triclass/protocol-go"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ChangeAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 厂商
	Manufacturer protocol_go.Manufacturer `protobuf:"varint,1,opt,name=manufacturer,proto3,enum=triclass.manufacturer.Manufacturer" json:"manufacturer,omitempty"`
	// 设备序列号
	Sn string `protobuf:"bytes,2,opt,name=sn,proto3" json:"sn,omitempty"`
	// 云视讯账号
	Account string `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *ChangeAccount) Reset() {
	*x = ChangeAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dm_change_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAccount) ProtoMessage() {}

func (x *ChangeAccount) ProtoReflect() protoreflect.Message {
	mi := &file_dm_change_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAccount.ProtoReflect.Descriptor instead.
func (*ChangeAccount) Descriptor() ([]byte, []int) {
	return file_dm_change_account_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeAccount) GetManufacturer() protocol_go.Manufacturer {
	if x != nil {
		return x.Manufacturer
	}
	return protocol_go.Manufacturer_UNKNOWN
}

func (x *ChangeAccount) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

func (x *ChangeAccount) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

var File_dm_change_account_proto protoreflect.FileDescriptor

var file_dm_change_account_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x69, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x2e, 0x64, 0x6d, 0x1a, 0x12, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0d, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x0c,
	0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x6d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x66,
	0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x1c, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e,
	0x64, 0x6d, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x64, 0x6d, 0x3b, 0x64, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dm_change_account_proto_rawDescOnce sync.Once
	file_dm_change_account_proto_rawDescData = file_dm_change_account_proto_rawDesc
)

func file_dm_change_account_proto_rawDescGZIP() []byte {
	file_dm_change_account_proto_rawDescOnce.Do(func() {
		file_dm_change_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_dm_change_account_proto_rawDescData)
	})
	return file_dm_change_account_proto_rawDescData
}

var file_dm_change_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dm_change_account_proto_goTypes = []interface{}{
	(*ChangeAccount)(nil),         // 0: triclass.dm.ChangeAccount
	(protocol_go.Manufacturer)(0), // 1: triclass.manufacturer.Manufacturer
}
var file_dm_change_account_proto_depIdxs = []int32{
	1, // 0: triclass.dm.ChangeAccount.manufacturer:type_name -> triclass.manufacturer.Manufacturer
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dm_change_account_proto_init() }
func file_dm_change_account_proto_init() {
	if File_dm_change_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dm_change_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeAccount); i {
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
			RawDescriptor: file_dm_change_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dm_change_account_proto_goTypes,
		DependencyIndexes: file_dm_change_account_proto_depIdxs,
		MessageInfos:      file_dm_change_account_proto_msgTypes,
	}.Build()
	File_dm_change_account_proto = out.File
	file_dm_change_account_proto_rawDesc = nil
	file_dm_change_account_proto_goTypes = nil
	file_dm_change_account_proto_depIdxs = nil
}
