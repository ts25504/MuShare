package main

import (
  "database/sql"
)

// Up is executed when this migration is applied
func Up_20160429102901(txn *sql.Tx) {
  var sql string = "create table if not exists users ( " +
    "id int(10) unsigned not null auto_increment, " +
    "created_at int(13) unsigned not null, " +
    "updated_at int(13) unsigned not null, " +
    "name varchar(100) not null unique, " +
    "screen_name varchar(100), " +
    "mail varchar(100) not null unique, " +
    "phone varchar(100) unique, " +
    "avatar varchar(200), " +
    "gender varchar(10), " +
    "birth int(13), " +
    "description text, " +
    "type varchar(16), " +
    "password varchar(100), " +
    "last_login_at int(13), " +
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
func Down_20160429102901(txn *sql.Tx) {
  var sql string = "drop table users"
  stmt, err := txn.Prepare(sql)
  if err != nil {
    panic(err.Error())
  }
  _, err = stmt.Exec()

  if err != nil {
    panic(err.Error())
  }
}
