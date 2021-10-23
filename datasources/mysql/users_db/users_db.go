package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_username = "mysql_username"
	mysql_password = "mysql_password"
	mysql_host     = "mysql_host"
	mysql_schema   = "mysql_schema"
)

var (
	Client *sql.DB

	username = "root"      //os.Getenv(mysql_username)
	password = "password"  //os.Getenv(mysql_password)
	host     = "localhost" //os.Getenv(mysql_host)
	schema   = "users_db"  //os.Getenv(mysql_schema)
)

func init() {

	// fmt.Println(username, password, host, schema)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured!!")
}
