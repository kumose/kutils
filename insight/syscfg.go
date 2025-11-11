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

package insight

import (
	"os"
	"strconv"
	"strings"

	sysctl "github.com/lorenzosaino/go-sysctl"
)

// SysCfg are extra system configs we collected
type SysCfg struct {
	SecLimit []SecLimitField   `json:"sec_limit,omitempty"`
	SysCtl   map[string]string `json:"sysctl,omitempty"`
}

// SecLimitField is the config field in security limit file
type SecLimitField struct {
	Domain string `json:"domain"`
	Type   string `json:"type"`
	Item   string `json:"item"`
	Value  int    `json:"value"`
}

//revive:disable:get-return
func (c *SysCfg) getSysCfg() {
	c.SysCtl = collectSysctl()
	c.SecLimit = collectSecLimit()
}

//revive:enable:get-return

func collectSysctl() map[string]string {
	msg, err := sysctl.GetAll()
	if err != nil {
		return nil
	}
	return msg
}

const limitFilePath = "/etc/security/limits.conf"

func collectSecLimit() []SecLimitField {
	result := make([]SecLimitField, 0)

	data, err := os.ReadFile(limitFilePath)
	if err != nil {
		return result
	}

	for line := range strings.SplitSeq(string(data), "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "#") {
			fields := strings.Fields(line)
			if len(fields) < 4 {
				continue
			}
			var field SecLimitField
			field.Domain = fields[0]
			field.Type = fields[1]
			field.Item = fields[2]
			v, err := strconv.Atoi(fields[3])
			if err != nil {
				continue
			}
			field.Value = v
			result = append(result, field)
		}
	}
	return result
}
