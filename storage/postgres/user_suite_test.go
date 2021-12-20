package postgres

import (
	"github.com/NafisaTojiboyeva/todo-service/config"
	pb "github.com/NafisaTojiboyeva/todo-service/genproto"
	"github.com/NafisaTojiboyeva/todo-service/pkg/db"
	"github.com/NafisaTojiboyeva/todo-service/storage/repo"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	CleanupFunc func()
	Repository  repo.TaskStorageI
}

func (suite *TaskRepositoryTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDBForSuite(config.Load())

	suite.Repository = NewTaskRepo(pgPool)
	suite.CleanupFunc = cleanup
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *TaskRepositoryTestSuite) TestTaskCRUD() {
	id := "0d512776-60ed-4980-b8a3-6904a2234fd4"
	assignee := "Lola"
	task := pb.Task{
		Id:       id,
		Assignee: "Lola",
		Title:    "Test",
		Summary:  "Just testing create function",
		Deadline: "2021-12-01",
		Status:   "Passed",
	}

	_ = suite.Repository.Delete(id)

	task, err := suite.Repository.Create(task)
	suite.Nil(err)

	getTask, err := suite.Repository.Get(task.Id)
	suite.Nil(err)
	suite.NotNil(getTask)
	suite.Equal(assignee, task.Assignee, "assignees must match")

	task.Title = "Suite Test"
	updatedTask, err := suite.Repository.Update(task)
	suite.Nil(err)

	getTask, err = suite.Repository.Get(task.Id)
	suite.Nil(err)
	suite.NotNil(getTask)
	suite.Equal(getTask.Title, updatedTask.Title)

	listTasks, _, err := suite.Repository.List(1, 5)
	suite.Nil(err)
	suite.NotEmpty(listTasks)
	suite.Equal(task.Title, listTasks[0].Title)

	overdueTasks, _, err := suite.Repository.ListOverdue("2021-12-19", 1, 2)
	suite.Nil(err)
	suite.NotEmpty(overdueTasks)
	suite.Equal(overdueTasks[0].Deadline, task.Deadline)

	err = suite.Repository.Delete(id)
	suite.Nil(err)
}

func (suite *TaskRepositoryTestSuite) TearDownSuite() {
	suite.CleanupFunc()
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
