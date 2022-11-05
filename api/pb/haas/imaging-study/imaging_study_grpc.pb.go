// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: haas/imaging-study/imaging_study.proto

package haas

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ImagingStudyServiceClient is the client API for ImagingStudyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImagingStudyServiceClient interface {
}

type imagingStudyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImagingStudyServiceClient(cc grpc.ClientConnInterface) ImagingStudyServiceClient {
	return &imagingStudyServiceClient{cc}
}

// ImagingStudyServiceServer is the server API for ImagingStudyService service.
// All implementations must embed UnimplementedImagingStudyServiceServer
// for forward compatibility
type ImagingStudyServiceServer interface {
	mustEmbedUnimplementedImagingStudyServiceServer()
}

// UnimplementedImagingStudyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImagingStudyServiceServer struct {
}

func (UnimplementedImagingStudyServiceServer) mustEmbedUnimplementedImagingStudyServiceServer() {}

// UnsafeImagingStudyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImagingStudyServiceServer will
// result in compilation errors.
type UnsafeImagingStudyServiceServer interface {
	mustEmbedUnimplementedImagingStudyServiceServer()
}

func RegisterImagingStudyServiceServer(s grpc.ServiceRegistrar, srv ImagingStudyServiceServer) {
	s.RegisterService(&ImagingStudyService_ServiceDesc, srv)
}

// ImagingStudyService_ServiceDesc is the grpc.ServiceDesc for ImagingStudyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImagingStudyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mnes.haas.imagingstudy.ImagingStudyService",
	HandlerType: (*ImagingStudyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "haas/imaging-study/imaging_study.proto",
}