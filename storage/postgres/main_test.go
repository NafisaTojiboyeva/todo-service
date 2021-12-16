package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/NafisaTojiboyeva/todo-service/config"
	"github.com/NafisaTojiboyeva/todo-service/pkg/db"
	"github.com/NafisaTojiboyeva/todo-service/pkg/logger"
)

var pgRepo *taskRepo

func TestMain(m *testing.M) {
	cfg := config.Load()

	connDb, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	pgRepo = NewTaskRepo(connDb)

	os.Exit(m.Run())
}
