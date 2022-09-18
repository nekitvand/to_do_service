package todo_service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

type Service struct {
	repository IRepository
}

type IRepository interface {
	CreateToDo(context.Context, *ToDo) (message string, err error)
	GetAllToDo(context.Context) (todoes []*ToDo, err error)
	GetToDoById(ctx context.Context, id int32)(todo *ToDo, err error)
	DeleteToDo(ctx context.Context,id int32)(string,error)
	UpdateFieldToDo(ctx context.Context, field string, value string,id int32)(string,error)
	UpdateToDo(ctx context.Context,id int32,title string,text string)(message string,err error)
}

func NewService(repo IRepository) *Service {
	return &Service{repository: repo}
}

var ErrNoToDo = errors.New("no todo")

func (s Service) GetToDoById(ctx context.Context, id int32) (*ToDo, error) {
	todo, err := s.repository.GetToDoById(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "repository.GetToDoById")
	}
	return todo, nil
}

func (s Service) GetAllToDo(ctx context.Context) ([]*ToDo, error) {
	todo, err := s.repository.GetAllToDo(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.GetAllToDo")
	}
	return todo, err
}

func (s Service) CreateToDo(ctx context.Context, id int32, name string, text string) (message string, err error) {
	todo := &ToDo{Id: id, Title: name, Text: text}
	str, err := s.repository.CreateToDo(ctx, todo)
	if err != nil {
		return "dont add to DB", err
	}
	return str, nil
}

func (s Service) DeleteToDo(ctx context.Context, id int32) (message string, err error) {
	str, err := s.repository.DeleteToDo(ctx, id)
	if err != nil {
		return fmt.Sprintf("dont delete todo with id: %v",id), err
	}
	return str, nil
}


func (s Service)UpdateFieldToDo(ctx context.Context, field string, value string,id int32)(message string,err error){
	str,err := s.repository.UpdateFieldToDo(ctx,field,value,id)
	if err != nil {
		return fmt.Sprintf("dont update field: %s in todo with id: %v",field,id),err
	}
	return str, nil

}

func (s Service)UpdateToDo(ctx context.Context,id int32,title string,text string)(message string,err error){
	str,err := s.repository.UpdateToDo(ctx,id,title,text)
	if err != nil{
		return fmt.Sprintf("dont update todo with id: %v",str),err
	}
	return fmt.Sprintf("update todo with id: %v",str),nil
}