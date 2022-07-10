/*
Copyright © 2022 Avishek <avishek@avishek.co.in>

*/
package cmd

import (
	"github.com/elliot40404/starman/env"
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

var rmCmd = &cobra.Command{
	Use:   "rm [app_name]",
	Short: "Remove an app from your startup list",
	Long: `Remove an app from your startup list
For example:
  remove spotify`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := env.READ_FILE()
		name := strings.ToUpper(args[0])
		var toBeRemoved string
		apps := regexp.MustCompile("(?m)"+`^@R.+\n.+`).FindAllString(data, -1)
		for _, app := range apps {
			if regexp.MustCompile(`@REM APP ` + name + "\n").MatchString(app) {
				toBeRemoved = app
			}
		}
		if toBeRemoved == "" {
			fmt.Println("App not found")
			return
		}
		data = strings.Replace(data, "\n\n"+toBeRemoved, "", -1)
		env.WRITE_FILE(data)
		lsCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
