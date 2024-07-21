package server

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/keito-isurugi/next-go-project/internal/infra/db"
	"github.com/keito-isurugi/next-go-project/internal/infra/env"
	"github.com/keito-isurugi/next-go-project/internal/infra/repository"
	presentationTodo "github.com/keito-isurugi/next-go-project/internal/presentation/todo"
	useCaseTodo "github.com/keito-isurugi/next-go-project/internal/usecase/todo"
)

func SetupRouter(ev *env.Values, dbClient db.Client, _ *zap.Logger, awsClient s3iface.S3API) *echo.Echo {
	e := echo.New()
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// handler
	todoHandler := presentationTodo.NewTodoHandler(
		repository.NewS3Repository(
			ev,
			awsClient,
		),
		useCaseTodo.NewListTodoUseCase(
			repository.NewTodoRepository(dbClient),
		),
		useCaseTodo.NewGetTodoUseCase(
			repository.NewTodoRepository(dbClient),
		),
		useCaseTodo.NewRegisterTodoUseCase(
			repository.NewTodoRepository(dbClient),
		),
		useCaseTodo.NewUpdateTodoUseCase(
			repository.NewTodoRepository(dbClient),
		),
		useCaseTodo.NewDeleteTodoUseCase(
			repository.NewTodoRepository(dbClient),
		),
	)

	todoGroup := e.Group("/todos")
	todoGroup.GET("", todoHandler.ListTodos)
	todoGroup.GET("/:id", todoHandler.GetTodo)
	todoGroup.POST("", todoHandler.RegisterTodo)
	todoGroup.PATCH("/:id", todoHandler.UpdateTodo)
	todoGroup.DELETE("/:id", todoHandler.DeleteTodo)

	return e
}
