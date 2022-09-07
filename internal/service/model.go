package todo_service

type ToDo struct {
	Id uint64
	Name string
	Text string 
}


type ToDos []*ToDo

func (t ToDos) FilterById(id uint64) *ToDo {
	for _, to := range t{
		if to.Id == id{
			return to
		}
	}
	return nil
}
