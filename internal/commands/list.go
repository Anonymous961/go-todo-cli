package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/anonymous961/todo-cli/internal/models"
	"github.com/anonymous961/todo-cli/internal/storage"
	"github.com/spf13/cobra"
)

func NewListCommand(store storage.ExcelStorage) *cobra.Command {
	var (
		category string
		priority int
	)

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List add todos",
		Run: func(cmd *cobra.Command, args []string) {
			todos, err := store.List()

			if err != nil {
				fmt.Printf("Error listing todos:%v\n", err)
				return
			}

			filtered := make([]*models.Todo, 0)
			for _, todo := range todos {
				if category != "" && todo.Category != category {
					continue
				}
				if priority > 0 && todo.Priority != priority {
					continue
				}
				filtered = append(filtered, todo)
			}

			if len(filtered) == 0 {
				fmt.Println("No todos found")
				return
			}

			// Format output as a table
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tTASK\tCATEGORY\tPRIORITY\tDUE DATE\tCOMPLETE")
			for _, todo := range filtered {
				complete := "No"
				if todo.Complete {
					complete = "Yes"
				}

				shortID := todo.ID
				if len(shortID) > 8 {
					shortID = todo.ID[:8] + "..."
				}
				fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%s\t%s\n",
					shortID, // Show shortened ID
					todo.Task,
					todo.Category,
					todo.Priority,
					todo.DueDate.Format("2006-01-02"),
					complete)
			}
			w.Flush()
		},
	}

	cmd.Flags().StringVarP(&category, "category", "c", "", "Filter by category")
	cmd.Flags().IntVarP(&priority, "priority", "p", 0, "Filter by priority (1-3)")

	return cmd
}
