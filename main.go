package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway/v2"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", tomatoHandler)

	log.Println("It's tomato time! 🍅")

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		log.Fatal(http.ListenAndServe(":3000", mux))
	} else {
		log.Fatal(gateway.ListenAndServe(":3000", mux))
	}
}

func tomatoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hit function!")

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Tom", "Hey Tom")

	msg := `
		<head>
		  <title>Tomato Page</title>
		<link rel="icon" href="data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.9em%22 font-size=%2290%22>🍅</text></svg>">

		</head>

		<h1>Tomato</h1>
		<p>yay for tomatoes</p>

		<style>
		body {
		  color: darkgreen;
		  background-color:tomato;
		}
		</style>
`

	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Println(err)
	}
}
