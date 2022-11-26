package jbeego

import (
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/pkg/errors"
)

func InitLog() error {
	var loglevel int
	switch config.DefaultString("loglevel", "debug") {
	case "debug":
		loglevel = logs.LevelDebug
	default:
		loglevel = logs.LevelInfo
	}

	logs.SetLevel(logs.LevelDebug)

	f := &logs.PatternLogFormatter{Pattern: "%w | %F:%n | %T >> %m\n\t"}

	logs.RegisterFormatter("pattern", f)
	_ = logs.SetGlobalFormatter("pattern")

	adapterStr := fmt.Sprintf(
		"{\"filename\":\"project.log\","+
			"\"level\":%d,"+
			"\"maxlines\":0,"+
			"\"maxsize\":0,"+
			"\"daily\":true,"+
			"\"maxdays\":10,"+
			"\"color\":true}",
		loglevel)

	err := logs.SetLogger(logs.AdapterFile, adapterStr)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
