package connection

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Column struct {
	Name string
	Type string
}

type Executer interface {
	Query(query string, args ...any) ([][]string, []string, error)
	ShowTables() ([][]string, []string, error)
	SelectTable(table string) ([][]string, []string, error)
	GetColumns(table string) ([]Column, error)
}

type service struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Executer {
	return &service{db: db}
}

func (s *service) Query(query string, args ...any) ([][]string, []string, error) {

	results, err := s.db.Query(query, args...)
	if err != nil {
		return nil, nil, err
	}

	defer results.Close()

	c, err := results.Columns()
	if err != nil {
		return nil, nil, err
	}

	// //Create a slice of interface{} to represent each column of data.
	values := make([]interface{}, len(c))

	// Create a slice of interface{} to represent each column of data.
	valuePtrs := make([]interface{}, len(c))

	all := [][]string{}

	for results.Next() {
		// // Scan each row into the interface{} slice.
		for i := range c {
			valuePtrs[i] = &values[i]
		}

		results.Scan(valuePtrs...)

		// Create a slice to hold the data of each row.
		row := make([]string, len(c))

		// Copy the data from the interface{} slice to
		// the string slice.
		for i := range c {
			row[i] = fmt.Sprintf("%v", values[i])
		}

		// Store the row in a slice of rows.
		all = append(all, row)

	}

	return all, c, nil
}

func (s *service) ShowTables() ([][]string, []string, error) {
	return s.Query("SELECT name FROM sqlite_schema WHERE type ='table' AND name NOT LIKE 'sqlite_%';")
}

func (s *service) SelectTable(table string) ([][]string, []string, error) {
	return s.Query(fmt.Sprintf("SELECT * FROM %s;", table))
}

func (s *service) GetColumns(table string) ([]Column, error) {

	var columns []Column
	all, _, err := s.Query(fmt.Sprintf("PRAGMA table_info(%s);", table))
	if err != nil {
		return nil, err
	}

	for _, row := range all {
		columns = append(columns, Column{
			Name: row[1],
			Type: row[2],
		})
	}

	return columns, nil
}
