package config

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func env(values map[string]string) func(string) string {
	return func(key string) string {
		return values[key]
	}
}

func TestResolveAuthEnvironmentPrecedence(t *testing.T) {
	auth, err := resolveAuth(context.Background(), env(map[string]string{
		"GITHUB_TOKEN": "github-token",
		"GH_TOKEN":     "gh-token",
	}), func(context.Context) (string, error) {
		t.Fatal("gh fallback should not be called")
		return "", nil
	})
	if err != nil {
		t.Fatalf("resolveAuth() error = %v", err)
	}
	if auth.Token != "github-token" || auth.Source != AuthSourceGitHubToken {
		t.Fatalf("auth = %+v, want GITHUB_TOKEN source", auth)
	}
}

func TestResolveAuthGHTokenFallback(t *testing.T) {
	auth, err := resolveAuth(context.Background(), env(map[string]string{
		"GH_TOKEN": "gh-token",
	}), nil)
	if err != nil {
		t.Fatalf("resolveAuth() error = %v", err)
	}
	if auth.Token != "gh-token" || auth.Source != AuthSourceGHToken {
		t.Fatalf("auth = %+v, want GH_TOKEN source", auth)
	}
}

func TestResolveAuthGHCLI(t *testing.T) {
	auth, err := resolveAuth(context.Background(), env(nil), func(context.Context) (string, error) {
		return "  cli-token\n", nil
	})
	if err != nil {
		t.Fatalf("resolveAuth() error = %v", err)
	}
	if auth.Token != "cli-token" || auth.Source != AuthSourceGHCLI {
		t.Fatalf("auth = %+v, want gh source", auth)
	}
}

func TestResolveAuthNone(t *testing.T) {
	auth, err := resolveAuth(context.Background(), env(map[string]string{
		"GITHUB_TOKEN": "   ",
	}), func(context.Context) (string, error) {
		return "", errors.New("not logged in")
	})
	if err != nil {
		t.Fatalf("resolveAuth() error = %v", err)
	}
	if auth.Authenticated() || auth.Source != AuthSourceNone {
		t.Fatalf("auth = %+v, want unauthenticated", auth)
	}
}

func TestResolveAuthContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := resolveAuth(ctx, env(nil), func(context.Context) (string, error) {
		return "", context.Canceled
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("resolveAuth() error = %v, want context.Canceled", err)
	}
}

func TestAuthStringRedactsToken(t *testing.T) {
	auth := Auth{Token: "super-secret", Source: AuthSourceGitHubToken}
	if got := auth.String(); strings.Contains(got, auth.Token) {
		t.Fatalf("Auth.String() leaked token: %q", got)
	}
}
