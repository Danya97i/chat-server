package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/Danya97i/chat-server/internal/api/chat"
	"github.com/Danya97i/chat-server/internal/service"
	serviceMock "github.com/Danya97i/chat-server/internal/service/mocks"
	pb "github.com/Danya97i/chat-server/pkg/chat_v1"
)

func TestCreateChat(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *pb.CreateChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id     = gofakeit.Int64()
		title  = gofakeit.Username()
		emails = []string{
			gofakeit.Email(),
			gofakeit.Email(),
		}

		serviceErr = fmt.Errorf("service error")

		req = &pb.CreateChatRequest{
			Title:      title,
			UserEmails: emails,
		}

		res = &pb.CreateChatResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *pb.CreateChatResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				chatServiceMock := serviceMock.NewChatServiceMock(mc)
				chatServiceMock.CreateMock.Expect(ctx, title, emails).Return(id, nil)
				return chatServiceMock
			},
		},
		{
			name: "service error",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				chatServiceMock := serviceMock.NewChatServiceMock(mc)
				chatServiceMock.CreateMock.Expect(ctx, title, emails).Return(0, serviceErr)
				return chatServiceMock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				chatServiceMock := tt.chatServiceMock(mc)
				server := chat.NewServer(chatServiceMock)

				resp, err := server.CreateChat(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
				require.Equal(t, tt.want, resp)
			})
		})
	}
}
