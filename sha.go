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
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"strings"

	"github.com/kumose/errors"
)

// CheckSHA256 returns an error if the hash of reader mismatches `sha`
func CheckSHA256(reader io.Reader, sha string) error {
	shaWriter := sha256.New()
	if _, err := io.Copy(shaWriter, reader); err != nil {
		return errors.Trace(err)
	}

	checksum := hex.EncodeToString(shaWriter.Sum(nil))
	if checksum != strings.TrimSpace(sha) {
		return &HashValidationErr{
			cipher: "sha256",
			expect: sha,
			actual: checksum,
		}
	}
	return nil
}

// SHA256 returns the hash of reader
func SHA256(reader io.Reader) (string, error) {
	shaWriter := sha256.New()
	if _, err := io.Copy(shaWriter, reader); err != nil {
		return "", errors.Trace(err)
	}

	checksum := hex.EncodeToString(shaWriter.Sum(nil))
	return checksum, nil
}

// SHA512 returns the hash of reader
func SHA512(reader io.Reader) (string, error) {
	shaWriter := sha512.New()
	if _, err := io.Copy(shaWriter, reader); err != nil {
		return "", errors.Trace(err)
	}

	checksum := hex.EncodeToString(shaWriter.Sum(nil))
	return checksum, nil
}
