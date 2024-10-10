
package cmd

import (
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "todolist",
    Short: "A simple CLI to-do list application",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {
  rootCmd.AddCommand(addCmd)
  rootCmd.AddCommand(deleteCmd)
  rootCmd.AddCommand(listCmd)
  rootCmd.AddCommand(truncCmd)
  rootCmd.AddCommand(didCmd)
  rootCmd.AddCommand(histoCmd)
}
