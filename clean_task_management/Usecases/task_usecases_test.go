package usecases_test

import (
	"errors"
	"testing"
	"time"
	"task-manager/Domain"
	"task-manager/Usecases"
	"task-manager/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskUsecaseTestSuite struct {
	suite.Suite
	taskRepo *mocks.TaskRepositoryMock
	tuc      usecases.TaskUsecase
}

func (s *TaskUsecaseTestSuite) SetupTest() {
	s.taskRepo = &mocks.TaskRepositoryMock{}
	s.tuc = usecases.NewTaskUsecase(s.taskRepo)
}

func (s *TaskUsecaseTestSuite) TestCreateTask_Success() {
	owner := "owner1"
	title := "T"
	desc := "D"

	// repository Create should be called with a *domain.Task
	s.taskRepo.On("Create", mockAnyTask()).Return(nil)

	task, err := s.tuc.CreateTask(title, desc, owner)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), title, task.Title)
	assert.Equal(s.T(), owner, task.OwnerID)
	s.taskRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseTestSuite) TestUpdateTask_Authorized() {
	owner := "owner1"
	task := &domain.Task{ID: "t1", Title: "T", Description: "D", OwnerID: owner}
	// GetByID returns existing
	s.taskRepo.On("GetByID", task.ID).Return(task, nil)
	// Update should succeed
	s.taskRepo.On("Update", mockAnyTask()).Return(nil)

	updated, err := s.tuc.UpdateTask(task, owner)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), task.ID, updated.ID)
	s.taskRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseTestSuite) TestUpdateTask_Unauthorized() {
	owner := "owner1"
	task := &domain.Task{ID: "t1", Title: "T", Description: "D", OwnerID: owner}
	s.taskRepo.On("GetByID", task.ID).Return(task, nil)

	_, err := s.tuc.UpdateTask(task, "other")
	assert.Error(s.T(), err)
}

func (s *TaskUsecaseTestSuite) TestGetTaskByID_Unauthorized() {
	owner := "owner1"
	task := &domain.Task{ID: "t1", Title: "T", Description: "D", OwnerID: owner}
	s.taskRepo.On("GetByID", task.ID).Return(task, nil)

	_, err := s.tuc.GetTaskByID(task.ID, "other")
	assert.Error(s.T(), err)
}

func (s *TaskUsecaseTestSuite) TestDeleteTask_Success() {
	owner := "owner1"
	task := &domain.Task{ID: "t1", Title: "T", Description: "D", OwnerID: owner}
	s.taskRepo.On("GetByID", task.ID).Return(task, nil)
	s.taskRepo.On("Delete", task.ID).Return(nil)

	err := s.tuc.DeleteTask(task.ID, owner)
	assert.NoError(s.T(), err)
	s.taskRepo.AssertExpectations(s.T())
}

func TestTaskUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseTestSuite))
}

// helpers
func mockAnyTask() func(interface{}) bool {
	return func(i interface{}) bool {
		_, ok := i.(*domain.Task)
		return ok
	}
}
