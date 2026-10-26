package github

import (
	"context"
	"time"
)

// CacheMeta is the conditional-request metadata stored for a URL.
type CacheMeta struct {
	ETag         string
	LastModified string
	FetchedAt    time.Time
}

// MetaStore persists conditional-request metadata. The SQLite store will
// implement this interface in a later phase.
type MetaStore interface {
	GetMeta(ctx context.Context, key string) (CacheMeta, bool, error)
	PutMeta(ctx context.Context, key string, meta CacheMeta) error
}
