package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the entry point
var (
	rootCmd = &cobra.Command{
		Use:   "{{service_name}}-sc",
		Short: "{{service_name}} consumer daemon command line interface",
	}
)

func main() {
	cobra.EnableCommandSorting = false

	rootCmd.AddCommand(startCmd())
	rootCmd.AddCommand(keysCmd)
	rootCmd.AddCommand(invokeCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
