package data

import (
	"context"
	"fmt"

	"github.com/yourusername/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var userCollection = InitDB().Database("taskdb").Collection("users")

func RegisterUser(user models.User) error {
	exists := userCollection.FindOne(context.TODO(), bson.M{"username": user.Username})
	if exists.Err() == nil {
		return fmt.Errorf("user already exists")
	}
	// First user becomes admin
	count, _ := userCollection.CountDocuments(context.TODO(), bson.M{})
	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hash)
	_, err := userCollection.InsertOne(context.TODO(), user)
	return err
}

func AuthenticateUser(username, password string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return user, fmt.Errorf("invalid credentials")
	}
	return user, nil
}

func PromoteUser(username string) error {
	_, err := userCollection.UpdateOne(context.TODO(), bson.M{"username": username}, bson.M{"$set": bson.M{"role": "admin"}})
	return err
}

func GetUserByID(id string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return user, err
}
