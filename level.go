// go-logging - Logging library for Go
//
// Copyright (c) 2014 Dmitry Prazdnichnov <dp@bambucha.org>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package logging

import (
	"strings"
)

const (
	// RFC 5424 http://tools.ietf.org/html/rfc5424 defines eight severity levels
	// System is unusable. A "panic" condition usually affecting multiple
	// apps/servers/sites. At this level it would usually notify all tech staff
	// on call.
	PANIC = iota
	// Action must be taken immediately. Should be corrected immediately,
	// therefore notify staff who can fix the problem. An example would be the
	// loss of a primary ISP connection.
	ALERT
	// Critical conditions. Should be corrected immediately, but indicates
	// failure in a secondary system, an example is a loss of a backup ISP
	// connection.
	CRITICAL
	// Error conditions. Non-urgent failures, these should be relayed to
	// developers or admins; each item must be resolved within a given time.
	ERROR
	// Warning conditions. Warning messages, not an error, but indication that
	// an error will occur if action is not taken, e.g. file system 85% full -
	// each item must be resolved within a given time.
	WARNING
	// Normal but significant condition. Events that are unusual but not error
	// conditions - might be summarized in an email to developers or admins to
	// spot potential problems - no immediate action required.
	NOTICE
	// Informational messages. Normal operational messages - may be harvested
	// for reporting, measuring throughput, etc. - no action required.
	INFO
	// Debug-level messages. Info useful to developers for debugging the
	// application, not useful during operations.
	DEBUG
	TRACE
	NOTSET
)

var levelNames = []string{
	"PANIC",
	"ALERT",
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
	"TRACE",
	"NOTSET",
}

var levelNumbers = map[string]int{
	"PANIC":    PANIC,
	"ALERT":    ALERT,
	"CRITICAL": CRITICAL,
	"ERROR":    ERROR,
	"WARNING":  WARNING,
	"NOTICE":   NOTICE,
	"INFO":     INFO,
	"DEBUG":    DEBUG,
	"TRACE":    TRACE,
	"NOTSET":   NOTSET,
}

func GetLevelName(level int) string {
	return levelNames[level]
}

func GetLevelNumber(level string) int {
	return levelNumbers[strings.ToUpper(level)]
}
