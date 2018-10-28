package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Вызовем метод HTML из Контекста Gin для обработки шаблона
	c.HTML(
		// Зададим HTTP статус 200 (OK)
		http.StatusOK,
		// Используем шаблон index.html
		"index.html",
		// Передадим данные в шаблон
		gin.H{
			"title":    "Home Page",
			"articles": articles,
		},
	)

}

func getArticle(c *gin.Context) {
	// Проверим валидность ID
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Проверим существование топика
		if article, err := getArticleByID(articleID); err == nil {
			// Вызовем метод HTML из Контекста Gin для обработки шаблона
			c.HTML(
				// Зададим HTTP статус 200 (OK)
				http.StatusOK,
				// Используем шаблон index.html
				"article.html",
				// Передадим данные в шаблон
				gin.H{
					"title":    article.Title,
					"article": article,
				},
			)

		} else {
			// Если топика нет, прервём с ошибкой
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// При некорректном ID в URL, прервём с ошибкой
		c.AbortWithStatus(http.StatusNotFound)
	}
}
