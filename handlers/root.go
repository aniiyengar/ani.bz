
// Route for /. For now, just redirect to aniruddh.co

package handlers

import (
    "net/http"
)

type RootHandler struct {}

func (h RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" && r.URL.Path == "/" {
        http.Redirect(w, r, "https://aniruddh.co", 302)
    } else {
        w.WriteHeader(404)
    }
}
