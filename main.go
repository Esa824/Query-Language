package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strings"

	functions "github.com/QueryLanguage/functions"

	_ "github.com/lib/pq"
)

func main() {
	// Get Database name
	fmt.Print("Database: ")
	database := readLine()
	database = strings.TrimSpace(database)

	// Establish a database connection
	db, err := sql.Open("postgres", "postgres://admin:password@localhost:5432/"+database+"?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for {
		fmt.Print("Query: ")
		query := readLine()
		query = strings.TrimSpace(query)
		goParser(db, query)
	}
}
func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim the newline character from the input
	input = strings.TrimSuffix(input, "\n")

	return input
}

func getColumnAndTableNames(str string) (string, string, string, bool) {
	pattern := `Get ((\w+) (\w+))( and (\w+) (\w+))* in (\w+)`

	if matches := regexp.MustCompile(pattern).FindStringSubmatch(str); len(matches) > 0 {
		columnName := strings.Title(matches[2])
		columnValue := strings.Title(matches[3])
		tableName := strings.Title(matches[len(matches)-1])
		return columnName, columnValue, tableName, true
	}

	return "", "", "", false
}

func goParser(db *sql.DB, query string) {
	if strings.HasPrefix(query, "Get All Rows in") {
		table := strings.TrimSpace(strings.TrimPrefix(query, "Get All Rows in"))
		functions.GetAllColumnsInTable(db, table)
	}
	if query == "List All Tables" {
		functions.GetAllTables(db)
	}
	if _, _, _, correctQuery := getColumnAndTableNames(query); correctQuery == true {
		ColumnName, ColumnValue, tableName, _ := getColumnAndTableNames(query)
		functions.WhereStatementWithTables(db, ColumnName, ColumnValue, tableName)
	}
	if strings.HasPrefix(query, "Create Row in") {
		words := strings.Fields(query)
		if len(words) > 4 {
			result := words[4:]
			functions.CreateRow(db, words[3], result)
		}
	}
	if strings.HasPrefix(query, "Delete Row in") {
		words := strings.Fields(query)
		functions.DeleteRow(db, words[4], words[3])
	}
	if strings.HasPrefix(query, "Update Row in") {
		words := strings.Fields(query)
		functions.UpdateRow(db, words[4], words[3], words[5:])
	}
	if query == "exit" {
		os.Exit(0)
		db.Close()
	}
	if strings.HasPrefix(query, "Delete Table") {
		words := strings.Fields(query)
		functions.DeleteTable(db, words[2:])
	}
	if strings.HasPrefix(query, "Create Table") {
		functions.CreateTable(db, query)
	}
	if strings.HasPrefix(query, "List All Columns in") {
		words := strings.Fields(query)
		functions.GetAllColumnsInTable(db, words[4])
	}
	if strings.HasSuffix(query, "--sql") {
		words := strings.Fields(query)
		functions.ExecuteActualSQLCommand(db, strings.Join(words[:len(words)-1], " "))
	}
}
