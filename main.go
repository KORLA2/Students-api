package main

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/storage/sqlite"
	"myapp/student"
	"net/http"
)

// "github.com/go-playground/validator/v10"

func main() {
	var cfg = config.MustLoad()
	fmt.Println(cfg.Address)

	db, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("This sis the database stuff", db)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/student", student.New(db))

	server := &http.Server{
		Addr:    cfg.Address,
		Handler: mux,
	}

	server.ListenAndServe()

}
