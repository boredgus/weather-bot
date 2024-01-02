package app

import (
	"subscription-bot/internal/db"
	"subscription-bot/internal/domain"
)

type ConversationService struct {
	Repo domain.ConversationRepository
}

func NewConversationService(repo domain.ConversationRepository) domain.ConversationSvc {
	return ConversationService{Repo: repo}
}

func ConversationSvc() domain.ConversationSvc {
	return NewConversationService(db.NewConversationRepository())
}

func (s ConversationService) Get(chatId int64) domain.Conversation {
	return s.Repo.Get(chatId)
}
func (s ConversationService) Create(chatId int64) domain.Conversation {
	return s.Repo.Create(chatId)
}
func (s ConversationService) Update(conv domain.Conversation) error {
	return s.Repo.Update(conv)
}
func (s ConversationService) RemoveAll(chatId int64) error {
	return s.Repo.RemoveAll(chatId)
}
