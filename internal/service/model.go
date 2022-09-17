package todo_service

type ToDo struct {
	Id int32 `db:"id"`
	Title string `db:"title"`
	Text string `db:"text"`
}
