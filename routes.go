package main

import (
	"gin-lessons/handlers"
	"gin-lessons/middleware"
)

func initRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(middleware.SetUserStatus())

	// определение роута главной страницы
	router.GET("/", handlers.ShowIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/user")
	{
		// Handle the GET requests at /user/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", middleware.EnsureNotLoggedIn(), handlers.ShowLoginPage)
		// Handle POST requests at /user/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)
		// Handle GET requests at /user/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", middleware.EnsureLoggedIn(), handlers.Logout)

		// Handle the GET requests at /user/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn(), handlers.ShowRegistrationPage)
		// Handle POST requests at /user/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn(), handlers.Register)
	}

	// Group article related routes together
	articleRoutes := router.Group("/article")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:article_id", handlers.GetArticle)
		// Handle the GET requests at /article/create
		// Show the article creation page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", middleware.EnsureLoggedIn(), handlers.ShowArticleCreationPage)
		// Handle POST requests at /article/create
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/create", middleware.EnsureLoggedIn(), handlers.CreateArticle)
	}

}
