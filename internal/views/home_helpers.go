package views

import (
	"strings"

	"github.com/a-h/templ"
	uiutils "github.com/fastygo/templ/utils"
)

func homePageLabel(value string) string {
	if value == "" {
		return "BuildY home"
	}
	return value
}

func homeIconToken(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "*"
	}
	return strings.ToUpper(value[:1])
}

func homeToneSurface(tone string) string {
	base := "flex h-10 w-10 shrink-0 items-center justify-center rounded-md text-xs font-semibold"
	switch tone {
	case "green":
		return uiutils.Cn(base, "bg-secondary text-secondary-foreground")
	case "amber":
		return uiutils.Cn(base, "bg-primary/10 text-primary")
	case "slate":
		return uiutils.Cn(base, "bg-muted text-muted-foreground")
	default:
		return uiutils.Cn(base, "bg-accent text-accent-foreground")
	}
}

func homeWorkflowStepClass(active bool) string {
	base := "items-start gap-3 rounded-md border border-border p-3"
	if active {
		return uiutils.Cn(base, "bg-card shadow-sm")
	}
	return uiutils.Cn(base, "bg-background")
}

func homePreviewSurfaceClass(tone string) string {
	base := "flex h-28 items-center justify-center rounded-md border border-border text-center text-xs font-semibold"
	switch tone {
	case "green":
		return uiutils.Cn(base, "bg-secondary text-secondary-foreground")
	case "amber":
		return uiutils.Cn(base, "bg-primary/10 text-primary")
	case "slate":
		return uiutils.Cn(base, "bg-muted text-muted-foreground")
	default:
		return uiutils.Cn(base, "bg-accent text-accent-foreground")
	}
}

func homeComposerAttrs(label string) templ.Attributes {
	return templ.Attributes{
		"aria-label":         label,
		"data-buildy-chat":   true,
		"autocomplete":       "off",
		"spellcheck":         "true",
		"data-min-rows":      "4",
		"data-max-height":    "220",
		"data-scroll-anchor": "bottom",
	}
}
