package main

import (
	"fmt"
	"net/http"
)

type params struct { //Allocates a memory space in the local memory

}

func main() {

	//custom ServeMux router
	router := http.NewServeMux()

	//handles root functionality GET /
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Hello"))
	})

	//handles POST functionality to shorten URL
	router.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is POST /post endpoint"))
	})

	//create custom server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Listening on port: 8080")
	server.ListenAndServe()
}
