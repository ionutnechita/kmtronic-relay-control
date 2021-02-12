// Copyright (c) 2021, Project KMtronic Relay Control, Ionut Nechita.
// All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(relayStdCfgCmd)
}

var relayStdCfgCmd = &cobra.Command{
	Use:   "using-std-cfg",
	Short: "KMtronic Relay Control - Standalone with configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("\n Under development")
	},
}
