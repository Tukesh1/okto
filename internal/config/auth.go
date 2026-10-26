package config

import (
	"context"
	"os"
	"os/exec"
	"strings"
)

// AuthSource identifies where a GitHub token came from.
type AuthSource string

const (
	AuthSourceNone        AuthSource = "none"
	AuthSourceGitHubToken AuthSource = "GITHUB_TOKEN"
	AuthSourceGHToken     AuthSource = "GH_TOKEN"
	AuthSourceGHCLI       AuthSource = "gh"
)

// Auth is the resolved GitHub credential. The token is never rendered by String.
type Auth struct {
	Token  string
	Source AuthSource
}

// Authenticated reports whether a token was found.
func (a Auth) Authenticated() bool {
	return a.Token != ""
}

// String returns a safe description of the credential source.
func (a Auth) String() string {
	return "auth(source=" + string(a.Source) + ")"
}

type tokenResolver func(context.Context) (string, error)

// ResolveAuth finds a token using environment variables first, then the GitHub CLI.
// Missing credentials are not an error; Okto supports unauthenticated requests.
func ResolveAuth(ctx context.Context) (Auth, error) {
	return resolveAuth(ctx, os.Getenv, ghAuthToken)
}

func resolveAuth(ctx context.Context, getenv func(string) string, gh tokenResolver) (Auth, error) {
	if token := strings.TrimSpace(getenv("GITHUB_TOKEN")); token != "" {
		return Auth{Token: token, Source: AuthSourceGitHubToken}, nil
	}
	if token := strings.TrimSpace(getenv("GH_TOKEN")); token != "" {
		return Auth{Token: token, Source: AuthSourceGHToken}, nil
	}

	if gh == nil {
		return Auth{Source: AuthSourceNone}, nil
	}

	token, err := gh(ctx)
	if err != nil {
		if ctx.Err() != nil {
			return Auth{}, ctx.Err()
		}
		return Auth{Source: AuthSourceNone}, nil
	}

	token = strings.TrimSpace(token)
	if token == "" {
		return Auth{Source: AuthSourceNone}, nil
	}
	return Auth{Token: token, Source: AuthSourceGHCLI}, nil
}

func ghAuthToken(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "gh", "auth", "token", "--hostname", "github.com")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
