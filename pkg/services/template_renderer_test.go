package services

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brian-dlee/lab/pkg/htmx"
	"github.com/brian-dlee/lab/pkg/page"
	"github.com/brian-dlee/lab/pkg/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplateRenderer_RenderPage(t *testing.T) {
	setup := func() (*httptest.ResponseRecorder, page.Page) {
		ctx, rec := tests.NewContext(c.Web, "/test/TestTemplateRenderer_RenderPage")
		tests.InitSession(ctx)

		p := page.New(ctx)
		p.Name = "home"
		p.Layout = "main"
		p.Cache.Enabled = false
		p.Headers["A"] = "b"
		p.Headers["C"] = "d"
		p.StatusCode = http.StatusCreated
		return rec, p
	}

	t.Run("missing name", func(t *testing.T) {
		// Rendering should fail if the Page has no name
		_, p := setup()
		p.Name = ""
		err := c.TemplateRenderer.RenderPage(p)
		assert.Error(t, err)
	})

	t.Run("no page cache", func(t *testing.T) {
		_, p := setup()
		err := c.TemplateRenderer.RenderPage(p)
		require.NoError(t, err)

		// Check status code and headers
		assert.Equal(t, http.StatusCreated, ctx.Response().Status)
		for k, v := range p.Headers {
			assert.Equal(t, v, ctx.Response().Header().Get(k))
		}
	})

	t.Run("htmx rendering", func(t *testing.T) {
		_, p := setup()
		p.HTMX.Request.Enabled = true
		p.HTMX.Response = &htmx.Response{
			Trigger: "trigger",
		}
		err := c.TemplateRenderer.RenderPage(p)
		require.NoError(t, err)

		// Check HTMX header
		assert.Equal(t, "trigger", ctx.Response().Header().Get(htmx.HeaderTrigger))
	})

	t.Run("page cache", func(t *testing.T) {
		rec, p := setup()
		p.Cache.Enabled = true
		p.Cache.Tags = []string{"tag1"}
		err := c.TemplateRenderer.RenderPage(p)
		require.NoError(t, err)

		// Fetch from the cache
		cp, err := c.TemplateRenderer.GetCachedPage(ctx, p.URL)
		require.NoError(t, err)

		// Compare the cached page
		assert.Equal(t, p.URL, cp.URL)
		assert.Equal(t, p.Headers, cp.Headers)
		assert.Equal(t, p.StatusCode, cp.StatusCode)
		assert.Equal(t, rec.Body.Bytes(), cp.HTML)

		// Clear the tag
		err = c.Cache.
			Flush().
			Tags(p.Cache.Tags[0]).
			Execute(context.Background())
		require.NoError(t, err)

		// Refetch from the cache and expect no results
		_, err = c.TemplateRenderer.GetCachedPage(ctx, p.URL)
		assert.Error(t, err)
	})
}
