package postgres

import (
	"github.com/keito-isurugi/next-go-project/internal/domain/entity"
	"github.com/keito-isurugi/next-go-project/internal/infra/db"
	"github.com/labstack/echo/v4"

)

type TodoPostgres interface {
	ListTodos(c echo.Context) (entity.ListTodo, error)
}

type todoPostgres struct {
	dbClient db.Client
}

func NewTodoPostgres(dbClient db.Client) *todoPostgres{
	return &todoPostgres{
		dbClient: dbClient,
	}
}

func (tp *todoPostgres) ListTodos(c echo.Context) (entity.ListTodo, error) {
	var lt entity.ListTodo
	if err := tp.dbClient.Conn(c.Request().Context()).Find(&lt).Error; err != nil {
		return nil, err
	}

	return lt, nil
}