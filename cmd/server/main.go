package main

import (
	"context"
	"github.com/romik1505/youtubeThumbnails/internal/app/config"
	"github.com/romik1505/youtubeThumbnails/internal/app/server"
	"github.com/romik1505/youtubeThumbnails/internal/app/service"
)

func main() {
	ctx := context.Background()
	conf := config.NewConfig()

	ts := service.NewThumbnailService(ctx, config.NewSqliteConnection(ctx, conf.GetValue(config.SqliteConnection)))
	s := server.NewServer(ctx, ts)
	s.Serve()
}
