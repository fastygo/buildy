package fixtures

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed site/*.json screens/*.json
var localeFS embed.FS

// SiteLocale holds embedded copy for app chrome in one language.
type SiteLocale struct {
	Brand               string      `json:"brand"`
	Footer              string      `json:"footer"`
	Theme               ThemeLocale `json:"theme"`
	LanguageToggleLabel string      `json:"language_toggle_label"`
	SidebarNav          []NavLocale `json:"sidebar_nav"`
	HeaderNav           []NavLocale `json:"header_nav"`
}

type ThemeLocale struct {
	Label             string `json:"label"`
	SwitchToDarkLabel string `json:"switch_to_dark"`
	SwitchToLight     string `json:"switch_to_light"`
}

type NavLocale struct {
	Label string `json:"label"`
	Path  string `json:"path"`
	Icon  string `json:"icon,omitempty"`
}

type HomeScreenLocale struct {
	NavLabel     string `json:"nav_label"`
	Title        string `json:"title"`
	Welcome      string `json:"welcome"`
	WelcomeBrand string `json:"welcome_brand"`
	Description  string `json:"description"`
}

type SampleScreenLocale struct {
	NavLabel    string `json:"nav_label"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

// LoadSiteLocale reads site/{code}.json (e.g. en, ru).
func LoadSiteLocale(code string) (SiteLocale, error) {
	var out SiteLocale
	if err := readLocaleJSON("site/"+code+".json", &out); err != nil {
		return SiteLocale{}, err
	}
	return out, nil
}

// LoadHomeScreen reads screens/home.{code}.json.
func LoadHomeScreen(code string) (HomeScreenLocale, error) {
	var out HomeScreenLocale
	if err := readLocaleJSON("screens/home."+code+".json", &out); err != nil {
		return HomeScreenLocale{}, err
	}
	return out, nil
}

// LoadSampleScreen reads screens/sample.{code}.json.
func LoadSampleScreen(code string) (SampleScreenLocale, error) {
	var out SampleScreenLocale
	if err := readLocaleJSON("screens/sample."+code+".json", &out); err != nil {
		return SampleScreenLocale{}, err
	}
	return out, nil
}

func readLocaleJSON(path string, out any) error {
	raw, err := localeFS.ReadFile(path)
	if err != nil {
		return fmt.Errorf("fixtures: read %q: %w", path, err)
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return fmt.Errorf("fixtures: parse %q: %w", path, err)
	}
	return nil
}
