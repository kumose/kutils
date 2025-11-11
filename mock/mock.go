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

package mock

import (
	"path"
	"reflect"
	"sync"

	"github.com/kumose/failpoint"
)

// Finalizer represent the function that clean a mock point
type Finalizer func()

type mockPoints struct {
	m map[string]any
	l sync.Mutex
}

func (p *mockPoints) set(fpname string, value any) {
	p.l.Lock()
	defer p.l.Unlock()

	p.m[fpname] = value
}

func (p *mockPoints) get(fpname string) any {
	p.l.Lock()
	defer p.l.Unlock()

	return p.m[fpname]
}

func (p *mockPoints) clr(fpname string) {
	p.l.Lock()
	defer p.l.Unlock()

	delete(p.m, fpname)
}

var points = mockPoints{m: make(map[string]any)}

// On inject a failpoint
func On(fpname string) any {
	var ret any
	failpoint.Inject(fpname, func() {
		ret = points.get(fpname)
	})
	return ret
}

// With enable failpoint and provide a value
func With(fpname string, value any) Finalizer {
	if err := failpoint.Enable(failpath(fpname), "return(true)"); err != nil {
		panic(err)
	}
	points.set(fpname, value)
	return func() {
		if err := Reset(fpname); err != nil {
			panic(err)
		}
	}
}

// Reset disable failpoint and remove mock value
func Reset(fpname string) error {
	if err := failpoint.Disable(failpath(fpname)); err != nil {
		return err
	}
	points.clr(fpname)
	return nil
}

func failpath(fpname string) string {
	type em struct{}
	return path.Join(reflect.TypeOf(em{}).PkgPath(), fpname)
}
