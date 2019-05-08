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

	"github.com/morgulbrut/color"
	"github.com/morgulbrut/timeLogger/consts"
	"github.com/spf13/cobra"
	"github.com/yunabe/easycsv"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Lists the projects in the log",
	Run: func(cmd *cobra.Command, args []string) {
		projects := getProjects()
		for pr := range projects {
			color.Green(projects[pr])
		}
	},
}

func init() {
	showCmd.AddCommand(projectsCmd)
}

func getProjects() []string {
	color.Yellow("Listing projects")
	r := easycsv.NewReaderFile(consts.TimeLogFile)
	set := make(map[string]struct{})
	projects := []string{}
	var entry struct {
		Project string `index:"0"`
	}
	for r.Read(&entry) {

		if entry.Project != "project" {
			set[entry.Project] = struct{}{}
		}
	}
	if err := r.Done(); err != nil {
		log.Fatalf("Failed to read a CSV file: %v", err)
	}
	for key := range set {
		projects = append(projects, key)
	}
	return projects
}
