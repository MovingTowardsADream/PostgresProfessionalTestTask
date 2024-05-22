package service

import (
	"PostgresProfessionalTestTask/internal/entity"
	"PostgresProfessionalTestTask/internal/repository"
	"context"
)

type CommandsService struct {
	repo repository.CommandsOperation
}

func NewCommandsService(repo repository.CommandsOperation) *CommandsService {
	return &CommandsService{repo: repo}
}

func (s *CommandsService) CreateCommand(ctx context.Context, command *entity.Command) (*entity.Command, error) {
	return s.repo.CreateCommand(ctx, command)
}
