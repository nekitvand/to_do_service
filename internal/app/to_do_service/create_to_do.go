package todo_service

import (
	"context"

	pb "github.com/nekitvand/to_do_service/pkg/to_do_service"
)


func (i *Implementation) CreateToDo(ctx context.Context, req *pb.CreateToDoRequest) (*pb.CreateToDoResponse, error) {
	todo := req.GetToDo()
	result, err := i.todoService.CreateToDo(ctx, todo.GetId(),todo.GetTitle(),todo.GetText())
	if err != nil{
		return &pb.CreateToDoResponse{Message: err.Error()}, err
	}
	res := &pb.CreateToDoResponse{Message: result}
	return res,err
}