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
	seed, err := LoadCMSSiteSeed("buildy-minimal-site.json")
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
