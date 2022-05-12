package service

import (
	"context"
	"github.com/romik1505/youtubeThumbnails/internal/app/store"
	"github.com/romik1505/youtubeThumbnails/internal/app/store/thumbnail"
	desc "github.com/romik1505/youtubeThumbnails/pkg/api/thumbnails"
	"net/http"
)

type ThumbnailService struct {
	Storage thumbnail.IThumbnailRepository
	Loader  IFileLoader
	desc.UnimplementedThumbnailServer
}

//go:generate mockgen -source=service.go -destination=../../pkg/mock/service/mock_service.go
type IThumbnailService interface {
	Get(context.Context, *desc.GetRequest) (*desc.GetResponse, error)
}

type FileLoader struct {
	Client http.Client
}

type IFileLoader interface {
	LoadImg(ctx context.Context, url string) ([]byte, error)
}

func NewThumbnailService(ctx context.Context, store store.Storage) *ThumbnailService {
	return &ThumbnailService{
		Loader: &FileLoader{
			Client: http.Client{},
		},
		Storage: thumbnail.NewThumbnailRepository(store),
	}
}
