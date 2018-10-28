package main

func initRoutes() {

	// определение роута главной страницы
	router.GET("/", showIndexPage)

	// Обработчик GET-запросов на /article/view/некоторый_article_id
	router.GET("/article/view/:article_id", getArticle)
}
