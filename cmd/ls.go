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

type App = env.APP

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all startup applications",
	Long: `list all startup applications and their status.
For example:
	starman list`,
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		data := env.READ_FILE()
		configs := regexp.MustCompile("(?m)"+`^@R.+\n.+`).FindAllString(data, -1)
		apps := make([]App, 0)
		for _, config := range configs {
			if regexp.MustCompile(`APP`).MatchString(config) {
				name := strings.Split(regexp.MustCompile(`@R.+`).FindString(config), "APP")[1]
				path := strings.Split(regexp.MustCompile(`ST.+`).FindString(config), "\"\"")[1]
				disabled := regexp.MustCompile(`@REM START`).MatchString(config)
				apps = append(apps, App{Name: strings.Trim(name, " "), IsDisabled: disabled, Path: strings.Trim(path, " ")})
			}
		}
		if len(apps) == 0 {
			fmt.Println("Startup list is empty! Start adding apps using `starman add`")
			return
		}
		env.TABLE(apps)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Aliases = []string{"l"}
	// TODO: add flags with ability to list all apps, enabled apps, disabled apps
}
