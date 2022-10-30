package cmd

import (
	"github.com/chalfel/rate-limiter/pkg/network"
	"github.com/spf13/cobra"
)

func NewServerCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "servier http application",
		RunE: func(cmd *cobra.Command, args []string) error {
			return StartServer(cmd, args)
		},
	}
	return command
}

func StartServer(cmd *cobra.Command, arg []string) error {
	router, err := network.NewRouter()
	if err != nil {
		return err
	}

	server := network.NewServer(router, ":3000")
	router.RegisterRoutes()

	if err := server.Init(); err != nil {
		return err
	}

	return nil
}
