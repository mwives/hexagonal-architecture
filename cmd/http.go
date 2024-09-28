package cmd

import (
	"fmt"

	"github.com/mwives/hexagonal-architecture/adapters/web/server"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start the HTTP server",
	Long:  "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewWebServer()
		server.Service = &productService
		fmt.Println("Starting HTTP server on port 8080")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
