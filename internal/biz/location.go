package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Location struct {
	Latitude float64
	Longitude float64
}

type LocationRepo interface {
	Search(ctx context.Context, country, province, city string) (Location, error)
}

type LocationUsecase struct {
	repo LocationRepo
}

func NewLocationUsecase(repo LocationRepo, logger log.Logger) *LocationUsecase {
	return &LocationUsecase{repo: repo}
}

func (uc *LocationUsecase) Search(ctx context.Context, country, province, city string) (Location, error) {
	return uc.repo.Search(ctx, country, province, city)
}
