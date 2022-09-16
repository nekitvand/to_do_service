package todo_service

import (
	"context"

	pb "github.com/nekitvand/to_do_service/pkg/to_do_service"
	"github.com/pkg/errors"
)

func (i *Implementation) GetToDoById(ctx context.Context, req *pb.GetToDoByIdRequest) (*pb.GetToDoByIdResponse, error) {
	t,err := i.todoService.GetToDoById(ctx,req.GetId())
	if err != nil {
		return nil, errors.Wrap(err,"todoService.GetToDoById")
	}
	pbToDo := &pb.ToDo{
				Id:   t.Id,
				Name: t.Name,
				Text: t.Text,
			}

	return &pb.GetToDoByIdResponse{ToDo: pbToDo},nil}