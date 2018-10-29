package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Home Page",
		"data":  articles}, "index.html")

}

func getArticle(c *gin.Context) {
	// Проверим валидность ID
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Проверим существование топика
		if article, err := getArticleByID(articleID); err == nil {
			// Call the render function with the name of the template to render
			render(c, gin.H{
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

func showArticleCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func createArticle(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")
	if a, err := createNewArticle(title, content); err == nil {
		// If the article is created successfully, show success message
		render(c, gin.H{
			"title": "Submission Successful",
			"data":  a}, "submission-successful.html")
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
