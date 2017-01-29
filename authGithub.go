package main

import (
  "time"
  "strconv"
  "github.com/google/uuid"
  "golang.org/x/oauth2"
  "github.com/google/go-github/github"
)

var githubConf *oauth2.Config;

func initGithubOAuth(clientId string, clientSecret string) {
  // https://developer.github.com/v3/oauth/#scopes
  githubConf = &oauth2.Config{
    ClientID: clientId,
    ClientSecret: clientSecret,
    Scopes: nil,
    Endpoint: oauth2.Endpoint{
      AuthURL: "https://github.com/login/oauth/authorize",
      TokenURL: "https://github.com/login/oauth/access_token",
    },
  };
}

func GetRandomString() string {
  if genUuid, err := uuid.NewRandom(); err == nil {
    return genUuid.String();
  } else {
    return strconv.FormatInt(time.Now().Unix(), 16);
  }
}

func GetGithubAuthUrl() (string, string) {
  state := GetRandomString()
  url := githubConf.AuthCodeURL(state, oauth2.AccessTypeOffline)
  return state, url;
}

func GetGithubUser(code string) (*github.User, error) {
  tok, err := githubConf.Exchange(oauth2.NoContext, code)
  if err != nil {
    return nil, err
  }

  tc := githubConf.Client(oauth2.NoContext, tok)
  client := github.NewClient(tc)
  user, _, err := client.Users.Get("")
  if err != nil {
    return nil, err
  }
  return user, nil
}
