package version

import (
	"strings"
	"testing"
)

func TestStringIncludesBuildMetadata(t *testing.T) {
	originalVersion, originalCommit, originalDate := Version, Commit, Date
	t.Cleanup(func() {
		Version, Commit, Date = originalVersion, originalCommit, originalDate
	})

	Version = "1.2.3"
	Commit = "abc1234"
	Date = "2026-10-26T00:00:00Z"

	got := String()
	for _, want := range []string{"okto", "1.2.3", "abc1234", "2026-10-26T00:00:00Z"} {
		if !strings.Contains(got, want) {
			t.Fatalf("String() = %q, want it to contain %q", got, want)
		}
	}
}
