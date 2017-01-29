package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	err := NewUser("bbirec", "bbirec")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("OK"))
}

func NewLinkHandler(w http.ResponseWriter, r *http.Request) {
	err := NewLink("http://google.com", []string{"test"}, "구글신", 1)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("OK"))
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

	// DB 연결
	err = Connect(dbUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// Routing
	r := mux.NewRouter()
	r.HandleFunc("/api/new-user", NewUserHandler)
	r.HandleFunc("/api/new-link", NewLinkHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", r)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
