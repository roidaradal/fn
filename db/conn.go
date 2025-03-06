package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type ConnParams struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewConnection(p *ConnParams) (*sql.DB, error) {
	dbAddr := fmt.Sprintf("%s:%s", p.Host, p.Port)
	dbCfg := mysql.Config{
		User:                 p.Username,
		Passwd:               p.Password,
		Net:                  "tcp",
		Addr:                 dbAddr,
		DBName:               p.Database,
		AllowNativePasswords: true,
	}
	dbc, err := sql.Open("mysql", dbCfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("cannot open db conn: %w", err)
	}
	err = dbc.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot ping db conn: %w", err)
	}
	return dbc, nil
}
