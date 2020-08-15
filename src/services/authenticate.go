 package services

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

type MongoFields struct {
    Username string `json:"Field Str"`
	Ciphertext string `json:"Field Str"`
	Email string `json:"Field Str"`
	Name string `json:"Field Str"`
	Age int `json:"Field Int"`
}

func Authenticate(ctx context.Context, users *mongo.Collection, username string, password string) (res bool) {

	result := MongoFields{}

	filterCursor := users.FindOne(ctx, bson.M{"username": username}).Decode(&result)
	_ = filterCursor
	fmt.Println(filterCursor)
	if filterCursor == nil {
		resultpass := Decrypt([]byte(result.Ciphertext), "so hungry")
		if string(resultpass) == password {
			res = true
		} else {
			res = false
		}	
	} else {
		res = false
	}

	return
}
