package command

import (
	"github.com/simonsargeant/telemeterm/internal/display"
	"github.com/spf13/cobra"
)

func NewRun() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Displays traces in the terminal",
		Long:  "Displays traces in the terminal",
		Run: func(cmd *cobra.Command, args []string) {
			display.Run()
		},
	}
}
