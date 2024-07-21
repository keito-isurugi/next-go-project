package handler

import "mime/multipart"

type registerTodoRequest struct {
	Title          string                `form:"title" json:"title" example:"サンプルタイトル" ja:"タイトル" validate:"required,max=255"`
	AttachmentFile *multipart.FileHeader `form:"attachmentFile" json:"attachmentFile" ja:"添付ファイル"`
}

type updateTodoRequest struct {
	ID             int    `param:"id" example:"1" ja:"TodoID" validate:"required"`
	UserID         int    `json:"userID" example:"1" ja:"ユーザーID" validate:"required"`
	Title          string `json:"title" example:"サンプルタイトル" ja:"タイトル" validate:"required,max=255"`
	AttachmentFile string `json:"attachmentFile" ja:"添付ファイル"`
	DoneFlag       bool   `json:"doneFlag" example:"true" ja:"完了フラグ" validate:"required"`
}
