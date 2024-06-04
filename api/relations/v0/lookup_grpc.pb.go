// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: relations/v0/lookup.proto

package v0

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KesselLookupService_LookupSubjects_FullMethodName = "/kessel.relations.v0.KesselLookupService/LookupSubjects"
)

// KesselLookupServiceClient is the client API for KesselLookupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KesselLookupServiceClient interface {
	LookupSubjects(ctx context.Context, in *LookupSubjectsRequest, opts ...grpc.CallOption) (KesselLookupService_LookupSubjectsClient, error)
}

type kesselLookupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKesselLookupServiceClient(cc grpc.ClientConnInterface) KesselLookupServiceClient {
	return &kesselLookupServiceClient{cc}
}

func (c *kesselLookupServiceClient) LookupSubjects(ctx context.Context, in *LookupSubjectsRequest, opts ...grpc.CallOption) (KesselLookupService_LookupSubjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &KesselLookupService_ServiceDesc.Streams[0], KesselLookupService_LookupSubjects_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &kesselLookupServiceLookupSubjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KesselLookupService_LookupSubjectsClient interface {
	Recv() (*LookupSubjectsResponse, error)
	grpc.ClientStream
}

type kesselLookupServiceLookupSubjectsClient struct {
	grpc.ClientStream
}

func (x *kesselLookupServiceLookupSubjectsClient) Recv() (*LookupSubjectsResponse, error) {
	m := new(LookupSubjectsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KesselLookupServiceServer is the server API for KesselLookupService service.
// All implementations must embed UnimplementedKesselLookupServiceServer
// for forward compatibility
type KesselLookupServiceServer interface {
	LookupSubjects(*LookupSubjectsRequest, KesselLookupService_LookupSubjectsServer) error
	mustEmbedUnimplementedKesselLookupServiceServer()
}

// UnimplementedKesselLookupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKesselLookupServiceServer struct {
}

func (UnimplementedKesselLookupServiceServer) LookupSubjects(*LookupSubjectsRequest, KesselLookupService_LookupSubjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method LookupSubjects not implemented")
}
func (UnimplementedKesselLookupServiceServer) mustEmbedUnimplementedKesselLookupServiceServer() {}

// UnsafeKesselLookupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KesselLookupServiceServer will
// result in compilation errors.
type UnsafeKesselLookupServiceServer interface {
	mustEmbedUnimplementedKesselLookupServiceServer()
}

func RegisterKesselLookupServiceServer(s grpc.ServiceRegistrar, srv KesselLookupServiceServer) {
	s.RegisterService(&KesselLookupService_ServiceDesc, srv)
}

func _KesselLookupService_LookupSubjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LookupSubjectsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KesselLookupServiceServer).LookupSubjects(m, &kesselLookupServiceLookupSubjectsServer{stream})
}

type KesselLookupService_LookupSubjectsServer interface {
	Send(*LookupSubjectsResponse) error
	grpc.ServerStream
}

type kesselLookupServiceLookupSubjectsServer struct {
	grpc.ServerStream
}

func (x *kesselLookupServiceLookupSubjectsServer) Send(m *LookupSubjectsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// KesselLookupService_ServiceDesc is the grpc.ServiceDesc for KesselLookupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KesselLookupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kessel.relations.v0.KesselLookupService",
	HandlerType: (*KesselLookupServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LookupSubjects",
			Handler:       _KesselLookupService_LookupSubjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "relations/v0/lookup.proto",
}
