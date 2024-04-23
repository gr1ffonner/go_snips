package db

import (
	"database/sql"
	"fmt"
	"pricer/internal/config"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// DB is a struct to hold the database connection.
type DB struct {
	conn *sql.DB
}

// NewDB creates a new instance of DB with the provided connection string.
func NewDB() (*DB, error) {
	host := config.GetDBCred()[0]
	user := config.GetDBCred()[1]
	password := config.GetDBCred()[2]
	port, err := strconv.Atoi(config.GetDBCred()[3])
	if err != nil {
		return nil, err
	}
	dbname := config.GetDBCred()[4]
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	return &DB{conn: db}, nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	return db.conn.Close()
}

// Ping pings the database to check the connection.
func (db *DB) Ping() error {
	return db.conn.Ping()
}

// CheckError checks for errors and panics if found.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (db *DB) GetTokenAddresses() ([]string, error) {
	rows, err := db.conn.Query("SELECT token_addr FROM public.tokens_prices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []string
	for rows.Next() {
		var addr string
		err := rows.Scan(&addr)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, addr)
	}

	return addresses, nil
}

func (db *DB) UpsertTokenPrice(tokenAddr string, price float64, priceChange24h float64, updateTime time.Time) error {
	query := `
    UPDATE public.tokens_prices
    SET price = $2, price_change24h = $3, dttm_update = $4
    WHERE token_addr = $1;
    `
	_, err := db.conn.Exec(query, tokenAddr, price, priceChange24h, updateTime)
	if err != nil {
		return err
	}
	return nil
}
