package service

import (
	"context"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/romik1505/youtubeThumbnails/internal/app/model"
	"github.com/romik1505/youtubeThumbnails/internal/app/store"
	mock_thumbnail "github.com/romik1505/youtubeThumbnails/internal/pkg/mock/repository"
	mock_service "github.com/romik1505/youtubeThumbnails/internal/pkg/mock/service"
	desc "github.com/romik1505/youtubeThumbnails/pkg/api/thumbnails"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestThumbnailService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rep := mock_thumbnail.NewMockIThumbnailRepository(ctrl)
	loader := mock_service.NewMockIFileLoader(ctrl)
	ctx := context.Background()

	ts := ThumbnailService{
		Storage: rep,
		Loader:  loader,
	}

	tests := []struct {
		name       string
		request    *desc.GetRequest
		want       *desc.GetResponse
		wantErr    bool
		hookBefore func()
	}{
		{
			name: "image returned from db",
			request: &desc.GetRequest{
				Url: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			},
			want: &desc.GetResponse{
				Image: []byte("6542"),
			},
			wantErr: false,
			hookBefore: func() {
				rep.EXPECT().GetThumbnail(gomock.Any(), "dQw4w9WgXcQ").Return(model.Thumbnail{
					IDVideo:        store.NewNullString("dQw4w9WgXcQ"),
					ThumbnailImage: []byte("6542"),
					CreatedAt:      store.NewNullTime(time.Now()),
				}, nil)
			},
		},
		{
			name: "image loaded from youtube and saved in db",
			request: &desc.GetRequest{
				Url: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			},
			want: &desc.GetResponse{
				Image: []byte("6542"),
			},
			wantErr: false,
			hookBefore: func() {
				rep.EXPECT().GetThumbnail(gomock.Any(), "dQw4w9WgXcQ").Return(model.Thumbnail{}, sql.ErrNoRows)

				loader.EXPECT().LoadImg(gomock.Any(), "dQw4w9WgXcQ").Return([]byte("6542"), nil)

				rep.EXPECT().InsertThumbnail(gomock.Any(), model.Thumbnail{
					IDVideo:        store.NewNullString("dQw4w9WgXcQ"),
					ThumbnailImage: []byte("6542"),
				}).Return(model.Thumbnail{
					IDVideo:        store.NewNullString("dQw4w9WgXcQ"),
					ThumbnailImage: []byte("6542"),
					CreatedAt:      store.NewNullTime(time.Now()),
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hookBefore()

			got, err := ts.Get(ctx, tt.request)
			if !tt.wantErr {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestThumbnailService_getIDByUrl(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		inputUrl string
		wantID   string
		wantErr  bool
	}{
		{
			name:     "ok",
			inputUrl: "https://www.youtube.com/watch?v=QeRNBV7aDJ4",
			wantID:   "QeRNBV7aDJ4",
			wantErr:  false,
		},
		{
			name:     "invalid url",
			inputUrl: "https://www.rutube.com/watch?v=QeRNBV7aDJ4",
			wantID:   "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := ThumbnailService{}
			got, err := ts.getIDByUrl(ctx, tt.inputUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIDByUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantID {
				t.Errorf("GetIDByUrl() got = %v, want %v", got, tt.wantID)
			}
		})
	}
}
