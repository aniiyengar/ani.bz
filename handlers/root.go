
// Route for /. For now, just redirect to aniruddh.co

package handlers

import (
    "fmt"
    "net/http"
    "html/template"

    "github.com/aniiyengar/ani.bz/utils"
)

type RootHandler struct {}

type FlashMessage struct {
    Message string
}

func (h RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" && r.URL.Path == "/" {
        t, _ := template.ParseFiles("views/index.html")
        msg, err := utils.GetFlash(w, r, "flash")
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(string(msg))
        t.Execute(w, FlashMessage{ Message: string(msg) })
    } else {
        w.WriteHeader(404)
    }
}
