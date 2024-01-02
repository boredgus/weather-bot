package db

import (
	"context"
	"subscription-bot/internal/domain"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type ConversationRepository struct {
	MongoStore
}

func NewConversationRepository() domain.ConversationRepository {
	return ConversationRepository{
		MongoStore: NewMongoStore(MongoClient().Database("db").Collection("conversation"), context.TODO()),
	}
}

func (r ConversationRepository) Get(chatId int64) domain.Conversation {
	result := r.collection.FindOne(r.context, bson.M{"chatId": chatId})
	var conv domain.Conversation
	if err := result.Decode(&conv); err != nil {
		logrus.Infof("failed to find conversation: %v", err.Error())
		return r.Create(chatId)
	}
	return conv
}

func (r ConversationRepository) Create(chatId int64) domain.Conversation {
	r.RemoveAll(chatId)
	conv := domain.BaseConversation(chatId)
	if _, err := r.collection.InsertOne(r.context, r.MongoStore.ToMap(conv)); err != nil {
		logrus.Warnf("failed to insert conversation: %v", err)
	}
	return conv
}

func (r ConversationRepository) Update(data domain.Conversation) error {
	res := r.collection.FindOneAndReplace(r.context, bson.M{"chatId": data.ChatId}, r.MongoStore.ToMap(data))
	var conv domain.Conversation
	err := res.Decode(&conv)
	if err != nil {
		logrus.Warnf("failed to replace conversation: %v", err.Error())
	}
	return err
}

func (r ConversationRepository) RemoveAll(chatId int64) error {
	_, err := r.collection.DeleteMany(r.context, bson.M{"chatId": chatId})
	if err != nil {
		logrus.Infof("failed to remove conversations with id=%v: %v", chatId, err.Error())
	}
	return err
}
