package commands

import (
	"github.com/anonymous961/todo-cli/internal/storage"
	"github.com/spf13/cobra"
)

func NewRootCommand(store storage.ExcelStorage) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "todo",
		Short: "A simple todo CLI with excel storage",
	}

	cmd.AddCommand(
		NewAddCommand(store),
		NewListCommand(store),
	)

	return cmd
}
