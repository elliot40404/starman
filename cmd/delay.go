/*
Copyright © 2022 Avishek <avishek@avishek.co.in>

*/
package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/elliot40404/starman/env"
	"github.com/spf13/cobra"
)

// delayCmd represents the delay command
var delayCmd = &cobra.Command{
	Use:   "delay [time]",
	Short: "Delay for starting apps on startup. 30 second is default",
	Long: `Delay for starting apps on startup. 30 second is default
For example:
  starman delay - outputs current delay
  starman delay N - sets delay to N seconds`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := env.READ_FILE()
		currentDelay := strings.Split(regexp.MustCompile(`(?m)^TI.+`).FindString(data), " ")[2]
		if len(args) == 0 {
			fmt.Println("Current delay:", currentDelay)
		} else {
			delay := args[0]
			if !regexp.MustCompile(`^[0-9]+$`).MatchString(delay) {
				fmt.Println("Delay must be a number")
				return
			}
			data = strings.Replace(data, "TIMEOUT /T "+currentDelay, "TIMEOUT /T "+delay, 1)
			env.WRITE_FILE(data)
			fmt.Println("Delay set to", delay)
		}
	},
}

func init() {
	rootCmd.AddCommand(delayCmd)
}
