// Package cli parses Okto's command-line arguments and dispatches commands.
package cli

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/tukesh1/okto/internal/version"
)

// Exit codes are part of Okto's public CLI contract.
const (
	ExitOK          = 0
	ExitError       = 1
	ExitUsage       = 2
	ExitAuth        = 3
	ExitNotFound    = 4
	ExitRateLimited = 5
	ExitNetwork     = 6
)

const usage = `Okto - GitHub intelligence in your terminal

Usage:
  okto <owner/repo> [flags]
  okto compare <owner/repo> <owner/repo> [flags]
  okto --help
  okto --version

Flags:
  --json        Emit machine-readable JSON on stdout
  --refresh     Bypass cache freshness checks
  --no-color    Disable ANSI colors
  -h, --help    Show this help
  -v, --version Show version

This is the Phase 0 scaffold. Overview and comparison are not implemented yet.
`

// Run executes the CLI and returns a process exit code.
func Run(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stdout, usage)
		return ExitOK
	}

	switch args[0] {
	case "-h", "--help", "help":
		fmt.Fprint(stdout, usage)
		return ExitOK
	case "-v", "--version", "version":
		fmt.Fprintln(stdout, version.String())
		return ExitOK
	case "compare":
		return runCompare(ctx, args[1:], stdout, stderr)
	default:
		if strings.HasPrefix(args[0], "-") {
			fmt.Fprintf(stderr, "okto: unknown flag %q\n\n", args[0])
			fmt.Fprint(stderr, usage)
			return ExitUsage
		}
		return runOverview(ctx, args, stdout, stderr)
	}
}

func runOverview(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	fmt.Fprintln(stderr, "okto: repository overview is not implemented yet (Phase 0 scaffold)")
	return ExitError
}

func runCompare(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) != 2 {
		fmt.Fprintln(stderr, "okto: compare requires exactly two repositories")
		return ExitUsage
	}
	fmt.Fprintln(stderr, "okto: repository comparison is not implemented yet (Phase 0 scaffold)")
	return ExitError
}
