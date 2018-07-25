
// Route to serve static css and images

package handlers

import (
    "net/http"
)

type StaticHandler struct {}

func (h StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    http.FileServer(http.Dir("./static/")).ServeHTTP(w, r)
}
