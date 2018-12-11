// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"log"
	"os"

	"github.com/caixw/gobuild"
	"github.com/issue9/term/colors"
)

// 是否显示 ignore 标记的日志
var showIgnore bool

var (
	succ   = log.New(newWriter(os.Stdout, colors.Green, "[SUCC]"), "", log.Ltime)
	info   = log.New(newWriter(os.Stdout, colors.Blue, "[INFO]"), "", log.Ltime)
	warn   = log.New(newWriter(os.Stderr, colors.Magenta, "[WARN]"), "", log.Ltime)
	erro   = log.New(newWriter(os.Stderr, colors.Red, "[ERRO]"), "", log.Ltime)
	ignore = log.New(newWriter(os.Stderr, colors.Default, "[IGNO]"), "", log.Ltime)

	logs = make(chan *gobuild.Log, 100)
)

func printLogs() {
	for log := range logs {
		switch log.Type {
		case gobuild.LogTypeError:
			erro.Println(log.Message)
		case gobuild.LogTypeIgnore:
			if !showIgnore {
				ignore.Println(log.Message)
			}
		case gobuild.LogTypeInfo:
			info.Println(log.Message)
		case gobuild.LogTypeSuccess:
			succ.Println(log.Message)
		case gobuild.LogTypeWarn:
			warn.Println(log.Message)
		default:
			panic("无效的日志类型")
		}
	}
}

// 带色彩输出的控制台。
type logWriter struct {
	out    io.Writer
	color  colors.Color
	prefix string
}

func newWriter(out io.Writer, color colors.Color, prefix string) io.Writer {
	return &logWriter{
		out:    out,
		color:  color,
		prefix: prefix,
	}
}

func (w *logWriter) Write(bs []byte) (int, error) {
	colors.Fprint(w.out, w.color, colors.Default, w.prefix)
	return colors.Fprint(w.out, colors.Default, colors.Default, string(bs))
}