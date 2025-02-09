package templates

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/mikestefanello/pagoda/config"
)

var (
	// CacheBuster is a random string that changes on each app restart
	// to force browsers to download new versions of static files
	CacheBuster string

	// web is the web interface for URL generation
	web interface {
		Reverse(name string, params ...any) string
	}
)

// SetWeb sets the web interface for URL generation
func SetWeb(w interface{ Reverse(name string, params ...any) string }) {
	web = w
}

// SetCacheBuster sets the cache buster string
func SetCacheBuster(cb string) {
	CacheBuster = cb
}

// url generates a URL from a given route name and optional parameters
func url(routeName string, params ...any) templ.SafeURL {
	if web == nil {
		panic("web interface not set")
	}
	return templ.URL(web.Reverse(routeName, params...))
}

// file appends a cache buster to a given filepath
func file(filepath string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("/%s/%s?v=%s", config.StaticPrefix, filepath, CacheBuster))
}

// link outputs HTML for a link element with active class support
func link(url, text, currentPath string, classes ...string) templ.Component {
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

// add adds two integers for use in templates
func add(a, b int) string {
	return fmt.Sprintf("%d", a+b)
}

// sub subtracts two integers for use in templates
func sub(a, b int) string {
	return fmt.Sprintf("%d", a-b)
}
