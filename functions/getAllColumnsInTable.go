package function

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetAllColumnsInTable(db *sql.DB, table string) {
	query := fmt.Sprintf("SELECT * FROM \"%s\"", table)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return
	}

	values := make([]interface{}, len(columns))
	pointers := make([]interface{}, len(columns))
	for i := range values {
		pointers[i] = &values[i]
	}

	for rows.Next() {
		if err := rows.Scan(pointers...); err != nil {
			fmt.Println(err)
			return
		}

		for i, value := range values {
			fmt.Printf("%s: %v\n", columns[i], value)
		}

		fmt.Println("-----------------------")
	}
}
