package function

import (
	"database/sql"
	"fmt"
	"strings"
)

func CreateTable(db *sql.DB, command string) error {
	parts := strings.SplitN(command, " with ", 2)
	if len(parts) != 2 {
		return fmt.Errorf("Invalid command format")
	}

	tableName := strings.TrimSpace(parts[0])
	columnDefs := strings.TrimSpace(parts[1])

	query := fmt.Sprintf("%s (%s)", tableName, columnDefs)

	_, err := db.Exec(query)
	return err
}
