/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package cmd

import (
	"log"
	"os"
	"runtime"

	"github.com/jedib0t/go-pretty/table"
	"github.com/morgulbrut/timeLogger/consts"
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
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
