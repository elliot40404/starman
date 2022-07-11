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
	Use:     "starman",
	Version: "v0.1.0-alpha",
	Short:   "A cli tool for managing your start up applications",
	Long: `A simple yet intutive and easy to use cli for managing startup applications on windows.

For example:
  starman add <app name> <app path>
  starman remove <app name>
  starman ls
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		env.WRITE_LOG(err.Error())
		os.Exit(1)
	}
}

func init() {
	// check for home directory
	if _, err := os.Stat(env.HOMEDIR()); os.IsNotExist(err) {
		err := os.Mkdir(env.HOMEDIR(), 0755)
		if err != nil {
			env.WRITE_LOG(err.Error())
			panic(err)
		}
	}
	// check for config file
	if _, err := os.Stat(env.CONFIG_FILE()); os.IsNotExist(err) {
		_, err := os.Create(env.CONFIG_FILE())
		if err != nil {
			env.WRITE_LOG(err.Error())
			panic(err)
		}
	}
	// check file contents - if empty, add default config
	if _, err := os.Stat(env.CONFIG_FILE()); err == nil {
		file, err := ioutil.ReadFile(env.CONFIG_FILE())
		if err != nil {
			env.WRITE_LOG(err.Error())
			panic(err)
		}
		if len(file) == 0 {
			err := ioutil.WriteFile(env.CONFIG_FILE(), []byte(env.CONFIG_FILE_DEFAULT()), 0644)
			if err != nil {
				env.WRITE_LOG(err.Error())
				panic(err)
			}
		}
	}
}
