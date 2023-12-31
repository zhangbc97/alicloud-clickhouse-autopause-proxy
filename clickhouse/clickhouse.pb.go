// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: clickhouse/clickhouse.proto

package clickhouse

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

type KeepAliveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionID     string `protobuf:"bytes,1,opt,name=RegionID,proto3" json:"RegionID,omitempty"`
	DBInstanceID string `protobuf:"bytes,2,opt,name=DBInstanceID,proto3" json:"DBInstanceID,omitempty"`
}

func (x *KeepAliveRequest) Reset() {
	*x = KeepAliveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clickhouse_clickhouse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepAliveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepAliveRequest) ProtoMessage() {}

func (x *KeepAliveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clickhouse_clickhouse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepAliveRequest.ProtoReflect.Descriptor instead.
func (*KeepAliveRequest) Descriptor() ([]byte, []int) {
	return file_clickhouse_clickhouse_proto_rawDescGZIP(), []int{0}
}

func (x *KeepAliveRequest) GetRegionID() string {
	if x != nil {
		return x.RegionID
	}
	return ""
}

func (x *KeepAliveRequest) GetDBInstanceID() string {
	if x != nil {
		return x.DBInstanceID
	}
	return ""
}

type KeepAliveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *KeepAliveResponse) Reset() {
	*x = KeepAliveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clickhouse_clickhouse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepAliveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepAliveResponse) ProtoMessage() {}

func (x *KeepAliveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clickhouse_clickhouse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepAliveResponse.ProtoReflect.Descriptor instead.
func (*KeepAliveResponse) Descriptor() ([]byte, []int) {
	return file_clickhouse_clickhouse_proto_rawDescGZIP(), []int{1}
}

func (x *KeepAliveResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_clickhouse_clickhouse_proto protoreflect.FileDescriptor

var file_clickhouse_clickhouse_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x63, 0x6c, 0x69,
	0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a,
	0x10, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a,
	0x0c, 0x44, 0x42, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x44, 0x42, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x44, 0x22, 0x2d, 0x0a, 0x11, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0x48, 0x0a, 0x10, 0x41, 0x6c, 0x69, 0x59, 0x75, 0x6e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76,
	0x65, 0x12, 0x11, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f,
	0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_clickhouse_clickhouse_proto_rawDescOnce sync.Once
	file_clickhouse_clickhouse_proto_rawDescData = file_clickhouse_clickhouse_proto_rawDesc
)

func file_clickhouse_clickhouse_proto_rawDescGZIP() []byte {
	file_clickhouse_clickhouse_proto_rawDescOnce.Do(func() {
		file_clickhouse_clickhouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_clickhouse_clickhouse_proto_rawDescData)
	})
	return file_clickhouse_clickhouse_proto_rawDescData
}

var file_clickhouse_clickhouse_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_clickhouse_clickhouse_proto_goTypes = []interface{}{
	(*KeepAliveRequest)(nil),  // 0: KeepAliveRequest
	(*KeepAliveResponse)(nil), // 1: KeepAliveResponse
}
var file_clickhouse_clickhouse_proto_depIdxs = []int32{
	0, // 0: AliYunClickhouse.KeepAlive:input_type -> KeepAliveRequest
	1, // 1: AliYunClickhouse.KeepAlive:output_type -> KeepAliveResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_clickhouse_clickhouse_proto_init() }
func file_clickhouse_clickhouse_proto_init() {
	if File_clickhouse_clickhouse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clickhouse_clickhouse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepAliveRequest); i {
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
		file_clickhouse_clickhouse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepAliveResponse); i {
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
			RawDescriptor: file_clickhouse_clickhouse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clickhouse_clickhouse_proto_goTypes,
		DependencyIndexes: file_clickhouse_clickhouse_proto_depIdxs,
		MessageInfos:      file_clickhouse_clickhouse_proto_msgTypes,
	}.Build()
	File_clickhouse_clickhouse_proto = out.File
	file_clickhouse_clickhouse_proto_rawDesc = nil
	file_clickhouse_clickhouse_proto_goTypes = nil
	file_clickhouse_clickhouse_proto_depIdxs = nil
}
