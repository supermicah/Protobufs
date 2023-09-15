// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: timeline.proto

package timeline

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

type Timeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NewsId     int64                  `protobuf:"varint,2,opt,name=news_id,json=newsId,proto3" json:"news_id,omitempty"`
	Content    string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	HappenTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=happen_time,json=happenTime,proto3" json:"happen_time,omitempty"`
	AttachIds  *int64                 `protobuf:"varint,5,opt,name=attach_ids,json=attachIds,proto3,oneof" json:"attach_ids,omitempty"`
	State      int32                  `protobuf:"varint,6,opt,name=state,proto3" json:"state,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Timeline) Reset() {
	*x = Timeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeline) ProtoMessage() {}

func (x *Timeline) ProtoReflect() protoreflect.Message {
	mi := &file_timeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeline.ProtoReflect.Descriptor instead.
func (*Timeline) Descriptor() ([]byte, []int) {
	return file_timeline_proto_rawDescGZIP(), []int{0}
}

func (x *Timeline) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Timeline) GetNewsId() int64 {
	if x != nil {
		return x.NewsId
	}
	return 0
}

func (x *Timeline) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Timeline) GetHappenTime() *timestamppb.Timestamp {
	if x != nil {
		return x.HappenTime
	}
	return nil
}

func (x *Timeline) GetAttachIds() int64 {
	if x != nil && x.AttachIds != nil {
		return *x.AttachIds
	}
	return 0
}

func (x *Timeline) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *Timeline) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Timeline) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_timeline_proto protoreflect.FileDescriptor

var file_timeline_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x02, 0x0a, 0x08,
	0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x73, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x68,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x68, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x49, 0x64, 0x73, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x73, 0x42, 0x40, 0x0a, 0x18, 0x77,
	0x69, 0x6b, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x61, 0x68, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x50, 0x01, 0x5a, 0x18, 0x6d, 0x69, 0x63, 0x61, 0x68, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x2f,
	0x6e, 0x65, 0x77, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_timeline_proto_rawDescOnce sync.Once
	file_timeline_proto_rawDescData = file_timeline_proto_rawDesc
)

func file_timeline_proto_rawDescGZIP() []byte {
	file_timeline_proto_rawDescOnce.Do(func() {
		file_timeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_timeline_proto_rawDescData)
	})
	return file_timeline_proto_rawDescData
}

var file_timeline_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_timeline_proto_goTypes = []interface{}{
	(*Timeline)(nil),              // 0: timeline.Timeline
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_timeline_proto_depIdxs = []int32{
	1, // 0: timeline.Timeline.happen_time:type_name -> google.protobuf.Timestamp
	1, // 1: timeline.Timeline.create_time:type_name -> google.protobuf.Timestamp
	1, // 2: timeline.Timeline.update_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_timeline_proto_init() }
func file_timeline_proto_init() {
	if File_timeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timeline); i {
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
	file_timeline_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_timeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_timeline_proto_goTypes,
		DependencyIndexes: file_timeline_proto_depIdxs,
		MessageInfos:      file_timeline_proto_msgTypes,
	}.Build()
	File_timeline_proto = out.File
	file_timeline_proto_rawDesc = nil
	file_timeline_proto_goTypes = nil
	file_timeline_proto_depIdxs = nil
}
