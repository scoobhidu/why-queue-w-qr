package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	bloomfilter_database "why-queue-w-qr/bloomfilter/database"
	bloomfilter_service "why-queue-w-qr/bloomfilter/service"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env data: %s", err.Error())
	}
	// TODO: setup bloomfilter Cache
}

func main() {
	postgreConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"),
	)

	studentsDB, err := sql.Open("postgres", postgreConn)
	if err != nil {
		log.Fatalf("Error Connecting to the students DB: %s", err.Error())
	}

	defer func(studentsDB *sql.DB) {
		err := studentsDB.Close()
		if err != nil {
			log.Fatalf("Error closing the studentsDB: %s", err.Error())
		}
	}(studentsDB)

	erNoCache := bloomfilter_database.BloomFilterCache{}
	erNoCache.OnBits = 0
	erNoCache.BitsArray = make([]byte, 100)

	// TODO get all the enrolment no from the table cache it in the bloomfilter and then start the api this will help in creating a new cache
	// ? to refresh the cache in case of a new batch simply stop the server and restart it
	var rehash = true
	rows, err := studentsDB.Query(`SELECT "enrolmentNo" FROM "students"`)
	if err != nil {
		log.Fatalf("Error getting enrolment no from students Table: %s", err.Error())
	}
	defer rows.Close()

	bloomfilter_service.SetCacheBits(rows, &erNoCache, &rehash)

	fmt.Println("Completed setting values in BloomFilter Cache\n")
	//
	//service.CheckInCache(checkArray, &cache, &database)

	app := fiber.New()
	v1 := app.Group("/v1")
	setupRoutes(v1)

	log.Fatal(app.Listen("127.0.0.1:9010"))
}
