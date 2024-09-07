package main

import (
	"log"
	"net"
import (
	"github.com/gustav1105/do-due/internal/proto" // Adjust the import path
	"google.golang.org/grpc"
)

func main() {
  lis, err := net.Listen("tcp", ":50051")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  s := grpc.NewServer()
  
  proto.RegisterTodoServiceServer(s, &server.Server{})

  log.Println("Server running on port 50051")

  if err := s.Serve(lis); err != nil {
    log.Fatalf("Failed to serve: %v", err)
  }
}

