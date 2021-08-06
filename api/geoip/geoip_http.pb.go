// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.3

package geoip

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type GeoipHTTPServer interface {
	GetGeoInfo(context.Context, *GetGeoInfoRequest) (*GetGeoInfoReply, error)
}

func RegisterGeoipHTTPServer(s *http.Server, srv GeoipHTTPServer) {
	r := s.Route("/")
	r.GET("/geoip", _Geoip_GetGeoInfo0_HTTP_Handler(srv))
}

func _Geoip_GetGeoInfo0_HTTP_Handler(srv GeoipHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGeoInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.geoip.Geoip/GetGeoInfo")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGeoInfo(ctx, req.(*GetGeoInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGeoInfoReply)
		return ctx.Result(200, reply)
	}
}

type GeoipHTTPClient interface {
	GetGeoInfo(ctx context.Context, req *GetGeoInfoRequest, opts ...http.CallOption) (rsp *GetGeoInfoReply, err error)
}

type GeoipHTTPClientImpl struct {
	cc *http.Client
}

func NewGeoipHTTPClient(client *http.Client) GeoipHTTPClient {
	return &GeoipHTTPClientImpl{client}
}

func (c *GeoipHTTPClientImpl) GetGeoInfo(ctx context.Context, in *GetGeoInfoRequest, opts ...http.CallOption) (*GetGeoInfoReply, error) {
	var out GetGeoInfoReply
	pattern := "/geoip"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.geoip.Geoip/GetGeoInfo"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
