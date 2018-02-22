package db

import (
	"../config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type SnsDB struct {
	db *sql.DB
}

func (this *SnsDB) Connect() {
	var err error
	var c config.Data = config.LoadConfig()
	var dbname string = c.Config.Database.Dbname
	var dbuser string = c.Config.Database.User
	var dbpassword string = c.Config.Database.Password
	this.db, err = sql.Open("mysql", dbuser+":"+dbpassword+"@/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
}

func (this *SnsDB) Close() {
	var err error
	err = this.db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
