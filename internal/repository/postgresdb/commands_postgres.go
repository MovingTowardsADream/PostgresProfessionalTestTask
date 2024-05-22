package postgresdb

import (
	"PostgresProfessionalTestTask/internal/entity"
	"PostgresProfessionalTestTask/pkg/postgres"
	"context"
)

type CommandsRepo struct {
	db *postgres.Postgres
}

func NewCommandsRepo(pg *postgres.Postgres) *CommandsRepo {
	return &CommandsRepo{pg}
}

func (r *CommandsRepo) CreateCommand(ctx context.Context, command *entity.Command) (*entity.Command, error) {
	return nil, nil
}
