package todo_service

import (
	"context"

	pb "github.com/nekitvand/to_do_service/pkg/to_do_service"
	"github.com/pkg/errors"
)

func (i *Implementation) UpdateFieldToDo(ctx context.Context, req *pb.UpdateFieldToDoRequest) (*pb.UpdateFieldToDoResponse, error) {
	var fieldName string
	var setValue string
	switch req.GetTodoOneof().(type){
	case *pb.UpdateFieldToDoRequest_Title:
		fieldName = "title"
		setValue = req.GetTitle()
	case *pb.UpdateFieldToDoRequest_Text:
		fieldName = "text"
		setValue = req.GetText()
	default:
		return &pb.UpdateFieldToDoResponse{Message:"Dont text in any field"},nil}
	
	str,err := i.todoService.UpdateFieldToDo(ctx,fieldName,setValue,req.GetId())
	if err != nil{
		return nil, errors.Wrap(err,"todoService.UpdateFieldToDo")}
	
	return &pb.UpdateFieldToDoResponse{Message: str},nil}