package views

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/fastygo/buildy/internal/ui/components/toggles"
	"github.com/fastygo/buildy/internal/ui/layout"
)

func TestSiteShell_homeRenders(t *testing.T) {
	d := LayoutData{
		PageTitle:  "Home",
		Lang:       "en",
		Brand:      "FastyGo",
		FooterText: "Made with 🤍 in FastyGo",
		Active:     "/",
		SidebarItems: []layout.NavItem{
			{Label: "Home", Path: "/", Icon: "home"},
		},
		HeaderNavItems: []layout.NavItem{
			{Label: "Documentation", Path: "#documentation"},
		},
		Assets: AssetPaths{
			CSS:     "/static/css/app.css",
			ThemeJS: "/static/js/theme.js",
			AppJS:   "/static/js/ui8kit.js",
			ChatJS:  "/static/js/chat-composer.js",
		},
		Theme: layout.ThemeToggleProps{},
		LanguageSwitch: toggles.LanguageSwitchProps{
			AriaLabel: "Language",
			Items: []toggles.LanguageSwitchItem{
				{Locale: "en", Label: "En", Href: "/?lang=en", Active: true},
				{Locale: "ru", Label: "Ru", Href: "/?lang=ru"},
			},
		},
	}
	body := HomePage(testHomeData())
	var buf bytes.Buffer
	if err := SiteShell(d, body).Render(context.Background(), &buf); err != nil {
		t.Fatal(err)
	}
	html := buf.String()
	if !strings.Contains(strings.ToLower(html), "<!doctype html>") {
		t.Fatal("expected full document with doctype")
	}
	if !strings.Contains(html, "<aside") {
		t.Fatal("expected desktop sidebar aside")
	}
	if !strings.Contains(html, `aria-label="Primary navigation"`) {
		t.Fatal("expected sidebar navigation group")
	}
	if !strings.Contains(html, `aria-label="Header navigation"`) {
		t.Fatal("expected header navigation group")
	}
	if !strings.Contains(html, `data-ui8kit="sheet"`) {
		t.Fatal("expected shell mobile sheet markup")
	}
	if !strings.Contains(html, `id="ui8kit-theme-toggle"`) {
		t.Fatal("expected theme toggle control")
	}
	if !strings.Contains(html, `id="ui8kit-mobile-sheet-trigger"`) {
		t.Fatal("expected mobile menu trigger")
	}
	if !strings.Contains(html, "Made with 🤍 in FastyGo") {
		t.Fatal("expected footer copy")
	}
	if !strings.Contains(html, "<title>Home · FastyGo</title>") {
		t.Fatal("expected brand in document title")
	}
	if !strings.Contains(html, "What do you want to build?") {
		t.Fatal("expected redesigned hero title")
	}
	if !strings.Contains(html, `id="buildy-home-prompt"`) || !strings.Contains(html, `data-buildy-chat`) {
		t.Fatal("expected AI chat composer textarea")
	}
	if !strings.Contains(html, "/static/js/chat-composer.js") {
		t.Fatal("expected chat composer enhancement script")
	}
	if !strings.Contains(html, "Prompt Plan") || !strings.Contains(html, "CMS Model") {
		t.Fatal("expected workflow tool cards")
	}
	if !strings.Contains(html, "AppCMS") || !strings.Contains(html, "Webmaster") {
		t.Fatal("expected showcase starter cards")
	}
	if !strings.Contains(html, "Prototype Mode") {
		t.Fatal("expected prototype mode copy")
	}
	if !strings.Contains(html, `role="group"`) || !strings.Contains(html, "En") {
		t.Fatal("expected language switch control")
	}
}

func testHomeData() HomeData {
	return HomeData{
		Workspace: HomeWorkspace{
			Label:         "My Workspace",
			PrototypeMode: "Prototype Mode",
		},
		Hero: HomeHero{
			Title:             "What do you want to build?",
			Subtitle:          "Describe your idea in natural language.",
			PromptPlaceholder: "Describe an app, landing page, CRM, CMS, portal...",
			PrimaryAction:     "Generate prototype",
			HelperText:        "BuildY turns intent into a plan, CMS model, preview, and source.",
			AttachmentLabel:   "Attach",
			AssistantLabel:    "AI",
			ShortcutLabel:     "Slash commands",
			KeyboardHint:      "Shift+Enter for a new line",
		},
		Suggestions: []HomeSuggestion{
			{Label: "Create SaaS landing with CMS content", Icon: "sparkles"},
		},
		Workflow: []HomeWorkflowStep{
			{Label: "Plan", Description: "Structured request", Icon: "list", Tone: "violet"},
			{Label: "Model", Description: "Content types", Icon: "database", Tone: "green"},
			{Label: "Preview", Description: "Live mock", Icon: "eye", Tone: "blue"},
		},
		Tools: []HomeToolCard{
			{Title: "Prompt Plan", Description: "Turn your request into a structured plan.", Icon: "list", Tone: "violet"},
			{Title: "CMS Model", Description: "Define content types and fields.", Icon: "database", Tone: "green"},
		},
		Showcase: HomeShowcase{
			Title:       "Explore powerful starting points",
			Description: "Use mock applications as starting points.",
			ActionLabel: "View all templates",
			Items: []HomeShowcaseItem{
				{
					Name:         "AppCMS",
					Category:     "CMS",
					Description:  "Content websites and pages.",
					PreviewNote:  "Stories",
					Capabilities: []string{"Content types", "Pages"},
					UseLabel:     "Use as starting point",
					PreviewLabel: "Preview",
					Tone:         "blue",
				},
				{
					Name:         "Webmaster",
					Category:     "SEO",
					Description:  "SEO and redirects.",
					PreviewNote:  "Site health",
					Capabilities: []string{"Monitoring", "Redirects"},
					UseLabel:     "Use as starting point",
					PreviewLabel: "Preview",
					Tone:         "green",
				},
			},
		},
		Notice: HomeNotice{
			Title:       "You're in Prototype Mode",
			Description: "Everything here is powered by mock fixtures.",
			ActionLabel: "Learn more about Prototype Mode",
		},
		Aria: HomeAccessibility{
			PageLabel:        "BuildY prototype workspace home",
			ChatFormLabel:    "Prototype generation prompt",
			ComposerLabel:    "Describe what you want BuildY to create",
			SendLabel:        "Generate prototype from prompt",
			AttachmentLabel:  "Attach a reference for the prompt",
			AssistantLabel:   "Ask AI to refine the prompt",
			ShortcutLabel:    "Open prompt command shortcuts",
			SuggestionsLabel: "Suggested prototype prompts",
			WorkflowLabel:    "Prototype generation workflow",
			ToolsLabel:       "BuildY workflow tools",
			ShowcaseLabel:    "Starting point templates",
			NoticeLabel:      "Prototype mode notice",
		},
	}
}
