// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: cqlcrude/cqlcrudepb/cqlcrude_services.proto

package cqlcrude

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_cqlcrude_cqlcrudepb_cqlcrude_services_proto protoreflect.FileDescriptor

var file_cqlcrude_cqlcrudepb_cqlcrude_services_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64, 0x65, 0x2f, 0x63, 0x71, 0x6c, 0x63, 0x72,
	0x75, 0x64, 0x65, 0x70, 0x62, 0x2f, 0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63,
	0x71, 0x6c, 0x63, 0x72, 0x75, 0x64, 0x65, 0x1a, 0x2b, 0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64,
	0x65, 0x2f, 0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64, 0x65, 0x70, 0x62, 0x2f, 0x63, 0x71, 0x6c,
	0x63, 0x72, 0x75, 0x64, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x95, 0x03, 0x0a, 0x0c, 0x43, 0x72, 0x75, 0x64, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x63,
	0x71, 0x6c, 0x63, 0x72, 0x75, 0x64, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64, 0x65, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x24, 0x2e, 0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x71, 0x6c,
	0x63, 0x72, 0x75, 0x64, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x12, 0x59, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x71, 0x6c, 0x63, 0x72,
	0x75, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c,
	0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64, 0x65, 0x2f, 0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64,
	0x65, 0x70, 0x62, 0x3b, 0x63, 0x71, 0x6c, 0x63, 0x72, 0x75, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_cqlcrude_cqlcrudepb_cqlcrude_services_proto_goTypes = []interface{}{
	(*SubmitTrackingEventRequest)(nil),  // 0: cqlcrude.SubmitTrackingEventRequest
	(*DeleteTrackingEventRequest)(nil),  // 1: cqlcrude.DeleteTrackingEventRequest
	(*ListTrackingEventRequest)(nil),    // 2: cqlcrude.ListTrackingEventRequest
	(*GetTrackingEventRequest)(nil),     // 3: cqlcrude.GetTrackingEventRequest
	(*SubmitTrackingEventResponse)(nil), // 4: cqlcrude.SubmitTrackingEventResponse
	(*DeleteTrackingEventResponse)(nil), // 5: cqlcrude.DeleteTrackingEventResponse
	(*ListTrackingEventResponse)(nil),   // 6: cqlcrude.ListTrackingEventResponse
	(*GetTrackingEventResponse)(nil),    // 7: cqlcrude.GetTrackingEventResponse
}
var file_cqlcrude_cqlcrudepb_cqlcrude_services_proto_depIdxs = []int32{
	0, // 0: cqlcrude.CrudeService.SubmitTrackingEvent:input_type -> cqlcrude.SubmitTrackingEventRequest
	1, // 1: cqlcrude.CrudeService.DeleteTrackingEvent:input_type -> cqlcrude.DeleteTrackingEventRequest
	2, // 2: cqlcrude.CrudeService.ListAllTrackingEvents:input_type -> cqlcrude.ListTrackingEventRequest
	3, // 3: cqlcrude.CrudeService.GetTrackingEvent:input_type -> cqlcrude.GetTrackingEventRequest
	4, // 4: cqlcrude.CrudeService.SubmitTrackingEvent:output_type -> cqlcrude.SubmitTrackingEventResponse
	5, // 5: cqlcrude.CrudeService.DeleteTrackingEvent:output_type -> cqlcrude.DeleteTrackingEventResponse
	6, // 6: cqlcrude.CrudeService.ListAllTrackingEvents:output_type -> cqlcrude.ListTrackingEventResponse
	7, // 7: cqlcrude.CrudeService.GetTrackingEvent:output_type -> cqlcrude.GetTrackingEventResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cqlcrude_cqlcrudepb_cqlcrude_services_proto_init() }
func file_cqlcrude_cqlcrudepb_cqlcrude_services_proto_init() {
	if File_cqlcrude_cqlcrudepb_cqlcrude_services_proto != nil {
		return
	}
	file_cqlcrude_cqlcrudepb_cqlcrude_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cqlcrude_cqlcrudepb_cqlcrude_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cqlcrude_cqlcrudepb_cqlcrude_services_proto_goTypes,
		DependencyIndexes: file_cqlcrude_cqlcrudepb_cqlcrude_services_proto_depIdxs,
	}.Build()
	File_cqlcrude_cqlcrudepb_cqlcrude_services_proto = out.File
	file_cqlcrude_cqlcrudepb_cqlcrude_services_proto_rawDesc = nil
	file_cqlcrude_cqlcrudepb_cqlcrude_services_proto_goTypes = nil
	file_cqlcrude_cqlcrudepb_cqlcrude_services_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CrudeServiceClient is the client API for CrudeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrudeServiceClient interface {
	// This is the service to create the event
	SubmitTrackingEvent(ctx context.Context, in *SubmitTrackingEventRequest, opts ...grpc.CallOption) (*SubmitTrackingEventResponse, error)
	// This is the service to delete the event
	DeleteTrackingEvent(ctx context.Context, in *DeleteTrackingEventRequest, opts ...grpc.CallOption) (*DeleteTrackingEventResponse, error)
	// This is to list all events
	ListAllTrackingEvents(ctx context.Context, in *ListTrackingEventRequest, opts ...grpc.CallOption) (CrudeService_ListAllTrackingEventsClient, error)
	GetTrackingEvent(ctx context.Context, in *GetTrackingEventRequest, opts ...grpc.CallOption) (*GetTrackingEventResponse, error)
}

type crudeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrudeServiceClient(cc grpc.ClientConnInterface) CrudeServiceClient {
	return &crudeServiceClient{cc}
}

func (c *crudeServiceClient) SubmitTrackingEvent(ctx context.Context, in *SubmitTrackingEventRequest, opts ...grpc.CallOption) (*SubmitTrackingEventResponse, error) {
	out := new(SubmitTrackingEventResponse)
	err := c.cc.Invoke(ctx, "/cqlcrude.CrudeService/SubmitTrackingEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudeServiceClient) DeleteTrackingEvent(ctx context.Context, in *DeleteTrackingEventRequest, opts ...grpc.CallOption) (*DeleteTrackingEventResponse, error) {
	out := new(DeleteTrackingEventResponse)
	err := c.cc.Invoke(ctx, "/cqlcrude.CrudeService/DeleteTrackingEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudeServiceClient) ListAllTrackingEvents(ctx context.Context, in *ListTrackingEventRequest, opts ...grpc.CallOption) (CrudeService_ListAllTrackingEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CrudeService_serviceDesc.Streams[0], "/cqlcrude.CrudeService/ListAllTrackingEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &crudeServiceListAllTrackingEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CrudeService_ListAllTrackingEventsClient interface {
	Recv() (*ListTrackingEventResponse, error)
	grpc.ClientStream
}

type crudeServiceListAllTrackingEventsClient struct {
	grpc.ClientStream
}

func (x *crudeServiceListAllTrackingEventsClient) Recv() (*ListTrackingEventResponse, error) {
	m := new(ListTrackingEventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crudeServiceClient) GetTrackingEvent(ctx context.Context, in *GetTrackingEventRequest, opts ...grpc.CallOption) (*GetTrackingEventResponse, error) {
	out := new(GetTrackingEventResponse)
	err := c.cc.Invoke(ctx, "/cqlcrude.CrudeService/GetTrackingEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrudeServiceServer is the server API for CrudeService service.
type CrudeServiceServer interface {
	// This is the service to create the event
	SubmitTrackingEvent(context.Context, *SubmitTrackingEventRequest) (*SubmitTrackingEventResponse, error)
	// This is the service to delete the event
	DeleteTrackingEvent(context.Context, *DeleteTrackingEventRequest) (*DeleteTrackingEventResponse, error)
	// This is to list all events
	ListAllTrackingEvents(*ListTrackingEventRequest, CrudeService_ListAllTrackingEventsServer) error
	GetTrackingEvent(context.Context, *GetTrackingEventRequest) (*GetTrackingEventResponse, error)
}

// UnimplementedCrudeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCrudeServiceServer struct {
}

func (*UnimplementedCrudeServiceServer) SubmitTrackingEvent(context.Context, *SubmitTrackingEventRequest) (*SubmitTrackingEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTrackingEvent not implemented")
}
func (*UnimplementedCrudeServiceServer) DeleteTrackingEvent(context.Context, *DeleteTrackingEventRequest) (*DeleteTrackingEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrackingEvent not implemented")
}
func (*UnimplementedCrudeServiceServer) ListAllTrackingEvents(*ListTrackingEventRequest, CrudeService_ListAllTrackingEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAllTrackingEvents not implemented")
}
func (*UnimplementedCrudeServiceServer) GetTrackingEvent(context.Context, *GetTrackingEventRequest) (*GetTrackingEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrackingEvent not implemented")
}

func RegisterCrudeServiceServer(s *grpc.Server, srv CrudeServiceServer) {
	s.RegisterService(&_CrudeService_serviceDesc, srv)
}

func _CrudeService_SubmitTrackingEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTrackingEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudeServiceServer).SubmitTrackingEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cqlcrude.CrudeService/SubmitTrackingEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudeServiceServer).SubmitTrackingEvent(ctx, req.(*SubmitTrackingEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudeService_DeleteTrackingEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTrackingEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudeServiceServer).DeleteTrackingEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cqlcrude.CrudeService/DeleteTrackingEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudeServiceServer).DeleteTrackingEvent(ctx, req.(*DeleteTrackingEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudeService_ListAllTrackingEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTrackingEventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrudeServiceServer).ListAllTrackingEvents(m, &crudeServiceListAllTrackingEventsServer{stream})
}

type CrudeService_ListAllTrackingEventsServer interface {
	Send(*ListTrackingEventResponse) error
	grpc.ServerStream
}

type crudeServiceListAllTrackingEventsServer struct {
	grpc.ServerStream
}

func (x *crudeServiceListAllTrackingEventsServer) Send(m *ListTrackingEventResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CrudeService_GetTrackingEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrackingEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudeServiceServer).GetTrackingEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cqlcrude.CrudeService/GetTrackingEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudeServiceServer).GetTrackingEvent(ctx, req.(*GetTrackingEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CrudeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cqlcrude.CrudeService",
	HandlerType: (*CrudeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTrackingEvent",
			Handler:    _CrudeService_SubmitTrackingEvent_Handler,
		},
		{
			MethodName: "DeleteTrackingEvent",
			Handler:    _CrudeService_DeleteTrackingEvent_Handler,
		},
		{
			MethodName: "GetTrackingEvent",
			Handler:    _CrudeService_GetTrackingEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAllTrackingEvents",
			Handler:       _CrudeService_ListAllTrackingEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cqlcrude/cqlcrudepb/cqlcrude_services.proto",
}