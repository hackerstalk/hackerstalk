package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	driver "github.com/lib/pq"
	"time"
)

var DB *sqlx.DB

// DB 연결
func Connect(dbUrl string) error {
	var err error
	DB, err = sqlx.Connect("postgres", dbUrl)
	return err
}

type User struct {
	Id          int             `db:"id"`
	Name        string          `db:"name"`
	GithubId    sql.NullString  `db:"github_id"`
	EditedTime  driver.NullTime `db:"edited_time"`
	CreatedTime time.Time       `db:"created_time"`
}

type Link struct {
	Id          int                `db:"id" json:"id"`
	Url         string             `db:"url" json:"url"`
	Tags        driver.StringArray `db:"tags" json:"tags"`
	Comment     string             `db:"comment" json:"comment"`
	UserId      int                `db:"user_id" json:"user_id"`
	EditedTime  driver.NullTime    `db:"edited_time" json:"edited_time"`
	CreatedTime time.Time          `db:"created_time" json:"created_time"`
}

// 새로운 user추가.
func UpsertUser(name string, githubId string) error {
	_, err := DB.Exec("insert into users(name, github_id) values($1, $2) on conflict (github_id) do update set name = EXCLUDED.name", name, githubId)
	return err
}

func GetUserByGithubId(githubId string) (*User, error) {
	user := User{}
	err := DB.Get(&user, "select * from users where github_id=$1 LIMIT 1", githubId)
	return &user, err
}

func GetLinks() ([]Link, error) {
	links := []Link{}
	err := DB.Select(&links, "select * from links")
	if err != nil {
		return nil, err
	}

	return links, nil
}

// 새로운 link추가.
func NewLink(url string, tags []string, comment string, userId int) error {
	_, err := DB.Exec("insert into links(url, tags, comment, user_id) values($1, $2, $3, $4)", url, driver.Array(tags), comment, userId)
	return err
}
