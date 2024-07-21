package todo

import (
	"github.com/labstack/echo/v4"

	"github.com/keito-isurugi/next-go-project/internal/domain/entity"
	domain "github.com/keito-isurugi/next-go-project/internal/domain/todo"
)

type GetTodoUseCase interface {
	Exec(c echo.Context, id int) (entity.Todo, error)
}

type getTodoUseCase struct {
	todoRepo domain.TodoRepository
}

func NewGetTodoUseCase(todoRepo domain.TodoRepository) GetTodoUseCase {
	return &getTodoUseCase{
		todoRepo: todoRepo,
	}
}

func (gtuc *getTodoUseCase) Exec(c echo.Context, id int) (entity.Todo, error) {
	t, err := gtuc.todoRepo.GetTodo(c.Request().Context(), id)
	if err != nil {
		return entity.Todo{}, err
	}

	return t, nil
}
