package main

import (
	"errors"
	"time"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrEmptyTitle = errors.New("title cannot be empty")
)

type Todo struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Done       bool      `json:"done"`
	CreateadAt time.Time `json:"created_at"`
}

type CreateTodoInput struct {
	Title string `json:"title"`
}

// Validate ensures the input is usable before touching the database.
func (i CreateTodoInput) Validate() error {
	if i.Title == "" {
		return ErrEmptyTitle
	}
	return nil
}
