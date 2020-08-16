 package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"models"
)

type User struct {
    Username string `json:"Field Str"`
	Ciphertext string `json:"Field Str"`
	Email string `json:"Field Str"`
	Name string `json:"Field Str"`
	Age int `json:"Field Int"`
}

func Authenticate(ctx context.Context, users *mongo.Collection, username string, password string) (res models.User) {

	result := models.User{}

	filterCursor := users.FindOne(ctx, bson.M{"username": username}).Decode(&result)
	_ = filterCursor

	if filterCursor == nil {
		resultpass := Decrypt([]byte(result.Ciphertext), "so hungry")
		if string(resultpass) == password {
			res = result
		} else {
			res = models.User{}
		}	
	} else {
		res = models.User{}
	}

	return
}
