package todo_service

import (
	"context"

	pb "github.com/nekitvand/to_do_service/pkg/to_do_service"
	"github.com/pkg/errors"
)


func (i *Implementation) UpdateToDo(ctx context.Context, req *pb.UpdateToDoRequest) (*pb.UpdateToDoResponse, error) {
	str,err := i.todoService.UpdateToDo(ctx,req.GetId(),req.GetTitle(),req.GetText())
	if err != nil{
		return nil,errors.Wrap(err,"todoService.UpdateToDo")
	}
	return &pb.UpdateToDoResponse{Message: str},nil}