package test

import (
  "testing"
  "MuShare/db/models"
)

func TestFriendsInsert(t *testing.T) {
  friends := models.Friends{
  }
  DB.Create(&friends)
}