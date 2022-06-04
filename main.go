package main

import (
	"fmt"

	"github.com/RyotaKITA-12/fuCalendar.git/app/controllers"
	"github.com/RyotaKITA-12/fuCalendar.git/app/models"
	// "log"
)

func main() {
    fmt.Println(models.Db)

    controllers.StartMainServer()

}
