package handler

type todoResponse struct {
	ID             int    `json:"id" example:"1"`
	Title          string `json:"title" example:"テストタイトル"`
	AttachmentFile string `json:"attachmentFile" example:"添付ファイル"`
	DoneFlag       bool   `json:"doneFlag" example:"false"`
}

type listTodosResponse []todoResponse
