package handler

type todoResponse struct {
	ID       int    `json:"id" example:"1"`
	Title    string `json:"title" example:"テストタイトル"`
	DoneFlag bool   `json:"doneFlag" example:"false"`
}

type listTodosResponse []todoResponse
