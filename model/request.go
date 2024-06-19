package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     int    `json:"due_date"`
}

type UpdateTaskRequest struct {
	TaskId      int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     int    `json:"due_date"`
	Status      string `json:"status"`
}

type GetTasksRequest struct {
	Title    string `form:"title"`
	Status   string `form:"status"`
	PageNum  int    `form:"page_num,default=0"`
	PageSize int    `form:"page_size,default=10"`
}
