package todo_service

import (
	"context"

	pb "github.com/nekitvand/to_do_service/pkg/to_do_service"
	"github.com/pkg/errors"
)


func (i *Implementation) DeleteToDo(ctx context.Context, req *pb.DeleteToDoRequest) (*pb.DeleteToDoResponse, error) {
	m,err := i.todoService.DeleteToDo(ctx,req.GetId())
	if err != nil {
		return nil, errors.Wrap(err,"todoService.DeleteToDo")}
	return &pb.DeleteToDoResponse{Message: m},err
}
