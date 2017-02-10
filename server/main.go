package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"

	"hackerstalk/server/auth"
	"hackerstalk/server/db"
	"hackerstalk/server/route"
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
		githubOauthClientId = "b06f9141609acfe076cc"
	}

	githubOauthClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	if githubOauthClientSecret == "" {
		githubOauthClientSecret = "c5b40b69fa4796673418d0c3f26806b3b5533b36"
	}

	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		sessionSecret = "hack3rsTa!kS3cr2t"
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
	store := sessions.NewCookieStore([]byte(sessionSecret))
	cacheOptions := sessions.Options{
		Path:     "/",
		MaxAge:   0,
		Secure:   false,
		HttpOnly: true,
	}
	store.Options(cacheOptions)
	router.Use(sessions.Sessions("ht_session", store))

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/api/link", route.GetLinks)
	router.POST("/api/link", route.NewLink)

	// GitHub 로그인
	router.GET("/auth/github", route.GithubAuth)
	router.GET("/auth/githubCallback", route.GithubAuthCallback)

	router.StaticFile("favicon.ico", "./static/favicon.ico")
	router.Static("/static", "static")
	router.Run(":" + port)
}
