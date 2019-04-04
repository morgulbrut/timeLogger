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

	"github.com/morgulbrut/color"
	"github.com/morgulbrut/timeLogger/consts"
	"github.com/morgulbrut/timeLogger/utils"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		Logout()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Logout() {
	if _, err := os.Stat(consts.TimeLockFile); err == nil {

		dat, err := ioutil.ReadFile(consts.TimeLockFile)
		if err != nil {
			utils.Error(err.Error())
		}
		log := strings.TrimRight(string(dat), "\r\n")
		logtime := time.Now()
		timeStrings := strings.Split(log, ";")
		timeString := timeStrings[len(timeStrings)-1]
		loginTime, err := time.Parse(consts.TimeFmtString, timeString)
		fmt.Println(loginTime)
		fmt.Println(logtime)
		duration := logtime.Sub(loginTime)
		durString := fmt.Sprintf("%d:%02d", int(duration.Minutes()/60), int(duration.Minutes())%60)
		msg := fmt.Sprintf("Projekt: %s, login:%s, logout: %s, duration: %s", timeStrings[1],
			logtime.Format(consts.TimeLogFmtString), loginTime.Format(consts.TimeLogFmtString), durString)
		color.Green(msg)
		if err := utils.AppendToFile(msg, "time.log"); err != nil {
			utils.Error(err.Error())
		}
		os.Remove("time.lck")
	} else {
		utils.Error("Not logged in")
	}
}
