
// URL lookup and redirect

package handlers

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/aniiyengar/ani.bz/db"
)

type UnshortenHandler struct {}

func lookupURL(slug string) (string, error) {
    link, err := db.QueryShortURL(slug)
    if err != nil {
        return "", err
    }

    return link, nil
}

func (h UnshortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        slug := strings.TrimPrefix(r.URL.Path, "/r/")

        if len(slug) == 0 {
            http.Redirect(w, r, "https://ani.bz/", 302)
            return
        }

        link, err := lookupURL(slug)
        if err != nil {
            http.Redirect(w, r, "https://ani.bz/", 302)
            return
        }

        http.Redirect(w, r, link, 302)
        return
    } else {
        w.WriteHeader(404)
        fmt.Fprintf(w, "Page not found.")
    }
}
