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
}

// Login starts time logging. If it already is logging, it stops it first"
func Login(proj string) {

	if _, err := os.Stat(consts.TimeLockFile); err == nil {
		Logout()
	}

	logtime := time.Now().Format(consts.TimeFmtString)
	msg := fmt.Sprintf("project;%s;login;%s", proj, logtime)
	color.Green(msg)
	if err := utils.AppendToFile(msg, consts.TimeLockFile); err != nil {
		utils.Error(err.Error())
	}
}
