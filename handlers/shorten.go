
// URL shortener API

package handlers

import (
    "fmt"
    "os"
    "net/http"
    "net/url"
    "errors"
    "encoding/json"
    "bytes"

    "github.com/aniiyengar/ani.bz/models"
    "github.com/aniiyengar/ani.bz/db"
)

type ShortenHandler struct {}

type verifyRecaptchaResponse struct {
    Success bool
}
type incomingShortenPostResponse struct {
    Link string
    Recaptcha string
}

var grecaptchaSecret = os.Getenv("ANI_BZ_GRECAPTCHA_SECRET")

func saveURL(linkTo string) (string, error) {
    link, err := db.WriteShortURL(models.ShortURL{ Link: linkTo })
    if err != nil {
        return "", err
    }

    return link, nil
}

func verifyRecaptcha(recaptcha string) error {
    urlEncoded := fmt.Sprintf(
        "response=%s&" +
        "secret=%s",
        recaptcha, grecaptchaSecret,
    )

    byteString := []byte(urlEncoded)

    resp, err := http.Post(
        "https://www.google.com/recaptcha/api/siteverify",
        "application/x-www-form-urlencoded",
        bytes.NewBuffer(byteString),
    )
    if err != nil {
        return err
    }

    var result verifyRecaptchaResponse
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return err
    }

    if result.Success != true {
        return errors.New("verifyRecaptcha: verification failed")
    }

    return nil
}

func (h ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" && r.URL.Path == "/s/" {
        var body incomingShortenPostResponse
        err := json.NewDecoder(r.Body).Decode(&body)
        if err != nil {
            w.WriteHeader(400)
            fmt.Fprintf(w, "Bad request.")
            return
        }

        linkTo := body.Link
        recaptcha := body.Recaptcha

        err := verifyRecaptcha(recaptcha)
        if err != nil {
            w.WriteHeader(400)
            fmt.Fprintf(w, "Recaptcha response was invalid.")
            return
        }

        u, validateErr := url.Parse(linkTo)
        if validateErr != nil {
            w.WriteHeader(400)
            fmt.Fprintf(w, "There was an error parsing the URL.")
            return
        } else if u.Scheme == "" || u.Host == "" {
            w.WriteHeader(400)
            fmt.Fprintf(w, "The URL must have a valid scheme and host.")
            return
        } else if !(u.Scheme == "http" || u.Scheme == "https") {
            w.WriteHeader(400)
            fmt.Fprintf(w, "The URL must be http(s)://.")
            return
        } else {
            result, err := saveURL(linkTo)
            if err != nil {
                w.WriteHeader(500)
                fmt.Fprintf(w, "Error creating new URL.")
                return
            }

            w.WriteHeader(200)
            fmt.Fprintf(w, result)
            return
        }
    } else {
        w.WriteHeader(404)
        fmt.Fprintf(w, "Page not found.")
    }
}
