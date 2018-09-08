package san

import (
	"runtime"
)

const (
	// Version is the san's parser version
	Version = "0.1.0"
)

// const set at build time
const (
	UTCBuildTime = "undefined"
	GitCommit    = "undefined"
	OS           = runtime.GOOS
	Arch         = runtime.GOARCH
	GoVersion    = "undefined"
)
