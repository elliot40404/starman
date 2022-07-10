/*
Copyright © 2022 Avishek <avishek@avishek.co.in>

*/
package cmd

import (
	"fmt"

	"github.com/elliot40404/starman/env"
	// "fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

type App = env.APP

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all startup applications",
	Long: `list all startup applications and their status.
For example:
	sm list`,
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("LISTING CONFIG")
		data := env.READ_FILE()
		configs := regexp.MustCompile("(?m)"+`^@R.+\n.+`).FindAllString(data, -1)
		apps := make([]App, 0)
		for _, config := range configs {
			if regexp.MustCompile(`APP`).MatchString(config) {
				name := strings.Split(regexp.MustCompile(`@R.+`).FindString(config), "APP")[1]
				path := strings.Split(regexp.MustCompile(`ST.+`).FindString(config), "\"\"")[1]
				apps = append(apps, App{Name: strings.Trim(name, " "), Path: strings.Trim(path, " ")})
			}
		}
		if len(apps) == 0 {
			fmt.Println("Startup list is empty! Start adding apps using `sm add`")
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
