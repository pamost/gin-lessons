package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// роутер по-умолчанию в Gin
	router = gin.Default()

	// Обработаем шаблоны вначале, так что их не нужно будет перечитывать
	// ещё раз. Из-за этого вывод HTML-страниц такой быстрый.
	router.LoadHTMLGlob("templates/*")

	// Инициализируем роуты
	initRoutes()

	// Запускаем приложение
	router.Run()

}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

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
