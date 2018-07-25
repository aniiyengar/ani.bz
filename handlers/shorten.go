
// URL shortener API

package handlers

import (
    "net/http"
    "net/url"

    "github.com/aniiyengar/ani.bz/models"
    "github.com/aniiyengar/ani.bz/utils"

    "github.com/globalsign/mgo"
)

type ShortenHandler struct {}

func (h ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" && r.URL.Path == "/s/" {
        r.ParseForm()

        linkTo := r.Form["link"][0]

        u, validateErr := url.Parse(linkTo)
        if validateErr != nil {
            utils.SetFlash(w, "flash", []byte("Error: not a valid URL."))
        } else if u.Scheme == "" || u.Host == "" {
            utils.SetFlash(w, "flash", []byte("Error: not a valid URL."))
        } else if !(u.Scheme == "http" || u.Scheme == "https") {
            utils.SetFlash(w, "flash", []byte("Error: not a valid URL."))
        } else {
            session, connectionErr := mgo.Dial("localhost")
            if connectionErr != nil {
                utils.SetFlash(w, "flash", []byte("Error: unable to connect to the DB."))
            }

            db := session.DB("urls")

            coll := db.C("url")
            err := coll.Insert(models.ShortURL{ Dest: r.Form["link"][0] })
            if err != nil {
                utils.SetFlash(w, "flash", []byte("Error: unable to create short URL."))
            }
        }

        http.Redirect(w, r, "/", 303)

    } else {
        w.WriteHeader(404)
    }
}
