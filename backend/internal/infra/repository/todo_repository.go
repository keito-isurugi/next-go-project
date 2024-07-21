package repository

import (
	"context"

	"github.com/keito-isurugi/next-go-project/internal/domain/entity"
	domain "github.com/keito-isurugi/next-go-project/internal/domain/todo"
	"github.com/keito-isurugi/next-go-project/internal/infra/db"
)

type todoRepository struct {
	dbClient db.Client
}

func NewTodoRepository(dbClient db.Client) domain.TodoRepository {
	return &todoRepository{
		dbClient: dbClient,
	}
}

func (tr *todoRepository) ListTodos(ctx context.Context) (entity.ListTodos, error) {
	var lt entity.ListTodos
	if err := tr.dbClient.Conn(ctx).Where("done_flag", false).Order("id").Find(&lt).Error; err != nil {
		return entity.ListTodos{}, err
	}

	return lt, nil
}

func (tr *todoRepository) GetTodo(ctx context.Context, id int) (entity.Todo, error) {
	var t entity.Todo
	if err := tr.dbClient.Conn(ctx).
		Where("id", id).
		First(&t).Error; err != nil {
		return entity.Todo{}, err
	}

	return t, nil
}

func (tr *todoRepository) RegisterTodo(ctx context.Context, todo *entity.Todo) (int, error) {
	if err := tr.dbClient.Conn(ctx).
		Create(todo).
		Error; err != nil {
		return 0, err
	}

	return todo.ID, nil
}

func (tr *todoRepository) UpdateTodo(ctx context.Context, todo *entity.Todo) error {
	var t entity.Todo
	if err := tr.dbClient.Conn(ctx).Where("id", todo.ID).First(&t).Error; err != nil {
		return err
	}

	todo.CreatedAt = t.CreatedAt

	if err := tr.dbClient.Conn(ctx).
		Updates(todo).
		Error; err != nil {
		return err
	}

	return nil
}

func (tr *todoRepository) DeleteTodo(ctx context.Context, id int) error {
	var t entity.Todo
	if err := tr.dbClient.Conn(ctx).Where("id", id).First(&t).Error; err != nil {
		return err
	}

	if err := tr.dbClient.Conn(ctx).
		Delete(t).
		Error; err != nil {
		return err
	}

	return nil
}
