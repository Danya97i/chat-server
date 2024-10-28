package app

import (
	"context"

	chatServer "github.com/Danya97i/chat-server/internal/api/chat"
	"github.com/Danya97i/chat-server/internal/client/db"
	"github.com/Danya97i/chat-server/internal/client/db/pg"
	"github.com/Danya97i/chat-server/internal/client/db/transaction"
	"github.com/Danya97i/chat-server/internal/closer"
	"github.com/Danya97i/chat-server/internal/config"
	"github.com/Danya97i/chat-server/internal/config/env"
	"github.com/Danya97i/chat-server/internal/repository"
	chatRepo "github.com/Danya97i/chat-server/internal/repository/chat"
	logRepo "github.com/Danya97i/chat-server/internal/repository/logs"
	"github.com/Danya97i/chat-server/internal/service"
	chatServ "github.com/Danya97i/chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig       config.PGConfig
	grpcConfig     config.GRPCConfig
	dbClient       db.Client
	txManager      db.TxManager
	chatRepository repository.ChatRepository
	chatService    service.ChatService
	chatServer     *chatServer.Server
	logRepository  repository.LogRepository
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

func (sp *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if sp.logRepository == nil {
		sp.logRepository = logRepo.NewRepository(sp.DBClient(ctx))
	}
	return sp.logRepository
}
