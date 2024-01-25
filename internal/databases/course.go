package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Course struct{
	db *sql.DB
	Id string
	Name string
	Description string
	CategoryId string
}

func NewCourse(db *sql.DB) *Course{
	return &Course{db:db}
}

func (c *Course) Create(name, description, categoryId string) (*Course, error){
	id:= uuid.New().String()
	query := "INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)"
	_, err := c.db.Exec(query, id, name, description, categoryId)

	if err != nil{
		return nil, err
	}

	return &Course{
		Id: id,
		Name: name,
		Description: description,
		CategoryId: categoryId,
	},nil
}

func (c *Course) FindAll() ([]Course, error){
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil{
		return nil, err
	}

	defer rows.Close()
	courses := []Course{}

	for rows.Next(){
		var id, name, description,categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil{
			return nil, err
		}

		courses = append(courses, Course{
			Id: id,
			Name: name,
			Description: description,
			CategoryId: categoryId,
		})
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error){
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryID)

	if err != nil{
		return nil, err
	}

	defer rows.Close()
	courses := []Course{}

	for rows.Next(){
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil{
			return nil, err
		}

		courses = append(courses, Course{
			Id: id,
			Name: name,
			Description: description,
			CategoryId: categoryId,
		})
	}

	return courses, nil
}