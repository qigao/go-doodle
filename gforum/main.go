package main

import (
	"fmt"
	"gforum/db"
	pb "gforum/grpc/gen/proto/forum/v1"
	"gforum/handler"
	"gforum/repository"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	gorm_logger "logger/gorm"
	grpc_logger "logger/grpc"
	"net"
)

const (
	port = ":50051"
)

func main() {
	l := grpc_logger.NewZeroLogger("logs", "trace.grpc")
	grpc.NewServer(grpc_logger.UnaryInterceptorWithLogger(&l))
	d, err := db.NewDbWithLogger(gorm_logger.NewWithLogger(l))
	if err != nil {
		err = fmt.Errorf("failed to connect database: %w", err)
		l.Fatal().Err(err).Msg("failed to connect the database")
	}
	l.Info().
		Msg("succeeded to connect to the database")

	err = db.AutoMigrate(d)
	if err != nil {
		l.Fatal().Err(err).Msg("failed to migrate database")
	}

	us := repository.NewUserRepository(d)
	as := repository.NewArticleRepository(d)

	h := handler.New(&l, us, as)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		l.Panic().Err(fmt.Errorf("failed to listen: %w", err))
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)
	pb.RegisterUsersServer(s, h)
	pb.RegisterArticlesServer(s, h)
	l.Info().Str("port", port).Msg("starting server")
	if err := s.Serve(lis); err != nil {
		l.Panic().Err(fmt.Errorf("failed to serve: %w", err))
	}
}
