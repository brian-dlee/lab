package models

import (
	"html/template"
)

type (
	Post struct {
		Title string
		Body  string
	}

	AboutData struct {
		ShowCacheWarning bool
		FrontendTabs     []AboutTab
		BackendTabs      []AboutTab
	}

	AboutTab struct {
		Title string
		Body  template.HTML
	}
)
