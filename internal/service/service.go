package todo_service

import (
	"context"
	"github.com/pkg/errors"
)

type Service struct {
	repository RepositoryInterface
}


type RepositoryInterface interface {
	FindAllToDoes(context.Context)(ToDos,error)
}

func NewService(repository RepositoryInterface) Service {
	return Service{repository: repository}
}

var ErrNoToDo = errors.New("no todo")

func (s Service) GetToDoById(ctx context.Context, id int32)(*ToDo,error) {
	todo, err := s.repository.FindAllToDoes(ctx)
	if err != nil {
		return nil, errors.Wrap(err,"repository.FindAllToDoes")
	}

	to := todo.FilterById(id)
	if to == nil{
		return nil,ErrNoToDo
	}
	return to,nil
}