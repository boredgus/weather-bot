package db

import (
	"context"
	"fmt"
	"subscription-bot/internal/domain"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubscriptionRepository struct {
	MongoStore
}

func NewSubscriptionRepository() SubscriptionRepository {
	return SubscriptionRepository{
		MongoStore: NewMongoStore(MongoClient().Database("db").Collection("subscription"), context.TODO()),
	}
}

func (r SubscriptionRepository) GetAll(limit, startFrom int64) (subs []domain.Subscription) {
	cur, err := r.collection.Find(r.context, bson.D{}, options.Find().SetSkip(startFrom).SetLimit(limit))
	if err != nil {
		return make([]domain.Subscription, 0)
	}
	err = cur.All(r.context, &subs)
	if err != nil {
		return make([]domain.Subscription, 0)
	}
	return
}

func (r SubscriptionRepository) GetAllFor(chatId int64) (subs []domain.Subscription) {
	cur, err := r.collection.Find(r.context, bson.M{"chatId": chatId})
	if err != nil {
		logrus.Infof("failed to find subscriptions: " + err.Error())
		return make([]domain.Subscription, 0)
	}

	err = cur.All(r.context, &subs)
	if err != nil {
		logrus.Infof("failed to read subs: " + err.Error())
		return make([]domain.Subscription, 0)
	}

	return subs
}

func (r SubscriptionRepository) GetById(subId string) (sub domain.Subscription, e error) {
	e = r.collection.FindOne(r.context, bson.M{"id": subId}).Decode(&sub)
	return
}

func (r SubscriptionRepository) Insert(sub domain.Subscription) error {
	_, err := r.collection.InsertOne(r.context, r.MongoStore.ToMap(sub))
	if err != nil {
		logrus.Infof("failed to insert subscription: %v", err.Error())
	}
	return err
}

func (r SubscriptionRepository) Update(sub domain.Subscription) error {
	r.collection.FindOneAndUpdate(r.context, bson.M{"id": sub.Id}, bson.M{"$set": r.MongoStore.ToMap(sub)})
	return nil
}

func (r SubscriptionRepository) Remove(subId string) error {
	res := r.collection.FindOneAndDelete(r.context, bson.M{"id": subId})
	var foundSub domain.Subscription
	if err := res.Decode(&foundSub); err != nil {
		return fmt.Errorf("failed to remove subscription: %v", err)
	}
	return nil
}
