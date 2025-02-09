package services

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/templates"
)

// cachedPageGroup stores the cache group for cached pages
const cachedPageGroup = "page"

type (
	// TemplateRenderer provides a flexible and easy to use method of rendering templ components
	// while providing caching depending on your current environment
	TemplateRenderer struct {
		// config stores application configuration
		config *config.Config

		// cache stores the cache client
		cache *CacheClient
	}

	// CachedPage is what is used to store a rendered Page in the cache
	CachedPage struct {
		// URL stores the URL of the requested page
		URL string

		// HTML stores the complete HTML of the rendered Page
		HTML []byte

		// StatusCode stores the HTTP status code
		StatusCode int

		// Headers stores the HTTP headers
		Headers map[string]string
	}
)

// NewTemplateRenderer creates a new TemplateRenderer
func NewTemplateRenderer(cfg *config.Config, cache *CacheClient) *TemplateRenderer {
	return &TemplateRenderer{
		config: cfg,
		cache:  cache,
	}
}

// RenderPage renders a Page as an HTTP response using templ components
func (t *TemplateRenderer) RenderPage(ctx echo.Context, page page.Page) error {
	// Page name is required
	if page.Name == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "page render failed due to missing name")
	}

	// Use the app name in configuration if a value was not set
	if page.AppName == "" {
		page.AppName = t.config.App.Name
	}

	// Create page context for templ components
	pageCtx := &pageContext{page: page}

	// Check if this is an HTMX non-boosted request which indicates that only partial
	// content should be rendered
	if page.HTMX.Request.Enabled && !page.HTMX.Request.Boosted {
		// Switch the layout which will only render the page content
		page.Layout = templates.LayoutHTMX
	}

	// Create a buffer to capture the rendered HTML for caching
	buf := new(bytes.Buffer)
	writer := io.MultiWriter(ctx.Response().Writer, buf)

	// Set the status code
	ctx.Response().Status = page.StatusCode

	// Set any headers
	for k, v := range page.Headers {
		ctx.Response().Header().Set(k, v)
	}

	// Apply the HTMX response, if one
	if page.HTMX.Response != nil {
		page.HTMX.Response.Apply(ctx)
	}

	// Render the appropriate component based on the page name
	var component templ.Component
	switch page.Name {
	case "home":
		component = pages.Home(pageCtx)
	case "about":
		component = pages.About(pageCtx)
	case "contact":
		component = pages.Contact(pageCtx)
	case "cache":
		component = pages.Cache(pageCtx)
	case "task":
		if page.HTMX.Request.Target != "" {
			component = pages.TaskProgress(pageCtx)
		} else {
			component = pages.Task(pageCtx)
		}
	case "search":
		component = pages.SearchResults(pageCtx)
	case "login":
		component = pages.Login(pageCtx)
	case "register":
		component = pages.Register(pageCtx)
	case "forgot_password":
		component = pages.ForgotPassword(pageCtx)
	case "reset_password":
		component = pages.ResetPassword(pageCtx)
	case "error":
		component = pages.Error(pageCtx)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unknown page: %s", page.Name))
	}

	// Render the component
	err := component.Render(ctx.Request().Context(), writer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to render template: %s", err))
	}

	// Cache this page, if caching was enabled
	t.cachePage(ctx, page, buf)

	return nil
}

// cachePage caches the HTML for a given Page if the Page has caching enabled
func (t *TemplateRenderer) cachePage(ctx echo.Context, page page.Page, html *bytes.Buffer) {
	if !page.Cache.Enabled || page.IsAuth {
		return
	}

	// If no expiration time was provided, default to the configuration value
	if page.Cache.Expiration == 0 {
		page.Cache.Expiration = t.config.Cache.Expiration.Page
	}

	// Extract the headers
	headers := make(map[string]string)
	for k, v := range ctx.Response().Header() {
		headers[k] = v[0]
	}

	// The request URL is used as the cache key so the middleware can serve the
	// cached page on matching requests
	key := ctx.Request().URL.String()
	cp := &CachedPage{
		URL:        key,
		HTML:       html.Bytes(),
		Headers:    headers,
		StatusCode: ctx.Response().Status,
	}

	err := t.cache.
		Set().
		Group(cachedPageGroup).
		Key(key).
		Tags(page.Cache.Tags...).
		Expiration(page.Cache.Expiration).
		Data(cp).
		Save(ctx.Request().Context())

	switch {
	case err == nil:
		log.Ctx(ctx).Debug("cached page")
	case !context.IsCanceledError(err):
		log.Ctx(ctx).Error("failed to cache page",
			"error", err,
		)
	}
}

// GetCachedPage attempts to fetch a cached page for a given URL
func (t *TemplateRenderer) GetCachedPage(ctx echo.Context, url string) (*CachedPage, error) {
	p, err := t.cache.
		Get().
		Group(cachedPageGroup).
		Key(url).
		Fetch(ctx.Request().Context())

	if err != nil {
		return nil, err
	}

	return p.(*CachedPage), nil
}

// getCacheKey gets a cache key for a given group and ID
func (t *TemplateRenderer) getCacheKey(group, key string) string {
	if group != "" {
		return fmt.Sprintf("%s:%s", group, key)
	}
	return key
}

// pageContext implements templates.PageContext
type pageContext struct {
	page page.Page
}

func (c *pageContext) IsAuth() bool {
	return c.page.IsAuth
}

func (c *pageContext) GetPath() string {
	return c.page.Path
}

func (c *pageContext) GetCSRF() string {
	return c.page.CSRF
}

func (c *pageContext) GetTitle() string {
	return c.page.Title
}

func (c *pageContext) GetAppName() string {
	return c.page.AppName
}

func (c *pageContext) GetMessages(typ msg.Type) []template.HTML {
	return c.page.GetMessages(typ)
}

func (c *pageContext) GetHTMXRequest() any {
	return c.page.HTMX.Request
}

func (c *pageContext) GetData() any {
	return c.page.Data
}
