package todo_service

import (
	"context"

	pb "github.com/nekitvand/to_do_service/pkg/to_do_service"
	"github.com/pkg/errors"
)


func (i *Implementation) GetAllToDo(ctx context.Context, req *pb.GetAllToDoRequest) (*pb.GetAllToDoResponse, error) {
	todoes := []*pb.ToDo{}
	t,err := i.todoService.GetAllToDo(ctx)
	if err != nil {
		return nil, errors.Wrap(err,"todoService.GetToDoById")
	}
	for _,v := range(t){
		todoes = append(todoes,&pb.ToDo{Id:v.Id,Name: v.Name,Text: v.Text})
	}
	return &pb.GetAllToDoResponse{Todo:todoes},err
}