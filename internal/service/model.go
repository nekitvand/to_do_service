package todo_service

type ToDo struct {
	Id int32 `db:"id"`
	Name string `db:"title"`
	Text string `db:"text"`
}


type ToDos []*ToDo

func (t ToDos) FilterById(id int32) *ToDo {
	for _, to := range t{
		if to.Id == id{
			return to
		}
	}
	return nil
}
