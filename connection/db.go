package connection

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	conn *sqlx.DB
	cmux sync.Mutex
)

type ConnFn func() (*sqlx.DB, error)

func ConnectionFn(url string) ConnFn {
	return func() (cx *sqlx.DB, err error) {
		cmux.Lock()
		defer cmux.Unlock()

		if conn != nil && conn.Ping() == nil {
			return conn, nil
		}

		fmt.Println("Connecting to database...")
		fmt.Println("Database URL: ", url)
		conn, err = sqlx.Connect("sqlite3", url)
		if err != nil {
			return nil, err
		}

		return conn, nil
	}
}
