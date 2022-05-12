package server

import (
	"context"
	"github.com/romik1505/youtubeThumbnails/internal/app/config"
	"github.com/romik1505/youtubeThumbnails/internal/app/service"
	desc "github.com/romik1505/youtubeThumbnails/pkg/api/thumbnails"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

type Server struct {
	config           config.Config
	ThumbnailService *service.ThumbnailService
}

func NewServer(ctx context.Context, ts *service.ThumbnailService) *Server {
	return &Server{
		config:           config.NewConfig(),
		ThumbnailService: ts,
	}
}

func (s *Server) Serve() {
	lis, err := net.Listen("tcp", s.config.GetValue(config.Port))
	if err != nil {
		log.Fatalf("Serve error %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	desc.RegisterThumbnailServer(grpcServer, s.ThumbnailService)

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Println(err.Error())
		}
	}()
	log.Println("Server started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done

	grpcServer.GracefulStop()
	log.Println("Server gracefully stopped")
}
