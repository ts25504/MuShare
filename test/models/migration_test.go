package test

import (
  "testing"
  "os"
  "MuShare/conf"
  "MuShare/db/models"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/jinzhu/gorm"
)

var DB *gorm.DB

func TestMain(m *testing.M) {
  DB = models.New(conf.Mysql{
    User:"root",
    Password:"",
    Database:"mushare_test",
  })
  os.Exit(m.Run())
}
