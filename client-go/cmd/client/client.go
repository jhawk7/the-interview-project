package main

import (
	"context"

	"interview-client/config"
	"interview-client/internal/consumer"
	"interview-client/internal/logger"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func makeAuthCall(c config.GrpcConfig) *grpc.ClientConn {
	ctx := context.Background()
	authconn, err := grpc.DialContext(
		ctx,
		c.Authserver,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		logger.LogError(errors.Wrap(err, "failed to connect to authserver"), true)
	}

	return authconn
}

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig()
	env := config.LoadEnv()

	apiconn, err := grpc.DialContext(
		ctx,
		cfg.Server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		logger.LogError(errors.Wrap(err, "failed to connect to service"), true)
	}

	authconn := makeAuthCall(cfg)
	consumer := consumer.New(apiconn, authconn)

	consumer.Authenticate(ctx, env)
	consumer.HelloWorld(ctx)
}
