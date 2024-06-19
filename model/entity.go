package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Task struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	DueDate     int    `db:"duedate"`
	Status      string `db:status`
	CreatedBy   int    `db:created_by`
}
