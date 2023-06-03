// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.1
// source: infotainment_message.proto

package pb

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

type Infotainment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Navigation         bool `protobuf:"varint,1,opt,name=navigation,proto3" json:"navigation,omitempty"`
	Bluetooth          bool `protobuf:"varint,2,opt,name=bluetooth,proto3" json:"bluetooth,omitempty"`
	Aux                bool `protobuf:"varint,3,opt,name=aux,proto3" json:"aux,omitempty"`
	TouchscreenDisplay bool `protobuf:"varint,4,opt,name=touchscreen_display,json=touchscreenDisplay,proto3" json:"touchscreen_display,omitempty"`
}

func (x *Infotainment) Reset() {
	*x = Infotainment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infotainment_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Infotainment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Infotainment) ProtoMessage() {}

func (x *Infotainment) ProtoReflect() protoreflect.Message {
	mi := &file_infotainment_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Infotainment.ProtoReflect.Descriptor instead.
func (*Infotainment) Descriptor() ([]byte, []int) {
	return file_infotainment_message_proto_rawDescGZIP(), []int{0}
}

func (x *Infotainment) GetNavigation() bool {
	if x != nil {
		return x.Navigation
	}
	return false
}

func (x *Infotainment) GetBluetooth() bool {
	if x != nil {
		return x.Bluetooth
	}
	return false
}

func (x *Infotainment) GetAux() bool {
	if x != nil {
		return x.Aux
	}
	return false
}

func (x *Infotainment) GetTouchscreenDisplay() bool {
	if x != nil {
		return x.TouchscreenDisplay
	}
	return false
}

var File_infotainment_message_proto protoreflect.FileDescriptor

var file_infotainment_message_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x66, 0x6f, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x8f, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x75, 0x65, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6c, 0x75, 0x65, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x75, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x75,
	0x78, 0x12, 0x2f, 0x0a, 0x13, 0x74, 0x6f, 0x75, 0x63, 0x68, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x74, 0x6f, 0x75, 0x63, 0x68, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x61, 0x72, 0x6b, 0x68, 0x61, 0x74, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x75, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infotainment_message_proto_rawDescOnce sync.Once
	file_infotainment_message_proto_rawDescData = file_infotainment_message_proto_rawDesc
)

func file_infotainment_message_proto_rawDescGZIP() []byte {
	file_infotainment_message_proto_rawDescOnce.Do(func() {
		file_infotainment_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_infotainment_message_proto_rawDescData)
	})
	return file_infotainment_message_proto_rawDescData
}

var file_infotainment_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infotainment_message_proto_goTypes = []interface{}{
	(*Infotainment)(nil), // 0: pb.Infotainment
}
var file_infotainment_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_infotainment_message_proto_init() }
func file_infotainment_message_proto_init() {
	if File_infotainment_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infotainment_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Infotainment); i {
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
			RawDescriptor: file_infotainment_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infotainment_message_proto_goTypes,
		DependencyIndexes: file_infotainment_message_proto_depIdxs,
		MessageInfos:      file_infotainment_message_proto_msgTypes,
	}.Build()
	File_infotainment_message_proto = out.File
	file_infotainment_message_proto_rawDesc = nil
	file_infotainment_message_proto_goTypes = nil
	file_infotainment_message_proto_depIdxs = nil
}
