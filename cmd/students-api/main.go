package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/himsrdr/students-api/internal/config"
	student "github.com/himsrdr/students-api/internal/http/handlers/students"
	"github.com/himsrdr/students-api/internal/storage/db"
)

func main() {
	cfg := config.Mustload()

	storage, err := db.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.DB.Close()
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.Create(storage))
	router.HandleFunc("GET /api/students/{id}", student.Get(storage))
	router.HandleFunc("PUT /api/students/{id}", student.Put(storage))
	router.HandleFunc("DELETE /api/students/{id}", student.Delete(storage))

	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}
	fmt.Printf("server started %s ", cfg.HttpServer.Address)

	// wg := sync.WaitGroup()
	// wg.add
	// go func(){
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}
	//}()

}
