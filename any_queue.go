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

// AnyQueue is a queue stores any
type AnyQueue struct {
	eq    func(a any, b any) bool
	slice []any
}

// NewAnyQueue builds a AnyQueue
func NewAnyQueue(eq func(a any, b any) bool, aa ...any) *AnyQueue {
	return &AnyQueue{eq, aa}
}

// Get returns previous stored value that equals to val and remove it from the queue, if not found, return nil
func (q *AnyQueue) Get(val any) any {
	for i, a := range q.slice {
		if q.eq(a, val) {
			q.slice = slices.Delete(q.slice, i, i+1)
			return a
		}
	}
	return nil
}

// Put inserts `val` into `q`.
func (q *AnyQueue) Put(val any) {
	q.slice = append(q.slice, val)
}
