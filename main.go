package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewMux()
	mux.Get("/", tomatoHandler)
	mux.Get("/truth", tomatoTruthHandler)
	mux.Get("/mystery", tomatoMysteryHandler)

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
	logRequest(r)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Tom", "Hey Tom")

	msg := `
		<head>
		  <title>Tomato Page</title>
     	  <meta charset="UTF-8">
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
	logRequest(r)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Tom", "Hey Tom")

	msg := `
		<head>
		  <title>Tomato Truth Page</title>
     	  <meta charset="UTF-8">
		<link rel="icon" href="data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.9em%22 font-size=%2290%22>🍅</text></svg>">

		</head>

		<h1>Tomato Truth</h1>
		<p>I'm not even that into Tomatoes</p>
	    <a href="/">home</a></br>
	    <a href="/mystery">mystery</a>

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

func tomatoMysteryHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Tom", "Hey Tom")

	count := rand.Intn(100) + 1

	tomatoes := "<p>" + strings.Repeat("🍅", count) + "</p>"

	msg := `
		<head>
		  <title>Tomato Mystery Page</title>
     	  <meta charset="UTF-8">
		<link rel="icon" href="data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.9em%22 font-size=%2290%22>🍅</text></svg>">

		</head>

		<style>
		body {
		  color: darkgreen;
		  background-color:tomato;
		}
		</style>

		<h1>Tomato Mystery</h1>
		<p>This page shows a random number of tomatoes.</p>
	    <a href="/">home</a>

`

	_, err := w.Write([]byte(msg + tomatoes))
	if err != nil {
		log.Println(err)
	}
}

func logRequest(r *http.Request) {
	log.Println("LOGGING REQUEST")
	log.Println(r)
	log.Println(r.URL)
	log.Println(r.RequestURI)
	log.Println(r.URL.Query())
}
