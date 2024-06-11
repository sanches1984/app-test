package service

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/sanches1984/referral-bot/internal/app/model"
	"github.com/sanches1984/referral-bot/internal/app/storage"
	tg "github.com/sanches1984/tg-client"
)

type Service struct {
	client  *tg.Client
	storage *storage.Storage
	logger  zerolog.Logger
}

func New(token string, storage *storage.Storage, logger zerolog.Logger) (*Service, error) {
	tgClient, err := tg.New(token, tg.NewLoggerMiddleware(logger))
	if err != nil {
		return nil, err
	}
	return &Service{
		client:  tgClient,
		storage: storage,
		logger:  logger,
	}, nil
}

func (s *Service) Run(ctx context.Context) {
	s.client.HandleCommand(model.CommandStart, s.processStart)
	s.client.HandleCommand(model.CommandAddReferral, s.addReferral)
	//s.client.HandleCommand(model.CommandDeleteReferral, s.users.Start)
	//s.client.HandleCommand(model.CommandGetReferrals, s.users.Start)
	s.client.Listen(ctx)
}

func (s *Service) processStart(ctx context.Context, msg *tg.IncomingMessage) []tg.OutgoingMessage {
	return []tg.OutgoingMessage{{Type: tg.MessageDefault, UserID: msg.UserID, ChatID: msg.ChatID, Message: model.TextWelcome}}
}

func (s *Service) addReferral(ctx context.Context, msg *tg.IncomingMessage) []tg.OutgoingMessage {
	return []tg.OutgoingMessage{{Type: tg.MessageDefault, UserID: msg.UserID, ChatID: msg.ChatID, Message: model.TextNewReferral}}
}
