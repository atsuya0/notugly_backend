package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tayusa/notugly_backend/app/config"
)

func NewMySql() *sql.DB {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		config.Data.DB.User,
		config.Data.DB.Password,
		config.Data.DB.Host,
		config.Data.DB.Port,
		config.Data.DB.Name,
	)

	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	return conn
}
