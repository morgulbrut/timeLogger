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
	"strconv"
	"strings"

	"github.com/morgulbrut/timeLogger/consts"

	"github.com/morgulbrut/color"
	"github.com/spf13/cobra"
	"github.com/yunabe/easycsv"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Shows the time accumulated on a project",
	Run: func(cmd *cobra.Command, args []string) {
		project := strings.Join(args, " ")
		time := getTime(project)
		color.Green("Total time for %q: %0.2f h", project, time)
	},
}

func init() {
	showCmd.AddCommand(timeCmd)
}

func getTime(project string) float64 {
	color.Yellow("Showing total time for %q", project)
	time := 0.0
	r := easycsv.NewReaderFile(consts.TimeLogFile)
	var entry struct {
		Project string `index:"0"`
		Time    string `index:"3"`
	}
	for r.Read(&entry) {
		if entry.Project == project {
			color.Yellow("%s", entry.Time)
			h, _ := strconv.ParseFloat(strings.Split(entry.Time, ":")[0], 64)
			m, _ := strconv.ParseFloat(strings.Split(entry.Time, ":")[1], 64)

			time += h + m/60
		}
	}
	if err := r.Done(); err != nil {
		log.Fatalf("Failed to read a CSV file: %v", err)
	}

	return time
}
