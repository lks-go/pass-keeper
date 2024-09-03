// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: pass_keeper.proto

package grpc_api

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

type RegisterUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterUserRequest) Reset() {
	*x = RegisterUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pass_keeper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserRequest) ProtoMessage() {}

func (x *RegisterUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pass_keeper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return file_pass_keeper_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterUserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *RegisterUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterUserResponse) Reset() {
	*x = RegisterUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pass_keeper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserResponse) ProtoMessage() {}

func (x *RegisterUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pass_keeper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return file_pass_keeper_proto_rawDescGZIP(), []int{1}
}

type AuthUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthUserRequest) Reset() {
	*x = AuthUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pass_keeper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserRequest) ProtoMessage() {}

func (x *AuthUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pass_keeper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserRequest.ProtoReflect.Descriptor instead.
func (*AuthUserRequest) Descriptor() ([]byte, []int) {
	return file_pass_keeper_proto_rawDescGZIP(), []int{2}
}

func (x *AuthUserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *AuthUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthUserResponse) Reset() {
	*x = AuthUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pass_keeper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserResponse) ProtoMessage() {}

func (x *AuthUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pass_keeper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserResponse.ProtoReflect.Descriptor instead.
func (*AuthUserResponse) Descriptor() ([]byte, []int) {
	return file_pass_keeper_proto_rawDescGZIP(), []int{3}
}

func (x *AuthUserResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AddDataLoginPassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Login string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Pass  string `protobuf:"bytes,3,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *AddDataLoginPassRequest) Reset() {
	*x = AddDataLoginPassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pass_keeper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDataLoginPassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDataLoginPassRequest) ProtoMessage() {}

func (x *AddDataLoginPassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pass_keeper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDataLoginPassRequest.ProtoReflect.Descriptor instead.
func (*AddDataLoginPassRequest) Descriptor() ([]byte, []int) {
	return file_pass_keeper_proto_rawDescGZIP(), []int{4}
}

func (x *AddDataLoginPassRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddDataLoginPassRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *AddDataLoginPassRequest) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

type AddDataLoginPassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddDataLoginPassResponse) Reset() {
	*x = AddDataLoginPassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pass_keeper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDataLoginPassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDataLoginPassResponse) ProtoMessage() {}

func (x *AddDataLoginPassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pass_keeper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDataLoginPassResponse.ProtoReflect.Descriptor instead.
func (*AddDataLoginPassResponse) Descriptor() ([]byte, []int) {
	return file_pass_keeper_proto_rawDescGZIP(), []int{5}
}

var File_pass_keeper_proto protoreflect.FileDescriptor

var file_pass_keeper_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x22, 0x47, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x43, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x59, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x41,
	0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8b, 0x02, 0x0a, 0x0a, 0x50, 0x61, 0x73, 0x73,
	0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x41,
	0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x12, 0x24, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pass_keeper_proto_rawDescOnce sync.Once
	file_pass_keeper_proto_rawDescData = file_pass_keeper_proto_rawDesc
)

func file_pass_keeper_proto_rawDescGZIP() []byte {
	file_pass_keeper_proto_rawDescOnce.Do(func() {
		file_pass_keeper_proto_rawDescData = protoimpl.X.CompressGZIP(file_pass_keeper_proto_rawDescData)
	})
	return file_pass_keeper_proto_rawDescData
}

var file_pass_keeper_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pass_keeper_proto_goTypes = []any{
	(*RegisterUserRequest)(nil),      // 0: pass.keeper.RegisterUserRequest
	(*RegisterUserResponse)(nil),     // 1: pass.keeper.RegisterUserResponse
	(*AuthUserRequest)(nil),          // 2: pass.keeper.AuthUserRequest
	(*AuthUserResponse)(nil),         // 3: pass.keeper.AuthUserResponse
	(*AddDataLoginPassRequest)(nil),  // 4: pass.keeper.AddDataLoginPassRequest
	(*AddDataLoginPassResponse)(nil), // 5: pass.keeper.AddDataLoginPassResponse
}
var file_pass_keeper_proto_depIdxs = []int32{
	0, // 0: pass.keeper.PassKeeper.RegisterUser:input_type -> pass.keeper.RegisterUserRequest
	2, // 1: pass.keeper.PassKeeper.AuthUser:input_type -> pass.keeper.AuthUserRequest
	4, // 2: pass.keeper.PassKeeper.AddDataLoginPass:input_type -> pass.keeper.AddDataLoginPassRequest
	1, // 3: pass.keeper.PassKeeper.RegisterUser:output_type -> pass.keeper.RegisterUserResponse
	3, // 4: pass.keeper.PassKeeper.AuthUser:output_type -> pass.keeper.AuthUserResponse
	5, // 5: pass.keeper.PassKeeper.AddDataLoginPass:output_type -> pass.keeper.AddDataLoginPassResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pass_keeper_proto_init() }
func file_pass_keeper_proto_init() {
	if File_pass_keeper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pass_keeper_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterUserRequest); i {
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
		file_pass_keeper_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterUserResponse); i {
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
		file_pass_keeper_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AuthUserRequest); i {
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
		file_pass_keeper_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AuthUserResponse); i {
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
		file_pass_keeper_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AddDataLoginPassRequest); i {
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
		file_pass_keeper_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AddDataLoginPassResponse); i {
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
			RawDescriptor: file_pass_keeper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pass_keeper_proto_goTypes,
		DependencyIndexes: file_pass_keeper_proto_depIdxs,
		MessageInfos:      file_pass_keeper_proto_msgTypes,
	}.Build()
	File_pass_keeper_proto = out.File
	file_pass_keeper_proto_rawDesc = nil
	file_pass_keeper_proto_goTypes = nil
	file_pass_keeper_proto_depIdxs = nil
}
