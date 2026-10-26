package model

import "time"

// RepoRef identifies a GitHub repository by owner and name.
type RepoRef struct {
	Owner string
	Name  string
}

// FullName returns the repository in owner/name form.
func (r RepoRef) FullName() string {
	if r.Owner == "" || r.Name == "" {
		return ""
	}
	return r.Owner + "/" + r.Name
}

// IsZero reports whether the reference is incomplete.
func (r RepoRef) IsZero() bool {
	return r.Owner == "" || r.Name == ""
}

// Repository is the current metadata and headline counts for a repository.
type Repository struct {
	ID                int64
	Ref               RepoRef
	FullName          string
	HTMLURL           string
	Description       string
	DefaultBranch     string
	Language          string
	LicenseSPDX       string
	Stars             int
	Forks             int
	Watchers          int
	OpenIssuesInclPRs int
	Archived          bool
	Private           bool
	CreatedAt         time.Time
	PushedAt          time.Time
	FetchedAt         time.Time
}

// Snapshot is an observation of counts collected by Okto itself.
type Snapshot struct {
	CapturedAt        time.Time
	Stars             int
	Forks             int
	Watchers          int
	OpenIssuesInclPRs int
}

// WeeklyActivity is one GitHub weekly commit-activity bucket.
type WeeklyActivity struct {
	WeekStart time.Time
	Commits   int
	Days      [7]int
}

// Release is a GitHub release.
type Release struct {
	ID          int64
	TagName     string
	Name        string
	Draft       bool
	Prerelease  bool
	CreatedAt   time.Time
	PublishedAt time.Time
	HTMLURL     string
}

// Contributor is one entry from the contributors list.
type Contributor struct {
	Login         string
	Contributions int
	Rank          int
}

// Commit is a recent commit entry.
type Commit struct {
	SHA         string
	AuthorLogin string
	AuthorName  string
	Message     string
	CommittedAt time.Time
	HTMLURL     string
}

// RateLimit is the latest rate-limit state observed from GitHub.
type RateLimit struct {
	Limit     int
	Remaining int
	ResetAt   time.Time
	Resource  string
}

// CacheState describes how much of a view came from fresh cached data.
type CacheState string

const (
	CacheMiss    CacheState = "miss"
	CacheFresh   CacheState = "fresh"
	CacheStale   CacheState = "stale"
	CachePartial CacheState = "partial"
)

// CacheInfo describes the cache status of a repository view.
type CacheInfo struct {
	State     CacheState
	FetchedAt time.Time
}

// Warning is a non-fatal data-quality or availability notice.
type Warning struct {
	Code     string
	Resource string
	Message  string
}

// Metrics contains the calculated headline metrics for a repository.
// Pointer fields distinguish unavailable data from an observed zero.
type Metrics struct {
	Stars               int
	Forks               int
	Watchers            int
	OpenIssuesInclPRs   int
	Commits52W          *int
	ReleasesLast12M     *int
	LatestRelease       *Release
	ContributorsSampled *int
}

// History contains metrics derived only from Okto's recorded snapshots.
type History struct {
	Observations    int
	FirstObservedAt *time.Time
	LastObservedAt  *time.Time
	StarsDelta      *int
	StarSeries      []Snapshot
}

// Overview is the shared view model consumed by the TUI, text, and JSON output.
type Overview struct {
	Repository    Repository
	Metrics       Metrics
	History       History
	Activity      []WeeklyActivity
	Releases      []Release
	Contributors  []Contributor
	RecentCommits []Commit
	Warnings      []Warning
	RateLimit     *RateLimit
	Cache         CacheInfo
}
