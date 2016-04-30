package models

type Sheet struct {
  Model
  Name      string   `gorm:""`
  Privilege string   `gorm:""`
  UserID    uint     `gorm:""`
  User      User     `gorm:"ForeignKey:UserID"`
  Music     []Music  `gorm:"ForeignKey:SheetID"`
}

func (Sheet) TableName() string{
  return "sheets"
}