package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.0.1"

func NewVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the current version",
		Long:  "Prints information about the current build version of oterml",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
}
