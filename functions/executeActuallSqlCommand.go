package function

import (
	"database/sql"
)

func ExecuteActualSQLCommand(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
