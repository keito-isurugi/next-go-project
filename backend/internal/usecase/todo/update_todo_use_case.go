package todo

import (
	"github.com/labstack/echo/v4"

	"github.com/keito-isurugi/next-go-project/internal/domain/entity"
	domain "github.com/keito-isurugi/next-go-project/internal/domain/todo"
)

type UpdateTodoUseCase interface {
	Exec(c echo.Context, todo *entity.Todo) error
}

type updateTodoUseCase struct {
	todoRepo domain.TodoRepository
}

func NewUpdateTodoUseCase(todoRepo domain.TodoRepository) UpdateTodoUseCase {
	return &updateTodoUseCase{
		todoRepo: todoRepo,
	}
}

func (utuc *updateTodoUseCase) Exec(c echo.Context, todo *entity.Todo) error {
	err := utuc.todoRepo.UpdateTodo(c.Request().Context(), todo)
	if err != nil {
		return err
	}

	return nil
}
