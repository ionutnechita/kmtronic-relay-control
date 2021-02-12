// Copyright (c) 2021, Project KMtronic Relay Control, Ionut Nechita.
// All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package relay

// Operation as struct
type Operation struct {
	cfg Config
}

// NewOperation as func, return *Operation
func NewOperation(cfg Config) *Operation {
	return &Operation{
		cfg: cfg,
	}
}

// StartRelay as func, return nil
func (o *Operation) StartRelay() (string, string, string, string, error) {
	return o.cfg.RelayIP, o.cfg.RelayPORT, o.cfg.RelayCHANNEL, o.cfg.RelaySTATUS, nil
}
