// Copyright (c) 2021, Project KMtronic Relay Control, Ionut Nechita.
// All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(relayAPICmd)
}

var relayAPICmd = &cobra.Command{
	Use:   "using-api",
	Short: "KMtronic Relay Control - API",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("\n Under development")
	},
}
