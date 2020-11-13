package datasources

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type Config struct {
	DriverName string
	DBHost     string
	DBName     string
	DBPort     string
	DBUserName string
	DBPassword string
}

type DataSource struct {
	Config *config.Config
}

func NewDB(d *DataSource) (*sql.DB, error) {
	ndb, err := sql.Open(d.Config.DriverName, d.connString())
	if err != nil {
		log.Panic(err)
	}

	if err = ndb.Ping(); err != nil {
		log.Panic(err)
	}
	return ndb, err
}

func (d *DataSource) connString() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", d.Config.DBHost, d.Config.DBUserName, d.Config.DBPassword, d.Config.DBName)

}
