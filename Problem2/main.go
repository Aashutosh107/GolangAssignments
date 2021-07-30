package main

import (
	"io"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	myHtml := `
           <!DOCTYPE html>
           <html>
           <head>
           </head>
           <body>
              <h1>`
	if queries["message"] != nil {
		myHtml += (queries["message"][0])
	} else {
		myHtml += ("Hello World!!!")
	}

	myHtml += (`
    </h1>
    </body>
    </html>
    `)
	io.WriteString(w, myHtml)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	handleRequests()
}
