// Copyright (c) 2021, Project KMtronic Relay Control, Ionut Nechita.
// All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"log"

	"../pkg/relay"
	"github.com/spf13/cobra"
)

var relayIP, relayPORT, relayCHANNEL, relaySTATUS string

func init() {
	rootCmd.AddCommand(relayStdCmd)
	relayStdCmd.Flags().StringVar(&relayIP, "relayIP", "", "insert relay ip")
	relayStdCmd.Flags().StringVar(&relayPORT, "relayPORT", "", "insert relay port")
	relayStdCmd.Flags().StringVar(&relayCHANNEL, "relayCHANNEL", "", "insert relay channel")
	relayStdCmd.Flags().StringVar(&relaySTATUS, "relaySTATUS", "", "insert relay status")
}

var relayStdCmd = &cobra.Command{
	Use:   "using-std",
	Short: "KMtronic Relay Control - Standalone",
	Run: func(cmd *cobra.Command, args []string) {
		if relayIP == "" {
			log.Fatalln("Please insert relayIP.")
		}
		if relayPORT == "" {
			log.Fatalln("Please insert relayPORT.")
		}
		if relayCHANNEL == "" {
			log.Fatalln("Please insert relayCHANNEL.")
		}
		if relaySTATUS == "" {
			log.Fatalln("Please insert relaySTATUS.")
		}
		o := relay.NewOperation(relay.Config{
			RelayIP:      relayIP,
			RelayPORT:    relayPORT,
			RelayCHANNEL: relayCHANNEL,
			RelaySTATUS:  relaySTATUS,
		})

		ip, port, channel, status, _ := o.StartRelay()
		log.Println(ip, port, channel, status)
	},
}
