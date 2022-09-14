package todo_service

import "context"


var todoes = ToDos{
	{
		Id: 1,
		Name: "new1",
		Text: "zadvcx",
	},
	{
		Id: 2,
		Name: "new2",
		Text: "cadvcx",
	},
	{
		Id: 3,
		Name: "new3",
		Text: "badvcx",
	},
}

type Repository struct {

}


func NewRepository() *Repository{
	return &Repository{}
}

func (r *Repository)FindAllToDoes(_ context.Context)(ToDos,error){
	return todoes,nil
}

func (r *Repository)SelectAllToDo(_ context.Context)([]struct{}){
	return nil
}