package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = "5432"
	USER     = "postgres"
	PASSWORD = "admin"
	DB       = "think_assignment_db"
)

type Test struct {
	TestRes string `json:"test"`
}

func main() {
	psqlInfo := "host=" + HOST + " port=" + PORT + " user=" + USER + " password=" + PASSWORD + " dbname=" + DB + " sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	router.GET("/details/:username", func(c *gin.Context) {
		getDetails(c, db)
	})

	router.Run(":3000")

}
