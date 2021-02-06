// Copyright (c) 2021, Project KMtronic Relay Control, Ionut Nechita.
// All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
    "log"

    "github.com/spf13/cobra"
    "../pkg/app"
)

func init() {
    rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "View application information",
    Run: func(cmd *cobra.Command, args []string) {
        log.Println(app.GetVERSIONAPP())
    },
}
