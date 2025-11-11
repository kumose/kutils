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

func TestAnySet(t *testing.T) {
	set := NewAnySet(reflect.DeepEqual)
	set.Insert(true)
	set.Insert(9527)

	require.Equal(t, true, set.slice[0])
	require.Equal(t, true, set.Slice()[0])

	require.Equal(t, 9527, set.slice[1])
	require.Equal(t, 9527, set.Slice()[1])
}
