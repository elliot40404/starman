/*
Copyright © 2022 Avishek <avishek@avishek.co.in>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/elliot40404/starman/env"
	"github.com/spf13/cobra"
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
			ioutil.WriteFile(env.STARTUP_FILE(), []byte(env.STARTUP_FILE_DEFAULT()), 0644)
			fmt.Println("Enabled starman to start apps on startup")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
