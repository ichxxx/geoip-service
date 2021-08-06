package data

import (
	"context"

	pb "geoip_serivce/api/geoip"
	"geoip_serivce/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type regionRepo struct {
	data *Data
	log *log.Helper
}

// NewRegionRepo .
func NewRegionRepo(data *Data, logger log.Logger) biz.RegionRepo {
	return &regionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func(rr *regionRepo) Search(ctx context.Context, ip string) (biz.Region, error) {
	info, err := rr.data.region.MemorySearch(ip)
	if err != nil {
		if err.Error() == "not found" {
			err = pb.ErrorGeoNotFound("region not found")
		}
		return biz.Region{}, err
	}
	return biz.Region{
		Country:  info.Country,
		Province: info.Province,
		City:     info.City,
		ISP:      info.ISP,
	}, nil
}