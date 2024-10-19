package main

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Danya97i/chat-server/pkg/chat_v1"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type server struct {
	pb.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

// CreateChat - метод для создания нового чата
func (s *server) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	log.Println("create chat request", req)

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	insertChatQueryBuilder := squirrel.Insert("chats").
		PlaceholderFormat(squirrel.Dollar).
		Columns("title").
		Values(req.Title).
		Suffix("returning id")
	insertChatQuery, args, err := insertChatQueryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	row := tx.QueryRow(ctx, insertChatQuery, args...)
	var chatID int64
	if err := row.Scan(&chatID); err != nil {
		return nil, err
	}

	insertUsersQueryBuilder := squirrel.Insert("chat_users").PlaceholderFormat(squirrel.Dollar).
		Columns("chat_id", "user_email")
	for _, email := range req.UserEmails {
		insertUsersQueryBuilder = insertUsersQueryBuilder.Values(chatID, email)
	}
	insertUsersQuery, args, err := insertUsersQueryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(ctx, insertUsersQuery, args...)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return &pb.CreateChatResponse{Id: chatID}, nil
}

// DeleteChat - метод для удаления чата
func (s *server) DeleteChat(ctx context.Context, req *pb.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Println("create chat request", req)

	deleteChatQueryBuilder := squirrel.Delete("chats").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": req.Id})
	deleteChatQuery, args, err := deleteChatQueryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = s.pool.Exec(ctx, deleteChatQuery, args...)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Метод для отправки сообщения в чат
func (s *server) SendMesage(_ context.Context, req *pb.SendMessageRequest) (*emptypb.Empty, error) {
	log.Println("send message request", req)
	return nil, nil
}
