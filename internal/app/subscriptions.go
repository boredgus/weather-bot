package app

import (
	"errors"
	"subscription-bot/internal/db"
	"subscription-bot/internal/domain"
)

type SubscriptionService struct {
	Repo domain.SubscriptionRepository
}

func NewSubscriptionService(repo domain.SubscriptionRepository) SubscriptionService {
	return SubscriptionService{Repo: repo}
}

func SubscriptionSvc() domain.SubscriptionSvc {
	return NewSubscriptionService(db.NewSubscriptionRepository())
}

var SubscriptionLimitError = errors.New("exceeded limit of subscriptions")

func (s SubscriptionService) GetAll(limit, startFrom int64) []domain.Subscription {
	return s.Repo.GetAll(limit, startFrom)
}
func (s SubscriptionService) GetAllFor(chatId int64) []domain.Subscription {
	return s.Repo.GetAllFor(chatId)
}
func (s SubscriptionService) GetById(subscriptionId string) (domain.Subscription, error) {
	return s.Repo.GetById(subscriptionId)
}
func (s SubscriptionService) Insert(sub domain.Subscription) error {
	if len(s.GetAllFor(sub.ChatId)) >= 50 {
		return SubscriptionLimitError
	}
	return s.Repo.Insert(sub)
}
func (s SubscriptionService) Update(sub domain.Subscription) error {
	return s.Repo.Update(sub)
}
func (s SubscriptionService) Remove(subId string) error {
	return s.Repo.Remove(subId)
}
