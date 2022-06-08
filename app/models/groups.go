package models

import (
	"log"
	"time"
)

type Group struct {
	ID        int
	Name      string
	HostID    int
	Gests     []User
	CreatedAt time.Time
}

func (u *User) CreateGroup(name string, gests []User) (err error) {
	cmd := `INSERT INTO groups (
                name,
                host_id,
                gest_id,
                created_at)
            VALUES ($1, $2, $3, $4)`
    for _, gest := range gests {
        _, err = Db.Exec(cmd, name, u.ID, gest.ID, time.Now())
        if err != nil {
            log.Fatalln(err)
        }
    }
	return err
}
