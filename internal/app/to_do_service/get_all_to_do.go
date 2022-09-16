package todo_service

import (
	"context"

	pb "github.com/nekitvand/to_do_service/pkg/to_do_service"
)


func (i *Implementation) GetAllToDo(ctx context.Context, req *pb.GetAllToDoRequest) (*pb.GetAllToDoResponse, error) {
	t,err := i.todoService.GetAllToDo(ctx)
	todoes := make([]*pb.ToDo,len(t))
	// if err != nil {
	// 	return nil, errors.Wrap(err,"todoService.GetToDoById")
	// }
	// for _,v := range(t){
	// 	todoes = append(todoes,&pb.ToDo{Id:v.Id,Name: v.Name,Text: v.Text})
	// }
	for idx,value := range t{
		todoes[idx] = &pb.ToDo{Id: value.Id,Name: value.Name,Text: value.Text}
	}
	return &pb.GetAllToDoResponse{Todo:todoes},err
}