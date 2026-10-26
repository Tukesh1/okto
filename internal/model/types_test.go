package model

import "testing"

func TestRepoRef(t *testing.T) {
	tests := []struct {
		name     string
		ref      RepoRef
		fullName string
		zero     bool
	}{
		{
			name:     "complete",
			ref:      RepoRef{Owner: "tukesh1", Name: "okto"},
			fullName: "tukesh1/okto",
		},
		{
			name: "missing owner",
			ref:  RepoRef{Name: "okto"},
			zero: true,
		},
		{
			name: "missing name",
			ref:  RepoRef{Owner: "tukesh1"},
			zero: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ref.FullName(); got != tt.fullName {
				t.Fatalf("FullName() = %q, want %q", got, tt.fullName)
			}
			if got := tt.ref.IsZero(); got != tt.zero {
				t.Fatalf("IsZero() = %v, want %v", got, tt.zero)
			}
		})
	}
}
