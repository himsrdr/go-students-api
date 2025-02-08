package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/himsrdr/students-api/internal/config"
	student "github.com/himsrdr/students-api/internal/http/handlers/students"
	"github.com/himsrdr/students-api/internal/storage/sqlite"
)

func main() {
	cfg := config.Mustload()

	_, err := sqlite.NewSqlite(cfg)
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.Create())

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
