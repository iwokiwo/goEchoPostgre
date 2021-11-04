package db

import (
	//"database/sql"
	"fmt"

	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
	//	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//	"github.com/pressly/goose"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "yespos"
)

//var DB *sql.DB
var db *gorm.DB
var err error

func Init() {
	//------------------------------MYSQL----------------------------
	//conf := config.GetConfig()
	//connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	//db, err = sql.Open("mysql", connectionString)
	//---------------------------------------------------------------

	//------------------------------Postgresql-----------------------
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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
