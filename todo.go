package todo_list

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}
