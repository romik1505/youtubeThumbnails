// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: thumbnails/thumbnails.proto

package thumbnails

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

// ThumbnailClient is the client API for Thumbnail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThumbnailClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type thumbnailClient struct {
	cc grpc.ClientConnInterface
}

func NewThumbnailClient(cc grpc.ClientConnInterface) ThumbnailClient {
	return &thumbnailClient{cc}
}

func (c *thumbnailClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/thumbnails.Thumbnail/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThumbnailServer is the server API for Thumbnail service.
// All implementations must embed UnimplementedThumbnailServer
// for forward compatibility
type ThumbnailServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedThumbnailServer()
}

// UnimplementedThumbnailServer must be embedded to have forward compatible implementations.
type UnimplementedThumbnailServer struct {
}

func (UnimplementedThumbnailServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedThumbnailServer) mustEmbedUnimplementedThumbnailServer() {}

// UnsafeThumbnailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThumbnailServer will
// result in compilation errors.
type UnsafeThumbnailServer interface {
	mustEmbedUnimplementedThumbnailServer()
}

func RegisterThumbnailServer(s grpc.ServiceRegistrar, srv ThumbnailServer) {
	s.RegisterService(&Thumbnail_ServiceDesc, srv)
}

func _Thumbnail_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThumbnailServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thumbnails.Thumbnail/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThumbnailServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Thumbnail_ServiceDesc is the grpc.ServiceDesc for Thumbnail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Thumbnail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thumbnails.Thumbnail",
	HandlerType: (*ThumbnailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Thumbnail_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thumbnails/thumbnails.proto",
}