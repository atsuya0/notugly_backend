package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tayusa/notugly_backend/configs"
)

func NewDB() *sql.DB {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		configs.Data.DB.User,
		configs.Data.DB.Password,
		configs.Data.DB.Host,
		configs.Data.DB.Port,
		configs.Data.DB.Name,
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	return db
}
