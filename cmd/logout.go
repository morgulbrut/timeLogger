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
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/timeLogger/consts"
	"github.com/morgulbrut/timeLogger/utils"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Stops logging and writes times to log file",
	Run: func(cmd *cobra.Command, args []string) {
		Logout()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

// Logout stops logging and writes times to log file
func Logout() {
	if _, err := os.Stat(consts.TimeLockFile); err == nil {

		dat, err := ioutil.ReadFile(consts.TimeLockFile)
		if err != nil {
			utils.Error(err.Error())
		}

		// if file dont exist add ne file
		_, err = os.Stat(consts.TimeLogFile)
		if err == nil {
			// empty because
		} else if os.IsNotExist(err) {
			if err := utils.AppendToFile("Project, Login, Logout, Duration", consts.TimeLogFile); err != nil {
				utils.Error(err.Error())
			}
		} else {
			color256.PrintHiRed("file %s stat error: %v", consts.TimeLogFile, err)
		}

		log := strings.TrimRight(string(dat), "\r\n")
		logtime := time.Now()
		timeStrings := strings.Split(log, ";")
		timeString := timeStrings[len(timeStrings)-1]
		loginTime, err := time.Parse(consts.TimeFmtString, timeString)
		duration := logtime.Sub(loginTime)
		durString := fmt.Sprintf("%d:%02d", int(duration.Minutes()/60), int(duration.Minutes())%60)
		msg := fmt.Sprintf("%s, %s, %s, %s", timeStrings[1],
			loginTime.Format(consts.TimeLogFmtString), logtime.Format(consts.TimeLogFmtString), durString)
		color256.PrintHiGreen(msg)
		if err := utils.AppendToFile(msg, consts.TimeLogFile); err != nil {
			utils.Error(err.Error())
		}
		os.Remove(consts.TimeLockFile)
	} else {
		utils.Error("Not logged in")
	}
}
