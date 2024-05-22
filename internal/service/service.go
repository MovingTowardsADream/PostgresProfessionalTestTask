package service

import (
	"PostgresProfessionalTestTask/internal/entity"
	"PostgresProfessionalTestTask/internal/repository"
	"context"
)

type CommandsOperations interface {
	CreateCommand(ctx context.Context, command *entity.Command) (*entity.Command, error)
}

type Service struct {
	CommandsOperations
}

func NewService(repos *repository.Repositories) *Service {
	return &Service{
		CommandsOperations: NewCommandsService(repos.CommandsOperation),
	}
}
