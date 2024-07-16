//go:generate mockgen -source=todo_repository.go -destination=./mock/todo_repository_mock.go
package domain

import (
	"context"

	"github.com/keito-isurugi/next-go-project/internal/domain/entity"
)

type TodoRepository interface {
	ListTodos(ctx context.Context) (entity.ListTodos, error)
}
