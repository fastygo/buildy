package site

import (
	"context"
	"net/http"
	"strings"

	"github.com/fastygo/buildy/internal/fixtures"
	"github.com/fastygo/buildy/internal/ui/components/toggles"
	"github.com/fastygo/buildy/internal/ui/layout"
	"github.com/fastygo/buildy/internal/views"
	"github.com/fastygo/framework/pkg/web/locale"
)

type fixtureProvider struct {
	available     []string
	defaultLocale string
}

func newFixtureProvider(available []string, defaultLocale string) *fixtureProvider {
	return &fixtureProvider{
		available:     available,
		defaultLocale: defaultLocale,
	}
}

func (p *fixtureProvider) Layout(ctx context.Context, r *http.Request, code string, title string, active string) views.LayoutData {
	site := p.siteLocale(code)
	return views.LayoutData{
		PageTitle:      title,
		Lang:           code,
		Brand:          site.Brand,
		FooterText:     site.Footer,
		Active:         active,
		SidebarItems:   navItems(site.SidebarNav),
		HeaderNavItems: navItems(site.HeaderNav),
		Assets: views.AssetPaths{
			CSS:     assetCSS,
			ThemeJS: assetThemeJS,
			AppJS:   assetAppJS,
			ChatJS:  assetChatJS,
		},
		Theme: layout.ThemeToggleProps{
			Label:              site.Theme.Label,
			SwitchToDarkLabel:  site.Theme.SwitchToDarkLabel,
			SwitchToLightLabel: site.Theme.SwitchToLight,
		},
		LanguageSwitch: p.languageSwitch(ctx, r, site),
	}
}

func (p *fixtureProvider) Home(code string) views.HomeData {
	screen := p.homeScreen(code)
	return views.HomeData{
		Welcome:      screen.Welcome,
		WelcomeBrand: screen.WelcomeBrand,
		Description:  screen.Description,
		Workspace: views.HomeWorkspace{
			Label:          screen.Workspace.Label,
			PrototypeMode:  screen.Workspace.PrototypeMode,
			ProfileName:    screen.Workspace.ProfileName,
			ProfileInitial: screen.Workspace.ProfileInitial,
		},
		Hero: views.HomeHero{
			Title:             screen.Hero.Title,
			Subtitle:          screen.Hero.Subtitle,
			PromptPlaceholder: screen.Hero.PromptPlaceholder,
			PrimaryAction:     screen.Hero.PrimaryAction,
			HelperText:        screen.Hero.HelperText,
			AttachmentLabel:   screen.Hero.AttachmentLabel,
			AssistantLabel:    screen.Hero.AssistantLabel,
			ShortcutLabel:     screen.Hero.ShortcutLabel,
			KeyboardHint:      screen.Hero.KeyboardHint,
		},
		Suggestions: homeSuggestions(screen.Suggestions),
		Workflow:    homeWorkflow(screen.Workflow),
		Tools:       homeTools(screen.Tools),
		Showcase: views.HomeShowcase{
			Title:       screen.Showcase.Title,
			Description: screen.Showcase.Description,
			ActionLabel: screen.Showcase.ActionLabel,
			Items:       homeShowcaseItems(screen.Showcase.Items),
		},
		Notice: views.HomeNotice{
			Title:       screen.Notice.Title,
			Description: screen.Notice.Description,
			ActionLabel: screen.Notice.ActionLabel,
		},
		Aria: views.HomeAccessibility{
			PageLabel:        screen.Aria.PageLabel,
			ChatFormLabel:    screen.Aria.ChatFormLabel,
			ComposerLabel:    screen.Aria.ComposerLabel,
			SendLabel:        screen.Aria.SendLabel,
			AttachmentLabel:  screen.Aria.AttachmentLabel,
			AssistantLabel:   screen.Aria.AssistantLabel,
			ShortcutLabel:    screen.Aria.ShortcutLabel,
			SuggestionsLabel: screen.Aria.SuggestionsLabel,
			WorkflowLabel:    screen.Aria.WorkflowLabel,
			ToolsLabel:       screen.Aria.ToolsLabel,
			ShowcaseLabel:    screen.Aria.ShowcaseLabel,
			NoticeLabel:      screen.Aria.NoticeLabel,
		},
	}
}

func (p *fixtureProvider) HomeTitle(code string) string {
	return p.homeScreen(code).Title
}

