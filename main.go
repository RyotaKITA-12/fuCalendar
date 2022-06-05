package main

import (
	"fmt"
	"time"

	"github.com/RyotaKITA-12/fuCalendar.git/app/controllers"
	"github.com/RyotaKITA-12/fuCalendar.git/app/models"
	// "log"
)

func main() {
	fmt.Println(models.Db)
    u := &models.User{}
    u.Name = "test_001"
    u.Email = "test_001@example.com"
    u.Password = "test_001"
    fmt.Println(u)
    u.CreateUser()
    u.CreateEvent("test-content", "渋谷", time.Now(), time.Now().AddDate(0, 0, 1), 1)
	controllers.StartMainServer()
    // e, _ := u.GetEventsByUser()

