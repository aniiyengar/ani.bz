
package main

import (
    "net/http"
    "fmt"
    "time"
    "math"

    "github.com/aniiyengar/ani.bz/db"
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
    retries := 0
    var err error

    http.Handle("/", cors(handlers.UnshortenHandler{}))
    http.Handle("/s/", cors(handlers.ShortenHandler{}))

    for retries < 10 {
        if err = db.Init(); err != nil {
            fmt.Println(err);
            retries += 1
            time.Sleep(time.Duration(math.Pow(1.5, float64(retries))) * time.Second)
        } else {
            http.ListenAndServe(":9003", nil)
        }
    }
}
