package db

import (
	"context"
	"time"

	"github.com/ishanshre/GoShortLinker/internals/database"
	"github.com/ishanshre/GoShortLinker/internals/repository"
)

type mongoDbRepo struct {
	DB  database.Database
	Ctx context.Context
}

func NewMongoDbRepo(db database.Database, ctx context.Context) repository.DbRepository {
	return &mongoDbRepo{
		DB:  db,
		Ctx: ctx,
	}
}

const timeout = 3 * time.Second
