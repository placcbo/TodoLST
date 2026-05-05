package main

import (
	"database/sql"
	"time"
)

type Todo struct {
	Title     string    `json:"title"`
	ID        int       `json:"id"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdat"`
}
type TodoStore interface {
	Create(title string) (Todo, error)
	GetAll() ([]Todo, error)
	GetByID(id int) (Todo, error)
	Update(id int, done bool) (Todo, error)
	Delete(id int) (Todo, error)
}
type PostgressStore struct{
	db *sql.DB
}
