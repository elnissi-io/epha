package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "epha"}
	rootCmd.AddCommand(renderCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
