package service

import (
	pb "github.com/NafisaTojiboyeva/todo-service/genproto"
	"google.golang.org/grpc"
	"log"
	"os"
	"testing"
)

var client pb.ToDoServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to get connection, %v", err)
	}

	client = pb.NewToDoServiceClient(conn)

	os.Exit(m.Run())
}
