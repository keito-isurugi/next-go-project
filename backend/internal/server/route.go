package server

import (
	// "github.com/keito-isurugi/next-go-project/internal/infra/db"
	"net/http"

	"github.com/keito-isurugi/next-go-project/internal/handler"
	"github.com/keito-isurugi/next-go-project/internal/infra/db"
	"github.com/keito-isurugi/next-go-project/internal/infra/env"
	"github.com/keito-isurugi/next-go-project/internal/infra/postgres"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func SetupRouter(ev *env.Values, dbClient db.Client, zapLogger *zap.Logger) *echo.Echo {
	e := echo.New()
	
	// TODO ここにmiddlewareを実装する

	e.GET("/health", func (c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// repositry
	todoRepo := postgres.NewTodoPostgres(dbClient)

	// handler
	todoHandler := handler.NewTodoHandler(todoRepo)

	todoGroup := e.Group("/todos")
	todoGroup.GET("", todoHandler.ListTodos)

	return e
}
