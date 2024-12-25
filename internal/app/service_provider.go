package app

import (
	"context"

	"github.com/Danya97i/auth/pkg/access_v1"
	"github.com/Danya97i/platform_common/pkg/closer"
	"github.com/Danya97i/platform_common/pkg/db"
	"github.com/Danya97i/platform_common/pkg/db/pg"
	"github.com/Danya97i/platform_common/pkg/db/transaction"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	chatServer "github.com/Danya97i/chat-server/internal/api/chat"
	"github.com/Danya97i/chat-server/internal/client/auth"
	"github.com/Danya97i/chat-server/internal/config"
	"github.com/Danya97i/chat-server/internal/config/env"
	"github.com/Danya97i/chat-server/internal/repository"
	chatRepo "github.com/Danya97i/chat-server/internal/repository/chat"
	logRepo "github.com/Danya97i/chat-server/internal/repository/logs"
	"github.com/Danya97i/chat-server/internal/service"
	chatServ "github.com/Danya97i/chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig         config.PGConfig
	grpcConfig       config.GRPCConfig
	authClientConfig config.AuthClientConfig

	dbClient  db.Client
	txManager db.TxManager

	chatRepository repository.ChatRepository
	logRepository  repository.LogRepository

	chatService service.ChatService

	chatServer *chatServer.Server

	accessClient auth.AccessClient
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig returns config for postgres
func (sp *serviceProvider) PGConfig() config.PGConfig {
	if sp.pgConfig == nil {
		pgConfig, err := env.NewPgConfig()
		if err != nil {
			panic(err)
		}
		sp.pgConfig = pgConfig
	}
	return sp.pgConfig
}

// GRPCConfig returns config for grpc
func (sp *serviceProvider) GRPCConfig() config.GRPCConfig {
	if sp.grpcConfig == nil {
		grpcConfig, err := env.NewGrpcConfig()
		if err != nil {
			panic(err)
		}
		sp.grpcConfig = grpcConfig
	}
	return sp.grpcConfig
}

// AuthClientConfig returns config for auth client
func (sp *serviceProvider) AuthClientConfig() config.AuthClientConfig {
	if sp.authClientConfig == nil {
		authClientConfig, err := env.NewAuthClientConfig()
		if err != nil {
			panic(err)
		}
		sp.authClientConfig = authClientConfig
	}
	return sp.authClientConfig
}

// DBClient returns db client
func (sp *serviceProvider) DBClient(ctx context.Context) db.Client {
	if sp.dbClient == nil {
		client, err := pg.NewPGClient(ctx, sp.PGConfig().DSN())
		if err != nil {
			panic(err)
		}
		if err := client.DB().Ping(ctx); err != nil {
			panic(err)
		}
		closer.Add(client.Close)
		sp.dbClient = client
	}
	return sp.dbClient
}

// TxManager returns transaction manager
func (sp *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if sp.txManager == nil {
		sp.txManager = transaction.NewTransactionManager(sp.DBClient(ctx).DB())
	}
	return sp.txManager
}

// ChatRepository returns chat repository
func (sp *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if sp.chatRepository == nil {
		sp.chatRepository = chatRepo.NewRepository(sp.DBClient(ctx))
	}
	return sp.chatRepository
}

// ChatService returns chat service
func (sp *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if sp.chatService == nil {
		sp.chatService = chatServ.NewService(sp.ChatRepository(ctx), sp.LogRepository(ctx), sp.TxManager(ctx))
	}
	return sp.chatService
}

// ChatServer returns chat server
func (sp *serviceProvider) ChatServer(ctx context.Context) *chatServer.Server {
	if sp.chatServer == nil {
		sp.chatServer = chatServer.NewServer(sp.ChatService(ctx))
	}
	return sp.chatServer
}

// LogRepository returns log repository
func (sp *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if sp.logRepository == nil {
		sp.logRepository = logRepo.NewRepository(sp.DBClient(ctx))
	}
	return sp.logRepository
}

// AccessClient returns access client
func (sp *serviceProvider) AccessClient(_ context.Context) auth.AccessClient {
	if sp.accessClient == nil {
		creds, err := credentials.NewClientTLSFromFile(sp.AuthClientConfig().CertFile(), "")
		if err != nil {
			panic(err)
		}

		conn, err := grpc.NewClient(sp.authClientConfig.Address(), grpc.WithTransportCredentials(creds))
		if err != nil {
			panic(err)
		}

		closer.Add(conn.Close)

		client := auth.NewAccessClient(access_v1.NewAccessV1Client(conn))
		sp.accessClient = client
	}
	return sp.accessClient
}
