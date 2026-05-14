package main

import "fmt"

type Todo struct {
	Title string
	ID    int64
	Done  bool
}

type MemoryStore struct {
	todos  []Todo
	NextID int64
}

func (m *MemoryStore) Create(title string) {
	m.NextID++
	todo := Todo{
		ID:    m.NextID,
		Title: title,
		Done:  false,
	}
	m.todos = append(m.todos, todo)
}

func main() {
	store := MemoryStore{}
	store.Create("learn golang")
	store.Create("learn API")
	store.Create("learn Todo")
	store.Create("learn postgres")
	store.Create("learn docker")


	for _, todo := range store.todos{
		fmt.Println(todo.ID, todo.Title)
	}

	

}
