package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	Url := strings.Split(r.URL.String(), ".html")
	fileText, err := ioutil.ReadFile("./views" + Url[0] + ".html")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var myHtml string
	if fileText != nil {
		myHtml = string(fileText)
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
