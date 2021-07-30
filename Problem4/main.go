package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	Url := strings.Split(r.URL.String(), ".html")
	fileText := make(chan string)
	go func() {
		data, err := ioutil.ReadFile("./views" + Url[0] + ".html")
		if err != nil {
			w.WriteHeader(404)
		}
		fileText <- string(data)
		close(fileText)
	}()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var myHtml string
	if fileText != nil {
		myHtml = string(<-fileText)
	}
	io.WriteString(w, myHtml)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	handleRequests()
}
