package globals

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
	"log"
)

type DataBaseResult struct {
	res     sql.Result
	rows    *sql.Rows
	columns []string
	row     map[string]string
}

func (r *DataBaseResult) GetLastInsertID() int64 {
	var lastID int64
	if r.res != nil {
		lastID, _ = r.res.LastInsertId()
	}
	return lastID
}

func (r *DataBaseResult) FetchRow() map[string]string {
	if r == nil || r.Fetch() == false {
		return make(map[string]string)
	}
	var row = r.row
	r.Free()
	return row
}

func (r *DataBaseResult) Fetch() bool {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	if r.rows.Next() == false {
		r.Free()
		return false
	}

	var pointers = make([]interface{}, len(r.columns))
	var container = make([][]byte, len(r.columns))
	for i, _ := range r.columns {
		pointers[i] = &container[i]
	}
	err := r.rows.Scan(pointers...)
	if err != nil {
		log.Println(err)
	}
	r.row = make(map[string]string)
	for key, value := range container {
		r.row[r.columns[key]] = strings.TrimSpace(string(value))
	}

	return true
}

func (r *DataBaseResult) Row() map[string]string {
	if r == nil {
		return make(map[string]string)
	}
	return r.row
}

func (r *DataBaseResult) Free() {
	if r == nil {
		return
	}
	r.row = make(map[string]string)
	r.rows.Close()
}
