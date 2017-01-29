package main

import (
  "github.com/jmoiron/sqlx"
  "github.com/lib/pq"
)

var DB *sqlx.DB

// DB 연결
func Connect(dbUrl string) error {
  var err error
  DB, err = sqlx.Connect("postgres", dbUrl)
  return err
}

// 새로운 user추가.
func NewUser(name string, githubId string) error {
  _, err := DB.Exec("insert into users(name, github_id) values($1, $2)", name, githubId)
  return err
}

// 새로운 link추가.
func NewLink(url string, tags []string, comment string, userId int) error {
  _, err := DB.Exec("insert into links(url, tags, comment, user_id) values($1, $2, $3, $4)", url, pq.Array(tags), comment, userId)
  return err
}
