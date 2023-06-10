package function

import (
	"database/sql"
	"fmt"
)

func WhereStatementWithTables(db *sql.DB, columnName string, columnValue string, tableName string) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = %s", tableName, columnName, columnValue)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	columnNames, err := rows.Columns()
	if err != nil {
		fmt.Println("Error retrieving column names:", err)
		return
	}

	columnValues := make([]interface{}, len(columnNames))
	columnPointers := make([]interface{}, len(columnNames))
	for i := range columnValues {
		columnPointers[i] = &columnValues[i]
	}

	for rows.Next() {
		if err := rows.Scan(columnPointers...); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		for i, value := range columnValues {
			fmt.Printf("%s: %v\n", columnNames[i], value)
		}

		fmt.Println("-----------------------")
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}
