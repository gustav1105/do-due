package server

import (
    "context"
    "database/sql"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gustav1105/do-due/internal/proto"
    "google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
    proto.UnimplementedTodoServiceServer
    db *sql.DB
}

func NewServer(dsn string) (*Server, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &Server{db: db}, nil
}

func (s *Server) AddTask(ctx context.Context, req *proto.AddTaskRequest) (*proto.AddTaskResponse, error) {
    task := req.GetTask()
    
    _, err := s.db.Exec(`INSERT INTO tasks (name, note, status, due_on) VALUES (?, ?, ?, ?)`,
        task.GetName(),
        task.GetNote(),
        task.GetStatus(),
        task.GetDueOn().AsTime(),
    )
    if err != nil {
        return nil, err
    }

    return &proto.AddTaskResponse{
        Task: task,
    }, nil
}

func (s *Server) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.GetTaskResponse, error) {
    taskID := req.GetId()
    task := &proto.Task{}

    row := s.db.QueryRow(`SELECT name, note, status, due_on FROM tasks WHERE id = ?`, taskID)
    var dueOn time.Time
    if err := row.Scan(&task.Name, &task.Note, &task.Status, &dueOn); err != nil {
        return nil, err
    }
    task.DueOn = timestamppb.New(dueOn)

    return &proto.GetTaskResponse{
        Task: task,
    }, nil
}

func (s *Server) GetTasks(ctx context.Context, req *proto.GetTasksRequest) (*proto.GetTasksResponse, error) {
    rows, err := s.db.Query(`SELECT name, note, status, due_on FROM tasks`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    tasks := []*proto.Task{}
    for rows.Next() {
        task := &proto.Task{}
        var dueOn time.Time
        if err := rows.Scan(&task.Name, &task.Note, &task.Status, &dueOn); err != nil {
            return nil, err
        }
        task.DueOn = timestamppb.New(dueOn)
        tasks = append(tasks, task)
    }

    return &proto.GetTasksResponse{
        Tasks: tasks,
    }, nil
}

func (s *Server) CompleteTask(ctx context.Context, req *proto.CompleteTaskRequest) (*proto.CompleteTaskResponse, error) {
    taskID := req.GetId()

    _, err := s.db.Exec(`UPDATE tasks SET status = ? WHERE id = ?`, proto.TaskStatus_COMPLETED, taskID)
    if err != nil {
        return nil, err
    }

    return &proto.CompleteTaskResponse{
        Task: &proto.Task{
            Name:   "Completed Task",
            Status: proto.TaskStatus_COMPLETED,
        },
    }, nil
}

func (s *Server) CancelTask(ctx context.Context, req *proto.CancelTaskRequest) (*proto.CancelTaskResponse, error) {
    taskID := req.GetId()

    _, err := s.db.Exec(`UPDATE tasks SET status = ? WHERE id = ?`, proto.TaskStatus_CANCELLED, taskID)
    if err != nil {
        return nil, err
    }

    return &proto.CancelTaskResponse{
        Task: &proto.Task{
            Name:   "Canceled Task",
            Status: proto.TaskStatus_CANCELLED,
        },
    }, nil
}

