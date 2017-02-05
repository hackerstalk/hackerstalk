package main

import (
  "github.com/jmoiron/sqlx"
  "github.com/lib/pq"
  "database/sql"
  "time"
)

var DB *sqlx.DB

// DB 연결
func Connect(dbUrl string) error {
  var err error
  DB, err = sqlx.Connect("postgres", dbUrl)
  return err
}

// users table
type User struct {
  Id          int             `db:"id"`
  Name        string          `db:"name"`
  GithubId    sql.NullString  `db:"github_id"`
  EditedTime  pq.NullTime     `db:"edited_time"`
  CreatedTime time.Time       `db:"created_time"`
}

// 새로운 user추가.
func UpsertUser(name string, githubId string) error {
  _, err := DB.Exec("insert into users(name, github_id) values($1, $2) on conflict (github_id) do update set name = EXCLUDED.name", name, githubId)
  return err
}

func GetUserByGithubId(githubId string) (User,error) {
  user := User{}
  err := DB.Get(&user, "select * from users where github_id=$1 LIMIT 1", githubId)
  return user, err
}

// 새로운 link추가.
func NewLink(url string, tags []string, comment string, userId int) error {
  _, err := DB.Exec("insert into links(url, tags, comment, user_id) values($1, $2, $3, $4)", url, pq.Array(tags), comment, userId)
  return err
}
