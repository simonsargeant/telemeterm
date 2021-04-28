package command

import "github.com/spf13/cobra"

func NewRoot() *cobra.Command {
	root := &cobra.Command{
		Use:   "telemeterm",
		Short: "Opentelemetry trace visualisation in your terminal",
		// TODO: embiggen long description
		Long: "Opentelemetry trace visualisation in your terminal",
	}

	root.AddCommand(
		NewRun(),
		NewVersion(),
	)

	return root
}
