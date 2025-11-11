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

import "slices"

// AnySet is a set stores any
type AnySet struct {
	eq    func(a any, b any) bool
	slice []any
}

// NewAnySet builds a AnySet
func NewAnySet(eq func(a any, b any) bool, aa ...any) *AnySet {
	slice := []any{}
out:
	for _, a := range aa {
		for _, b := range slice {
			if eq(a, b) {
				continue out
			}
		}
		slice = append(slice, a)
	}
	return &AnySet{eq, slice}
}

// Exist checks whether `val` exists in `s`.
func (s *AnySet) Exist(val any) bool {
	for _, a := range s.slice {
		if s.eq(a, val) {
			return true
		}
	}
	return false
}

// Insert inserts `val` into `s`.
func (s *AnySet) Insert(val any) {
	if !s.Exist(val) {
		s.slice = append(s.slice, val)
	}
}

// Intersection returns the intersection of two sets
func (s *AnySet) Intersection(rhs *AnySet) *AnySet {
	newSet := NewAnySet(s.eq)
	for elt := range rhs.slice {
		if s.Exist(elt) {
			newSet.Insert(elt)
		}
	}
	return newSet
}

// Remove removes `val` from `s`
func (s *AnySet) Remove(val any) {
	for i, a := range s.slice {
		if s.eq(a, val) {
			s.slice = slices.Delete(s.slice, i, i+1)
			return
		}
	}
}

// Difference returns the difference of two sets
func (s *AnySet) Difference(rhs *AnySet) *AnySet {
	newSet := NewAnySet(s.eq)
	diffSet := NewAnySet(s.eq, rhs.slice...)
	for elt := range s.slice {
		if !diffSet.Exist(elt) {
			newSet.Insert(elt)
		}
	}
	return newSet
}

// Slice converts the set to a slice
func (s *AnySet) Slice() []any {
	return slices.Clone(s.slice)
}