func (p *fixtureProvider) Sample(code string) views.SampleData {
	screen := p.sampleScreen(code)
	return views.SampleData{
		Title:       screen.Title,
		Description: screen.Description,
		Body:        screen.Body,
	}
}

func (p *fixtureProvider) SampleTitle(code string) string {
	return p.sampleScreen(code).Title
}

func (p *fixtureProvider) ResolveCode(ctx context.Context) string {
	code := locale.From(ctx)
	if code == "" {
		return p.defaultLocale
	}
	return code
}

func (p *fixtureProvider) siteLocale(code string) fixtures.SiteLocale {
	out, err := fixtures.LoadSiteLocale(code)
	if err == nil {
		return out
	}
	out, _ = fixtures.LoadSiteLocale(p.defaultLocale)
	return out
}

func (p *fixtureProvider) homeScreen(code string) fixtures.HomeScreenLocale {
	out, err := fixtures.LoadHomeScreen(code)
	if err == nil {
		return out
	}
	out, _ = fixtures.LoadHomeScreen(p.defaultLocale)
	return out
}

func (p *fixtureProvider) sampleScreen(code string) fixtures.SampleScreenLocale {
	out, err := fixtures.LoadSampleScreen(code)
	if err == nil {
		return out
	}
	out, _ = fixtures.LoadSampleScreen(p.defaultLocale)
	return out
}

func (p *fixtureProvider) languageSwitch(ctx context.Context, r *http.Request, site fixtures.SiteLocale) toggles.LanguageSwitchProps {
	current := strings.ToLower(strings.TrimSpace(locale.From(ctx)))
	if current == "" {
		current = strings.ToLower(strings.TrimSpace(p.defaultLocale))
	}
	labels := map[string]string{"en": "En", "ru": "Ru"}
	var items []toggles.LanguageSwitchItem
	for _, loc := range p.available {
		loc = strings.ToLower(strings.TrimSpace(loc))
		label := labels[loc]
		if label == "" {
			label = strings.ToUpper(loc)
		}
		items = append(items, toggles.LanguageSwitchItem{
			Locale: loc,
			Label:  label,
			Href:   locale.BuildLangHref(r, loc, p.defaultLocale),
			Active: loc == current,
		})
	}
	return toggles.LanguageSwitchProps{
		AriaLabel: site.LanguageToggleLabel,
		Items:     items,
	}
}

func navItems(items []fixtures.NavLocale) []layout.NavItem {
	out := make([]layout.NavItem, 0, len(items))
	for _, item := range items {
		out = append(out, layout.NavItem{
			Label: item.Label,
			Path:  item.Path,
			Icon:  item.Icon,
		})
	}
	return out
}

func homeSuggestions(items []fixtures.HomeSuggestionLocale) []views.HomeSuggestion {
	out := make([]views.HomeSuggestion, 0, len(items))
	for _, item := range items {
		out = append(out, views.HomeSuggestion{
			Label: item.Label,
			Icon:  item.Icon,
		})
	}
	return out
}

func homeWorkflow(items []fixtures.HomeWorkflowStepLocale) []views.HomeWorkflowStep {
	out := make([]views.HomeWorkflowStep, 0, len(items))
	for _, item := range items {
		out = append(out, views.HomeWorkflowStep{
			Label:       item.Label,
			Description: item.Description,
			Icon:        item.Icon,
			Tone:        item.Tone,
		})
	}
	return out
}

func homeTools(items []fixtures.HomeToolCardLocale) []views.HomeToolCard {
	out := make([]views.HomeToolCard, 0, len(items))
	for _, item := range items {
		out = append(out, views.HomeToolCard{
			Title:       item.Title,
			Description: item.Description,
			Icon:        item.Icon,
			Tone:        item.Tone,
		})
	}
	return out
}

func homeShowcaseItems(items []fixtures.HomeShowcaseItemLocale) []views.HomeShowcaseItem {
	out := make([]views.HomeShowcaseItem, 0, len(items))
	for _, item := range items {
		out = append(out, views.HomeShowcaseItem{
			Name:         item.Name,
			Category:     item.Category,
			Description:  item.Description,
			PreviewNote:  item.PreviewNote,
			Capabilities: append([]string(nil), item.Capabilities...),
			UseLabel:     item.UseLabel,
			PreviewLabel: item.PreviewLabel,
			Tone:         item.Tone,
		})
	}
	return out
}
