package extractor

import (
	"os"
	"testing"

	"go.uber.org/goleak"

	"github.com/reapertechlabs/zeno/internal/pkg/config"
)

func TestMain(m *testing.M) {
	config.InitConfig()
	goleak.VerifyTestMain(m)
	os.Exit(m.Run())
}
