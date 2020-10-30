// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: bookmerce/bookmerce.proto

package bookmerce

import (
	proto "github.com/golang/protobuf/proto"
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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author   string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Language string `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmerce_bookmerce_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_bookmerce_bookmerce_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_bookmerce_bookmerce_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type BookID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BookID) Reset() {
	*x = BookID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmerce_bookmerce_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookID) ProtoMessage() {}

func (x *BookID) ProtoReflect() protoreflect.Message {
	mi := &file_bookmerce_bookmerce_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookID.ProtoReflect.Descriptor instead.
func (*BookID) Descriptor() ([]byte, []int) {
	return file_bookmerce_bookmerce_proto_rawDescGZIP(), []int{1}
}

func (x *BookID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_bookmerce_bookmerce_proto protoreflect.FileDescriptor

var file_bookmerce_bookmerce_proto_rawDesc = []byte{
	0x0a, 0x19, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x6f, 0x6f,
	0x6b, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x22, 0x5e, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x69, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x4d, 0x65,
	0x72, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0f,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a,
	0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x49, 0x44, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x11, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44,
	0x1a, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4f, 0x6c, 0x6f, 0x72, 0x75, 0x6e, 0x74, 0x6f, 0x62, 0x69, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_bookmerce_bookmerce_proto_rawDescOnce sync.Once
	file_bookmerce_bookmerce_proto_rawDescData = file_bookmerce_bookmerce_proto_rawDesc
)

func file_bookmerce_bookmerce_proto_rawDescGZIP() []byte {
	file_bookmerce_bookmerce_proto_rawDescOnce.Do(func() {
		file_bookmerce_bookmerce_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookmerce_bookmerce_proto_rawDescData)
	})
	return file_bookmerce_bookmerce_proto_rawDescData
}

var file_bookmerce_bookmerce_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bookmerce_bookmerce_proto_goTypes = []interface{}{
	(*Book)(nil),   // 0: bookmerce.Book
	(*BookID)(nil), // 1: bookmerce.BookID
}
var file_bookmerce_bookmerce_proto_depIdxs = []int32{
	0, // 0: bookmerce.BookMerce.AddBook:input_type -> bookmerce.Book
	1, // 1: bookmerce.BookMerce.GetBook:input_type -> bookmerce.BookID
	1, // 2: bookmerce.BookMerce.AddBook:output_type -> bookmerce.BookID
	0, // 3: bookmerce.BookMerce.GetBook:output_type -> bookmerce.Book
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bookmerce_bookmerce_proto_init() }
func file_bookmerce_bookmerce_proto_init() {
	if File_bookmerce_bookmerce_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bookmerce_bookmerce_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_bookmerce_bookmerce_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookID); i {
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
			RawDescriptor: file_bookmerce_bookmerce_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookmerce_bookmerce_proto_goTypes,
		DependencyIndexes: file_bookmerce_bookmerce_proto_depIdxs,
		MessageInfos:      file_bookmerce_bookmerce_proto_msgTypes,
	}.Build()
	File_bookmerce_bookmerce_proto = out.File
	file_bookmerce_bookmerce_proto_rawDesc = nil
	file_bookmerce_bookmerce_proto_goTypes = nil
	file_bookmerce_bookmerce_proto_depIdxs = nil
}
