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

	"github.com/joomcode/errorx"
)

var (
	// ErrPropSuggestion is a property of an Error that will be printed as the suggestion.
	ErrPropSuggestion = errorx.RegisterProperty("suggestion")

	// ErrTraitPreCheck means that the Error is a pre-check error so that no error logs will be outputted directly.
	ErrTraitPreCheck = errorx.RegisterTrait("pre_check")
)

var (
	// ErrValidateChecksum is an empty HashValidationErr object, useful for type checking
	ErrValidateChecksum = &HashValidationErr{}
)

// HashValidationErr is the error indicates a failed hash validation
type HashValidationErr struct {
	cipher string
	expect string // expected hash
	actual string // input hash
}

// Error implements the error interface
func (e *HashValidationErr) Error() string {
	return fmt.Sprintf(
		"%s checksum mismatch, expect: %v, got: %v",
		e.cipher, e.expect, e.actual,
	)
}

// Unwrap implements the error interface
func (e *HashValidationErr) Unwrap() error { return nil }

// Is implements the error interface
func (e *HashValidationErr) Is(target error) bool {
	t, ok := target.(*HashValidationErr)
	if !ok {
		return false
	}

	return (e.cipher == t.cipher || t.cipher == "") &&
		(e.expect == t.expect || t.expect == "") &&
		(e.actual == t.actual || t.actual == "")
}
