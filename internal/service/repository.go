package todo_service

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateToDo(ctx context.Context, todo *ToDo) (message string, err error) {
	query := sq.Insert("todo").PlaceholderFormat(sq.Dollar).Columns("id", "title", "text").Values(todo.Id, todo.Title, todo.Text).RunWith(r.DB)

	_, err = query.ExecContext(ctx)
	if err != nil {
		fmt.Println(err)
		return "dont added", err
	}
	return fmt.Sprint("added todo with id: ", todo.Id), nil

}

func (r *Repository) GetAllToDo(ctx context.Context) (todoes []*ToDo, err error) {
	query := sq.Select("*").From("todo")

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var todos []*ToDo
	err = r.DB.SelectContext(ctx, &todos, s, args...)
	if err != nil {
		return nil, err
	}
	return todos, err
}

func (r *Repository) GetToDoById(ctx context.Context, id int32)(*ToDo, error){
	query := sq.Select("*").PlaceholderFormat(sq.Dollar).From("todo").Where(sq.Eq{"id":id}).RunWith(r.DB)
	todoByID := ToDo{}
	rows, err := query.Query()
	if err != nil{
		return nil,err
	}
	for rows.Next(){
		err = rows.Scan(&todoByID.Id,&todoByID.Title,&todoByID.Text)
	}
	if err != nil{
		return nil,err
	}
	return &todoByID, nil
}

func (r *Repository)DeleteToDo(ctx context.Context,id int32)(message string,err error){
	query := sq.Delete("todo").Where(sq.Eq{"id":id}).PlaceholderFormat(sq.Dollar).RunWith(r.DB)
	_, err = query.ExecContext(ctx)
	if err != nil{
		return fmt.Sprintf("dont delete todo with id: %v",id),err
	}
	return fmt.Sprintf("delete todo with id: %v",id),nil
}