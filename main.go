
package main

import (
    "net/http"

    "github.com/aniiyengar/ani.bz/handlers"
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

func main() {
    http.Handle("/s/", cors(handlers.ShortenHandler{}))
    http.Handle("/r/", cors(handlers.UnshortenHandler{}))
    http.Handle("/", cors(handlers.RootHandler{}))

    http.Handle("/static/", cors(http.StripPrefix("/static/", handlers.StaticHandler{})))

    http.ListenAndServe(":9003", nil)
}
