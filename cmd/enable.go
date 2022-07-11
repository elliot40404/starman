/*
Copyright © 2022 Avishek <avishek@avishek.co.in>

*/
package cmd

import (
	"fmt"
	"github.com/elliot40404/starman/env"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a app to load on startup",
	Long: `Enable a app to load on startup
For example:
  enable app1`,
	Aliases: []string{"e"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.ToUpper(args[0])
		data := env.READ_FILE()
		var toBeEnabled string
		apps := regexp.MustCompile("(?m)"+`^@R.+\n.+`).FindAllString(data, -1)
		for _, app := range apps {
			if regexp.MustCompile(`@REM APP ` + name + "\n").MatchString(app) {
				toBeEnabled = app
			}
		}
		if toBeEnabled == "" {
			fmt.Println("App not found")
			return
		}
		if strings.Contains(strings.Split(toBeEnabled, "\n")[1], "@REM") {
			newCmd := strings.Replace(toBeEnabled, "@REM START", "START", -1)
			env.WRITE_FILE(strings.Replace(data, toBeEnabled, newCmd, -1))
			lsCmd.Run(cmd, args)
		}
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
	enableCmd.Aliases = []string{"e"}
	// TODO: implement enable all flag -> enables all disabled apps
}
