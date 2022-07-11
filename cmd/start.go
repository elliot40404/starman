/*
Copyright © 2022 Avishek <avishek@avishek.co.in>

*/
package cmd

import (
	"fmt"
	"github.com/elliot40404/starman/env"
	"github.com/spf13/cobra"
	"os"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Enable starman to start apps on startup",
	Long: `Enable starman to start apps on startup
For example:
  starman start`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(env.STARTUP_FILE()); err != nil {
			data := env.READ_FILE()
			env.WRITE_FILE(data)
			fmt.Println("Enabled starman to start apps on startup")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
