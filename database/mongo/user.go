package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"main/models/user"
)

func (mb *MongoBase) IsExists(user userModel.User) bool {
	filter := bson.D{{Key: "nickname", Value: user.Nickname}}
	searchUser := new(userModel.User)
	err := mb.collection.FindOne(context.TODO(), filter).Decode(searchUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	if searchUser.Nickname == user.Nickname {
		return true
	}
	return false
}

func (mb *MongoBase) CreateUser(user userModel.User) {
	user.ObjectID = primitive.NewObjectID()
	insertsQueue["users"] = append(insertsQueue["users"], user)
}

func (mb *MongoBase) CreateUserNow(user userModel.User) {
	mb.collection.InsertOne(context.TODO(), user)
}

func (mb *MongoBase) tickerUser() {
	if mb.collection == nil {
		return
	}
	if len(insertsQueue["users"]) > 0 {
		_, err := mb.collection.InsertMany(context.TODO(), insertsQueue["users"])
		if err != nil {
			log.Printf("❌ Error: fail register user: (%v), len(%v)", err, len(insertsQueue["users"]))
			return
		}
		insertsQueue["users"] = insertsQueue["users"][:len(insertsQueue["users"])-1]
	}
}
