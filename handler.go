package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Document(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getDocuments(w, r)
	case "POST":
		postDocument(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getDocuments(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ここにデータベースから情報を取得して返すコードを書く
	fmt.Fprintf(w, "get document")
}

func postDocument(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ここにデータベースに情報を登録するコードを書く
	fmt.Fprintf(w, "create document")
}

func JsonDocument(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getJsonDocuments(w, r)
	case "POST":
		postJsonDocument(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

type Documents struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	Title string `json:"title"`
}

func getJsonDocuments(w http.ResponseWriter, r *http.Request) {
	sample := &Documents{
		Articles: []Article{
			{Title: "title1"},
		},
	}

	output, _ := json.Marshal(sample)
	fmt.Println(string(output))
	fmt.Fprintf(w, string(output))
}

func postJsonDocument(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	article := Article{}
	if err := json.Unmarshal(body, &article); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", article)

	output, _ := json.Marshal(article)
	fmt.Fprintf(w, string(output))
}
