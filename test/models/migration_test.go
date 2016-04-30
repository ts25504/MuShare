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

func BenchmarkDropUser(b *testing.B) {
  DB.DropTable(&models.User{})
  b.Skip()
}

func BenchmarkCreateUser(b *testing.B) {
  DB.CreateTable(&models.User{})
  b.Skip()
}

func BenchmarkDropSheet(b *testing.B) {
  DB.DropTable(&models.Sheet{})
  b.Skip()
}

func BenchmarkCreateSheet(b *testing.B) {
  DB.CreateTable(&models.Sheet{})
  b.Skip()
}

func BenchmarkDropFriends(b *testing.B) {
  DB.DropTable(&models.Friends{})
  b.Skip()
}

func BenchmarkCreateFriends(b *testing.B) {
  DB.CreateTable(&models.Friends{})
  b.Skip()
}

func BenchmarkDropSubscribe(b *testing.B) {
  DB.DropTable(&models.Subscribe{})
  b.Skip()
}

func BenchmarkCreateSubscribe(b *testing.B) {
  DB.CreateTable(&models.Subscribe{})
  b.Skip()
}

func TestMain(m *testing.M) {
  DB = models.New(conf.Mysql{
    User:"root",
    Password:"w19920610",
    Database:"mushare_test",
  })
  os.Exit(m.Run())
}
