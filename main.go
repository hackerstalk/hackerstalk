package main

import (
  "log"
  "net/http"
  "os"
  "time"
  "errors"

  "github.com/gin-contrib/gzip"
  "github.com/gin-contrib/sessions"
  "gopkg.in/gin-gonic/gin.v1"
)

type NewLinkForm struct {
  Url      string    `form:"url" json:"url" binding:"required"`
  Tags     []string  `form:"tags" json:"tags"`
  Comment  string    `form:"comment" json:"comment"`
}

func setLoginSession(c *gin.Context, session sessions.Session, userId int, userName string) {
  session.Set("userId", userId)
  session.Set("salt", time.Now().Unix())
  c.SetCookie("name", userName, 0, "/", "", !gin.IsDebugging(), false)
}

func getUserIdFromSession(session sessions.Session) (int, error) {
  value := session.Get("userId")
  if value == nil {
    return -1, errors.New("Not logged in")
  } else {
    return value.(int), nil
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

  // Session
  store := sessions.NewCookieStore([]byte("hack3rsTa!kS3cr2t"))
  cacheOptions := sessions.Options{
    Path: "/",
    MaxAge: 0,
    Secure: !gin.IsDebugging(),
    HttpOnly: true,
  }
  store.Options(cacheOptions)
  router.Use(sessions.Sessions("ht_session", store))

  router.LoadHTMLGlob("templates/*.html")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
  })

  router.POST("/api/login", func(c *gin.Context) {
    githubId := c.PostForm("githubId")

    user, err := GetUserByGithubId(githubId)
    if err != nil {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "msg":    err.Error(),
      })
      return
    }

    session := sessions.Default(c)
    setLoginSession(c, session, user.Id, user.Name)
    err = session.Save()
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

  router.POST("/api/link/add", func(c *gin.Context) {
    session := sessions.Default(c)
    userId, err := getUserIdFromSession(session)
    if err != nil {
      c.JSON(401, gin.H{
        "status": "FAIL",
        "msg":    err.Error(),
      })
      return
    }

    var form NewLinkForm
    if c.Bind(&form) == nil {
      err := NewLink(form.Url, form.Tags, form.Comment, userId)
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
    } else {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "msg":    "Bind failed??",
      })
    }
  })

  router.GET("/auth/github", func(c *gin.Context) {
    session := sessions.Default(c)
    state, url := GetGithubAuthUrl()
    session.Set("state", state)
    err := session.Save()
    if err != nil {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "msg":    err.Error(),
      })
      return
    }
    c.Redirect(http.StatusFound, url)
  })

  router.GET("/auth/githubCallback", func(c *gin.Context) {
    session := sessions.Default(c)
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

    savedState := GetDefault(session, "state", "")
    if savedState != c.Query("state") {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "msg":    "state mismatches.",
      })
      return
    }

    githubUser, err := GetGithubUser(code)
    if err != nil {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "msg":    err.Error(),
      })
      return
    }

    userName := *githubUser.Name
    githubId := *githubUser.Login
    err = UpsertUser(userName, githubId)
    if err != nil {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "msg":    err.Error(),
      })
      return
    }
    session.Delete("state")

    var user *User;
    user, err = GetUserByGithubId(githubId)
    if err != nil {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "msg":    err.Error(),
      })
      return
    }
    setLoginSession(c, session, user.Id, userName)
    err = session.Save()
    if err != nil {
      c.JSON(500, gin.H{
        "status": "FAIL",
        "msg":    err.Error(),
      })
      return
    }

    c.Redirect(http.StatusFound, "/")
  })

  if gin.IsDebugging() {
    router.GET("/signin", func(c *gin.Context) {
      c.HTML(http.StatusOK, "signin.html", nil)
    })

    router.POST("/signin", func(c *gin.Context) {
      name := c.PostForm("name")
      githubId := c.PostForm("githubId")

      if name == "" || githubId == "" {
        c.JSON(500, gin.H{
          "status": "FAIL",
          "msg":    "name of githubId is missing",
        })
        return
      }

      err := UpsertUser(name, githubId)
      if err != nil {
        c.JSON(500, gin.H{
          "status": "FAIL",
          "msg":    err.Error(),
        })
        return
      }
      c.Redirect(http.StatusFound, "/");
    })

    router.GET("/debug/session", func(c *gin.Context) {
      session := sessions.Default(c)
      name := GetDefault(session, "name", "")
      githubId := GetDefault(session, "githubId", "")
      c.JSON(200, gin.H{
        "name": name,
        "githubId": githubId,
      })
    })
  }

  router.StaticFile("favicon.ico", "./static/favicon.ico")
  router.Static("/static", "static")
  router.Run(":" + port)
}
