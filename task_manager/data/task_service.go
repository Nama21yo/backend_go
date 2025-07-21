package data

import (
	"errors"
	"task_manager/models"
	"time"
)

var tasks = make(map[int]models.Task)
var nextID = 1

func GetAllTasks() []models.Task {
	list := []models.Task{}
	for _, task := range tasks {
		list = append(list, task)
	}
	return list
}

func GetTaskByID(id int) (models.Task, error) {
	task, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func CreateTask(task models.Task) models.Task {
	task.ID = nextID
	nextID++
	if task.DueDate.IsZero() {
		task.DueDate = time.Now().Add(24 * time.Hour)
	}
	tasks[task.ID] = task
	return task
}

func UpdateTask(id int, updated models.Task) (models.Task, error) {
	_, exists := tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	updated.ID = id
	tasks[id] = updated
	return updated, nil
}

func DeleteTask(id int) error {
	if _, exists := tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}
