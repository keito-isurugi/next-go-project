package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/keito-isurugi/next-go-project/internal/infra/db"
	useCaseTodo "github.com/keito-isurugi/next-go-project/internal/usecase/todo"
)

type Todo struct {
	ID        int
	UserID    int
	Title     string
	DoneFlag  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ListTodos []Todo

func NewTodo(
	id int,
	userID int,
	title string,
	doneFlag bool,
	createdAt time.Time,
	updatedAt time.Time,
) *Todo {
	return &Todo{
		ID:        id,
		UserID:    userID,
		Title:     title,
		DoneFlag:  doneFlag,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type TodoHandler interface {
	ListTodos(c echo.Context) error
	GetTodo(c echo.Context) error
	RegisterTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoHnadler struct {
	dbClient         db.Client
	listTodosUseCase useCaseTodo.ListTodosUseCase
}

func NewTodoHandler(
	dbClient db.Client,
	listTodosUseCase useCaseTodo.ListTodosUseCase,
) TodoHandler {
	return &todoHnadler{
		dbClient:         dbClient,
		listTodosUseCase: listTodosUseCase,
	}
}

func (th *todoHnadler) ListTodos(c echo.Context) error {
	lt, err := th.listTodosUseCase.Exec(c)
	if err != nil {
		return err
	}

	res := make(listTodosResponse, len(lt))
	for i := range lt {
		res[i] = todoResponse{
			ID:       lt[i].ID,
			Title:    lt[i].Title,
			DoneFlag: lt[i].DoneFlag,
		}
	}

	return c.JSON(http.StatusOK, res)
}

func (th *todoHnadler) GetTodo(c echo.Context) error {
	id := c.Param("id")

	var t Todo
	if err := th.dbClient.Conn(c.Request().Context()).Where("id", id).Find(&t).Error; err != nil {
		return err
	}

	res := todoResponse{
		ID:       t.ID,
		Title:    t.Title,
		DoneFlag: t.DoneFlag,
	}

	return c.JSON(http.StatusOK, res)
}

func (th *todoHnadler) RegisterTodo(c echo.Context) error {
	var req registerTodoRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	reqTodo := NewTodo(
		0,
		1,
		req.Title,
		false,
		time.Now(),
		time.Now(),
	)

	if err := th.dbClient.Conn(c.Request().Context()).Create(reqTodo).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, reqTodo.ID)
}

func (th *todoHnadler) UpdateTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var req updateTodoRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	var todo Todo
	if err := th.dbClient.Conn(c.Request().Context()).Where("id", id).First(&todo).Error; err != nil {
		return err
	}

	reqTodo := NewTodo(
		todo.ID,
		todo.UserID,
		req.Title,
		req.DoneFlag,
		todo.CreatedAt,
		time.Now(),
	)

	if err := th.dbClient.Conn(c.Request().Context()).Updates(reqTodo).Error; err != nil {
		return err
	}
	return nil
}

func (th *todoHnadler) DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var todo Todo
	if err := th.dbClient.Conn(c.Request().Context()).Where("id", id).First(&todo).Error; err != nil {
		return err
	}

	if err := th.dbClient.Conn(c.Request().Context()).Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}
