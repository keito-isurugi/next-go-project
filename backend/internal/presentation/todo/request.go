package handler

type registerTodoRequest struct {
	Title string `json:"title" example:"サンプルタイトル" ja:"タイトル" validate:"required,max=255"`
}

type updateTodoRequest struct {
	Title    string `json:"title" example:"サンプルタイトル" ja:"タイトル" validate:"required,max=255"`
	DoneFlag bool   `json:"doneFlag" example:"true" ja:"完了フラグ" validate:"required"`
}
