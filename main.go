package main

import (
	"fmt"
	"time"

	// "github.com/RyotaKITA-12/fuCalendar.git/app/controllers"
	"github.com/RyotaKITA-12/fuCalendar.git/app/models"
	// "log"
)

func main() {
	fmt.Println(models.Db)

	// controllers.StartMainServer()

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.Password = "testtest"
	fmt.Println(u)

	u.CreateUser()

    u.CreateEvent("test-content", "渋谷", time.Now(), time.Now().AddDate(0, 0, 1), 1)

    e, _ := u.GetEventsByUser()
    fmt.Println(e)
}
