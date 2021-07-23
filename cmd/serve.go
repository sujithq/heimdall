package cmd

import (
	"fmt"
	"github.com/moducate/heimdall/internal/server"
	"github.com/moducate/x/osx"
	"github.com/spf13/cobra"
)

func MakeServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Serves Heimdall's HTTP API",
		Run: func(cmd *cobra.Command, args []string) {
			srv := server.New()
			srv.ListenAndServe(fmt.Sprintf(":%s", osx.Getenv("PORT", "1470")))
		},
	}
}

var serveCmd = MakeServeCmd()

func init() {
	rootCmd.AddCommand(serveCmd)
}
