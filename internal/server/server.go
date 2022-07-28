package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/diyliv/manga-service/configs"
	mangaService "github.com/diyliv/manga-service/internal/manga/delivery/grpc"
	mangapb "github.com/diyliv/manga-service/proto/manga"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	logger *zap.Logger
	cfg    *configs.Config
}

func NewServer(logger *zap.Logger, cfg *configs.Config) *server {
	return &server{
		logger: logger,
		cfg:    cfg,
	}
}

func (s *server) RungRPC() {
	s.logger.Info(fmt.Sprintf("Starting GRPC server on port %s", s.cfg.Server.Port))

	lis, err := net.Listen("tcp", s.cfg.Server.Host+s.cfg.Server.Port)
	if err != nil {
		s.logger.Error("Error while listening: " + err.Error())
	}

	mangaService := mangaService.NewMangaService(s.logger)

	serv := grpc.NewServer()
	mangapb.RegisterMangaServiceServer(serv, mangaService)

	go func() {
		if err := serv.Serve(lis); err != nil {
			s.logger.Error("Error while serving: " + err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
	s.logger.Info("Exiting was successful")
}
