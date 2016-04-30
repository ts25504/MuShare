
package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20160429131456(txn *sql.Tx) {
	var sql string = "create table if not exists friends ( " +
	  "id int(10) unsigned not null auto_increment, " +
	  "created_at int(13) unsigned not null, " +
	  "from_id int(13) unsigned not null, " +
	  "to_id int(10) unsigned not null, " +
	  "state varchar(100), " +
	  "PRIMARY KEY(id) " + ")"
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
func Down_20160429131456(txn *sql.Tx) {
  var sql string = "drop table friends"
  stmt, err := txn.Prepare(sql)
  if err != nil {
    panic(err.Error())
  }
  _, err = stmt.Exec()

  if err != nil {
    panic(err.Error())
  }
}
