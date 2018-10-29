package main

func initRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// определение роута главной страницы
	router.GET("/", showIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/user")
	{
		// Handle the GET requests at /user/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		// Handle POST requests at /user/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		// Handle GET requests at /user/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", ensureLoggedIn(), logout)

		// Handle the GET requests at /user/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		// Handle POST requests at /user/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	// Group article related routes together
	articleRoutes := router.Group("/article")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:article_id", getArticle)
		// Handle the GET requests at /article/create
		// Show the article creation page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)
		// Handle POST requests at /article/create
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}

}
