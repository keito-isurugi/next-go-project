package repository

import (
	"context"

	"github.com/keito-isurugi/next-go-project/internal/domain/entity"
	domain "github.com/keito-isurugi/next-go-project/internal/domain/todo"
	"github.com/keito-isurugi/next-go-project/internal/infra/db"
)

type todoRepository struct {
	dbClient db.Client
}

func NewTodoRepository(dbClient db.Client) domain.TodoRepository {
	return &todoRepository{
		dbClient: dbClient,
	}
}

func (tr *todoRepository) ListTodos(ctx context.Context) (entity.ListTodos, error) {
	var lt entity.ListTodos
	if err := tr.dbClient.Conn(ctx).Where("done_flag", false).Order("id").Find(&lt).Error; err != nil {
		return entity.ListTodos{}, err
	}
	return entity.ListTodos{}, nil
}
