package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kk-no/csview/csv"
	"github.com/kk-no/csview/executor"
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

	if _, err := csv.Load(*f, *h); err != nil {
		log.Fatalf("failed to load file: %s", err)
	}

	t, err := template.New()
	if err != nil {
		log.Fatalf("failed to parse template: %s", err)
	}

	e := executor.NewCSVExecutor(t)

	http.Handle("/", handler.NewTemplateHandler(e))
	http.Handle("/query", handler.NewQueryHandler(e))

	log.Printf("listen serve on http://localhost:%s/", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("failed to listen server %s: %s", port, err)
	}
}
