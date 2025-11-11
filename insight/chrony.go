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

// Use ntpq to get basic info of chrony on the system

package insight

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// ChronyStat is holding the chrony statistics
type ChronyStat struct {
	ReferenceID    string  `json:"referenceid,omitempty"`
	Stratum        int     `json:"stratum,omitempty"`
	RefTime        string  `json:"ref_time,omitempty"`
	SystemTime     string  `json:"system_time,omitempty"`
	LastOffset     float64 `json:"last_offset,omitempty"` // millisecond
	RMSOffset      float64 `json:"rms_offset,omitempty"`  // millisecond
	Frequency      float64 `json:"frequency,omitempty"`   // millisecond
	ResidualFreq   string  `json:"residual_freq,omitempty"`
	Skew           string  `json:"skew,omitempty"`
	RootDelay      float64 `json:"root_delay,omitempty"`      // millisecond
	RootDispersion float64 `json:"root_dispersion,omitempty"` // millisecond
	UpdateInterval float64 `json:"update_interval,omitempty"` // millisecond
	LeapStatus     string  `json:"leap_status,omitempty"`
}

//revive:disable:get-return
func (cs *ChronyStat) getChronyInfo() {
	// try common locations first, then search PATH, this could cover some
	// contitions when PATH is not correctly set on calling `collector`
	var syncdBinPaths = []string{"/usr/sbin/chronyc", "/usr/bin/chronyc", "chronyc"}
	var syncd string
	var err error
	for _, syncdPath := range syncdBinPaths {
		if syncd, err = exec.LookPath(syncdPath); err == nil {
			// use the first found exec
			break
		}
		cs.LeapStatus = err.Error()
	}
	// when no `chrony` found, just return
	if syncd == "" {
		return
	}

	cmd := exec.Command(syncd, "tracking")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		cs.LeapStatus = "none"
		return
	}

	// set default sync status to none
	cs.LeapStatus = "none"

	output := strings.FieldsFunc(out.String(), multiSplit)
	for _, kv := range output {
		tmp := strings.Split(strings.TrimSpace(kv), " : ")
		switch {
		case strings.HasPrefix(tmp[0], "Reference ID"):
			cs.ReferenceID = tmp[1]
		case strings.HasPrefix(tmp[0], "Stratum"):
			cs.Stratum, err = strconv.Atoi(tmp[1])
			if err != nil {
				log.Fatal(err)
			}
		case strings.HasPrefix(tmp[0], "Ref time"):
			cs.RefTime = tmp[1]
		case strings.HasPrefix(tmp[0], "System time"):
			cs.SystemTime = tmp[1]
		case strings.HasPrefix(tmp[0], "Last offset"):
			cs.LastOffset, err = strconv.ParseFloat(strings.Split(tmp[1], " ")[0], 64)
			if err != nil {
				log.Fatal(err)
			}
			// second -> millisecond
			cs.LastOffset *= 1000
		case strings.HasPrefix(tmp[0], "RMS offset"):
			cs.RMSOffset, err = strconv.ParseFloat(strings.Split(tmp[1], " ")[0], 64)
			if err != nil {
				log.Fatal(err)
			}
			// second -> millisecond
			cs.RMSOffset *= 1000
		case strings.HasPrefix(tmp[0], "Frequency"):
			cs.Frequency, err = strconv.ParseFloat(strings.Split(tmp[1], " ")[0], 64)
			if err != nil {
				log.Fatal(err)
			}
			// second -> millisecond
			cs.Frequency *= 1000
		case strings.HasPrefix(tmp[0], "Residual freq"):
			cs.ResidualFreq = tmp[1]
		case strings.HasPrefix(tmp[0], "Skew"):
			cs.Skew = tmp[1]
		case strings.HasPrefix(tmp[0], "Root delay"):
			cs.RootDelay, err = strconv.ParseFloat(strings.Split(tmp[1], " ")[0], 64)
			if err != nil {
				log.Fatal(err)
			}
			// second -> millisecond
			cs.RootDelay *= 1000
		case strings.HasPrefix(tmp[0], "Root dispersion"):
			cs.RootDispersion, err = strconv.ParseFloat(strings.Split(tmp[1], " ")[0], 64)
			if err != nil {
				log.Fatal(err)
			}
			// second -> millisecond
			cs.RootDispersion *= 1000
		case strings.HasPrefix(tmp[0], "Update interval"):
			cs.UpdateInterval, err = strconv.ParseFloat(strings.Split(tmp[1], " ")[0], 64)
			if err != nil {
				log.Fatal(err)
			}
			cs.UpdateInterval *= 1000
		case strings.HasPrefix(tmp[0], "Leap status"):
			// none,  normal, close
			cs.LeapStatus = strings.ToLower(tmp[1])
		default:
			continue
		}
	}
}

//revive:enable:get-return
