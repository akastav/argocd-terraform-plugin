// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: yandex/cloud/mdb/elasticsearch/v1/extension.proto

package elasticsearch

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type Extension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the extension.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique ID of the extension.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the Elasticsearch cluster the extension belongs to.
	ClusterId string `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Version of the extension.
	Version int64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	// The flag shows whether the extension is active.
	Active bool `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Extension) Reset() {
	*x = Extension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension) ProtoMessage() {}

func (x *Extension) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension.ProtoReflect.Descriptor instead.
func (*Extension) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDescGZIP(), []int{0}
}

func (x *Extension) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Extension) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Extension) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *Extension) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Extension) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type ExtensionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the extension.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// URI of the zip archive to create the new extension from. Currently only supports links that are stored in Yandex Object Storage.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// The flag shows whether to create the extension in disabled state.
	Disabled bool `protobuf:"varint,3,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

func (x *ExtensionSpec) Reset() {
	*x = ExtensionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionSpec) ProtoMessage() {}

func (x *ExtensionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionSpec.ProtoReflect.Descriptor instead.
func (*ExtensionSpec) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDescGZIP(), []int{1}
}

func (x *ExtensionSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExtensionSpec) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *ExtensionSpec) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

var File_yandex_cloud_mdb_elasticsearch_v1_extension_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDesc = []byte{
	0x0a, 0x31, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x5f, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31,
	0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x7c, 0x0a, 0x25, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64,
	0x62, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDescData = file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDesc
)

func file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDescData
}

var file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_goTypes = []interface{}{
	(*Extension)(nil),     // 0: yandex.cloud.mdb.elasticsearch.v1.Extension
	(*ExtensionSpec)(nil), // 1: yandex.cloud.mdb.elasticsearch.v1.ExtensionSpec
}
var file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_init() }
func file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_init() {
	if File_yandex_cloud_mdb_elasticsearch_v1_extension_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extension); i {
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
		file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionSpec); i {
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
			RawDescriptor: file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_elasticsearch_v1_extension_proto = out.File
	file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_rawDesc = nil
	file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_goTypes = nil
	file_yandex_cloud_mdb_elasticsearch_v1_extension_proto_depIdxs = nil
}
