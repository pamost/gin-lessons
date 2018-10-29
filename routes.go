package main

func initRoutes() {

	// определение роута главной страницы
	router.GET("/", showIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/user")
	{
		// Handle the GET requests at /user/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", showRegistrationPage)
		// Handle POST requests at /user/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", register)
	}

	// Обработчик GET-запросов на /article/view/некоторый_article_id
	router.GET("/article/view/:article_id", getArticle)
}
