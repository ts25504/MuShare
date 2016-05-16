package main

import (
  "database/sql"
)

// Up is executed when this migration is applied
func Up_20160515111342(txn *sql.Tx) {
  var sql string = "alter table friends change " +
    "from_id user_id int(13) unsigned not null"
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
func Down_20160515111342(txn *sql.Tx) {
  var sql string = "alter table friends change " +
    "user_id from_id int(13) unsigned not null"
  stmt, err := txn.Prepare(sql)
  if err != nil {
    panic(err.Error())
  }
  _, err = stmt.Exec()

  if err != nil {
    panic(err.Error())
  }
}
