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
