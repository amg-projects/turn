// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package server

import (
	"github.com/pion/stun/v3"
	"github.com/amg-projects/turn/v4/internal/ipnet"
)

func handleBindingRequest(r Request, m *stun.Message) error {
	r.Log.Debugf("(STUN) Received BindingRequest from %s", r.SrcAddr)

	ip, port, err := ipnet.AddrIPPort(r.SrcAddr)
	if err != nil {
		r.Log.Errorf("Failed to extract IP and Port from %s: %v", r.SrcAddr, err)
		return err
	}

	attrs := buildMsg(m.TransactionID, stun.BindingSuccess, &stun.XORMappedAddress{
		IP:   ip,
		Port: port,
	}, stun.Fingerprint)

	r.Log.Debugf("(STUN) Responding to BindingRequest from %s", r.SrcAddr)
	return buildAndSend(r.Conn, r.SrcAddr, attrs...)
}
