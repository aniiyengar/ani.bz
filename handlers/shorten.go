
// URL shortener API

package handlers

import (
    "net/http"
)

type ShortenHandler struct {}

func (h ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(501) // Not implemented yet
}
