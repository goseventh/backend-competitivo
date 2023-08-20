package mongodb

import (
	"context"
	"log"
	userModel "main/models/user"
)

func (mb *MongoBase) CreateUser(user userModel.User) {
	insertsQueue["users"] = append(insertsQueue["users"], user)
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
