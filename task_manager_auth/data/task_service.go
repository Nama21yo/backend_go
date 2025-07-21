package data

import (
	"context"

	"github.com/yourusername/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection = InitDB().Database("taskdb").Collection("tasks")

func CreateTask(task models.Task) error {
	_, err := taskCollection.InsertOne(context.TODO(), task)
	return err
}

func GetTasks() ([]models.Task, error) {
	cursor, err := taskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	if err := cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {
	var task models.Task
	err := taskCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&task)
	return task, err
}

func UpdateTask(id string, updated models.Task) error {
	_, err := taskCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": updated})
	return err
}

func DeleteTask(id string) error {
	_, err := taskCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
