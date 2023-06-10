package function

import (
	"database/sql"
)

func CreateRow(db *sql.DB, tablename string, values []string) {
	query := `insert into ` + tablename + ` Values (`
	for i, value := range values {
		if i != len(values)-1 {
			query += value + ", "
		} else {
			query += value
		}
	}
	query += ");"
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}
