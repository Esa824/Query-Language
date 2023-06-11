package function

import (
	"database/sql"
)

func UpdateRow(db *sql.DB, id string, tablename string, values []string) {
	query := `UPDATE ` + tablename + ` SET (`
	columnNames, err := GetColumnNames(db, tablename)
	if err != nil {
		panic(err)
	}

	for i := range columnNames {
		if i != len(columnNames)-1 {
			query += columnNames[i] + `,`
		} else {
			query += columnNames[i] + `) = (`
		}
	}

	for i := range values {
		if values[i] != "_" {
			if i != len(values)-1 {
				query += "" + values[i] + "" + `,`
			} else {
				query += "" + values[i] + "" + `) WHERE id = ` + id
			}
		} else {
			if i != len(values)-1 {
				query += columnNames[i] + `,`
			} else {
				query += columnNames[i] + `) WHERE id = ` + id
			}
		}
	}
	query += `;`
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
