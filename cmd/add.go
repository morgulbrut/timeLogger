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
	Short: "Adds a new entry to the time.csv",
	Long:  "Format schould be in the format <PROJECTNAME> <DATE>(2019-02-18) <TIME>(1:30)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			dur := args[len(args)-1]
			date := args[len(args)-2]
			pr := args[:len(args)-2]
			proj := strings.Join(pr, " ")

			msg := fmt.Sprintf("%s, , %s, %s", proj, date, dur)
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
}
