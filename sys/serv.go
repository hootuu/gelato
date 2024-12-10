package sys

import (
	"github.com/hootuu/gelato/configure"
	"github.com/hootuu/gelato/errors"
	"github.com/rs/xid"
	"os"
	"strings"
)

var ServerID string
var RunMode Mode
var WorkingDirectory string

func Exit(err *errors.Error) {
	if err != nil {
		Error("Crash error: ", err.Error())
	}
	os.Exit(0)
}

func init() {
	ServerID = strings.ToUpper(xid.New().String())
	RunMode = ModeValueOf(configure.GetString(ConfigSysRunMode))
	wd, nErr := os.Getwd()
	if nErr != nil {
		Error("Get Current Working Directory Failed: ", nErr.Error())
		Exit(errors.System("Get Current Working Directory Failed!"))
		return
	}
	WorkingDirectory = wd
	Warn("# Server ID: ", ServerID)
	Warn("# Run Mode: ", strings.ToUpper(string(RunMode)))
	Warn("# Working Directory: ", string(WorkingDirectory))
}
