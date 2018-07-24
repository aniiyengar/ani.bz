
package main

import (
    "fmt"
    "net/http"
)

// Deal with preflight requests accordingly
func cors(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "OPTIONS" {
            headers := w.Header()
            headers.Add("Access-Control-Allow-Origin", "*")
            headers.Add("Access-Control-Allow-Headers", "GET, POST, OPTIONS")
            headers.Add("Vary", "Origin")
            headers.Add("Vary", "Access-Control-Request-Headers")
        } else {
            h.ServeHTTP(w, r)
        }
    })
}

// For basic, non-shortening stuff
type BasicHandler struct {}

func (h BasicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        http.Redirect(w, r, "https://aniruddh.co", 302)
    } else {
        fmt.Fprintf(w, "Cannot %s %s", r.Method, r.URL.Path)
    }
}

func main() {
    http.Handle("/", cors(BasicHandler{}))
    http.ListenAndServe(":9003", nil)
}
