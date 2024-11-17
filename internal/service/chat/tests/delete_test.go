package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/Danya97i/platform_common/pkg/db"
	dbMocks "github.com/Danya97i/platform_common/pkg/db/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/Danya97i/chat-server/internal/models"
	"github.com/Danya97i/chat-server/internal/repository"
	repoMocks "github.com/Danya97i/chat-server/internal/repository/mocks"
	"github.com/Danya97i/chat-server/internal/service/chat"
)

func TestDelete(t *testing.T) {
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type logRepositoryMockFunc func(mc *minimock.Controller) repository.LogRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		logInfo = models.LogInfo{ChatID: id, Action: models.ActionDelete}

		deleteChatError = errors.New("Delete chat error")
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		chatRepositoryMock chatRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		txManagerMock      txManagerMockFunc
	}{{
		name: "success",
		args: args{
			ctx: ctx,
			id:  id,
		},
		err: nil,

		chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
			mock := repoMocks.NewChatRepositoryMock(mc)
			mock.DeleteMock.Expect(ctx, id).Return(nil)
			return mock
		},

		logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
			mock := repoMocks.NewLogRepositoryMock(mc)
			mock.SaveMock.Expect(ctx, logInfo).Return(nil)
			return mock
		},

		txManagerMock: func(mc *minimock.Controller) db.TxManager {
			mock := dbMocks.NewTxManagerMock(mc)
			mock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
				return f(ctx)
			})
			return mock
		},
	}, {
		name: "deletes chat error case",
		args: args{
			ctx: ctx,
			id:  id,
		},
		err: deleteChatError,

		chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
			mock := repoMocks.NewChatRepositoryMock(mc)
			mock.DeleteMock.Expect(ctx, id).Return(deleteChatError)
			return mock
		},

		logRepositoryMock: func(mc *minimock.Controller) repository.LogRepository {
			mock := repoMocks.NewLogRepositoryMock(mc)
			return mock
		},

		txManagerMock: func(mc *minimock.Controller) db.TxManager {
			mock := dbMocks.NewTxManagerMock(mc)
			mock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
				return f(ctx)
			})
			return mock
		},
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			chatRepoMock := tt.chatRepositoryMock(mc)
			logRepoMock := tt.logRepositoryMock(mc)
			txManager := tt.txManagerMock(mc)
			service := chat.NewMockService(chatRepoMock, logRepoMock, txManager)
			err := service.Delete(tt.args.ctx, tt.args.id)
			require.Equal(t, tt.err, err)
		})
	}
}
