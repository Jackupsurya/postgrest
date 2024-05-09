package repository

import (
	"context"
	"myProject/postgrest/models"
)

type ToDoRepositoryInterface interface {
	GetAllTasks(ctx context.Context) (resp []models.Todo, err error)
	GetTaskById(ctx context.Context, id string) (resp []models.Todo, err error)
	CreateTask(ctx context.Context, task models.Todo) (resp models.Todo, err error)
}
