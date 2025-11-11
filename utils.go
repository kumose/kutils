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
	"math"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

// JoinInt joins a slice of int to string
func JoinInt(nums []int, delim string) string {
	result := ""
	for _, i := range nums {
		result += strconv.Itoa(i)
		result += delim
	}
	return strings.TrimSuffix(result, delim)
}

// IsFlagSetByUser check if the a flag is set by user explicitly
func IsFlagSetByUser(flagSet *pflag.FlagSet, flagName string) bool {
	setByUser := false
	flagSet.Visit(func(f *pflag.Flag) {
		if f.Name == flagName {
			setByUser = true
		}
	})
	return setByUser
}

// MustAtoI calls strconv.Atoi and ignores error
func MustAtoI(a string) int {
	v, _ := strconv.Atoi(a)
	return v
}

// Base62Tag returns a tag based on time
func Base62Tag() string {
	const base = 62
	const sets = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 0)
	num := time.Now().UnixNano() / int64(time.Millisecond)
	for num > 0 {
		r := math.Mod(float64(num), float64(base))
		num /= base
		b = append([]byte{sets[int(r)]}, b...)
	}
	return string(b)
}

// Ternary operator
func Ternary(condition bool, a, b any) any {
	if condition {
		return a
	}
	return b
}

// JoinHostPort return host and port
func JoinHostPort(host string, port int) string {
	return net.JoinHostPort(host, strconv.Itoa(port))
}

// ParseHostPort Prase host and port
func ParseHostPort(hostport string) (host, port string) {
	colon := strings.LastIndex(hostport, ":")

	host = strings.TrimSuffix(strings.TrimPrefix(hostport[:colon], "["), "]")
	port = hostport[colon+1:]
	return
}
