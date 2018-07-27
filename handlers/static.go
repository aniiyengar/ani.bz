
// Route to serve static css and images

package handlers

import (
    // "fmt"
    "net/http"
)

type StaticHandler struct {}

func (h StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    http.FileServer(http.Dir("./client/")).ServeHTTP(w, r)
}
