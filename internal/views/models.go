package views

import (
	"github.com/fastygo/buildy/internal/ui/components/toggles"
	"github.com/fastygo/buildy/internal/ui/layout"
)

// AssetPaths are URLs for CSS and JS bundles.
type AssetPaths struct {
	CSS     string
	ThemeJS string
	AppJS   string
	ChatJS  string
}

// LayoutData drives the app shell (header, main, footer).
type LayoutData struct {
	PageTitle      string
	Lang           string
	Brand          string
	FooterText     string
	Active         string
	SidebarItems   []layout.NavItem
	HeaderNavItems []layout.NavItem
	Assets         AssetPaths
	Theme          layout.ThemeToggleProps
	LanguageSwitch toggles.LanguageSwitchProps
}

// DocumentTitle returns the SEO document title for <title>.
func (d LayoutData) DocumentTitle() string {
	return FormatDocumentTitle(d.PageTitle, d.Brand)
}

// FormatDocumentTitle builds "Page · Brand" for the document head.
func FormatDocumentTitle(pageTitle, brand string) string {
	if brand == "" {
		brand = "FastyGo"
	}
	return pageTitle + " · " + brand
}

// HomeData is the home page hero body inside the shell.
type HomeData struct {
	Welcome      string
	WelcomeBrand string
	Description  string
	Workspace    HomeWorkspace
	Hero         HomeHero
	Suggestions  []HomeSuggestion
	Workflow     []HomeWorkflowStep
	Tools        []HomeToolCard
	Showcase     HomeShowcase
	Notice       HomeNotice
	Aria         HomeAccessibility
}

type HomeWorkspace struct {
	Label          string
	PrototypeMode  string
	ProfileName    string
	ProfileInitial string
}

type HomeHero struct {
	Title             string
	Subtitle          string
	PromptPlaceholder string
	PrimaryAction     string
	HelperText        string
	AttachmentLabel   string
	AssistantLabel    string
	ShortcutLabel     string
	KeyboardHint      string
}

type HomeSuggestion struct {
	Label string
	Icon  string
}

type HomeWorkflowStep struct {
	Label       string
	Description string
	Icon        string
	Tone        string
}

type HomeToolCard struct {
	Title       string
	Description string
	Icon        string
	Tone        string
}

type HomeShowcase struct {
	Title       string
	Description string
	ActionLabel string
	Items       []HomeShowcaseItem
}

type HomeShowcaseItem struct {
	Name         string
	Category     string
	Description  string
	PreviewNote  string
	Capabilities []string
	UseLabel     string
	PreviewLabel string
	Tone         string
}

type HomeNotice struct {
	Title       string
	Description string
	ActionLabel string
}

type HomeAccessibility struct {
	PageLabel        string
	ChatFormLabel    string
	ComposerLabel    string
	SendLabel        string
	AttachmentLabel  string
	AssistantLabel   string
	ShortcutLabel    string
	SuggestionsLabel string
	WorkflowLabel    string
	ToolsLabel       string
	ShowcaseLabel    string
	NoticeLabel      string
}

// SampleData is a second stub route for onboarding new pages.
type SampleData struct {
	Title       string
	Description string
	Body        string
}
