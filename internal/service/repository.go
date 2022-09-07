package todo_service

import "context"


var todoes = ToDos{
	{
		Id: uint64(1),
		Name: "new1",
		Text: "zadvcx",
	},
	{
		Id: uint64(2),
		Name: "new2",
		Text: "cadvcx",
	},
	{
		Id: uint64(3),
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