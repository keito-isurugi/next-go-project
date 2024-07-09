package handler

import (
	"net/http"

	// "github.com/keito-isurugi/next-go-project/internal/domain/entity"
	// "github.com/keito-isurugi/next-go-project/internal/infra/db"
	"github.com/labstack/echo/v4"

	"github.com/keito-isurugi/next-go-project/internal/infra/postgres"
)

type TodoHandler interface {
	ListTodos(c echo.Context) error
}

type todoHnadler struct {
	todoRepo postgres.TodoPostgres
}

func NewTodoHandler(todoRepo postgres.TodoPostgres) *todoHnadler {
	return &todoHnadler{
		todoRepo: todoRepo,
	}
}

func (th *todoHnadler) ListTodos(c echo.Context) error {
	lt, err := th.todoRepo.ListTodos(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lt)
}
