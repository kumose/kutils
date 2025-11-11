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

package insight

import "github.com/vishvananda/netlink"

// Socket is a network socket
type Socket struct {
	Family     uint8  `json:"family"`
	State      uint8  `json:"state"`
	SourceAddr string `json:"source_addr"`
	SourcePort uint16 `json:"source_port"`
	DestAddr   string `json:"dest_addr"`
	DestPort   uint16 `json:"dest_port"`
}

func (info *Info) collectSockets() error {
	sockets, err := GetIPV4Sockets(netlink.TCP_ESTABLISHED)
	info.Sockets = sockets
	return err
}
