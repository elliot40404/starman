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

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable an app from starting on boot",
	Long: `Disable an app from starting on boot
For example:
  starman disable app1`,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.ToUpper(args[0])
		data := env.READ_FILE()
		var toBeDisabled string
		apps := regexp.MustCompile("(?m)"+`^@R.+\n.+`).FindAllString(data, -1)
		for _, app := range apps {
			if regexp.MustCompile(`@REM APP ` + name + "\n").MatchString(app) {
				toBeDisabled = app
			}
		}
		if toBeDisabled == "" {
			fmt.Println("App not found")
			return
		}
		if !strings.Contains(strings.Split(toBeDisabled, "\n")[1], "@REM") {
			newCmd := strings.Replace(toBeDisabled, "START", "@REM START", -1)
			env.WRITE_FILE(strings.Replace(data, toBeDisabled, newCmd, -1))
			lsCmd.Run(cmd, args)
		}
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
	disableCmd.Aliases = []string{"d"}
}
