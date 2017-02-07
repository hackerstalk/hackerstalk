package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"

	"hackerstalk/auth"
	"hackerstalk/db"
	"hackerstalk/route"
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

	githubOauthClientId := os.Getenv("GITHUB_CLIENT_ID")
	if githubOauthClientId == "" {
		log.Print("GITHUB_CLIENT_ID is missing")
	}

	githubOauthClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	if githubOauthClientSecret == "" {
		log.Print("GITHUB_CLIENT_SECRET is missing")
	}

	// Github OAuth 초기화
	auth.InitGithubOAuth(githubOauthClientId, githubOauthClientSecret)

	// DB 연결
	err = db.Connect(dbUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// GIN Routing
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Debug에서만 log를 출력
	if gin.IsDebugging() {
		router.Use(gin.Logger())
	}

	// Session
	store := sessions.NewCookieStore([]byte("hack3rsTa!kS3cr2t"))
	cacheOptions := sessions.Options{
		Path:     "/",
		MaxAge:   0,
		Secure:   !gin.IsDebugging(),
		HttpOnly: true,
	}
	store.Options(cacheOptions)
	router.Use(sessions.Sessions("ht_session", store))

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/api/login", route.LoginPost)
	router.GET("/api/link", route.GetLinks)
	router.POST("/api/link/add", route.NewLink)
	router.GET("/auth/github", route.GithubAuth)
	router.GET("/auth/githubCallback", route.GithubAuthCallback)

	if gin.IsDebugging() {
		router.GET("/signin", route.DebugSignIn)
		router.POST("/signin", route.DebugSignInPost)
		router.GET("/debug/session", route.DebugSession)
	}

	router.StaticFile("favicon.ico", "./static/favicon.ico")
	router.Static("/static", "static")
	router.Run(":" + port)
}
