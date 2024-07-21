//go:generate mockgen -source=todo_repository.go -destination=./mock/todo_repository_mock.go
package domain

import (
	"context"

	"github.com/keito-isurugi/next-go-project/internal/domain/entity"
)

type TodoRepository interface {
	ListTodos(ctx context.Context) (entity.ListTodos, error)
	GetTodo(ctx context.Context, id int) (entity.Todo, error)
	RegisterTodo(ctx context.Context, todo *entity.Todo) (int, error)
	UpdateTodo(ctx context.Context, todo *entity.Todo) error
	DeleteTodo(ctx context.Context, id int) error
}
