package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrEmptyTitle = errors.New("title cannot be empty")
)

type Todo struct {
	Title       string    `json:"title"`
	ID          int64     `json:"id"`
	Done        bool      `json:"done"`
	CompletedAt time.Time `json:"completed_at"`
}

type CreateTodoInput struct {
	Title string `json:"title"`
}

func (t CreateTodoInput) Validate() error {
	if t.Title == "" {
		return ErrEmptyTitle
	}
	return nil
}

type UpdateTodoInput struct {
	Done bool `json:"done"`
}

type TodoStore interface {
	Create(ctx context.Context, input CreateTodoInput) (Todo, error)
	GetAll(ctx context.Context) ([]Todo, error)
	GetByID(ctx context.Context, id int64) (Todo, error)
	Update(ctx context.Context, id int64, input UpdateTodoInput) (Todo, error)
	Delete(ctx context.Context, id int64) error
}

func main() {
	fmt.Println("====CHECKPOINT1: Create a todo manually")
	t1 := Todo{ID: 1, Title: "learn go", Done: false}
	fmt.Println("Todo:", t1)
	fmt.Println("ID:", t1.ID)
	fmt.Println("Done:", t1.Done)

	fmt.Println("===CHECKPOINT2=== validate a valid input")
	input1 := CreateTodoInput{
		Title: "learn postgress",
	}
	err := input1.Validate()
	fmt.Println("Input:", input1)
	fmt.Println("validation error: ", err)

	fmt.Println("===CHECKPOINT3=== validate a valid input")
	input2 := CreateTodoInput{
		Title: "",
	}
	err1 := input2.Validate()
	fmt.Println("input:", input2)
	fmt.Println("validation error:", err1)

	fmt.Println("\n===== CHECKPOINT 4: UpdateTodoInput =====")

	update1 := UpdateTodoInput{
		Done: true,
	}
	fmt.Println("Update input:", update1)
	fmt.Println("New done value:", update1.Done)

	fmt.Println("\n===== CHECKPOINT 5: Simulate updating t1 =====")
	t1.Done = update1.Done
	t1.CompletedAt = time.Now()

	fmt.Println("Updated Todo:", t1)
	fmt.Println("Done:", t1.Done)
	fmt.Println("CompletedAt:", t1.CompletedAt)
}
