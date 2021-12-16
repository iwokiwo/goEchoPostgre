package db

import (
	//"database/sql"
	"fmt"
	"log"
	// "os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
	//	"log"
	"github.com/joho/godotenv"
	//	"github.com/pressly/goose"
)

const (
	host     = "147.139.139.7"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "yespos"
)

//var DB *sql.DB
var db *gorm.DB
var err error

func Init() {
	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbUsername := os.Getenv("DB_USERNAME")
	// dbPass := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// fmt.Println(host)	
	//------------------------------MYSQL----------------------------
	//conf := config.GetConfig()
	//connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	//db, err = sql.Open("mysql", connectionString)
	//---------------------------------------------------------------

	//------------------------------Postgresql-----------------------
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
//	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPass, dbName)
	fmt.Println(psqlconn)
	db, err = gorm.Open("postgres", psqlconn)
	//---------------------------------------------------------------

	//psqlconn2 := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//DB, err = sql.Open("postgres", psqlconn2)

	if err != nil {
		panic("connectionString error..")
	}

	// 	err = db.Ping()
	// 	if err != nil {
	// 		panic("DSN Invalid")
	// 	}
	// }

	// func CreateCon() *sql.DB {
	// 	return db
	// }
	//db.AutoMigrate(&models.Users{})
	//db.AutoMigrate(&models.Products{})

	//log.Printf("Start reloading database \n")
	//err := goose.DownTo(DB, ".", 0)

	//log.Printf("Start migrating database \n")
	//err = goose.Up(DB, ".")

}

func DbManager() *gorm.DB {
	return db
}
