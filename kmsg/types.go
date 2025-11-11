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

package kmsg

import (
	"fmt"
)

// Ref: https://www.kernel.org/doc/Documentation/ABI/testing/dev-kmsg
// The kmsg lines have prefix of the following format:
// | priority | sequence | monotonic timestamp | flag | message
//          6 ,      339 ,             5140900 ,    - ; NET: Registered protocol family 10
//         30 ,      340 ,             5690716 ,    - ; udevd[80]: starting version 181
// where the flag is not necessary for us, so we only parse the prefix
// for the first 3 fields: priority, sequence and timestamp

// the device to read kernel log from
const kmsgFile = "/dev/kmsg"

const severityMask = 0x07
const facilityMask = 0xf8

// Severity is part of the log priority
type Severity int

//revive:disable
const (
	// From /usr/include/sys/syslog.h.
	// These are the same on Linux, BSD, and OS X.
	LOG_EMERG Severity = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

//revive:enable

// String implements the string interface
func (p Severity) String() string {
	return []string{
		"emerg", "alert", "crit", "err",
		"warning", "notice", "info", "debug",
	}[p]
}

// Facility is part of the log priority
type Facility int

//revive:disable
const (
	// From /usr/include/sys/syslog.h.
	// These are the same up to LOG_FTP on Linux, BSD, and OS X.
	LOG_KERN Facility = iota << 3
	LOG_USER
	LOG_MAIL
	LOG_DAEMON
	LOG_AUTH
	LOG_SYSLOG
	LOG_LPR
	LOG_NEWS
	LOG_UUCP
	LOG_CRON
	LOG_AUTHPRIV
	LOG_FTP
	_ // unused
	_ // unused
	_ // unused
	_ // unused
	LOG_LOCAL0
	LOG_LOCAL1
	LOG_LOCAL2
	LOG_LOCAL3
	LOG_LOCAL4
	LOG_LOCAL5
	LOG_LOCAL6
	LOG_LOCAL7
)

//revive:enable

// String implements the string interface
func (p Facility) String() string {
	return []string{
		"kern", "user", "mail", "daemon",
		"auth", "syslog", "lpr", "news",
		"uucp", "cron", "authpriv", "ftp",
		"", "", "", "",
		"local0", "local1", "local2", "local3",
		"local4", "local5", "local6", "local7",
	}[p]
}

func decodeSeverity(p int) Severity {
	return Severity(p) & severityMask
}

func decodeFacility(p int) Facility {
	return Facility(p) & facilityMask
}

// Msg is the type of kernel message
type Msg struct {
	Severity  Severity
	Facility  Facility
	Sequence  int // Sequence is the 64 bit message sequence number
	Timestamp int // Timestamp is the monotonic timestamp in microseconds
	Message   string
}

// String implements the string interface
func (m *Msg) String() string {
	return fmt.Sprintf("%s:%s: [%.6f] %s", m.Facility, m.Severity, float64(m.Timestamp)/1e6, m.Message)
}
