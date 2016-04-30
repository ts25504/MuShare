
package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20160429124010(txn *sql.Tx) {
	var sql string = "create table if not exists albums ( " +
	  "id int(10) unsigned not null auto_increment, " +
	  "created_at int(13) unsigned not null, " +
	  "update_at int(13) unsigned not null, " +
	  "name varchar(100) not null , " +
	  "cover varchar(200), " +
	  "user_id int(10) unsigned, " +
	  "PRIMARY KEY(id) " + ")"
	stmt, err := txn.Prepare(sql)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}

// Down is executed when this migration is rolled back
func Down_20160429124010(txn *sql.Tx) {
	var sql string = "drop table albums"
	stmt, err := txn.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()

	if err != nil {
		panic(err.Error())
	}
}
