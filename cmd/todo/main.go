package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/anonymous961/todo-cli/internal/storage"
)

func main() {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	storagePath := filepath.Join(homeDir, ".todo", "todos.xlsx")

	if err := os.MkdirAll(filepath.Dir(storagePath), 0755); err != nil {
		log.Fatal(err)
	}

	fmt.Println(storagePath)

	store := storage.NewExcelStorage(storagePath)
	fmt.Println(store)
}
