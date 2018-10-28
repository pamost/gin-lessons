package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

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
