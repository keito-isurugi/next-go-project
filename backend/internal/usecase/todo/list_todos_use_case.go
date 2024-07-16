package todo

import (
	"github.com/labstack/echo/v4"

	"github.com/keito-isurugi/next-go-project/internal/domain/entity"
	domain "github.com/keito-isurugi/next-go-project/internal/domain/todo"
)

type ListTodosUseCase interface {
	Exec(c echo.Context) (entity.ListTodos, error)
}

type listTodosUseCase struct {
	todoRepo domain.TodoRepository
}

func NewListTodoUseCase(todoRepo domain.TodoRepository) ListTodosUseCase {
	return &listTodosUseCase{
		todoRepo: todoRepo,
	}
}

func (ltuc *listTodosUseCase) Exec(c echo.Context) (entity.ListTodos, error) {
	lt, err := ltuc.todoRepo.ListTodos(c.Request().Context())
	if err != nil {
		return entity.ListTodos{}, err
	}

	return lt, nil
}
