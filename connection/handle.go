package connection

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Executer interface {
	Query(query string, args ...any) ([][]string, []string, error)
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
			row[i] = fmt.Sprintf("%s", values[i])
		}

		// Store the row in a slice of rows.
		all = append(all, row)

	}

	return all, c, nil
}
