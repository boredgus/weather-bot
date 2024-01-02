package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Subscription struct {
	Id       string   `json:"id"`
	ChatId   int64    `json:"chatId"`
	Time     string   `json:"time"`
	Location Location `json:"location"`
}

func NewSubscription(chatId int64, t string, loc Location) Subscription {
	return Subscription{
		Id:       uuid.New().String(),
		ChatId:   chatId,
		Time:     t,
		Location: loc,
	}
}

func (s Subscription) String() string {
	return fmt.Sprintf("%v %v", s.Time, s.Location.String())
}

type SubscriptionRepository interface {
	GetAll(limit, startFrom int64) []Subscription
	GetAllFor(chatId int64) []Subscription
	GetById(subscriptionId string) (Subscription, error)
	Insert(sub Subscription) error
	Update(sub Subscription) error
	Remove(subId string) error
}

type SubscriptionSvc interface {
	GetAll(limit, startFrom int64) []Subscription
	GetAllFor(chatId int64) []Subscription
	GetById(subscriptionId string) (Subscription, error)
	Insert(sub Subscription) error
	Update(sub Subscription) error
	Remove(subId string) error
}
