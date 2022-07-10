/*
Copyright © 2022 Avishek <avishek@avishek.co.in>

*/
package cmd

import (
	"github.com/elliot40404/starman/env"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "sm",
	Version: "v0.1.0-alpha",
	Short: "A cli startup application manager for windows",
	Long: `A simple yet intutive and easy to use cli for managing startup applications on windows.

For example:
  sm add <app name> <app path>
  sm remove <app name>
  sm ls
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//  TODO: add logging w/ log file in same home directory
	// check for home directory
	if _, err := os.Stat(env.HOMEDIR()); os.IsNotExist(err) {
		err := os.Mkdir(env.HOMEDIR(), 0755)
		if err != nil {
			panic(err)
		}
	}
	// check for config file
	if _, err := os.Stat(env.CONFIG_FILE()); os.IsNotExist(err) {
		_, err := os.Create(env.CONFIG_FILE())
		if err != nil {
			panic(err)
		}
		// TODO: add privilege escalation for symlink
		// err = os.Symlink(env.CONFIG_FILE(), env.STARTUP_DIR())
		// if err != nil {
		// 	panic(err)
		// }
	}
	// check file contents - if empty, add default config
	if _, err := os.Stat(env.CONFIG_FILE()); err == nil {
		file, err := ioutil.ReadFile(env.CONFIG_FILE())
		if err != nil {
			panic(err)
		}
		if len(file) == 0 {
			err := ioutil.WriteFile(env.CONFIG_FILE(), []byte(env.CONFIG_FILE_DEFAULT()), 0644)
			if err != nil {
				panic(err)
			}
		}
	}
}
