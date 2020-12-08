// +build !windows

/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package xlog
// syslog

import 	"log/syslog"

func NewSysLog(opts ...Option) *Log {
	w, err := syslog.New(syslog.LOG_DEBUG, "")
	if err != nil {
		panic(err)
	}
	return NewXLog(w, opts...)
}
