package dbutil

import (
	"backend/config"
	"backend/pkg/logutil"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DatabaseConnections struct {
	MyDB *sqlx.DB
}

func OpenDatabases() (*DatabaseConnections, error) {
	mySchemaConfig := config.ConfigData.DatabaseConfigs["my_schema"]

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		mySchemaConfig.User,
		mySchemaConfig.Password,
		mySchemaConfig.Host,
		mySchemaConfig.Port,
		mySchemaConfig.DBName,
		mySchemaConfig.Charset,
		mySchemaConfig.ParseTime,
		mySchemaConfig.Loc,
	)

	var myDB *sqlx.DB
	var err error
	maxRetries := 10
	retryDelay := time.Second * 5

	for i := 0; i < maxRetries; i++ {
		myDB, err = sqlx.Connect("mysql", dsn)
		if err == nil {
			break
		}
		logutil.Error("Database connection failed (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to my_schema DB after retries: %v", err)
	}

	configureDatabase(myDB)
	return &DatabaseConnections{MyDB: myDB}, nil
}


func configureDatabase(db *sqlx.DB) {
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
}
