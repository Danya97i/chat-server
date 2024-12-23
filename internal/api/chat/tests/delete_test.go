package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Danya97i/chat-server/internal/api/chat"
	"github.com/Danya97i/chat-server/internal/service"
	serviceMock "github.com/Danya97i/chat-server/internal/service/mocks"
	pb "github.com/Danya97i/chat-server/pkg/chat_v1"
)

func TestDeleteChat(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *pb.DeleteChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("service error")

		req = &pb.DeleteChatRequest{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				chatServiceMock := serviceMock.NewChatServiceMock(mc)
				chatServiceMock.DeleteMock.Expect(ctx, id).Return(nil)
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
				chatServiceMock.DeleteMock.Expect(ctx, id).Return(serviceErr)
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

				_, err := server.DeleteChat(tt.args.ctx, tt.args.req)
				require.Equal(t, tt.err, err)
			})
		})
	}
}
