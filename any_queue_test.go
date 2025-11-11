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
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnyQueue(t *testing.T) {
	q := NewAnyQueue(reflect.DeepEqual)

	q.Put(true)
	q.Put(9527)

	require.Equal(t, true, q.slice[0])
	require.Equal(t, 9527, q.slice[1])

	q.Put(true)
	q.Put(9527)

	require.Equal(t, true, q.slice[2])
	require.Equal(t, 9527, q.slice[3])

	require.Equal(t, true, q.Get(true))
	require.Equal(t, true, q.Get(true))
	require.Nil(t, q.Get(true))

	require.Equal(t, 9527, q.Get(9527))
	require.Equal(t, 9527, q.Get(9527))
	require.Nil(t, q.Get(9527))
}
