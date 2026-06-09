package layout

import "github.com/a-h/templ"

// NavItem describes a single navigation link.
type NavItem struct {
	Path  string
	Label string
	Icon  string
}

// NavProps configures horizontal or vertical navigation links.
type NavProps struct {
	Items    []NavItem
	Active   string
	Vertical bool
}

// SidebarProps configures the desktop sidebar and primary mobile sheet group.
type SidebarProps struct {
	Items  []NavItem
	Active string
}

// HeaderProps configures the top header bar.
type HeaderProps struct {
	BrandName       string
	HeaderNavItems  []NavItem
	Active          string
	ShowMenuTrigger bool
	ShowBrand       bool
	Trailing        templ.Component
	ThemeToggle     ThemeToggleProps
}

// ThemeToggleProps configures copy for the theme toggle button.
type ThemeToggleProps struct {
	Label              string
	SwitchToDarkLabel  string
	SwitchToLightLabel string
}

// ShellProps configures the full page shell (sidebar + header + main + footer).
type ShellProps struct {
	Title          string
	Lang           string
	BrandName      string
	Active         string
	SidebarItems   []NavItem
	HeaderNavItems []NavItem
	FooterText     string
	HeadExtra      templ.Component
	HeaderTrailing templ.Component
	ThemeToggle    ThemeToggleProps
}
