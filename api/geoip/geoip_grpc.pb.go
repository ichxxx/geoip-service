// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package geoip

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

// GeoipClient is the client API for Geoip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeoipClient interface {
	GetGeoInfo(ctx context.Context, in *GetGeoInfoRequest, opts ...grpc.CallOption) (*GetGeoInfoReply, error)
}

type geoipClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoipClient(cc grpc.ClientConnInterface) GeoipClient {
	return &geoipClient{cc}
}

func (c *geoipClient) GetGeoInfo(ctx context.Context, in *GetGeoInfoRequest, opts ...grpc.CallOption) (*GetGeoInfoReply, error) {
	out := new(GetGeoInfoReply)
	err := c.cc.Invoke(ctx, "/api.geoip.Geoip/GetGeoInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoipServer is the server API for Geoip service.
// All implementations must embed UnimplementedGeoipServer
// for forward compatibility
type GeoipServer interface {
	GetGeoInfo(context.Context, *GetGeoInfoRequest) (*GetGeoInfoReply, error)
	mustEmbedUnimplementedGeoipServer()
}

// UnimplementedGeoipServer must be embedded to have forward compatible implementations.
type UnimplementedGeoipServer struct {
}

func (UnimplementedGeoipServer) GetGeoInfo(context.Context, *GetGeoInfoRequest) (*GetGeoInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeoInfo not implemented")
}
func (UnimplementedGeoipServer) mustEmbedUnimplementedGeoipServer() {}

// UnsafeGeoipServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeoipServer will
// result in compilation errors.
type UnsafeGeoipServer interface {
	mustEmbedUnimplementedGeoipServer()
}

func RegisterGeoipServer(s grpc.ServiceRegistrar, srv GeoipServer) {
	s.RegisterService(&Geoip_ServiceDesc, srv)
}

func _Geoip_GetGeoInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGeoInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoipServer).GetGeoInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.geoip.Geoip/GetGeoInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoipServer).GetGeoInfo(ctx, req.(*GetGeoInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Geoip_ServiceDesc is the grpc.ServiceDesc for Geoip service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Geoip_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.geoip.Geoip",
	HandlerType: (*GeoipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGeoInfo",
			Handler:    _Geoip_GetGeoInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/geoip/geoip.proto",
}
