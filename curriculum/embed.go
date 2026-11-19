// Package curriculum embeds the canonical Golings curriculum: the manifest,
// the Go by Example topic inventory, and every exercise's starter, checks,
// solution, fixtures and mutants.
//
// Exercise sources live under _exercises/, which the go tool ignores, so
// intentionally broken starters never affect `go build ./...` or
// `go vet ./...` for the golings module itself.
package curriculum

import "embed"

// FS holds manifest.json, topics.json, migrations.json and _exercises/.
//
//go:embed manifest.json topics.json migrations.json all:_exercises
var FS embed.FS
