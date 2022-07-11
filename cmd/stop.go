/*
Copyright © 2022 Avishek <avishek@avishek.co.in>

*/
package cmd

import (
	"fmt"
	"os"
	"github.com/elliot40404/starman/env"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop starman from starting apps on startup",
	Long: `Stops starman from starting apps on startup
For example:
  starman stop`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(env.STARTUP_FILE()); err == nil {
			os.Remove(env.STARTUP_FILE())
			fmt.Println("Stopped starman from starting apps on startup")
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
