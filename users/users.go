package users

import (
	"context"
	"sync"

	"github.com/gofrs/uuid"
	userspb "github.com/johanbrandhorst/buf-example/proto/users/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	*userspb.UnimplementedUserServiceServer
	mu sync.Mutex

	users []*userspb.User
}

func (s *Service) AddUser(ctx context.Context, req *userspb.AddUserRequest) (*userspb.AddUserResponse, error) {
	user := &userspb.User{
		CreateTime: timestamppb.Now(),
		Name:       req.Name,
		Id:         uuid.Must(uuid.NewV4()).String(),
	}
	s.mu.Lock()
	s.users = append(s.users, user)
	s.mu.Unlock()

	return &userspb.AddUserResponse{
		User: user,
	}, nil
}
