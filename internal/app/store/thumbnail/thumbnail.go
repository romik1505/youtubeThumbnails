package thumbnail

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/romik1505/youtubeThumbnails/internal/app/model"
	"github.com/romik1505/youtubeThumbnails/internal/app/store"
)

type Repository struct {
	Storage store.Storage
}

func NewThumbnailRepository(storage store.Storage) IThumbnailRepository {
	return &Repository{
		Storage: storage,
	}
}

//go:generate mockgen -source=thumbnail.go -destination=../../../pkg/mock/repository/mock_rep.go
type IThumbnailRepository interface {
	InsertThumbnail(ctx context.Context, thumbnail model.Thumbnail) (model.Thumbnail, error)
	GetThumbnail(ctx context.Context, id string) (model.Thumbnail, error)
}

func (r Repository) InsertThumbnail(ctx context.Context, thumbnail model.Thumbnail) (model.Thumbnail, error) {
	query := r.Storage.Builder().
		Insert("thumbnails").
		Columns("id", "thumbnail").
		Values(thumbnail.IDVideo, thumbnail.ThumbnailImage).Suffix("RETURNING created_at")

	err := query.QueryRow().Scan(&thumbnail.CreatedAt)
	if err != nil {
		return model.Thumbnail{}, err
	}
	return thumbnail, nil
}

func (r Repository) GetThumbnail(ctx context.Context, id string) (model.Thumbnail, error) {
	query := r.Storage.Builder().Select("*").
		From("thumbnails").
		Where(sq.Eq{"id": id})

	ret := model.Thumbnail{}
	return ret, query.QueryRow().Scan(&ret.IDVideo, &ret.ThumbnailImage, &ret.CreatedAt)
}
