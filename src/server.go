// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"text/template"

	"github.com/Naoji1984/go-rana8/config"
)

func main() {
	showInitialMessage()
	config.LoadConfig()

	http.HandleFunc("/", handlerPage)

	// Determine port for HTTP service.
	port := config.Config.Port

	log.Println(config.Config.RDB.Host)

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handlerPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/root.html")
	if err != nil {
		fmt.Printf("%w", err)
	}
	t.Execute(w, "Hello Rana8")
}

func showInitialMessage() {
	log.Println("*****************************************************")
	log.Println("                  Running on Rana8                   ")
	log.Println("")
	log.Printf("Go version: %s\n", runtime.Version())
	log.Printf("Number of cpu: %d\n", runtime.NumCPU())
	log.Println("*****************************************************")
}
