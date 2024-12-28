package version

import (
	"fmt"
)

const (
	VersionFirst  = 1
	VersionSecond = 0
	VersionThird  = 0
)

func VersionString() string {
	return fmt.Sprintf("%d.%d.%d", VersionFirst, VersionSecond, VersionThird)
}
