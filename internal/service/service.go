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

type ServicesDependencies struct {
	Repos *repository.Repositories

	SignKey string
}

func NewService(deps ServicesDependencies) *Service {
	return &Service{
		CommandsOperations: NewCommandsService(deps.Repos.CommandsOperation, deps.SignKey),
	}
}
