/*
Copyright © 2022 Avishek <avishek@avishek.co.in>

*/
package cmd

import (
	"fmt"
	"github.com/elliot40404/starman/env"
	"github.com/spf13/cobra"
	"os/exec"
	"regexp"
	"strings"
)

// delayCmd represents the delay command
var runCmd = &cobra.Command{
	Use:   "run [app_name]",
	Short: "Run any app added to starman",
	Long: `Run any app added to starman
For example:
  starman run app - starts app
  starman r app - starts app`,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.ToUpper(args[0])
		data := env.READ_FILE()
		var toRun string
		apps := regexp.MustCompile("(?m)"+`^@R.+\n.+`).FindAllString(data, -1)
		for _, app := range apps {
			if regexp.MustCompile(`@REM APP ` + name + "\n").MatchString(app) {
				toRun = app
			}
		}
		if toRun == "" {
			fmt.Println("App not found")
			lsCmd.Run(cmd, args)
			return
		}
		app := strings.Split(toRun, "\n")[1]
		appPath := strings.Split(app, " ")[2]
		run := exec.Command("powershell", "start", appPath)
		run.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Aliases = []string{"r"}
}
