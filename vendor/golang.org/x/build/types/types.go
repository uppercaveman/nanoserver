// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package types contains common types used by the Go continuous build
// system.
package types

// BuildStatus is the data structure that's marshalled as JSON
// for the https://build.golang.org/?mode=json page.
type BuildStatus struct {
	// Builders is a list of all known builders.
	// The order that builders appear is the same order as the build results for a revision.
	Builders []string `json:"builders"`

	// Revisions are the revisions shown on the front page of build.golang.org,
	// in the same order. It starts with the "go" repo, from recent to old, and then
	// it has 1 each of the subrepos, with only their most recent commit.
	Revisions []BuildRevision `json:"revisions"`
}

// BuildRevision is the status of a commit across all builders.
// It corresponds to a single row of https://build.golang.org/
type BuildRevision struct {
	// Repo is "go" for the main repo, else  "tools", "crypto", "net", etc.
	// These are repos as listed at https://go.googlesource.com/
	Repo string `json:"repo"`

	// Revision is the full git hash of the repo.
	Revision string `json:"revision"`

	// ParentRevisions is the full git hashes of the parents of
	// Revision.
	ParentRevisions []string `json:"parentRevisions"`

	// GoRevision is the full git hash of the "go" repo, if Repo is not "go" itself.
	// Otherwise this is empty.
	GoRevision string `json:"goRevision,omitempty"`

	// Date is the commit date of this revision, formatted in RFC3339.
	Date string `json:"date"`

	// Branch is the branch of this commit, e.g. "master" or "dev.ssa".
	Branch string `json:"branch"`

	// Author is the author of this commit in standard git form
	// "Name <email>".
	Author string `json:"author"`

	// Desc is the commit message of this commit. It may be
	// truncated.
	Desc string `json:"desc"`

	// Results are the build results for each of the builders in
	// the same length slice BuildStatus.Builders.
	// Each string is either "" (if no data), "ok", or the URL to failure logs.
	Results []string `json:"results"`
}
