package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewMux()
	mux.Get("/", tomatoHandler)
	mux.Get("/truth", tomatoTruthHandler)

	log.Println("It's tomato time! 🍅")

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		// local
		log.Fatal(http.ListenAndServe(":3000", mux))
	} else {
		// lambda
		lambda.Start(httpadapter.New(mux).ProxyWithContext)
	}
}

func tomatoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Tom", "Hey Tom")

	msg := `
		<head>
		  <title>Tomato Page</title>
		<link rel="icon" href="data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.9em%22 font-size=%2290%22>🍅</text></svg>">

		</head>

		<h1>Tomato</h1>
		<p>yay for tomatoes</p>
	    <a href="/truth">see truth</a>

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

func tomatoTruthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Tom", "Hey Tom")

	msg := `
		<head>
		  <title>Tomato Truth Page</title>
		<link rel="icon" href="data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.9em%22 font-size=%2290%22>🍅</text></svg>">

		</head>

		<h1>Tomato Truth</h1>
		<p>I'm not even that into Tomatoes</p>
	    <a href="/">home</a>

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
