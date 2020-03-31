package main

import (
	"log"
	"os"

	"github.com/asaskevich/govalidator"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"

	"github.com/velopaymentsapi/payor-example-go/internal/payor"
)

func main() {
	govalidator.SetFieldsRequiredByDefault(true)

	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	db, err := gorm.Open("postgres", os.Getenv("DB_URI"))
	if err != nil {
		panic(err)
	}

	r := payor.InitRoutes(db)

	payor.VeloOAuthRefresh()

	c := cron.New()
	c.AddFunc("* * * * *", payor.VeloOAuthRefresh)
	c.Start()

	r.Run(":" + port)
	log.Println("Listening on 0.0.0.0:" + port)
}
