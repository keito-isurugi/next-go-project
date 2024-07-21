package todo

import (
	"github.com/labstack/echo/v4"

	"github.com/keito-isurugi/next-go-project/internal/domain/entity"
	domain "github.com/keito-isurugi/next-go-project/internal/domain/todo"
)

type RegisterTodoUseCase interface {
	Exec(c echo.Context, todo *entity.Todo) (int, error)
}

type registerTodoUseCase struct {
	todoRepo domain.TodoRepository
}

func NewRegisterTodoUseCase(todoRepo domain.TodoRepository) RegisterTodoUseCase {
	return &registerTodoUseCase{
		todoRepo: todoRepo,
	}
}

func (rtuc *registerTodoUseCase) Exec(c echo.Context, todo *entity.Todo) (int, error) {
	id, err := rtuc.todoRepo.RegisterTodo(c.Request().Context(), todo)
	if err != nil {
		return id, err
	}

	return id, nil
}
