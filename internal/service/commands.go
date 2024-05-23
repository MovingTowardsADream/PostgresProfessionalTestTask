package service

import (
	"PostgresProfessionalTestTask/internal/entity"
	"PostgresProfessionalTestTask/internal/repository"
	"context"
	"fmt"
)

type CommandsService struct {
	repo    repository.CommandsOperation
	signKey string
}

func NewCommandsService(repo repository.CommandsOperation, signKey string) *CommandsService {
	return &CommandsService{repo: repo, signKey: signKey}
}

func (s *CommandsService) CreateCommand(ctx context.Context, command *entity.Command) (*entity.Command, error) {
	fmt.Println(s.signKey)
	return s.repo.CreateCommand(ctx, command)
}
