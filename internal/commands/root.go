package commands

import (
	"github.com/anonymous961/todo-cli/internal/storage"
	"github.com/spf13/cobra"
)

func NewRootCommand(store storage.ExcelStorage) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "todo",
		Short: "A simple todo CLI with excel storage",
		Long: `Todo CLI helps you manage your tasks with features like:
- Adding tasks with categories and priorities
- Listing and filtering tasks
- Marking tasks as complete`,
	}

	cmd.AddCommand(
		NewAddCommand(store),
		NewListCommand(store),
		NewDeleteCommand(store),
		NewCompleteCommand(store),
	)

	return cmd
}
