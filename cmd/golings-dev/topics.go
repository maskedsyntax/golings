package main

import (
	"context"
	"flag"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/maskedsyntax/golings/internal/catalog"
)

const gobyexampleIndex = "https://gobyexample.com/"

// indexEntry is one topic link on the Go by Example index page.
type indexEntry struct {
	Slug, Title string
}

var indexLinkRE = regexp.MustCompile(`<li>\s*<a href="([a-z0-9-]+)">([^<]+)</a>`)

// parseIndex extracts the topic links, in page order.
func parseIndex(page string) []indexEntry {
	var out []indexEntry
	for _, m := range indexLinkRE.FindAllStringSubmatch(page, -1) {
		out = append(out, indexEntry{Slug: m[1], Title: strings.TrimSpace(html.UnescapeString(m[2]))})
	}
	return out
}

// compareTopics diffs topics.json against the live index. It is the only
// golings command that uses the network, and it never writes anything.
func compareTopicsCmd(ctx context.Context, cat *catalog.Catalog, args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("compare-topics", flag.ContinueOnError)
	fs.SetOutput(stderr)
	from := fs.String("from", "", "read the index from this saved HTML file instead of "+gobyexampleIndex)
	if err := fs.Parse(args); err != nil {
		return 2
	}
	var page []byte
	var err error
	if *from != "" {
		page, err = os.ReadFile(*from)
	} else {
		page, err = fetchIndex(ctx)
	}
	if err != nil {
		fmt.Fprintln(stderr, "compare-topics:", err)
		return 1
	}
	live := parseIndex(string(page))
	if len(live) == 0 {
		fmt.Fprintln(stderr, "compare-topics: found no topic links; the page layout may have changed")
		return 1
	}
	diffs := diffTopics(cat.Topics.Topics, live)
	if len(diffs) == 0 {
		fmt.Fprintf(stdout, "ok: topics.json matches the live index (%d topics)\n", len(live))
		return 0
	}
	fmt.Fprintf(stdout, "topics.json (snapshot %s) differs from the live index:\n", cat.Topics.SnapshotDate)
	for _, d := range diffs {
		fmt.Fprintln(stdout, "  "+d)
	}
	return 1
}

func fetchIndex(ctx context.Context) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, gobyexampleIndex, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "golings-dev compare-topics")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %s", gobyexampleIndex, resp.Status)
	}
	return io.ReadAll(io.LimitReader(resp.Body, 4<<20))
}

// diffTopics reports topics added to or removed from the index, title
// changes, and order changes.
func diffTopics(ours []catalog.Topic, live []indexEntry) []string {
	var diffs []string
	slug := func(t catalog.Topic) string { return strings.TrimPrefix(t.URL, catalog.TopicURLPrefix) }
	liveBySlug := map[string]indexEntry{}
	liveOrder := map[string]int{}
	for i, e := range live {
		liveBySlug[e.Slug] = e
		liveOrder[e.Slug] = i
	}
	oursBySlug := map[string]bool{}
	var common []string
	for _, t := range ours {
		s := slug(t)
		oursBySlug[s] = true
		e, ok := liveBySlug[s]
		switch {
		case !ok:
			diffs = append(diffs, fmt.Sprintf("removed upstream: %s (%q)", s, t.Title))
		case e.Title != t.Title:
			diffs = append(diffs, fmt.Sprintf("title changed: %s: %q -> %q", s, t.Title, e.Title))
			common = append(common, s)
		default:
			common = append(common, s)
		}
	}
	for _, e := range live {
		if !oursBySlug[e.Slug] {
			diffs = append(diffs, fmt.Sprintf("new upstream: %s (%q) at position %d", e.Slug, e.Title, liveOrder[e.Slug]+1))
		}
	}
	for i := 1; i < len(common); i++ {
		if liveOrder[common[i-1]] > liveOrder[common[i]] {
			diffs = append(diffs, fmt.Sprintf("order changed: %s now comes before %s", common[i], common[i-1]))
		}
	}
	return diffs
}
