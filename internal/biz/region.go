package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Region struct {
	Country  string
	Province string
	City     string
	ISP      string
}

type RegionRepo interface {
	Search(ctx context.Context, ip string) (Region, error)
}

type RegionUsecase struct {
	repo RegionRepo
}

func NewRegionUsecase(repo RegionRepo, logger log.Logger) *RegionUsecase {
	return &RegionUsecase{repo: repo}
}

func (uc *RegionUsecase) Search(ctx context.Context, ip string) (region Region, err error) {
	return uc.repo.Search(ctx, ip)
}