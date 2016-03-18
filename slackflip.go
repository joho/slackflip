package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		portNumber = "5000"
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

		fmt.Fprintf(w, "(╯°□°）╯︵ ┻━┻")
	}

	http.HandleFunc("/", handler)

	log.Printf("(╯°□°）╯︵ ┻━┻ on port %v", portNumber)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", portNumber), nil))
}
