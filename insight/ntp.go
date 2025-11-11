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
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// TimeStat is holding the NTP time statistics
type TimeStat struct {
	Ver       string  `json:"version,omitempty"`
	Sync      string  `json:"sync,omitempty"`
	Stratum   int     `json:"stratum,omitempty"`
	Precision int     `json:"precision,omitempty"`
	Rootdelay float64 `json:"rootdelay,omitempty"`
	Rootdisp  float64 `json:"rootdisp,omitempty"`
	Refid     string  `json:"refid,omitempty"`
	Peer      int     `json:"peer,omitempty"`
	TC        int     `json:"tc,omitempty"`
	Mintc     int     `json:"mintc,omitempty"`
	Offset    float64 `json:"offset,omitempty"`
	Frequency float64 `json:"frequency,omitempty"`
	Jitter    float64 `json:"jitter,omitempty"`
	ClkJitter float64 `json:"clk_jitter,omitempty"`
	ClkWander float64 `json:"clk_wander,omitempty"`
	Status    string  `json:"status,omitempty"`
}

//revive:disable:get-return
func (ts *TimeStat) getNTPInfo() {
	// try common locations first, then search PATH, this could cover some
	// contitions when PATH is not correctly set on calling `collector`
	var syncdBinPaths = []string{"/usr/sbin/ntpq", "/usr/bin/ntpq", "ntpq"}
	var syncd string
	var err error
	for _, syncdPath := range syncdBinPaths {
		if syncd, err = exec.LookPath(syncdPath); err == nil {
			// use the first found exec
			break
		}
		ts.Ver = err.Error()
	}
	// when no `ntpq` found, just return
	if syncd == "" {
		return
	}

	cmd := exec.Command(syncd, "-c rv", "127.0.0.1")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// set default sync status to none
	ts.Sync = "none"

	output := strings.FieldsFunc(out.String(), multiSplit)
	for _, kv := range output {
		tmp := strings.Split(strings.TrimSpace(kv), "=")
		switch {
		case tmp[0] == "version":
			ts.Ver = strings.Trim(tmp[1], "\"")
		case tmp[0] == "stratum":
			ts.Stratum, err = strconv.Atoi(tmp[1])
			if err != nil {
				log.Fatal(err)
			}
		case tmp[0] == "precision":
			ts.Precision, err = strconv.Atoi(tmp[1])
			if err != nil {
				log.Fatal(err)
			}
		case tmp[0] == "rootdelay":
			ts.Rootdelay, err = strconv.ParseFloat(tmp[1], 64)
			if err != nil {
				log.Fatal(err)
			}
		case tmp[0] == "rootdisp":
			ts.Rootdisp, err = strconv.ParseFloat(tmp[1], 64)
			if err != nil {
				log.Fatal(err)
			}
		case tmp[0] == "refid":
			ts.Refid = tmp[1]
		case tmp[0] == "peer":
			ts.Peer, err = strconv.Atoi(tmp[1])
			if err != nil {
				log.Fatal(err)
			}
		case tmp[0] == "tc":
			ts.TC, err = strconv.Atoi(tmp[1])
			if err != nil {
				log.Fatal(err)
			}
		case tmp[0] == "mintc":
			ts.Mintc, err = strconv.Atoi(tmp[1])
			if err != nil {
				log.Fatal(err)
			}
		case tmp[0] == "offset":
			ts.Offset, err = strconv.ParseFloat(tmp[1], 64)
			if err != nil {
				log.Fatal(err)
			}
		case tmp[0] == "frequency":
			ts.Frequency, err = strconv.ParseFloat(tmp[1], 64)
			if err != nil {
				log.Fatal(err)
			}
		case tmp[0] == "sys_jitter":
			ts.Jitter, err = strconv.ParseFloat(tmp[1], 64)
			if err != nil {
				log.Fatal(err)
			}
		case tmp[0] == "clk_jitter":
			ts.ClkJitter, err = strconv.ParseFloat(tmp[1], 64)
			if err != nil {
				log.Fatal(err)
			}
		case tmp[0] == "clk_wander":
			ts.ClkWander, err = strconv.ParseFloat(tmp[1], 64)
			if err != nil {
				log.Fatal(err)
			}
		case strings.Contains(tmp[0], "sync"):
			ts.Sync = tmp[0]
		case len(tmp) > 2 && strings.Contains(tmp[1], "status"):
			// sample line of tmp: ["associd", "0 status", "0618 leap_none"]
			ts.Status = strings.Split(tmp[2], " ")[0]
		default:
			continue
		}
	}
}

//revive:enable:get-return

func multiSplit(r rune) bool {
	switch r {
	case ',':
		return true
	case '\n':
		return true
	default:
		return false
	}
}
