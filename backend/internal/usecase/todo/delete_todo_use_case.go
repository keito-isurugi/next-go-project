package todo

import (
	"github.com/labstack/echo/v4"

	storageDomain "github.com/keito-isurugi/next-go-project/internal/domain/storage"
	domain "github.com/keito-isurugi/next-go-project/internal/domain/todo"
)

type DeleteTodoUseCase interface {
	Exec(c echo.Context, id int) error
}

type deleteTodoUseCase struct {
	todoRepo    domain.TodoRepository
	storageRepo storageDomain.StorageRepository
}

func NewDeleteTodoUseCase(
	todoRepo domain.TodoRepository,
	storageRepo storageDomain.StorageRepository,
) DeleteTodoUseCase {
	return &deleteTodoUseCase{
		todoRepo:    todoRepo,
		storageRepo: storageRepo,
	}
}

func (dtuc *deleteTodoUseCase) Exec(c echo.Context, id int) error {
	todo, err := dtuc.todoRepo.GetTodo(c.Request().Context(), id)
	if err != nil {
		return err
	}

	err = dtuc.storageRepo.DeleteObject(todo.AttachmentFile)
	if err != nil {
		return err
	}

	err = dtuc.todoRepo.DeleteTodo(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return nil
}
