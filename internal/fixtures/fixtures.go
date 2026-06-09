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
	NavLabel     string                   `json:"nav_label"`
	Title        string                   `json:"title"`
	Welcome      string                   `json:"welcome"`
	WelcomeBrand string                   `json:"welcome_brand"`
	Description  string                   `json:"description"`
	Workspace    HomeWorkspaceLocale      `json:"workspace"`
	Hero         HomeHeroLocale           `json:"hero"`
	Suggestions  []HomeSuggestionLocale   `json:"prompt_suggestions"`
	Workflow     []HomeWorkflowStepLocale `json:"workflow_steps"`
	Tools        []HomeToolCardLocale     `json:"tool_cards"`
	Showcase     HomeShowcaseLocale       `json:"showcase"`
	Notice       HomeNoticeLocale         `json:"notice"`
	Aria         HomeAccessibilityLocale  `json:"aria"`
}

type HomeWorkspaceLocale struct {
	Label          string `json:"label"`
	PrototypeMode  string `json:"prototype_mode"`
	ProfileName    string `json:"profile_name"`
	ProfileInitial string `json:"profile_initial"`
}

type HomeHeroLocale struct {
	Title             string `json:"title"`
	Subtitle          string `json:"subtitle"`
	PromptPlaceholder string `json:"prompt_placeholder"`
	PrimaryAction     string `json:"primary_action"`
	HelperText        string `json:"helper_text"`
	AttachmentLabel   string `json:"attachment_label"`
	AssistantLabel    string `json:"assistant_label"`
	ShortcutLabel     string `json:"shortcut_label"`
	KeyboardHint      string `json:"keyboard_hint"`
}

type HomeSuggestionLocale struct {
	Label string `json:"label"`
	Icon  string `json:"icon"`
}

type HomeWorkflowStepLocale struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Tone        string `json:"tone"`
}

type HomeToolCardLocale struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Tone        string `json:"tone"`
}

type HomeShowcaseLocale struct {
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	ActionLabel string                   `json:"action_label"`
	Items       []HomeShowcaseItemLocale `json:"items"`
}

type HomeShowcaseItemLocale struct {
	Name         string   `json:"name"`
	Category     string   `json:"category"`
	Description  string   `json:"description"`
	PreviewNote  string   `json:"preview_note"`
	Capabilities []string `json:"capabilities"`
	UseLabel     string   `json:"use_label"`
	PreviewLabel string   `json:"preview_label"`
	Tone         string   `json:"tone"`
}

type HomeNoticeLocale struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ActionLabel string `json:"action_label"`
}

type HomeAccessibilityLocale struct {
	PageLabel        string `json:"page_label"`
	ChatFormLabel    string `json:"chat_form_label"`
	ComposerLabel    string `json:"composer_label"`
	SendLabel        string `json:"send_label"`
	AttachmentLabel  string `json:"attachment_label"`
	AssistantLabel   string `json:"assistant_label"`
	ShortcutLabel    string `json:"shortcut_label"`
	SuggestionsLabel string `json:"suggestions_label"`
	WorkflowLabel    string `json:"workflow_label"`
	ToolsLabel       string `json:"tools_label"`
	ShowcaseLabel    string `json:"showcase_label"`
	NoticeLabel      string `json:"notice_label"`
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
