package headless

import "github.com/reapertechlabs/zeno/internal/pkg/log"

var logger = log.NewFieldedLogger(&log.Fields{
	"component": "archiver.headless",
})
