package main

func initRoutes() {

	// определение роута главной страницы
	router.GET("/", showIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/user")
	{
		// Handle the GET requests at /user/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", showLoginPage)
		// Handle POST requests at /user/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", performLogin)
		// Handle GET requests at /user/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", logout)

		// Handle the GET requests at /user/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", showRegistrationPage)
		// Handle POST requests at /user/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", register)
	}

	// Group article related routes together
	articleRoutes := router.Group("/article")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:article_id", getArticle)
		// Handle the GET requests at /article/create
		// Show the article creation page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", showArticleCreationPage)
		// Handle POST requests at /article/create
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/create", createArticle)
	}

}
