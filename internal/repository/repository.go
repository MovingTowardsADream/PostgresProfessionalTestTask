package repository

import (
	"PostgresProfessionalTestTask/internal/entity"
	"PostgresProfessionalTestTask/internal/repository/postgresdb"
	"PostgresProfessionalTestTask/pkg/postgres"
	"context"
)

type CommandsOperation interface {
	CreateCommand(ctx context.Context, command *entity.Command) (*entity.Command, error)
}

type Repositories struct {
	CommandsOperation
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		CommandsOperation: postgresdb.NewCommandsRepo(pg),
	}
}
