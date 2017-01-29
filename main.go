package main

import (
  "log"
  "net/http"
  "os"

  "github.com/gin-contrib/gzip"
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

  githubOauthClientId := os.Getenv("GITHUB_CLIENT_ID")
  if githubOauthClientId == "" {
    log.Print("GITHUB_CLIENT_ID is missing")
  }

  githubOauthClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
  if githubOauthClientSecret == "" {
    log.Print("GITHUB_CLIENT_SECRET is missing")
  }

  // Github OAuth 초기화
  initGithubOAuth(githubOauthClientId, githubOauthClientSecret)

  // DB 연결
  err = Connect(dbUrl)
  if err != nil {
    log.Fatalln(err)
  }

  // GIN Routing
  router := gin.New()
  router.Use(gin.Recovery())
  router.Use(gzip.Gzip(gzip.DefaultCompression))

  if gin.IsDebugging() {
    router.Use(gin.Logger())
  }

  router.LoadHTMLGlob("templates/*.html")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
  })

  router.POST("/api/login", func(c *gin.Context) {
    name := c.DefaultPostForm("name", "bbirec")
    githubId := c.DefaultPostForm("githubId", "bbirec")

    err := NewUser(name, githubId)
    if err != nil {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "msg":    err.Error(),
      })
      return
    }

    c.JSON(200, gin.H{
      "status": "OK",
    })
  })

  router.GET("/api/link", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "status": "OK",
    })
  })

  router.POST("/api/link", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "status": "OK",
    })
  })

  router.GET("/auth/github", func(c *gin.Context) {
    _, url := GetGithubAuthUrl()
    // TODO: save state to session and check it on callback
    c.Redirect(http.StatusMovedPermanently, url)
  })

  router.GET("/auth/githubCallback", func(c *gin.Context) {
    apiError := c.Query("error")
    apiErrorDescription  := c.Query("error_description")
    code := c.Query("code")

    if apiError != "" {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "error" : apiError,
        "msg":    apiErrorDescription,
      })
      return
    }
    // TODO: check c.Query("state") against session.state
    githubUser, err := GetGithubUser(code)
    if err != nil {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "msg":    err.Error(),
      })
      return
    }

    err = NewUser(*githubUser.Name, *githubUser.Login)
    if err != nil {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "msg":    err.Error(),
      })
      return
    }

    c.JSON(200, gin.H{
      "status": "OK",
    })
  })

  router.Static("/static", "static")
  router.Run(":" + port)
}
