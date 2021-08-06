package service

import (
	"context"

	pb "geoip_serivce/api/geoip"
	"geoip_serivce/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type GeoipService struct {
	pb.UnimplementedGeoipServer
	log *log.Helper
	region *biz.RegionUsecase
	location *biz.LocationUsecase
}

// NewGeoipService new a geoip service.
func NewGeoipService(ruc *biz.RegionUsecase, luc *biz.LocationUsecase, logger log.Logger) *GeoipService {
	return &GeoipService{region: ruc, location: luc, log: log.NewHelper(logger)}
}

// GetGeoInfo implements geoip.GeoipServer
func (s *GeoipService) GetGeoInfo(ctx context.Context, req *pb.GetGeoInfoRequest) (*pb.GetGeoInfoReply, error) {
	region, err := s.region.Search(ctx, req.Ip)
	if err != nil {
		return &pb.GetGeoInfoReply{}, err
	}
	reply := &pb.GetGeoInfoReply{
		City:     zeroToEmpty(region.City),
		Country:  zeroToEmpty(region.Country),
		Province: zeroToEmpty(region.Province),
		Isp:      zeroToEmpty(region.ISP),
	}

	location, err := s.location.Search(ctx, region.Country, region.Province, region.City)
	if err != nil {
		return reply, err
	}

	reply.Location = &pb.GetGeoInfoReply_Location{
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}
	return reply, nil
}

func zeroToEmpty(s string) string {
	if s == "0" {
		return ""
	}
	return s
}