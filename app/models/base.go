package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/RyotaKITA-12/Fu-calendar.git/config"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

var Db *sql.DB

var err error

// const (
// 	tableNameUser    = "users"
//     tableNameTodo    = "todos"
//     tableNameSession = "sessions"
// )

func init() {
    url := os.Getenv("DATABASE_URL")
    connection, _ := pq.ParseURL(url)
    // connection += "sslmode=disable"
    connection += "sslmode=require"
    Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}
}

func createUUID() (uuidobj uuid.UUID) {
    uuidobj, _ = uuid.NewUUID()
    return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
    cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
    return cryptext
}
