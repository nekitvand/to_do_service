package todo_service

import (
	"context"

	pb "github.com/nekitvand/to_do_service/pkg/to_do_service"
	"github.com/pkg/errors"
)


func (i *Implementation) GetAllToDo(ctx context.Context, req *pb.GetAllToDoRequest) (*pb.GetAllToDoResponse, error) {
	t,err := i.todoService.GetAllToDo(ctx)
	todoes := make([]*pb.ToDo,len(t))
	if err != nil {
		return nil, errors.Wrap(err,"todoService.GetToDoById")}
	for idx,value := range t{
		todoes[idx] = &pb.ToDo{Id: value.Id,Name: value.Name,Text: value.Text}
	}
	return &pb.GetAllToDoResponse{Todo:todoes},err
}
