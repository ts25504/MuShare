package main

import (
  "database/sql"
)

// Up is executed when this migration is applied
func Up_20160429131051(txn *sql.Tx) {
  var sql string = "create table if not exists subscribe ( " +
  "id int(10) unsigned not null auto_increment, " +
  "created_at int(13) unsigned not null, " +
  "updated_at int(13) unsigned not null, " +
  "user_id int(10) unsigned not null, " +
  "sheet_id int(10) unsigned not null, " +
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
func Down_20160429131051(txn *sql.Tx) {
  var sql string = "drop table subscribe"
  stmt, err := txn.Prepare(sql)
  if err != nil {
    panic(err.Error())
  }
  _, err = stmt.Exec()

  if err != nil {
    panic(err.Error())
  }
}
