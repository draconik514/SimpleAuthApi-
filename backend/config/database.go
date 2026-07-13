package config 

import ( 
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB(){
	err := godotenv.Load()
	if err != nil {
		log.Println("there is no .env file here")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPassword, dbHost, dbPort, dbName)

	//output dns

	log.Println("Connecting to MySQL.....")
	log.Printf("DNS : %s:%s@tcp(%s:%s)/%s",
        dbUser, "******", dbHost, dbPort, dbName )

	// open mysql
	db, err := sql.Open("mysql", dns)
	if err != nil{
		log.Fatal("Failed to open database connection :", err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatal("Database ping failed :", err)
	}

	DB = db
	log.Println("Database connected successfully")
}

func GetDB() *sql.DB{
	return DB
}