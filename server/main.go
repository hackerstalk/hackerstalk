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

// HTTPS강제 처리. Heroku router에서 X-Forwarded-Proto 헤더로 http인 경우 https로 redirect.
func RedirectToHTTPS() gin.HandlerFunc {
	return func(c *gin.Context) {
		proto := c.Request.Header["X-Forwarded-Proto"]
		if len(proto) > 0 && proto[0] == "http" {
			// HTTPS로 Redirect
			c.Redirect(302, "https://"+c.Request.Host+c.Request.RequestURI)
		} else {
			c.Next()
		}

	}
}

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

	// Static Key를 환경변수로 받는다.
	// Cache-busting을 위해 static folder를 build마다 unique한 static key생성.
	staticKey := os.Getenv("STATIC_KEY")
	if staticKey == "" {
		staticKey = "static"
	}

	log.Printf("Static Key : %s\n", staticKey)

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
	router.Use(RedirectToHTTPS())

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
		c.HTML(http.StatusOK, "index.html", gin.H{
			"staticKey": staticKey,
		})
	})

	router.GET("/api/link", route.GetLinks)
	router.POST("/api/link", route.NewLink)

	// GitHub 로그인
	router.GET("/auth/github", route.GithubAuth)
	router.GET("/auth/githubCallback", route.GithubAuthCallback)

	router.StaticFile("favicon.ico", "./static/favicon.ico")
	router.Static(staticKey, "static")
	router.Run(":" + port)
}
