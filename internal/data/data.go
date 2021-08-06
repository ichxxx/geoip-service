package data

import (
	"encoding/gob"
	"os"

	"geoip_serivce/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

type location map[string]map[string]map[string][2]float64

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRegionRepo, NewLocationRepo)

// Data .
type Data struct {
	region *ip2region.Ip2Region
	location location
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	regionDb, err := ip2region.New(c.Region.File)
	if err != nil {
		return &Data{}, nil, err
	}
	locationFile, err := os.Open(c.Location.File)
	if err != nil {
		return &Data{}, nil, err
	}
	defer locationFile.Close()
	location := location{}
	dec := gob.NewDecoder(locationFile)
	err = dec.Decode(&location)
	d := &Data{
		region:  regionDb,
		location: location,
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		d.region.Close()
	}
	return d, cleanup, err
}
