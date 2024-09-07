package server

import (
	"context"
	"testing"

	"github.com/gustav1105/do-due/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

const bufconnSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufconnSize)
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAddTodo(t *testing.T) {
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

	resp, err := client.AddTodo(context.Background(), &proto.AddTodoRequest{
	})

	if err != nil {
		t.Fatalf("AddTodo failed: %v", err)
	}

	if resp == nil {
		t.Fatal("Response is nil")
	}
}

