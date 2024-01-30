package main

import "net/http"

func main() {
	srv := http.Server{
		Addr: ":80",
	}

	http.HandleFunc("/", helloWorld)

	srv.ListenAndServe()
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}
