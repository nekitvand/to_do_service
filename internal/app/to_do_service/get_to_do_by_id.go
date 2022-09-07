package todo_service

import (
	"context"

	proto_to_do_service "github.com/nekitvand/to_do_service/pkg/to_do_service"
	todo "github.com/nekitvand/to_do_service/internal/service"
	"github.com/pkg/errors"
)

func (i *Implementation) GetToDoById(ctx context.Context, req *proto_to_do_service.GetToDoByIdRequest) (*proto_to_do_service.GetToDoByIdResponse, error) {
	t,err := i.todoService.GetToDoById(ctx,req.GetId())
	if err != nil {
		return nil, errors.Wrap(err,"todoService.GetToDoById")
	}

	return makeGetToDoById(t), nil
}

func makeGetToDoById(to_do *todo.ToDo) *proto_to_do_service.GetToDoByIdResponse {

	pbToDo := &proto_to_do_service.ToDo{
		Id:   uint64(to_do.Id),
		Name: to_do.Name,
		Text: to_do.Text,
	}

	return &proto_to_do_service.GetToDoByIdResponse{
		ToDo: pbToDo,

	}
}