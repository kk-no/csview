package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kk-no/csview/csv"
	"github.com/kk-no/csview/handler"
	"github.com/kk-no/csview/template"
)

var port = "8888"

var f = flag.String("f", "", "Read file path")
var h = flag.Bool("h", true, "With or without csv header")

func main() {
	flag.Parse()

	log.Printf("Load [%s]\n", *f)
	log.Printf("Header exists flag [%t]", *h)

	rows, err := csv.Load(*f, *h)
	if err != nil {
		log.Fatalf("failed to load file: %s", err)
	}

	t, err := template.New()
	if err != nil {
		log.Fatalf("failed to parse template: %s", err)
	}

	http.Handle("/", handler.NewTemplateHandler(t, rows))

	log.Printf("listen serve on http://localhost:%s/", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("failed to listen server %s: %s", port, err)
	}
}
