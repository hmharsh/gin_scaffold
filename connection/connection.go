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

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
	//	dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)
	dsn := fmt.Sprintf("%s:%s@/%s", dbConfig.User, dbConfig.Password, dbConfig.Name) //"user:password@/dbname"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to open connection to Database, err: %v", err)
	}

	return db, nil

}
