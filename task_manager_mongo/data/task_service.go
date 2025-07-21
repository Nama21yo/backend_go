package data

import (
	"context"
	"errors"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection

func InitMongoDB(client *mongo.Client, dbName string, collectionName string) {
	taskCollection = client.Database(dbName).Collection(collectionName)
}

func CreateTask(ctx context.Context, task models.Task) (*models.Task, error) {
	task.ID = primitive.NewObjectID()
	_, err := taskCollection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func GetTasks(ctx context.Context) ([]models.Task, error) {
	cursor, err := taskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	tasks := []models.Task{}
	for cursor.Next(ctx) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var task models.Task
	err = taskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(ctx context.Context, id string, updated models.Task) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{
			"title":       updated.Title,
			"description": updated.Description,
			"due_date":    updated.DueDate,
			"status":      updated.Status,
		},
	}
	_, err = taskCollection.UpdateByID(ctx, objID, update)
	if err != nil {
		return nil, err
	}
	return GetTaskByID(ctx, id)
}

func DeleteTask(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := taskCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
