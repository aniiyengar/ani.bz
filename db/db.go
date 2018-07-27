
package db

import (
    "fmt"
    "os"
    "database/sql"

    _ "github.com/lib/pq"

    "github.com/aniiyengar/ani.bz/models"
    "github.com/aniiyengar/ani.bz/utils"
)

var (
    host = os.Getenv("ANI_BZ_DB_HOST")
    user = os.Getenv("ANI_BZ_DB_USER")
    password = os.Getenv("ANI_BZ_DB_PASSWORD")
    database = os.Getenv("ANI_BZ_DB_NAME")

    port = 5432
)

func connect(fn func(d *sql.DB) error) error {
    connectInfo := fmt.Sprintf(
        "host=%s " +
        "port=%d " +
        "user=%s " +
        "password=%s " +
        "dbname=%s " +
        "sslmode=disable",
        host, port, user, password, database,
    )

    db, err := sql.Open("postgres", connectInfo)
    if err != nil {
        return err
    }

    defer db.Close()

    err = fn(db)
    if err != nil {
        return err
    }

    return nil
}

func WriteShortURL(s models.ShortURL) (string, error) {
    statement := "INSERT INTO urls (link) VALUES ($1) RETURNING id;"
    var newId int

    fn := func(db *sql.DB) error {
        err := db.QueryRow(statement, s.Link).Scan(&newId)
        if err != nil {
            return err
        }

        return nil
    }

    err := connect(fn)
    if err != nil {
        return "", err
    }

    slug, err := utils.IntToBase62(newId)
    if err != nil {
        return "", err
    }

    return "https://ani.bz/r/" + slug, nil
}

func QueryShortURL(slug string) (string, error) {
    statement := "SELECT link FROM urls WHERE id=$1"
    var link string

    fn := func(db *sql.DB) error {
        id, err := utils.Base62ToInt(slug)
        if err != nil {
            return err
        }

        err = db.QueryRow(statement, id).Scan(&link)
        if err != nil {
            return err
        }

        return nil
    }

    err := connect(fn)
    if err != nil {
        return "", err
    }

    return link, nil
}
