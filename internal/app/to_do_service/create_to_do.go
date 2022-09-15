package todo_service

import (
	"context"
	"log"

	pb "github.com/nekitvand/to_do_service/pkg/to_do_service"
)


func (i *Implementation) CreateToDo(ctx context.Context, req *pb.CreateToDoRequest) (*pb.CreateToDoResponse, error) {
	todo := req.GetToDo()
	result, err := i.todoService.CreateToDo(ctx, todo.GetId(),todo.GetName(),todo.GetText())
	if err != nil{
		log.Fatal(err)
	}
	res := &pb.CreateToDoResponse{Message: result}
	return res,err
}