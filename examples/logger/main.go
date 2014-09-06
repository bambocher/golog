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

package main

import log "github.com/bambocher/golog"

func main() {
	test_formatter := log.GetFormatter("test_formatter")
	test_formatter.SetFormat("[{date}][{levelName}][{fileName}:{lineNo}] {message}")
	test_formatter.SetDateFormat("2006-01-02 15:04:05")

	test_log := log.GetLogger("test")
	test_log.SetLevel(log.NOTSET)

	for handler := range test_log.GetHandlers() {
		test_log.GetHandlers()[handler].SetFormatter(test_formatter)
	}

	test_log.Print("Notset message.")
	test_log.Trace("Trace message.")
	test_log.Debug("Debug message.")
	test_log.Info("Informational message.")
	test_log.Notice("Notice message.")
	test_log.Warning("Warning message.")
	test_log.Error("Error message.")
	test_log.Critical("Critical message.")
	test_log.Alert("Alert message.")
	test_log.Panic("Panic message.")

	test2_formatter := log.GetFormatter("test2_formatter")
	test2_formatter.SetFormat("[{date}][{levelNo}][{loggerName}] {message}\n")
	test2_formatter.SetDateFormat("2006-01-02")

	test2_log := log.GetLogger("test2")
	test2_log.SetLevel("NOTICE")

	for handler := range test2_log.GetHandlers() {
		test2_log.GetHandlers()[handler].SetFormatter(test2_formatter)
	}

	test2_log.Print("Notset message.")
	test2_log.Trace("Trace message.")
	test2_log.Debug("Debug message.")
	test2_log.Info("Informational message.")
	test2_log.Notice("Notice message.")
	test2_log.Warning("Warning message.")
	test2_log.Error("Error message.")
	test2_log.Critical("Critical message.")
	test2_log.Alert("Alert message.")
	test2_log.Panic("Panic message.")
}
