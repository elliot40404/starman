package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"github.com/elliot40404/starman/env"
)

var addCmd = &cobra.Command{
	Use:   "add [app_name] [app_path]",
	Short: "Add a new app to your startup list",
	Long: `Add a new app to your startup list
For example:
  add spotify C:/Users/Elliot/Music/Spotify/Spotify.exe`,
	Args:    cobra.ExactArgs(2),
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.ToUpper(args[0])
		path := args[1]
		data := env.READ_FILE()
		newCmd := env.FORMAT_CMD(name, path)
		if strings.Contains(data, newCmd) {
			fmt.Println("App already exists")
			return
		}
		if strings.Contains(data, name) {
			fmt.Println("App name already exists! Please choose a different name")
			return
		}
		if strings.Contains(data, path) {
			// TODO: Check if the path is the same by removing slashes
			fmt.Println("App path already exists! Please choose a different path")
			return
		}
		data += "\n\n" + newCmd
		env.WRITE_FILE(data)
		lsCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Aliases = []string{"a"}
}
