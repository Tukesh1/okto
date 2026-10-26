package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/tukesh1/okto/internal/version"
)

func run(t *testing.T, args ...string) (code int, stdout, stderr string) {
	t.Helper()

	var out, errOut bytes.Buffer
	code = Run(context.Background(), args, &out, &errOut)
	return code, out.String(), errOut.String()
}

func TestHelp(t *testing.T) {
	for _, arg := range []string{"--help", "-h", "help"} {
		t.Run(arg, func(t *testing.T) {
			code, stdout, stderr := run(t, arg)
			if code != ExitOK {
				t.Fatalf("exit code = %d, want %d", code, ExitOK)
			}
			if stderr != "" {
				t.Fatalf("stderr = %q, want empty", stderr)
			}
			for _, want := range []string{"Usage:", "compare", "--json", "--refresh"} {
				if !strings.Contains(stdout, want) {
					t.Fatalf("stdout missing %q:\n%s", want, stdout)
				}
			}
		})
	}
}

func TestVersion(t *testing.T) {
	for _, arg := range []string{"--version", "-v", "version"} {
		t.Run(arg, func(t *testing.T) {
			code, stdout, stderr := run(t, arg)
			if code != ExitOK {
				t.Fatalf("exit code = %d, want %d", code, ExitOK)
			}
			if stderr != "" {
				t.Fatalf("stderr = %q, want empty", stderr)
			}
			if !strings.Contains(stdout, "okto "+version.Version) {
				t.Fatalf("stdout = %q, want it to contain okto version", stdout)
			}
		})
	}
}

func TestNoArgumentsShowsHelp(t *testing.T) {
	code, stdout, stderr := run(t)
	if code != ExitOK {
		t.Fatalf("exit code = %d, want %d", code, ExitOK)
	}
	if stderr != "" {
		t.Fatalf("stderr = %q, want empty", stderr)
	}
	if !strings.Contains(stdout, "Usage:") {
		t.Fatalf("stdout = %q, want usage", stdout)
	}
}

func TestUnknownFlag(t *testing.T) {
	code, stdout, stderr := run(t, "--wat")
	if code != ExitUsage {
		t.Fatalf("exit code = %d, want %d", code, ExitUsage)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
	if !strings.Contains(stderr, "unknown flag") {
		t.Fatalf("stderr = %q, want unknown flag message", stderr)
	}
}

func TestCompareRequiresTwoRepositories(t *testing.T) {
	code, stdout, stderr := run(t, "compare", "owner/one")
	if code != ExitUsage {
		t.Fatalf("exit code = %d, want %d", code, ExitUsage)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
	if !strings.Contains(stderr, "exactly two repositories") {
		t.Fatalf("stderr = %q, want arity message", stderr)
	}
}

func TestOverviewIsScaffolded(t *testing.T) {
	code, stdout, stderr := run(t, "owner/repo")
	if code != ExitError {
		t.Fatalf("exit code = %d, want %d", code, ExitError)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
	if !strings.Contains(stderr, "not implemented yet") {
		t.Fatalf("stderr = %q, want scaffold message", stderr)
	}
}

func TestCompareIsScaffolded(t *testing.T) {
	code, stdout, stderr := run(t, "compare", "owner/one", "owner/two")
	if code != ExitError {
		t.Fatalf("exit code = %d, want %d", code, ExitError)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
	if !strings.Contains(stderr, "not implemented yet") {
		t.Fatalf("stderr = %q, want scaffold message", stderr)
	}
}
