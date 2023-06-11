package function

import (
	"database/sql"
)

func DeleteTable(db *sql.DB, tablename []string) {
	query := `DROP TABLE  `
	for i := range tablename {
		if i != len(tablename)-1 {
			query += tablename[i] + `,`
		} else {
			query += tablename[i] + `;`
		}
	}

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
