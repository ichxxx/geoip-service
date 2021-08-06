package data

import (
	"context"

	pb "geoip_serivce/api/geoip"
	"geoip_serivce/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type locationRepo struct {
	data *Data
	log *log.Helper
}

// NewLocationRepo .
func NewLocationRepo(data *Data, logger log.Logger) biz.LocationRepo {
	return &locationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func(lr *locationRepo) Search(ctx context.Context, country, province, city string) (biz.Location, error) {
	if pMap, ok := lr.data.location[country]; ok {
		if cMap, ok := pMap[province]; ok {
			if location, ok := cMap[city]; ok {
				return biz.Location{
					Latitude: location[0],
					Longitude: location[1],
				}, nil
			}
		}
	}
	return biz.Location{}, pb.ErrorGeoNotFound("location not found")
}
