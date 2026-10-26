# Okto contracts

This document describes the planned MVP contract. It is intentionally written
before the feature implementation so the CLI, JSON output, and tests share one
definition.

**Phase 0 status:** only `--help` and `--version` are implemented. The
repository commands are stubs.

## CLI

```text
okto <owner/repo> [flags]
okto compare <owner/repo> <owner/repo> [flags]
okto --help
okto --version
```

### Flags

| Flag | Behavior |
|---|---|
| `--json` | Write only schema-v1 JSON to stdout. Diagnostics go to stderr. |
| `--refresh` | Bypass cache freshness checks. Conditional requests may still return `304`. |
| `--no-color` | Disable ANSI colors and styling. |
| `-h`, `--help` | Print help and exit `0`. |
| `-v`, `--version` | Print version metadata and exit `0`. |

Flags may appear before or after positional arguments once the full parser is
implemented.

### Default output

- stdout is a TTY: interactive Bubble Tea dashboard.
- stdout is not a TTY: plain-text summary.
- `--json`: JSON only, regardless of TTY status.

### Comparison scope

The MVP accepts exactly two repositories. More than two is a usage error and is
deferred.

## Exit codes

| Code | Meaning |
|---:|---|
| `0` | Usable output was produced, including cached output with warnings. |
| `1` | Unexpected internal error. |
| `2` | Invalid usage, flags, or repository reference. |
| `3` | Authentication failed or the token is invalid. |
| `4` | Repository not found or not accessible. |
| `5` | GitHub rate limit reached and no usable cached data exists. |
| `6` | Network failure and no usable cached data exists. |

Partial data with cached fallback exits `0` and reports warnings in the output.

## JSON schema v1

JSON is UTF-8, uses `snake_case`, and contains no ANSI or progress output.
`null` means unavailable or unobserved; `0` means an observed zero.

```json
{
  "schema_version": 1,
  "command": "overview",
  "generated_at": "2026-10-26T12:00:00Z",
  "repositories": [
    {
      "owner": "tukesh1",
      "name": "okto",
      "full_name": "tukesh1/okto",
      "html_url": "https://github.com/tukesh1/okto",
      "description": null,
      "default_branch": "main",
      "language": "Go",
      "license_spdx": null,
      "archived": false,
      "private": false,
      "metrics": {
        "stars": 0,
        "forks": 0,
        "watchers": 0,
        "open_issues_incl_prs": 0,
        "commits_52w": null,
        "releases_last_12m": null,
        "latest_release": null,
        "contributors_sampled": null
      },
      "history": {
        "observations": 0,
        "first_observed_at": null,
        "last_observed_at": null,
        "stars_delta": null,
        "star_series": []
      },
      "activity": {
        "weekly_commits": [],
        "recent_commits": []
      },
      "releases": [],
      "contributors": [],
      "warnings": [],
      "cache": {
        "state": "miss",
        "fetched_at": null
      }
    }
  ],
  "comparison": null,
  "warnings": [],
  "rate_limit": {
    "limit": null,
    "remaining": null,
    "reset_at": null,
    "resource": null
  }
}
```

### Fields

- `command`: `overview` or `compare`.
- `generated_at`: UTC RFC 3339 timestamp.
- `repositories`: one entry for overview, two for comparison.
- `comparison`: `null` for overview; a comparison object for `compare`.
- `warnings`: repository-level and command-level non-fatal notices.
- `cache.state`: `miss`, `fresh`, `stale`, or `partial`.
- `rate_limit`: latest observed GitHub rate-limit state; fields may be `null`.

### Warning codes

Planned warning codes include:

| Code | Meaning |
|---|---|
| `unauthenticated` | No token was found; unauthenticated limits apply. |
| `rate_limit_low` | Optional requests were skipped to conserve quota. |
| `commit_activity_pending` | GitHub returned `202`; statistics were unavailable. |
| `releases_truncated` | The release list hit the MVP fetch cap. |
| `contributors_truncated` | The contributor list hit the MVP fetch cap. |
| `snapshot_history_insufficient` | Fewer than two observations exist for a trend. |
| `optional_resource_failed` | A non-required resource could not be refreshed. |
| `cache_stale` | Cached data was returned after a refresh failure. |

## Metric definitions

| Metric | Source and window |
|---|---|
| Stars | `stargazers_count` |
| Forks | `forks_count` |
| Watchers | `subscribers_count` |
| Open issues + PRs | `open_issues_count`; combined and labeled as such |
| Commits, 52 weeks | Sum of `stats/commit_activity` weekly totals; default branch, excludes merge commits |
| Weekly commits | GitHub weekly buckets; the current week may be partial |
| Releases, trailing 12 months | Published, non-draft, non-prerelease releases in the window |
| Latest release | Most recent published stable release |
| Contributors | Top fetched contributors; not a total count |
| Star/fork trend | Delta between the first and last Okto snapshot; requires at least two observations |

## Data-quality rules

- Never fabricate historical values.
- Never treat unavailable data as zero.
- Label sampled or truncated data.
- Compare equivalent time windows.
- Do not describe every count increase as growth.
- Keep credentials out of the database, logs, errors, fixtures, and JSON.
