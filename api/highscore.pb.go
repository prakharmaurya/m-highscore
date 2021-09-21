// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: highscore.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SetHighScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighScore float64 `protobuf:"fixed64,1,opt,name=high_score,json=highScore,proto3" json:"high_score,omitempty"`
}

func (x *SetHighScoreRequest) Reset() {
	*x = SetHighScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_highscore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHighScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHighScoreRequest) ProtoMessage() {}

func (x *SetHighScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_highscore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHighScoreRequest.ProtoReflect.Descriptor instead.
func (*SetHighScoreRequest) Descriptor() ([]byte, []int) {
	return file_highscore_proto_rawDescGZIP(), []int{0}
}

func (x *SetHighScoreRequest) GetHighScore() float64 {
	if x != nil {
		return x.HighScore
	}
	return 0
}

type SetHighScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set bool `protobuf:"varint,1,opt,name=set,proto3" json:"set,omitempty"`
}

func (x *SetHighScoreResponse) Reset() {
	*x = SetHighScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_highscore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHighScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHighScoreResponse) ProtoMessage() {}

func (x *SetHighScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_highscore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHighScoreResponse.ProtoReflect.Descriptor instead.
func (*SetHighScoreResponse) Descriptor() ([]byte, []int) {
	return file_highscore_proto_rawDescGZIP(), []int{1}
}

func (x *SetHighScoreResponse) GetSet() bool {
	if x != nil {
		return x.Set
	}
	return false
}

type GetHighScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHighScoreRequest) Reset() {
	*x = GetHighScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_highscore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHighScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHighScoreRequest) ProtoMessage() {}

func (x *GetHighScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_highscore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHighScoreRequest.ProtoReflect.Descriptor instead.
func (*GetHighScoreRequest) Descriptor() ([]byte, []int) {
	return file_highscore_proto_rawDescGZIP(), []int{2}
}

type GetHighScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighScore float64 `protobuf:"fixed64,1,opt,name=high_score,json=highScore,proto3" json:"high_score,omitempty"`
}

func (x *GetHighScoreResponse) Reset() {
	*x = GetHighScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_highscore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHighScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHighScoreResponse) ProtoMessage() {}

func (x *GetHighScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_highscore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHighScoreResponse.ProtoReflect.Descriptor instead.
func (*GetHighScoreResponse) Descriptor() ([]byte, []int) {
	return file_highscore_proto_rawDescGZIP(), []int{3}
}

func (x *GetHighScoreResponse) GetHighScore() float64 {
	if x != nil {
		return x.HighScore
	}
	return 0
}

var File_highscore_proto protoreflect.FileDescriptor

var file_highscore_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x34, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x48, 0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x68, 0x69,
	0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x28, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x48, 0x69,
	0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x65,
	0x74, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x48, 0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x48,
	0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x68, 0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x32,
	0x80, 0x01, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x48,
	0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x2e, 0x53, 0x65, 0x74, 0x48, 0x69,
	0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x53, 0x65, 0x74, 0x48, 0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x48, 0x69, 0x67, 0x68,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x67, 0x68, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x67, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_highscore_proto_rawDescOnce sync.Once
	file_highscore_proto_rawDescData = file_highscore_proto_rawDesc
)

func file_highscore_proto_rawDescGZIP() []byte {
	file_highscore_proto_rawDescOnce.Do(func() {
		file_highscore_proto_rawDescData = protoimpl.X.CompressGZIP(file_highscore_proto_rawDescData)
	})
	return file_highscore_proto_rawDescData
}

var file_highscore_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_highscore_proto_goTypes = []interface{}{
	(*SetHighScoreRequest)(nil),  // 0: SetHighScoreRequest
	(*SetHighScoreResponse)(nil), // 1: SetHighScoreResponse
	(*GetHighScoreRequest)(nil),  // 2: GetHighScoreRequest
	(*GetHighScoreResponse)(nil), // 3: GetHighScoreResponse
}
var file_highscore_proto_depIdxs = []int32{
	0, // 0: Game.SetHighScore:input_type -> SetHighScoreRequest
	2, // 1: Game.GetHighScore:input_type -> GetHighScoreRequest
	1, // 2: Game.SetHighScore:output_type -> SetHighScoreResponse
	3, // 3: Game.GetHighScore:output_type -> GetHighScoreResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_highscore_proto_init() }
func file_highscore_proto_init() {
	if File_highscore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_highscore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHighScoreRequest); i {
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
		file_highscore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHighScoreResponse); i {
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
		file_highscore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHighScoreRequest); i {
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
		file_highscore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHighScoreResponse); i {
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
			RawDescriptor: file_highscore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_highscore_proto_goTypes,
		DependencyIndexes: file_highscore_proto_depIdxs,
		MessageInfos:      file_highscore_proto_msgTypes,
	}.Build()
	File_highscore_proto = out.File
	file_highscore_proto_rawDesc = nil
	file_highscore_proto_goTypes = nil
	file_highscore_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameClient interface {
	SetHighScore(ctx context.Context, in *SetHighScoreRequest, opts ...grpc.CallOption) (*SetHighScoreResponse, error)
	GetHighScore(ctx context.Context, in *GetHighScoreRequest, opts ...grpc.CallOption) (*GetHighScoreResponse, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) SetHighScore(ctx context.Context, in *SetHighScoreRequest, opts ...grpc.CallOption) (*SetHighScoreResponse, error) {
	out := new(SetHighScoreResponse)
	err := c.cc.Invoke(ctx, "/Game/SetHighScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetHighScore(ctx context.Context, in *GetHighScoreRequest, opts ...grpc.CallOption) (*GetHighScoreResponse, error) {
	out := new(GetHighScoreResponse)
	err := c.cc.Invoke(ctx, "/Game/GetHighScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
type GameServer interface {
	SetHighScore(context.Context, *SetHighScoreRequest) (*SetHighScoreResponse, error)
	GetHighScore(context.Context, *GetHighScoreRequest) (*GetHighScoreResponse, error)
}

// UnimplementedGameServer can be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (*UnimplementedGameServer) SetHighScore(context.Context, *SetHighScoreRequest) (*SetHighScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHighScore not implemented")
}
func (*UnimplementedGameServer) GetHighScore(context.Context, *GetHighScoreRequest) (*GetHighScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHighScore not implemented")
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_SetHighScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetHighScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).SetHighScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Game/SetHighScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).SetHighScore(ctx, req.(*SetHighScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetHighScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHighScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetHighScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Game/GetHighScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetHighScore(ctx, req.(*GetHighScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetHighScore",
			Handler:    _Game_SetHighScore_Handler,
		},
		{
			MethodName: "GetHighScore",
			Handler:    _Game_GetHighScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "highscore.proto",
}