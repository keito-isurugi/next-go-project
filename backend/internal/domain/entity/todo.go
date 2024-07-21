package entity

import "time"

type Todo struct {
	ID             int
	UserID         int
	Title          string
	AttachmentFile string
	DoneFlag       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

type ListTodos []Todo

func NewTodo(
	id int,
	userID int,
	title string,
	attachmentFile string,
	doneFlag bool,
	createdAt time.Time,
	updatedAt time.Time,
) *Todo {
	return &Todo{
		ID:             id,
		UserID:         userID,
		Title:          title,
		AttachmentFile: attachmentFile,
		DoneFlag:       doneFlag,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}
