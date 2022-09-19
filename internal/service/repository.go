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
	query := sq.Insert("todo").PlaceholderFormat(sq.Dollar).Columns("id", "title", "text").Values(todo.Id, todo.Title, todo.Text).Suffix("RETURNING id").RunWith(r.DB)
	rows, err := query.QueryContext(ctx)
	for rows.Next(){
		rows.Scan(&message)
	}
	if err != nil {
		return "",err
	}
	return message, nil

}

func (r *Repository) GetAllToDo(ctx context.Context) (todoes []*ToDo, err error) {
	query := sq.Select("*").From("todo")

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = r.DB.SelectContext(ctx, &todoes, s, args...)
	if err != nil {
		return nil, err
	}
	return todoes, err
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
	query := sq.Delete("todo").Where(sq.Eq{"id":id}).PlaceholderFormat(sq.Dollar).Suffix("RETURNING id").RunWith(r.DB)
	rows, err := query.QueryContext(ctx)
	for rows.Next(){
		rows.Scan(&message)
	}

	if err != nil{
		return "",err
	}
	return message,nil
}

func (r *Repository)UpdateFieldToDo(ctx context.Context,field string, value string,id int32)(message map[string]interface{},err error){
	query,args,_ := sq.Update("todo").Set(field,value).Where(sq.Eq{"id":id}).PlaceholderFormat(sq.Dollar).Suffix(fmt.Sprintf("RETURNING id,%v",field)).ToSql()
	rows,_  := r.DB.QueryxContext(ctx,query,args...)
	results := make(map[string]interface{})
	for rows.Next(){
		err := rows.MapScan(results)
		fmt.Println(results)
		if err != nil{
			return nil,err
		}
	}
	return results,nil
}

func (r *Repository)UpdateToDo(ctx context.Context,id int32,title string,text string)(message string,err error){
	query,args,_ := sq.Update("todo").Set("title",title).Set("text",text).Where(sq.Eq{"id":id}).PlaceholderFormat(sq.Dollar).Suffix("RETURNING id").ToSql()
	rows,err := r.DB.QueryxContext(ctx,query,args...)
	for rows.Next(){
		rows.Scan(&message)
	}
	if err != nil{
		return "",err
	}
	return message, nil
}