// Package version holds build metadata injected at link time.
package version

import "fmt"

var (
	// Version is the release version. It is replaced by the build.
	Version = "dev"
	// Commit is the source revision. It is replaced by the build.
	Commit = "none"
	// Date is the UTC build timestamp. It is replaced by the build.
	Date = "unknown"
)

// String returns a human-readable version string.
func String() string {
	return fmt.Sprintf("okto %s (commit %s, built %s)", Version, Commit, Date)
}
