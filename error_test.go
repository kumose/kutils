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
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashValidationErr(t *testing.T) {
	err0 := &HashValidationErr{
		cipher: "sha256",
		expect: "hash111",
		actual: "hash222",
	}
	// identical errors are equal
	require.True(t, errors.Is(err0, err0))
	require.True(t, errors.Is(ErrValidateChecksum, ErrValidateChecksum))
	require.True(t, errors.Is(ErrValidateChecksum, &HashValidationErr{}))
	require.True(t, errors.Is(&HashValidationErr{}, ErrValidateChecksum))
	// not equal for different error types
	require.False(t, errors.Is(err0, errors.New("")))
	// default Value matches any error
	require.True(t, errors.Is(err0, ErrValidateChecksum))
	// error with values are not matching default ones
	require.False(t, errors.Is(ErrValidateChecksum, err0))

	err1 := &HashValidationErr{
		cipher: "sha256",
		expect: "hash111",
		actual: "hash222",
	}
	require.True(t, errors.Is(err1, ErrValidateChecksum))
	// errors with same values are equal
	require.True(t, errors.Is(err0, err1))
	require.True(t, errors.Is(err1, err0))
	// errors with different ciphers are not equal
	err1.cipher = "sha512"
	require.False(t, errors.Is(err0, err1))
	require.False(t, errors.Is(err1, err0))
	// errors with different expected hashes are not equal
	err1.cipher = err0.cipher
	require.True(t, errors.Is(err0, err1))
	err1.expect = "hash1112"
	require.False(t, errors.Is(err0, err1))
	require.False(t, errors.Is(err1, err0))
	// errors with different actual hashes are not equal
	err1.expect = err0.expect
	require.True(t, errors.Is(err0, err1))
	err1.actual = "hash2223"
	require.False(t, errors.Is(err0, err1))
	require.False(t, errors.Is(err1, err0))
}
