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

// StringSet is a string set.
type StringSet map[string]struct{}

// NewStringSet builds a string set.
func NewStringSet(ss ...string) StringSet {
	set := make(StringSet)
	for _, s := range ss {
		set.Insert(s)
	}
	return set
}

// Exist checks whether `val` exists in `s`.
func (s StringSet) Exist(val string) bool {
	_, ok := s[val]
	return ok
}

// Insert inserts `val` into `s`.
func (s StringSet) Insert(val string) {
	s[val] = struct{}{}
}

// Join add all elements of `add` to `s`.
func (s StringSet) Join(add StringSet) StringSet {
	for elt := range add {
		s.Insert(elt)
	}
	return s
}

// Intersection returns the intersection of two sets
func (s StringSet) Intersection(rhs StringSet) StringSet {
	newSet := NewStringSet()
	for elt := range s {
		if rhs.Exist(elt) {
			newSet.Insert(elt)
		}
	}
	return newSet
}

// Remove removes `val` from `s`
func (s StringSet) Remove(val string) {
	delete(s, val)
}

// Difference returns the difference of two sets
func (s StringSet) Difference(rhs StringSet) StringSet {
	newSet := NewStringSet()
	for elt := range s {
		if !rhs.Exist(elt) {
			newSet.Insert(elt)
		}
	}
	return newSet
}

// Slice converts the set to a slice
func (s StringSet) Slice() []string {
	res := make([]string, 0)
	for val := range s {
		res = append(res, val)
	}
	return res
}
