package commands

import (
	"fmt"
	"time"

	"github.com/anonymous961/todo-cli/internal/models"
	"github.com/anonymous961/todo-cli/internal/storage"
	"github.com/spf13/cobra"
)

func NewAddCommand(store storage.Storage) *cobra.Command {
	var (
		category string
		priority int
		dueDate  string
	)

	cmd := &cobra.Command{
		Use:   "add [task]",
		Short: "Add a new todo",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			task := args[0]

			var parsedDueDate time.Time
			if dueDate != "" {
				var err error
				parsedDueDate, err = time.Parse("2006-01-02", dueDate)
				if err != nil {
					fmt.Printf("Invalid date format: %v\n", err)
					return
				}
			} else {
				parsedDueDate = time.Now().AddDate(0, 0, 7)
			}

			todo := models.NewTodo(task, category, priority, parsedDueDate)
			if err := store.Add(todo); err != nil {
				fmt.Printf("Error adding todo: %v\n", err)
				return
			}

			fmt.Printf("Added todo: %s (Due: %s)\n", todo.Task, todo.DueDate.Format("Jan 2, 2006"))
		},
	}
	cmd.Flags().StringVarP(&category, "category", "c", "general", "Todo category")
	cmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority (1=high, 2=medium, 3=low)")
	cmd.Flags().StringVarP(&dueDate, "due", "d", "", "Due date (YYYY-MM-DD)")

	return cmd
}
