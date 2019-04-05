/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/morgulbrut/color"
	"github.com/spf13/cobra"
)

var cfgFile string
var logo = `
 _______           __                         
/_  __(_)_ _  ___ / /  ___  ___ ____ ____ ____
 / / / /  ' \/ -_) /__/ _ \/ _ '/ _ '/ -_) __/
/_/ /_/_/_/_/\__/____/\___/\_, /\_, /\__/_/   
                          /___//___/          
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "timelogger",
	Short: "A simple time logger",
	Long: `Let's you log times for your projects, one at the time.

timelogger login <project> starts logging
timelogger logout stops logging

a login while it's logging, stops logging first and then change the project.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		color.HiCyan("%s", logo)
		color.Yellow("Please enter project to login to:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			proj := scanner.Text()
			if proj == "logout" {
				Logout()
			} else {
				Login(scanner.Text())
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
