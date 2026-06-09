package site

import (
	"net/http"

	"github.com/fastygo/buildy/internal/views"
	"github.com/fastygo/framework/pkg/app"
	"github.com/fastygo/framework/pkg/web"
)

const (
	assetCSS     = "/static/css/app.css"
	assetThemeJS = "/static/js/theme.js"
	assetAppJS   = "/static/js/ui8kit.js"
	assetChatJS  = "/static/js/chat-composer.js"
)

// Feature wires public HTTP routes for the app shell (header nav, i18n, theme).
type Feature struct {
	available     []string
	defaultLocale string
	content       *fixtureProvider
}

// NewFeature constructs the site feature.
func NewFeature(available []string, defaultLocale string) *Feature {
	return &Feature{
		available:     available,
		defaultLocale: defaultLocale,
		content:       newFixtureProvider(available, defaultLocale),
	}
}

// ID implements app.Feature.
func (f *Feature) ID() string {
	return "site"
}

// NavItems implements app.Feature.
// Navigation labels are resolved per request from fixtures (see siteNav).
func (f *Feature) NavItems() []app.NavItem {
	return nil
}

// Routes implements app.Feature.
func (f *Feature) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", f.getHome)
	mux.HandleFunc("GET /sample", f.getSample)
}

func (f *Feature) getHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := f.content.ResolveCode(ctx)
	home := f.content.Home(code)
	data := f.content.Layout(ctx, r, code, f.content.HomeTitle(code), "/")
	_ = web.Render(ctx, w, views.SiteShell(data, views.HomePage(home)))
}

func (f *Feature) getSample(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := f.content.ResolveCode(ctx)
	sample := f.content.Sample(code)
	data := f.content.Layout(ctx, r, code, f.content.SampleTitle(code), "/sample")
	_ = web.Render(ctx, w, views.SiteShell(data, views.SamplePage(sample)))
}
