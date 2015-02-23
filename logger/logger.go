package logger

import (
	"github.com/op/go-logging"

	"os"
)

var (
	Log    = logging.MustGetLogger("prism")
	format = logging.MustStringFormatter(
		"%{time:2006/01/02 15:04:05.000000} %{shortpkg:-6.6s} %{shortfunc:-12.12s} | %{color:bold}%{level:.4s}%{color:reset} %{color}%{message}%{color:reset}",
	)
)

func init() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(logging.DEBUG, "")
	Log.SetBackend(backendLeveled)
}
