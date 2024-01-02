package db

import (
	"context"
	"reflect"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var mongoClient *mongo.Client

func MongoClient() *mongo.Client {
	once.Do(func() {
		uri, err := NewMongoDbUri().Generate()
		if err != nil {
			panic(err)
		}

		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		mongoClient = client
	})
	return mongoClient
}

type MongoStore struct {
	collection *mongo.Collection
	context    context.Context
}

func NewMongoStore(coll *mongo.Collection, ctx context.Context) MongoStore {
	return MongoStore{collection: coll, context: ctx}
}

func (s MongoStore) ToMap(instance any) bson.M {
	res := bson.M{}
	for i, val := 0, reflect.ValueOf(instance); i < val.Type().NumField(); i++ {
		res[val.Type().Field(i).Tag.Get("json")] = val.Field(i).Interface()
	}
	return res
}
