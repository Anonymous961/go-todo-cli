package models

import (
	"log"
	"os/exec"
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
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(newUUID)
}
