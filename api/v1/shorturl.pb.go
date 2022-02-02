// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: v1/shorturl.proto

package v1

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

type CreateShortUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	TtlSeconds  int64  `protobuf:"varint,2,opt,name=ttl_seconds,json=ttlSeconds,proto3" json:"ttl_seconds,omitempty"`
}

func (x *CreateShortUrlRequest) Reset() {
	*x = CreateShortUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shorturl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortUrlRequest) ProtoMessage() {}

func (x *CreateShortUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shorturl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortUrlRequest.ProtoReflect.Descriptor instead.
func (*CreateShortUrlRequest) Descriptor() ([]byte, []int) {
	return file_v1_shorturl_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShortUrlRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

func (x *CreateShortUrlRequest) GetTtlSeconds() int64 {
	if x != nil {
		return x.TtlSeconds
	}
	return 0
}

type CreateShortUrlReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortKey string `protobuf:"bytes,1,opt,name=short_key,json=shortKey,proto3" json:"short_key,omitempty"`
}

func (x *CreateShortUrlReply) Reset() {
	*x = CreateShortUrlReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shorturl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortUrlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortUrlReply) ProtoMessage() {}

func (x *CreateShortUrlReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shorturl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortUrlReply.ProtoReflect.Descriptor instead.
func (*CreateShortUrlReply) Descriptor() ([]byte, []int) {
	return file_v1_shorturl_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShortUrlReply) GetShortKey() string {
	if x != nil {
		return x.ShortKey
	}
	return ""
}

type GetOriginalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortKey string `protobuf:"bytes,1,opt,name=short_key,json=shortKey,proto3" json:"short_key,omitempty"`
}

func (x *GetOriginalRequest) Reset() {
	*x = GetOriginalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shorturl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOriginalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginalRequest) ProtoMessage() {}

func (x *GetOriginalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shorturl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginalRequest.ProtoReflect.Descriptor instead.
func (*GetOriginalRequest) Descriptor() ([]byte, []int) {
	return file_v1_shorturl_proto_rawDescGZIP(), []int{2}
}

func (x *GetOriginalRequest) GetShortKey() string {
	if x != nil {
		return x.ShortKey
	}
	return ""
}

type GetOriginalReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *GetOriginalReply) Reset() {
	*x = GetOriginalReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shorturl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOriginalReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginalReply) ProtoMessage() {}

func (x *GetOriginalReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shorturl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginalReply.ProtoReflect.Descriptor instead.
func (*GetOriginalReply) Descriptor() ([]byte, []int) {
	return file_v1_shorturl_proto_rawDescGZIP(), []int{3}
}

func (x *GetOriginalReply) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type RedirectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortKey string `protobuf:"bytes,1,opt,name=short_key,json=shortKey,proto3" json:"short_key,omitempty"`
}

func (x *RedirectRequest) Reset() {
	*x = RedirectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shorturl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectRequest) ProtoMessage() {}

func (x *RedirectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shorturl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectRequest.ProtoReflect.Descriptor instead.
func (*RedirectRequest) Descriptor() ([]byte, []int) {
	return file_v1_shorturl_proto_rawDescGZIP(), []int{4}
}

func (x *RedirectRequest) GetShortKey() string {
	if x != nil {
		return x.ShortKey
	}
	return ""
}

type RedirectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *RedirectReply) Reset() {
	*x = RedirectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shorturl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectReply) ProtoMessage() {}

func (x *RedirectReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shorturl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectReply.ProtoReflect.Descriptor instead.
func (*RedirectReply) Descriptor() ([]byte, []int) {
	return file_v1_shorturl_proto_rawDescGZIP(), []int{5}
}

func (x *RedirectReply) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

var File_v1_shorturl_proto protoreflect.FileDescriptor

var file_v1_shorturl_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x74, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x22, 0x32, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x31, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c,
	0x22, 0x2e, 0x0a, 0x0f, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79,
	0x22, 0x32, 0x0a, 0x0d, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x72, 0x6c, 0x32, 0x8d, 0x02, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72,
	0x6c, 0x12, 0x5f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22,
	0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x56, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x48, 0x0a, 0x08, 0x52, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x7b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x7d, 0x42, 0x41, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x19, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_shorturl_proto_rawDescOnce sync.Once
	file_v1_shorturl_proto_rawDescData = file_v1_shorturl_proto_rawDesc
)

func file_v1_shorturl_proto_rawDescGZIP() []byte {
	file_v1_shorturl_proto_rawDescOnce.Do(func() {
		file_v1_shorturl_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_shorturl_proto_rawDescData)
	})
	return file_v1_shorturl_proto_rawDescData
}

var file_v1_shorturl_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_shorturl_proto_goTypes = []interface{}{
	(*CreateShortUrlRequest)(nil), // 0: v1.CreateShortUrlRequest
	(*CreateShortUrlReply)(nil),   // 1: v1.CreateShortUrlReply
	(*GetOriginalRequest)(nil),    // 2: v1.GetOriginalRequest
	(*GetOriginalReply)(nil),      // 3: v1.GetOriginalReply
	(*RedirectRequest)(nil),       // 4: v1.RedirectRequest
	(*RedirectReply)(nil),         // 5: v1.RedirectReply
}
var file_v1_shorturl_proto_depIdxs = []int32{
	0, // 0: v1.ShortUrl.CreateShortUrl:input_type -> v1.CreateShortUrlRequest
	2, // 1: v1.ShortUrl.GetOriginalUrl:input_type -> v1.GetOriginalRequest
	4, // 2: v1.ShortUrl.Redirect:input_type -> v1.RedirectRequest
	1, // 3: v1.ShortUrl.CreateShortUrl:output_type -> v1.CreateShortUrlReply
	3, // 4: v1.ShortUrl.GetOriginalUrl:output_type -> v1.GetOriginalReply
	5, // 5: v1.ShortUrl.Redirect:output_type -> v1.RedirectReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_shorturl_proto_init() }
func file_v1_shorturl_proto_init() {
	if File_v1_shorturl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_shorturl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortUrlRequest); i {
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
		file_v1_shorturl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortUrlReply); i {
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
		file_v1_shorturl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOriginalRequest); i {
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
		file_v1_shorturl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOriginalReply); i {
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
		file_v1_shorturl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectRequest); i {
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
		file_v1_shorturl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectReply); i {
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
			RawDescriptor: file_v1_shorturl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_shorturl_proto_goTypes,
		DependencyIndexes: file_v1_shorturl_proto_depIdxs,
		MessageInfos:      file_v1_shorturl_proto_msgTypes,
	}.Build()
	File_v1_shorturl_proto = out.File
	file_v1_shorturl_proto_rawDesc = nil
	file_v1_shorturl_proto_goTypes = nil
	file_v1_shorturl_proto_depIdxs = nil
}
