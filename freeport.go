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

package kutils

import (
	"net"
	"sync"
	"time"
)

// To avoid the same port be generated twice in a short time
var portCache sync.Map

func getFreePort(host string, defaultPort int) (int, error) {
	//revive:disable
	if port, err := getPort(host, defaultPort); err == nil {
		return port, nil
	} else if port, err := getPort(host, 0); err == nil {
		return port, nil
	} else {
		return 0, err
	}
	//revive:enable
}

// MustGetFreePort asks the kernel for a free open port that is ready to use, if fail, panic
func MustGetFreePort(host string, defaultPort int, portOffset int) int {
	bestPort := defaultPort + portOffset
	if port, err := getFreePort(host, bestPort); err == nil {
		return port
	}
	panic("can't get a free port")
}

func getPort(host string, port int) (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", JoinHostPort(host, port))
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}

	port = l.Addr().(*net.TCPAddr).Port
	l.Close()

	key := JoinHostPort(host, port)
	if t, ok := portCache.Load(key); ok && t.(time.Time).Add(time.Minute).After(time.Now()) {
		return getPort(host, (port+1)%65536)
	}
	portCache.Store(key, time.Now())
	return port, nil
}
