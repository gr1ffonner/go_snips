package db

import (
	"database/sql"
	"encoding/json"

	config "imggen/internal/config"
	models "imggen/internal/nft"

	_ "github.com/lib/pq"
)

// DB is a struct to hold the database connection.
type DB struct {
	conn *sql.DB
}

// NewDB creates a new instance of DB with the provided connection string.
func NewDB() (*DB, error) {
	host := config.GetHostName()
	psqlconn := host

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

// UpdateTariffData updates the tarifs_nft_metadata for the given token_addr with the provided TariffData.
func (db *DB) UpdateTariffData(tokenAddr string, data models.TariffsData) error {
	query := `
		UPDATE public.projects
		SET tarifs_nft_metadata = $1
		WHERE token_addr = $2;
	`

	// Convert TariffData to JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Execute the SQL statement
	_, err = db.conn.Exec(query, string(jsonData), tokenAddr)
	if err != nil {
		return err
	}

	return nil
}

// CheckError checks for errors and panics if found.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
