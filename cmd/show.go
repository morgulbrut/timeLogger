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
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/timeLogger/consts"
	"github.com/morgulbrut/timeLogger/utils"
	"github.com/spf13/cobra"
	"github.com/yunabe/easycsv"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows data from the log",
	Run: func(cmd *cobra.Command, args []string) {

		r := easycsv.NewReaderFile(consts.TimeLogFile)
		var entry struct {
			Project string `index:"0"`
			Date    string `index:"2"`
			Time    string `index:"3"`
		}
		if _, err := os.Stat(consts.TimeLockFile); err == nil {
			dat, err := ioutil.ReadFile(consts.TimeLockFile)
			if err != nil {
				utils.Error(err.Error())
			}
			log := strings.TrimRight(string(dat), "\r\n")
			timeStrings := strings.Split(log, ";")
			timeString := timeStrings[len(timeStrings)-1]
			loginTime, err := time.Parse(consts.TimeFmtString, timeString)
			logtime := time.Now()
			duration := logtime.Sub(loginTime)
			durString := fmt.Sprintf("%d:%02d", int(duration.Minutes()/60), int(duration.Minutes())%60)
			color256.PrintHiOrange("Logged in to: %s for %s", timeStrings[1], durString)
		} else {
			color256.PrintHiMagenta("Not logged in")
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		if runtime.GOOS == "windows" {
			t.SetStyle(table.StyleDouble)
		} else {
			t.SetStyle(table.StyleColoredDark)
		}
		t.AppendHeader(table.Row{"Project", "Date", "Time"})
		for r.Read(&entry) {
			if entry.Project != "Project" {
				t.AppendRow(table.Row{entry.Project, entry.Date, entry.Time})
			}
			if err := r.Done(); err != nil {
				log.Fatalf("Failed to read a CSV file: %v", err)
			}
		}
		t.Render()
		color256.PrintCyan("Press any key to close")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
