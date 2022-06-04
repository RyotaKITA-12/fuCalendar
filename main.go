package main

import (
	"fmt"

	"github.com/RyotaKITA-12/fuCalendar.git/app/controllers"
	"github.com/RyotaKITA-12/fuCalendar.git/app/models"
	// "log"
)

func main() {
    fmt.Println(models.Db)
    u := &models.User{}
    u.Name = "test"
    u.Email = "test@example.com"
    u.Password = "test"
    fmt.Println(u)

    u.CreateUser()

    controllers.StartMainServer()

}
