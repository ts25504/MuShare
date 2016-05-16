
package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20160515112257(txn *sql.Tx) {
	var sql string = "alter table friends change " +
	  "to_id friend_id int(13) unsigned not null"
	stmt, err := txn.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()

	if err != nil {
		panic(err.Error())
	}
}

// Down is executed when this migration is rolled back
func Down_20160515112257(txn *sql.Tx) {
	var sql string = "alter table friends change " +
	  "friend_id to_id int(13) unsigned not null"
	stmt, err := txn.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()

	if err != nil {
		panic(err.Error())
	}
}
