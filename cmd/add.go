/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package cmd

import (
	"fmt"
	"strings"

	"github.com/morgulbrut/color"
	"github.com/morgulbrut/timeLogger/consts"
	"github.com/morgulbrut/timeLogger/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			dur := args[len(args)-1]
			pr := args[:len(args)-1]
			proj := strings.Join(pr, " ")

			msg := fmt.Sprintf("%s, , , %s", proj, dur)
			color.Green(msg)
			if err := utils.AppendToFile(msg, consts.TimeLogFile); err != nil {
				utils.Error(err.Error())
			}
		} else {
			utils.Error("Time and Project as argument needed")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
