// Copyright (C) Kumo inc. and its affiliates.
// Author: Jeff.li lijippy@163.com
// All rights reserved.
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

//go:build linux
// +build linux

package insight

import (
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

// GetIPV4Sockets is getting sockets from states
func GetIPV4Sockets(states ...uint8) ([]Socket, error) {
	netSockets, err := netlink.SocketDiagTCP(unix.AF_INET)
	if err != nil {
		return nil, err
	}

	tcpStates := make(map[uint8]bool, len(states))
	for _, state := range states {
		tcpStates[state] = true
	}
	sockets := make([]Socket, 0, len(netSockets))
	for _, socket := range netSockets {
		if len(tcpStates) > 0 && !tcpStates[socket.State] {
			continue
		}
		sockets = append(sockets, Socket{
			Family:     socket.Family,
			State:      socket.State,
			SourceAddr: socket.ID.Source.String(),
			SourcePort: socket.ID.SourcePort,
			DestAddr:   socket.ID.Destination.String(),
			DestPort:   socket.ID.DestinationPort,
		})
	}

	return sockets, nil
}
