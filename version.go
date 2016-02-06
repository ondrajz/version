package version

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var (
	Number      = "<undefined>"
	CommitHash  = "<undefined>"
	CommitStamp = "<undefined>"
	BuildUser   = "<undefined>"
	BuildHost   = "<undefined>"
	BuildStamp  = "<undefined>"

	Long string
)

func init() {
	stamp, _ := strconv.Atoi(CommitStamp)
	commitDate := time.Unix(int64(stamp), 0).UTC().Format("2006/01/02-15:04")
	stamp, _ = strconv.Atoi(BuildStamp)
	buildDate := time.Unix(int64(stamp), 0).UTC().Format("Mon Jan 02 15:04:05 MST 2006")
	Long = fmt.Sprintf(`%s-%s (%s) built on %s@%s (%s %s-%s) at %s`,
		Number, CommitHash, commitDate,
		BuildUser, BuildHost,
		runtime.Version(), runtime.GOOS, runtime.GOARCH,
		buildDate,
	)
}
