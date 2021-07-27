package connection

import (
	"database/sql"
	"fmt"

	"ginapp/configuration"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// NewConnection creates a new connection from the given configs
func NewConnection(dbConfig configuration.DatabaseConfiguration) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name) //"user:password@/dbname"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection to the database, err: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database, err: %v", err)
	}
	return db, nil
}
