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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringSet(t *testing.T) {
	set := NewStringSet()
	vals := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for i := range vals {
		set.Insert(vals[i])
		set.Insert(vals[i])
		set.Insert(vals[i])
		set.Insert(vals[i])
		set.Insert(vals[i])
	}

	require.Equal(t, len(vals), len(set))
	for i := range vals {
		require.True(t, set.Exist(vals[i]))
	}

	require.False(t, set.Exist("11"))

	set = NewStringSet("1", "2", "3", "4", "5", "6")
	for i := 1; i < 7; i++ {
		require.True(t, set.Exist(fmt.Sprintf("%d", i)))
	}
	require.False(t, set.Exist("7"))

	s1 := NewStringSet("1", "2", "3")
	s2 := NewStringSet("4", "2", "3")
	s3 := s1.Intersection(s2)
	require.Equal(t, NewStringSet("2", "3"), s3)

	s4 := NewStringSet("4", "5", "3")
	require.Equal(t, NewStringSet("3"), s3.Intersection(s4))

	s5 := NewStringSet("4", "5")
	require.Equal(t, NewStringSet(), s3.Intersection(s5))

	s6 := NewStringSet()
	require.Equal(t, NewStringSet(), s3.Intersection(s6))
}
