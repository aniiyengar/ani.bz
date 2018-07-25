
package utils

import (
    "encoding/base64"
    "net/http"
    "time"
)

func strtob(str string) ([]byte, error) {
    return base64.URLEncoding.DecodeString(str)
}

func btostr(b []byte) string {
    return base64.URLEncoding.EncodeToString(b)
}

//https://www.alexedwards.net/blog/simple-flash-messages-in-golang
func SetFlash(w http.ResponseWriter, name string, value []byte) {
    c := &http.Cookie{ Name: name, Value: btostr(value) }
    http.SetCookie(w, c)
}

func GetFlash(w http.ResponseWriter, r *http.Request, name string) ([]byte, error) {
    c, err := r.Cookie(name)
    if err != nil {
        switch err {
        case http.ErrNoCookie:
            return nil, nil
        default:
            return nil, err
        }
    }

    value, err := strtob(c.Value)
    if err != nil {
        return nil, err
    }

    dc := &http.Cookie{ Name: name, MaxAge: -1, Expires: time.Unix(1, 0) }
    http.SetCookie(w, dc)
    return value, nil
}
