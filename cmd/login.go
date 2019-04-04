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
	"os"
	"strings"
	"time"

	"github.com/morgulbrut/color"
	"github.com/morgulbrut/timeLogger/consts"
	"github.com/morgulbrut/timeLogger/utils"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Starts time logging. If it already is logging, it stops it first",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			proj := strings.Join(args, " ")
			Login(proj)
		} else {
			utils.Error("Project as argument needed")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Login(proj string) {

	if _, err := os.Stat(consts.TimeLockFile); err == nil {
		Logout()
	}

	logtime := time.Now().Format(consts.TimeFmtString)
	msg := fmt.Sprintf("Projekt;%s;login;%s", proj, logtime)
	color.Green(msg)
	if err := utils.AppendToFile(msg, consts.TimeLockFile); err != nil {
		utils.Error(err.Error())
	}
}
