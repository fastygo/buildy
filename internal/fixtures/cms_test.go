package fixtures

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

func TestLoadCMSSiteSeed(t *testing.T) {
	seed, err := LoadCMSSiteSeed("buildy-minimal-site.en.json")
	if err != nil {
		t.Fatalf("load cms site seed: %v", err)
	}
	if seed.CodexVersion != "codex.v1" {
		t.Fatalf("codex version = %q, want codex.v1", seed.CodexVersion)
	}
	if seed.Site.Title != "BuildY" {
		t.Fatalf("site title = %q, want BuildY", seed.Site.Title)
	}
	if len(seed.ContentTypes) == 0 || len(seed.ContentEntries) == 0 || len(seed.Menus) == 0 {
		t.Fatalf("cms seed must include content types, entries, and menus: %#v", seed)
	}
}

func TestCMSFixturesHaveSemanticIntegrity(t *testing.T) {
	files, err := filepath.Glob(filepath.Join("cms", "*.json"))
	if err != nil {
		t.Fatalf("glob cms fixtures: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected cms fixtures")
	}
	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			seed, err := LoadCMSSiteSeed(filepath.Base(file))
			if err != nil {
				t.Fatalf("load seed: %v", err)
			}
			seenContentTypes := map[string]bool{}
			for _, item := range seed.ContentTypes {
				requireUnique(t, seenContentTypes, item.ID, "content type id")
			}
			seenEntries := map[string]bool{}
			entrySlugs := map[string]bool{}
			for _, item := range seed.ContentEntries {
				requireUnique(t, seenEntries, item.ID, "content entry id")
				if !seenContentTypes[item.Kind] {
					t.Fatalf("content entry %q references missing kind/content type %q", item.ID, item.Kind)
				}
				if item.Slug != "" {
					entrySlugs[item.Slug] = true
				}
			}
			for _, menu := range seed.Menus {
				seenItems := map[string]bool{}
				for _, item := range menu.Items {
					requireUnique(t, seenItems, item.ID, "menu item id")
					if strings.HasPrefix(item.URL, "/") && item.URL != "/" {
						slug := strings.TrimPrefix(item.URL, "/")
						if !entrySlugs[slug] {
							t.Fatalf("menu item %q points to missing page slug %q", item.ID, slug)
						}
					}
				}
			}
			seenSettings := map[string]bool{}
			for _, item := range seed.Settings {
				requireUnique(t, seenSettings, item.Key, "setting key")
			}
		})
	}
}

func TestCMSFixturesValidateAgainstAppCMSCodexSchema(t *testing.T) {
	schema := loadCodexSchema(t, "seed-site.schema.json")
	files, err := filepath.Glob(filepath.Join("cms", "*.json"))
	if err != nil {
		t.Fatalf("glob cms fixtures: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected cms fixtures")
	}
	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			raw, err := os.ReadFile(file)
			if err != nil {
				t.Fatalf("read fixture: %v", err)
			}
			var value any
			if err := json.Unmarshal(raw, &value); err != nil {
				t.Fatalf("decode fixture: %v", err)
			}
			if err := schema.Validate(value); err != nil {
				t.Fatalf("fixture failed AppCMS codex validation: %v", err)
			}
		})
	}
}

func loadCodexSchema(t *testing.T, name string) *jsonschema.Schema {
	t.Helper()
	compiler := jsonschema.NewCompiler()
	compiler.DefaultDraft(jsonschema.Draft2020)
	dir := appCMSCodexSchemaDir(t)
	if err := registerCodexSchemaDir(compiler, dir); err != nil {
		t.Fatalf("register codex schemas: %v", err)
	}
	schema, err := compiler.Compile(codexSchemaURL(name))
	if err != nil {
		t.Fatalf("compile codex schema %q: %v", name, err)
	}
	return schema
}

func appCMSCodexSchemaDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve test path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", "..", "..", "@AppCMS", "schema", "codex", "v1"))
}

func registerCodexSchemaDir(compiler *jsonschema.Compiler, dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".schema.json") {
			continue
		}
		raw, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			return err
		}
		var doc any
		if err := json.Unmarshal(raw, &doc); err != nil {
			return fmt.Errorf("decode %s: %w", entry.Name(), err)
		}
		if err := compiler.AddResource(codexSchemaURL(entry.Name()), doc); err != nil {
			return err
		}
	}
	return nil
}

func codexSchemaURL(name string) string {
	return fmt.Sprintf("https://fastygo.dev/schema/codex/v1/%s", name)
}

func requireUnique(t *testing.T, seen map[string]bool, value string, label string) {
	t.Helper()
	if value == "" {
		t.Fatalf("%s is required", label)
	}
	if seen[value] {
		t.Fatalf("duplicate %s %q", label, value)
	}
	seen[value] = true
}
