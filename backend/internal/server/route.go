package server

import (
	// "github.com/keito-isurugi/next-go-project/internal/infra/db"
	"net/http"

	// "github.com/keito-isurugi/next-go-project/internal/handler"
	"github.com/keito-isurugi/next-go-project/internal/presentation/todos"
	"github.com/keito-isurugi/next-go-project/internal/infra/db"
	"github.com/keito-isurugi/next-go-project/internal/infra/env"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func SetupRouter(ev *env.Values, dbClient db.Client, zapLogger *zap.Logger) *echo.Echo {
	e := echo.New()
	
	e.GET("/health", func (c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})


	// handler
	todoHandler := todos.NewTodoHandler(dbClient)

	todoGroup := e.Group("/todos")
	todoGroup.GET("", todoHandler.ListTodos)
	todoGroup.GET("/:id", todoHandler.GetTodo)
	todoGroup.POST("", todoHandler.RegisterTodo)
	todoGroup.PATCH("/:id", todoHandler.UpdateTodo)
	todoGroup.DELETE("/:id", todoHandler.DeleteTodo)

	return e
}
