package models

type RawTodo struct {
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed string `json:"completed"`
}

type Todo struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"createdAt"`
	Description string `json:"description"`
	RawTodo
}
