package fixtures

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed cms/*.json
var cmsFS embed.FS

// CMSSiteSeed mirrors AppCMS codex.v1 seed-site JSON without importing AppCMS.
type CMSSiteSeed struct {
	CodexVersion   string            `json:"codex_version"`
	Site           CMSSite           `json:"site"`
	ContentTypes   []CMSContentType  `json:"content_types"`
	ContentEntries []CMSContentEntry `json:"content_entries"`
	Menus          []CMSMenu         `json:"menus"`
	Settings       []CMSSetting      `json:"settings"`
}

type CMSSite struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Locale      string `json:"locale"`
}

type CMSContentType struct {
	ID               string             `json:"id"`
	Label            string             `json:"label"`
	PermalinkPattern string             `json:"permalink_pattern"`
	Public           bool               `json:"public"`
	Supports         CMSContentSupports `json:"supports"`
}

type CMSContentSupports struct {
	Title         bool `json:"title"`
	Editor        bool `json:"editor"`
	Excerpt       bool `json:"excerpt"`
	FeaturedMedia bool `json:"featured_media"`
}

type CMSContentEntry struct {
	ID              string            `json:"id"`
	Kind            string            `json:"kind"`
	Title           map[string]string `json:"title"`
	Slug            string            `json:"slug"`
	Content         string            `json:"content"`
	Excerpt         string            `json:"excerpt"`
	Status          string            `json:"status"`
	Visibility      string            `json:"visibility"`
	AuthorID        string            `json:"author_id"`
	FeaturedMediaID string            `json:"featured_media_id"`
	TermIDs         []string          `json:"term_ids"`
	Metadata        map[string]any    `json:"metadata"`
	PublishedAt     string            `json:"published_at"`
	ScheduledFor    string            `json:"scheduled_for"`
	CreatedAt       string            `json:"created_at"`
	UpdatedAt       string            `json:"updated_at"`
}

type CMSMenu struct {
	ID       string        `json:"id"`
	Location string        `json:"location"`
	Items    []CMSMenuItem `json:"items"`
}

type CMSMenuItem struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	URL      string `json:"url"`
	ParentID string `json:"parent_id"`
}

type CMSSetting struct {
	Key    string `json:"key"`
	Group  string `json:"group"`
	Value  any    `json:"value"`
	Public bool   `json:"public"`
}

// LoadCMSSiteSeed reads cms/{name}.json (for example, buildy-minimal-site.en.json).
func LoadCMSSiteSeed(name string) (CMSSiteSeed, error) {
	raw, err := cmsFS.ReadFile("cms/" + name)
	if err != nil {
		return CMSSiteSeed{}, fmt.Errorf("fixtures: read cms seed %q: %w", name, err)
	}
	var out CMSSiteSeed
	if err := json.Unmarshal(raw, &out); err != nil {
		return CMSSiteSeed{}, fmt.Errorf("fixtures: parse cms seed %q: %w", name, err)
	}
	return out, nil
}
