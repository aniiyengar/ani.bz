
// URL lookup and redirect

package handlers

import (
    "net/http"
)

type UnshortenHandler struct {}

func (h UnshortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(501) // Not implemented yet
}
