package models

import (
	"math/rand"
	"time"
)

type Todo struct {
	ID        string
	Task      string
	Complete  bool
	Category  string
	DueDate   time.Time
	Priority  int // 1=High, 2=Medium, 3=Low
	CreatedAt time.Time
}

func NewTodo(task string, category string, priority int, dueDate time.Time) *Todo {
	return &Todo{
		ID:        generateID(),
		Task:      task,
		Complete:  false,
		Category:  category,
		DueDate:   dueDate,
		Priority:  priority,
		CreatedAt: time.Now(),
	}
}

func generateID() string {
	// rand.Seed(time.Now().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
