package fixtures

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLocaleLoaders(t *testing.T) {
	for _, code := range []string{"en", "ru"} {
		t.Run(code, func(t *testing.T) {
			site, err := LoadSiteLocale(code)
			if err != nil {
				t.Fatalf("load site locale: %v", err)
			}
			if site.Brand == "" || site.Footer == "" || site.LanguageToggleLabel == "" {
				t.Fatalf("site locale %q has incomplete chrome copy: %#v", code, site)
			}
			if len(site.SidebarNav) == 0 || len(site.HeaderNav) == 0 {
				t.Fatalf("site locale %q must include sidebar and header nav", code)
			}
			home, err := LoadHomeScreen(code)
			if err != nil {
				t.Fatalf("load home screen: %v", err)
			}
			if home.Title == "" || home.NavLabel == "" || home.Description == "" {
				t.Fatalf("home locale %q has incomplete copy: %#v", code, home)
			}
			sample, err := LoadSampleScreen(code)
			if err != nil {
				t.Fatalf("load sample screen: %v", err)
			}
			if sample.Title == "" || sample.NavLabel == "" || sample.Body == "" {
				t.Fatalf("sample locale %q has incomplete copy: %#v", code, sample)
			}
		})
	}
}

func TestLocaleJSONShapeParity(t *testing.T) {
	for _, pair := range [][2]string{
		{"site/en.json", "site/ru.json"},
		{"screens/home.en.json", "screens/home.ru.json"},
		{"screens/sample.en.json", "screens/sample.ru.json"},
	} {
		t.Run(pair[0], func(t *testing.T) {
			left := loadJSONShape(t, pair[0])
			right := loadJSONShape(t, pair[1])
			if !reflect.DeepEqual(left, right) {
				t.Fatalf("fixture JSON shape mismatch:\n%s: %#v\n%s: %#v", pair[0], left, pair[1], right)
			}
		})
	}
}

func loadJSONShape(t *testing.T, path string) any {
	t.Helper()
	raw, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	var value any
	if err := json.Unmarshal(raw, &value); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	return jsonShape(value)
}

func jsonShape(value any) any {
	switch typed := value.(type) {
	case map[string]any:
		out := map[string]any{}
		for key, child := range typed {
			out[key] = jsonShape(child)
		}
		return out
	case []any:
		if len(typed) == 0 {
			return []any{}
		}
		return []any{jsonShape(typed[0])}
	default:
		if value == nil {
			return "null"
		}
		return reflect.TypeOf(value).String()
	}
}
