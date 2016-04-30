package models

type Album  struct {
  Model
  Name   string `gorm:""`
  Cover  string `gorm:""`
  UserID uint   `gorm:""`
  User   User   `gorm:"ForeignKey:UserID"`
}

func (Album) TableName() string{
  return "Albums"
}