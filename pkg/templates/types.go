package templates

import (
	"html/template"

	"github.com/mikestefanello/pagoda/pkg/msg"
)

// BaseContext defines the minimum interface required by all templates
type BaseContext interface {
	// IsAuth returns whether the user is authenticated
	IsAuth() bool

	// GetPath returns the current request path
	GetPath() string

	// GetCSRF returns the CSRF token for the current request
	GetCSRF() string
}

// PageContext defines the interface required by page templates
type PageContext interface {
	BaseContext

	// GetTitle returns the page title
	GetTitle() string

	// GetAppName returns the application name
	GetAppName() string

	// GetMessages gets all flash messages for a given type
	GetMessages(typ msg.Type) []template.HTML

	// GetHTMXRequest returns the HTMX request information
	GetHTMXRequest() any

	// GetData returns the page-specific data
	GetData() any
}
