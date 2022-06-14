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

func (u *User) GetGroupByName(name string) (group Group, err error) {
    group = Group{}
    cmd := `SELECT id, name, host_id, gest_id, created_at
            FROM groups
            WHERE name = $1 AND host_id = $2`
    err = Db.QueryRow(cmd, name, u.ID).Scan(
        &group.ID,
        &group.Name,
        &group.ID,
        &group.ID,
        &group.ID,
        &group.ID,
    )

}
