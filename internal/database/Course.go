package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, categoryId string) (Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, categoryId) VALUES ($1, $2, $3, $4)",
		id,
		name,
		description,
		categoryId,
	)

	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}

		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryId: categoryId})
	}
	return courses, nil
}

func (c *Course) FindByCategoryId(idCategory string) ([]Course, error) {
	row, err := c.db.Query("SELECT * from courses where categoryId = $1", idCategory)

	if err != nil {
		return nil, err
	}
	defer row.Close()

	courses := []Course{}
	for row.Next() {
		var id, name, description, categoryId string
		if err := row.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}

		courses = append(courses, Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryId:  categoryId,
		})

	}

	return courses, nil
}
