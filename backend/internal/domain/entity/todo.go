package entity

import "time"

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
