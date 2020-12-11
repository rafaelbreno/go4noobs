package query

import "fmt"

// type Builder interface {}

type DB struct {
	table  string
	query  string
	action []Action
}

type Action struct {
	Type string
	Body string
}

func Table(table string) DB {
	return DB{
		table: table,
	}
}

func (db *DB) Where(index, cond, value string) *DB {
	var act Action

	actStr := fmt.Sprintf("%s %s '%s'", index, cond, value)

	act.Type = "WHERE"
	act.Body = actStr

	db.action = append(db.action, act)

	return db
}

func (db *DB) Get() string {
	item := fmt.Sprintf("SELECT * from %s", db.table)

	db.append(item)

	if len(db.action) > 1 {
		for key, value := range db.action {
			if key == 0 {
				db.append(value.Type)
				db.append(value.Body)
			} else {
				db.append("AND")
				db.append(value.Body)
			}
		}
	}

	return db.query
}

func (db *DB) append(str string) {
	db.query = fmt.Sprintf("%s %s", db.query, str)
}
