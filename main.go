package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const name = "partysvc"

// These variables are assigned values during the build process using the -ldflags="-X ..." linker option.
var version = "Not set"
var origin = "Not set"
var commit = "Not set"
var branch = "Not set"
var built = "Not set"

type info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Origin  string `json:"origin"`
	Commit  string `json:"commit"`
	Branch  string `json:"branch"`
	Built   string `json:"built"`
}

type sampleLink struct {
	CollectionExerciseID string `json:"collectionExerciseId"`
	SampleSummaryID      string `json:"sampleSummaryId"`
}

func main() {
	port, overridden := os.LookupEnv("PORT")
	if !overridden {
		port = ":8081"
	}

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		json, _ := json.Marshal(info{Name: name, Version: version, Origin: origin, Commit: commit, Branch: branch, Built: built})
		fmt.Fprintf(w, "%s\n", json)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json, _ := json.Marshal(sampleLink{CollectionExerciseID: "00000000-0000-0000-0000-000000000000", SampleSummaryID: "00000000-0000-0000-0000-000000000000"})
		fmt.Fprintf(w, "%s\n", json)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
