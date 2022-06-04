package main

import (
	"fmt"

	"github.com/RyotaKITA-12/fuCalendar.git/app/models"
)

func main() {
    fmt.Println(models.Db)

    u := &models.User{}
    u.Name = "test"
    u.Email = "test@example.com"
    u.Password = "testtest"
    fmt.Println(u)
    //
    u.CreateUser()
}
