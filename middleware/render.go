package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func Render(c *gin.Context, data gin.H, templateName string) {

	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {

	case "application/json":
		c.JSON(http.StatusOK, data["data"]) // Respond with JSON

	case "application/xml":
		c.XML(http.StatusOK, data["data"]) // Respond with XML

	default:
		c.HTML(http.StatusOK, templateName, data) // Respond with HTML
	}

}
