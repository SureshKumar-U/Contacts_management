package config

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDatabase() error {

	var err error

	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "contacts_management",
		AllowNativePasswords: true,
	}

	Db, err = sql.Open("mysql", cfg.FormatDSN()) // creates a database object

	if err != nil {
		return err
	}
	err = Db.Ping() //---> connecting to database
	if err != nil {
		return err
	}

	err = CreateUserTable()

	if err != nil {
		return err
	}

	err = CreateContactsTable()

	if err != nil {
		return err
	}

	return nil

}

func CreateUserTable() error {
	var err error

	query := `CREATE TABLE IF NOT EXISTS userdetails(
		id int AUTO_INCREMENT PRIMARY KEY,
		firstname varchar(100) NOT NULL,
		lastname  varchar(100) NOT NULL,
		email varchar(100) NOT NULL,
		password varchar(100) NOT NULL,
		admin varchar(100) NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP

	)`

	_, err = Db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func CreateContactsTable() error {
	var err error

	query := `CREATE TABLE IF NOT EXISTS contacts(
		id int AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		phoneNumber VARCHAR(20) NOT NULL,
        company VARCHAR(100),
        jobTitle VARCHAR(100),
        address VARCHAR(200),
		website VARCHAR(100),
		FOREIGN KEY (user_id) REFERENCES userdetails(id)
	)`
	_, err = Db.Exec(query)
	
    

	if err != nil {
		return err

	}

	return nil

}
