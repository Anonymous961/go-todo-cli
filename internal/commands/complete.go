package commands

import (
	"fmt"

	"github.com/anonymous961/todo-cli/internal/storage"
	"github.com/spf13/cobra"
)

func NewCompleteCommand(store storage.Storage) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "complete <id>",
		Short: "update task completion",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			err := store.Complete(id)
			if err != nil {
				fmt.Printf("Error updating todo with %v\n", id)
				return
			}
			fmt.Printf("Updated todo with ID %v\n", id)
		},
	}

	return cmd
}
