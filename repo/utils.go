package repo

import (
	"database/sql"
	"laundry/warehouse"
)

type RowsScanner func(rows *sql.Rows) error
type RowScanner func(row *sql.Row) error

func preparedQuery(query string, params ...interface{}) (err error) {
	db := warehouse.Conn.DB

	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(params...)
	if err != nil {
		return
	}

	return
}

func preparedQueryTx(db *sql.Tx, query string, params ...interface{}) (err error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(params...)
	if err != nil {
		return
	}

	return
}

func fetchRows(query string, params []interface{}, scanner RowsScanner) (err error) {
	db := warehouse.Conn.DB
	rows, err := db.Query(query, params...)
	if err != nil {
		return
	}

	for rows.Next() {
		err = scanner(rows)
		if err != nil {
			return
		}
	}

	return
}

func fetchRowsTx(db *sql.Tx, query string, params []interface{}, scanner RowsScanner) (err error) {
	rows, err := db.Query(query, params...)
	if err != nil {
		return
	}

	for rows.Next() {
		err = scanner(rows)
		if err != nil {
			return
		}
	}

	return
}

func fetchRow(query string, params []interface{}, scanner RowScanner) (err error) {
	db := warehouse.Conn.DB
	row := db.QueryRow(query, params...)
	return scanner(row)
}

func fetchRowTx(db *sql.Tx, query string, params []interface{}, scanner RowScanner) (err error) {
	row := db.QueryRow(query, params...)
	return scanner(row)
}
