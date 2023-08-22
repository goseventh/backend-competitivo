package mongodb

import (
	"context"
	userModel "main/models/user"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoBase struct {
	database   *mongo.Database
	collection *mongo.Collection
	connection *mongo.Client
	channel    chan channelMSG
}
type channelMSG struct {
	collection *mongo.Collection
	context    *userModel.User
}

func Builder() *MongoBase {
	mb := new(MongoBase)
	mb.connection = GetConn()
	mb.channel = make(chan channelMSG)
	go mb.tickerInsert()
	for i := 0; i < 30124; i++ {
		go worker(mb.channel)
	}
	return mb
}
func worker(channel <-chan channelMSG) {
	for msg := range channel {
		msg.collection.InsertOne(context.TODO(), msg.context)
	}
}

func (mb *MongoBase) UseDatabase(db string) *MongoBase {
	mb.database = mb.connection.Database(db)
	return mb
}

func (mb *MongoBase) UseCollection(collection string) *MongoBase {
	mb.collection = mb.database.Collection(collection)
	return mb
}
