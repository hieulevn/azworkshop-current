package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	const appName = "demo"
	const appVersion = "0.1.0"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		respHtml := fmt.Sprintf(`[%s v%s] Hello, %q<br>Now is: %s`, appName, appVersion,
			html.EscapeString(r.URL.Path), time.Now().Format(time.RFC3339))

		respHtml += "<br><br><b>Env:</b><br>"
		respHtml += "<pre>"
		for _, e := range os.Environ() {
			respHtml += e + "\n"
		}
		respHtml += "</pre>"

		respHtml += "<br><br><b>Headers:</b><br>"
		respHtml += "<pre>"
		for k, v := range r.Header {
			respHtml += fmt.Sprintf("%s=%s\n", k, v)
		}
		respHtml += "</pre>"

		fmt.Fprintf(w, respHtml)
	})

	listenPort := 8080
	log.Printf("Running application on port %d\n", listenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil))
}
