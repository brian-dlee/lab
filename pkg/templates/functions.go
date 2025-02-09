package templates

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"

	"github.com/brian-dlee/lab/config"
)

var (
	DefaultCacheBuster = random.String(10)
)

type FuncMap struct {
	web         *echo.Echo
	cacheBuster string
}

// NewFuncMap provides a template function map
func NewFuncMap(web *echo.Echo, cacheBuster string) *FuncMap {
	return &FuncMap{web: web, cacheBuster: cacheBuster}
}

// URL generates a URL from a given route name and optional parameters
func (fm *FuncMap) URL(routeName string, params ...any) templ.SafeURL {
	return templ.URL(fm.web.Reverse(routeName, params...))
}

// File appends a cache buster to a given filepath
func (fm *FuncMap) File(filepath string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/%s/%s?v=%s", config.StaticPrefix, filepath, fm.cacheBuster))
}

// Link outputs HTML for a Link element with active class support
func (fm *FuncMap) Link(url, text, currentPath string, classes ...string) templ.Component {
	activeClass := ""
	if currentPath == url {
		activeClass = " is-active"
	}

	return templ.Raw(fmt.Sprintf(`<a href="%s" class="%s%s">%s</a>`,
		url,
		templ.EscapeString(joinClasses(classes...)),
		activeClass,
		templ.EscapeString(text)))
}

// joinClasses joins CSS classes with spaces
func joinClasses(classes ...string) string {
	result := ""
	for i, class := range classes {
		if i > 0 {
			result += " "
		}
		result += class
	}
	return result
}

// Add adds two integers for use in templates
func Add(a, b int) string {
	return fmt.Sprintf("%d", a+b)
}

// Sub subtracts two integers for use in templates
func Sub(a, b int) string {
	return fmt.Sprintf("%d", a-b)
}
