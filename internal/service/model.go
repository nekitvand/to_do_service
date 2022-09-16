package todo_service

type ToDo struct {
	Id int32 `db:"id"`
	Name string `db:"title"`
	Text string `db:"text"`
}
