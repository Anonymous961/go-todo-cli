package commands

import (
	"fmt"

	"github.com/anonymous961/todo-cli/internal/storage"
	"github.com/spf13/cobra"
)

func NewDeleteCommand(store storage.Storage) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete a todo",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			err := store.Delete(id)
			if err != nil {
				fmt.Printf("Error Deleting todo with ID %v\n", id)
				return
			}
			fmt.Printf("Deleted todo with ID %v\n", id)
		},
	}

	return cmd
}
