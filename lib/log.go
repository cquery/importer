package lib

import (
	stdlog "log"
	"os"

	"github.com/go-kit/kit/log"
)

var Logger log.Logger

func init() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.NewContext(Logger).With("ts", log.DefaultTimestampUTC).With("caller", log.DefaultCaller)
	stdlog.SetFlags(0) // flags are handled by Go kit's logger
	stdlog.SetOutput(log.NewStdlibAdapter(Logger))
}
