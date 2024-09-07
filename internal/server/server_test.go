package server

import (
	"context"
	"testing"
  "net"

	"github.com/gustav1105/do-due/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const bufconnSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufconnSize)
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAddTask(t *testing.T) {
	s := grpc.NewServer()
	proto.RegisterTodoServiceServer(s, &Server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := proto.NewTodoServiceClient(conn)

	// Create a sample task
	task := &proto.Task{
		Name:   "Sample Task",
		Note:   "This is a sample task.",
		Status: proto.TaskStatus_TODO,
		DueOn:  timestamppb.New(time.Now().Add(24 * time.Hour)),
	}

	resp, err := client.AddTask(context.Background(), &proto.AddTaskRequest{
		Task: task,
	})

	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}

	if resp == nil {
		t.Fatal("Response is nil")
	}

	if resp.Task == nil {
		t.Fatal("Task in response is nil")
	}

	if resp.Task.Name != task.Name {
		t.Errorf("Expected task name %s, got %s", task.Name, resp.Task.Name)
	}
}

