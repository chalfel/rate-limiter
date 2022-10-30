package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "rate-limiter",
		Short: "Rate Limiter CLI",
	}

	rootCmd.AddCommand(NewServerCmd())

	return rootCmd
}
