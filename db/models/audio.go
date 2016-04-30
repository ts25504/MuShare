package models

type Audio struct {
  Model
  Url  string `gorm:""`
  Hash string `gorm:""`
}

func (Audio) TableName() string{
  return "audio"
}
