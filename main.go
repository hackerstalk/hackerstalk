package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	var err error

	// 환경변수에서 DB, PORT정보 가져옴
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgresql://localhost/ht?sslmode=disable"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// DB 연결
	err = Connect(dbUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// GIN Routing
	router := gin.Default()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/api/login", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "OK",
			"message": message,
			"nick":    nick,
		})
	})

	router.GET("/api/link", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	router.POST("/api/link", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "OK",
			"message": message,
			"nick":    nick,
		})
	})

	router.Static("/static", "static")
	router.Run(":" + port)
}
