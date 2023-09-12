package main

import (
	"github.com/rl404/mal-cover/internal/utils"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "mal-cover",
		Short: "MAL Cover API",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Run API server",
		RunE: func(*cobra.Command, []string) error {
			return server()
		},
	})

	if err := cmd.Execute(); err != nil {
		utils.Fatal(err.Error())
	}
}
