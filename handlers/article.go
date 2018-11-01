package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-lessons/middleware"
	"gin-lessons/models"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the middleware.Render function with the name of the template to middleware.Render
	middleware.Render(c, gin.H{
		"title": "Home Page",
		"data":  articles}, "index.html")

}

func GetArticle(c *gin.Context) {
	// Проверим валидность ID
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Проверим существование топика
		if article, err := models.GetArticleByID(articleID); err == nil {
			// Call the middleware.Render function with the name of the template to middleware.Render
			middleware.Render(c, gin.H{
				"title": article.Title,
				"data":  article}, "article.html")

		} else {
			// Если топика нет, прервём с ошибкой
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// При некорректном ID в URL, прервём с ошибкой
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func ShowArticleCreationPage(c *gin.Context) {
	// Call the middleware.Render function with the name of the template to middleware.Render
	middleware.Render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func CreateArticle(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")
	if a, err := models.CreateNewArticle(title, content); err == nil {
		// If the article is created successfully, show success message
		middleware.Render(c, gin.H{
			"title": "Submission Successful",
			"data":  a}, "submission-successful.html")
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
