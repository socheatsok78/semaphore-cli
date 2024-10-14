package internals

import (
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

var Logger = log.NewLogfmtLogger(os.Stderr)

func init() {
	Logger = level.NewFilter(Logger, level.AllowAll())
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	// Logger = log.With(Logger, "caller", log.DefaultCaller)
}
