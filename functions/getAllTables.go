package function

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetAllTables(db *sql.DB) {
	query := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public' AND table_type = 'BASE TABLE';
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tableName)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
