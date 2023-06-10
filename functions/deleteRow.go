package function

import (
	"database/sql"
)

func DeleteRow(db *sql.DB, id string, tablename string) {
	stmt, err := db.Prepare("Delete from " + tablename + " where id = " + id + ";")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}
