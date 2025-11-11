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
	"fmt"
	"strings"
)

const (
	space = "0123456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
	base  = len(space)
)

// Encode returns a string by encoding the id over a 51 characters space
func Encode(id int64) string {
	var short []byte
	for id > 0 {
		i := id % int64(base)
		short = append(short, space[i])
		id /= int64(base)
	}
	for i, j := 0, len(short)-1; i < j; i, j = i+1, j-1 {
		short[i], short[j] = short[j], short[i]
	}
	return string(short)
}

// Decode will decode the string and return the id
// The input string should be a valid one with only characters in the space
func Decode(encoded string) (int64, error) {
	if len(encoded) != len([]rune(encoded)) {
		return 0, fmt.Errorf("invalid encoded string: '%s'", encoded)
	}
	var id int64
	for i := 0; i < len(encoded); i++ {
		idx := strings.IndexByte(space, encoded[i])
		if idx < 0 {
			return 0, fmt.Errorf("invalid encoded string: '%s' contains invalid character", encoded)
		}
		id = id*int64(base) + int64(idx)
	}
	return id, nil
}
