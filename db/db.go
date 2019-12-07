
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
    host = os.Getenv("POSTGRES_HOST")
    user = os.Getenv("POSTGRES_USER")
    password = os.Getenv("POSTGRES_PASSWORD")
    database = os.Getenv("POSTGRES_DB")

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

func Init() error {
    create_url := "CREATE TABLE IF NOT EXISTS urls (link VARCHAR, id SERIAL);"
    fn := func(db *sql.DB) error {
        _, err := db.Query(create_url)
        return err
    }

    return connect(fn)
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

    slug, err := utils.IntToBase62(utils.MagicHashForward(uint64(newId)))
    if err != nil {
        return "", err
    }

    return "https://ani.bz/" + slug, nil
}

func QueryShortURL(slug string) (string, error) {
    statement := "SELECT link FROM urls WHERE id=$1"
    var link string

    fn := func(db *sql.DB) error {
        id, err := utils.Base62ToInt(slug)
        if err != nil {
            return err
        }

        err = db.QueryRow(
            statement,
            utils.MagicHashBackward(uint64(id)),
        ).Scan(&link)

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
