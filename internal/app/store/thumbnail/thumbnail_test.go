package thumbnail

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/romik1505/youtubeThumbnails/internal/app/config"
	"github.com/romik1505/youtubeThumbnails/internal/app/model"
	"github.com/romik1505/youtubeThumbnails/internal/app/store"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var Storage store.Storage

func TestMain(m *testing.M) {
	conf := config.NewConfig()
	Storage = config.NewSqliteConnection(context.Background(), conf.GetValue(config.SqliteConnection))
	os.Exit(m.Run())
}

func mustTruncate() {
	Storage.Exec("DELETE FROM thumbnails") // nolint
}

func TestRepository_InsertThumbnail(t *testing.T) {
	r := NewThumbnailRepository(Storage)
	ctx := context.Background()

	mustTruncate()

	tests := []struct {
		name    string
		input   model.Thumbnail
		want    model.Thumbnail
		wantErr bool
	}{
		{
			name: "ok",
			input: model.Thumbnail{
				IDVideo:        store.NewNullString("1234"),
				ThumbnailImage: []byte{1, 2, 3},
			},
			want: model.Thumbnail{
				IDVideo:        store.NewNullString("1234"),
				ThumbnailImage: []byte{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "second insert",
			input: model.Thumbnail{
				IDVideo:        store.NewNullString("1234"),
				ThumbnailImage: []byte{1, 2, 3},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.InsertThumbnail(ctx, tt.input)
			if !tt.wantErr {
				require.NoError(t, err)
				require.False(t, got.CreatedAt.Time.IsZero())
			}
			require.Empty(t, cmp.Diff(tt.want, got, cmpopts.IgnoreFields(model.Thumbnail{}, "CreatedAt")))
		})
	}
}

func TestRepository_GetThumbnail(t *testing.T) {
	r := NewThumbnailRepository(Storage)
	ctx := context.Background()
	mustTruncate()

	tests := []struct {
		name       string
		inputId    string
		want       model.Thumbnail
		wantErr    bool
		hookBefore func()
	}{
		{
			name:    "ok case",
			inputId: "5678",
			want: model.Thumbnail{
				IDVideo:        store.NewNullString("5678"),
				ThumbnailImage: []byte{8, 7, 6, 5},
			},
			wantErr: false,
			hookBefore: func() {
				r.InsertThumbnail(ctx, model.Thumbnail{ //nolint
					IDVideo:        store.NewNullString("5678"),
					ThumbnailImage: []byte{8, 7, 6, 5},
				})
			},
		},
		{
			name:    "object not found",
			inputId: "5678",
			want:    model.Thumbnail{},
			wantErr: true,
			hookBefore: func() {
				mustTruncate()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hookBefore()
			got, err := r.GetThumbnail(ctx, tt.inputId)
			if !tt.wantErr {
				require.NoError(t, err)
				require.False(t, got.CreatedAt.Time.IsZero())
			}
			require.Empty(t, cmp.Diff(tt.want, got, cmpopts.IgnoreFields(model.Thumbnail{}, "CreatedAt")))
		})
	}
}
