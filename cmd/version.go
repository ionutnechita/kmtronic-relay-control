// Copyright (c) 2021, Project KMtronic Relay Control, Ionut Nechita.
// All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"log"

	"../pkg/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "KMtronic Relay Control - View application information",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(app.GetVERSIONAPP())
	},
}
