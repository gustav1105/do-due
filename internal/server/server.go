package server
import (
	"context"

	"github.com/gustav1105/do-due/internal/proto" // Adjust the import path
)

type Server struct {
  proto.UnimplementedTodoServiceServer
}
