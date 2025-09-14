//go:build windows

package watchers

import (
	"fmt"

	"golang.org/x/sys/windows"

	"github.com/reapertechlabs/zeno/internal/pkg/config"
)

func CheckDiskUsage(path string) error {
	var freeBytesAvailable uint64
	var totalNumberOfBytes uint64
	var totalNumberOfFreeBytes uint64
	if err := windows.GetDiskFreeSpaceEx(windows.StringToUTF16Ptr(path), &freeBytesAvailable, &totalNumberOfBytes, &totalNumberOfFreeBytes); err != nil {
		panic(fmt.Sprintf("Error retrieving disk stats: %v\n", err))
	}
	return checkThreshold(totalNumberOfBytes, freeBytesAvailable, config.Get().MinSpaceRequired)
}
