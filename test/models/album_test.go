package test

import (
  "testing"
  "MuShare/db/models"
)

func TestAlbumInsert(t *testing.T){
  album := models.Album{
    Name:"test",
    UserID:1,
  }
  DB.Create(&album)
}
