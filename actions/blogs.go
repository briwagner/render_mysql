package actions

import (
	"fmt"
	"log"
	"net/http"
	"render_mysql/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

// BlogsNew default implementation.
func BlogsNew(c buffalo.Context) error {
	b := &models.Blog{}
	c.Set("blog", b)
	return c.Render(http.StatusOK, r.HTML("blogs/new.html"))
}

// BlogsShow default implementation.
func BlogsShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	b := &models.Blog{}
	err := tx.Find(b, c.Param("blog_id"))
	if err != nil {
		c.Flash().Add("info", "Not found")
		return c.Redirect(404, "/")
	}

	c.Set("blog", b)
	return c.Render(http.StatusOK, r.HTML("blogs/show.html"))
}

// BlogsShow default implementation.
func BlogsIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	b := &models.Blogs{}
	err := tx.All(b)
	if err != nil {
		c.Flash().Add("info", "Not found")
		return c.Redirect(404, "/")
	}

	c.Set("blogs", b)
	return c.Render(http.StatusOK, r.HTML("blogs/index.html"))
}

// BlogsCreate default implementation.
func BlogsCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	b := &models.Blog{}
	err := c.Bind(b)
	if err != nil {
		log.Print(err)
		return c.Redirect(301, "/")
	}

	err = tx.Create(b)
	if err != nil {
		log.Print(err)
		return c.Redirect(301, "/")
	}

	c.Set("blog", b)
	return c.Redirect(301, fmt.Sprintf("/blogs/%s", b.ID))
}
