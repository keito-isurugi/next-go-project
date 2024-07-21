package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/keito-isurugi/next-go-project/internal/domain/entity"
	domain "github.com/keito-isurugi/next-go-project/internal/domain/storage"
	useCaseTodo "github.com/keito-isurugi/next-go-project/internal/usecase/todo"
)

const (
	S3ObjectKey = "todo-images"
)

type TodoHandler interface {
	ListTodos(c echo.Context) error
	GetTodo(c echo.Context) error
	RegisterTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoHnadler struct {
	storageRepo         domain.StorageRepository
	listTodosUseCase    useCaseTodo.ListTodosUseCase
	getTodoUseCase      useCaseTodo.GetTodoUseCase
	registerTodoUseCase useCaseTodo.RegisterTodoUseCase
	updateTodoUseCase   useCaseTodo.UpdateTodoUseCase
	deleteTodoUseCase   useCaseTodo.DeleteTodoUseCase
}

func NewTodoHandler(
	storageRepo domain.StorageRepository,
	listTodosUseCase useCaseTodo.ListTodosUseCase,
	getTodoUseCase useCaseTodo.GetTodoUseCase,
	registerTodoUseCase useCaseTodo.RegisterTodoUseCase,
	updateTodoUseCase useCaseTodo.UpdateTodoUseCase,
	deleteTodoUseCase useCaseTodo.DeleteTodoUseCase,
) TodoHandler {
	return &todoHnadler{
		storageRepo:         storageRepo,
		listTodosUseCase:    listTodosUseCase,
		getTodoUseCase:      getTodoUseCase,
		registerTodoUseCase: registerTodoUseCase,
		updateTodoUseCase:   updateTodoUseCase,
		deleteTodoUseCase:   deleteTodoUseCase,
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
			ID:             lt[i].ID,
			Title:          lt[i].Title,
			AttachmentFile: lt[i].AttachmentFile,
			DoneFlag:       lt[i].DoneFlag,
		}
	}

	return c.JSON(http.StatusOK, res)
}

func (th *todoHnadler) GetTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	t, err := th.getTodoUseCase.Exec(c, id)
	if err != nil {
		return err
	}

	res := todoResponse{
		ID:             t.ID,
		Title:          t.Title,
		AttachmentFile: t.AttachmentFile,
		DoneFlag:       t.DoneFlag,
	}

	return c.JSON(http.StatusOK, res)
}

func (th *todoHnadler) RegisterTodo(c echo.Context) error {
	// リクエストの内容をデバッグ出力
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Printf("Error parsing multipart form: %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid multipart form"})
	}

	var attachmentPath string
	// ファイルを取得
	files := form.File["attachmentFile"]
	if len(files) > 0 {
		file := files[0]
		filename := fmt.Sprintf("%s/%d", S3ObjectKey, time.Now().UnixNano())

		attachmentPath, err = th.storageRepo.PutObject(file, "next-go-images", filename)
		if err != nil {
			fmt.Printf("Error uploading to S3: %v\n", err)
			return err
		}

		// 諸々デバッグ
		fmt.Printf("File uploaded successfully. S3 URL: %s\n", attachmentPath)
		fmt.Printf("Received file: %s\n", file.Filename)
		fmt.Printf("File size: %d bytes\n", file.Size)
		fmt.Println("3. Header: ")
		for key, values := range file.Header {
			for _, value := range values {
				fmt.Printf("   %s: %s\n", key, value)
			}
		}
		fmt.Printf("File saved to: %s\n", attachmentPath)
	}

	var req registerTodoRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	reqTodo := entity.NewTodo(
		0,
		1,
		req.Title,
		attachmentPath,
		false,
		time.Now(),
		time.Now(),
	)

	id, err := th.registerTodoUseCase.Exec(c, reqTodo)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, id)
}

func (th *todoHnadler) UpdateTodo(c echo.Context) error {
	var req updateTodoRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	reqTodo := entity.NewTodo(
		req.ID,
		req.UserID,
		req.Title,
		req.AttachmentFile,
		req.DoneFlag,
		time.Now(),
		time.Now(),
	)

	err := th.updateTodoUseCase.Exec(c, reqTodo)
	if err != nil {
		return err
	}

	return nil
}

func (th *todoHnadler) DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := th.deleteTodoUseCase.Exec(c, id)
	if err != nil {
		return err
	}

	return nil
}
