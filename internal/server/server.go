package server

import (
    "context"
    "time"

    "github.com/gustav1105/do-due/internal/proto"
    "google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
    proto.UnimplementedTodoServiceServer
}

// AddTask implements TodoService.AddTask
func (s *Server) AddTask(ctx context.Context, req *proto.AddTaskRequest) (*proto.AddTaskResponse, error) {
    // Implement your logic here
    task := req.GetTask()
    // For demo purposes, we'll just return the task we received
    return &proto.AddTaskResponse{
        Task: task,
    }, nil
}

// GetTask implements TodoService.GetTask
func (s *Server) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.GetTaskResponse, error) {
    // Implement your logic here
    taskID := req.GetId()
    // For demo purposes, we'll return a dummy task
    return &proto.GetTaskResponse{
        Task: &proto.Task{
            Name:   "Sample Task",
            Note:   "This is a sample task.",
            Status: proto.TaskStatus_TODO,
            DueOn:  timestamppb.New(time.Now().Add(24 * time.Hour)), // Example due date
        },
    }, nil
}

// GetTasks implements TodoService.GetTasks
func (s *Server) GetTasks(ctx context.Context, req *proto.GetTasksRequest) (*proto.GetTasksResponse, error) {
    // Implement your logic here
    tasks := []*proto.Task{
        {
            Name:   "Sample Task 1",
            Note:   "This is a sample task 1.",
            Status: proto.TaskStatus_TODO,
            DueOn:  timestamppb.New(time.Now().Add(48 * time.Hour)), // Example due date
        },
        {
            Name:   "Sample Task 2",
            Note:   "This is a sample task 2.",
            Status: proto.TaskStatus_DUE,
            DueOn:  timestamppb.New(time.Now().Add(72 * time.Hour)), // Example due date
        },
    }
    return &proto.GetTasksResponse{
        Tasks: tasks,
    }, nil
}

// CompleteTask implements TodoService.CompleteTask
func (s *Server) CompleteTask(ctx context.Context, req *proto.CompleteTaskRequest) (*proto.CompleteTaskResponse, error) {
    // Implement your logic here
    taskID := req.GetId()
    // For demo purposes, we'll return a dummy completed task
    return &proto.CompleteTaskResponse{
        Task: &proto.Task{
            Name:   "Completed Task",
            Status: proto.TaskStatus_COMPLETED,
        },
    }, nil
}

// CancelTask implements TodoService.CancelTask
func (s *Server) CancelTask(ctx context.Context, req *proto.CancelTaskRequest) (*proto.CancelTaskResponse, error) {
    // Implement your logic here
    taskID := req.GetId()
    // For demo purposes, we'll return a dummy canceled task
    return &proto.CancelTaskResponse{
        Task: &proto.Task{
            Name:   "Canceled Task",
            Status: proto.TaskStatus_CANCELLED,
        },
    }, nil
}

