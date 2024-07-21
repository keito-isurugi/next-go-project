package todo

import (
	"github.com/labstack/echo/v4"

	domain "github.com/keito-isurugi/next-go-project/internal/domain/todo"
)

type DeleteTodoUseCase interface {
	Exec(c echo.Context, id int) error
}

type deleteTodoUseCase struct {
	todoRepo domain.TodoRepository
}

func NewDeleteTodoUseCase(todoRepo domain.TodoRepository) DeleteTodoUseCase {
	return &deleteTodoUseCase{
		todoRepo: todoRepo,
	}
}

func (dtuc *deleteTodoUseCase) Exec(c echo.Context, id int) error {
	err := dtuc.todoRepo.DeleteTodo(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return nil
}
