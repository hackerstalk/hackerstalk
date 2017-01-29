package main

import (
  "github.com/gin-contrib/sessions"
)

func GetDefault(session sessions.Session, key string, defaultValue string) string {
  value := session.Get(key)
  if value == nil {
    return defaultValue
  } else {
    return value.(string)
  }
}
