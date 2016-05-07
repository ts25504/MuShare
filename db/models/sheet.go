package models

type Sheet struct {
  Model
  Name      string   `gorm:"" json:"name"`
  Privilege string   `gorm:"" json:"privilege"`
  UserID    uint     `gorm:"" json:"userId"`
  User      User     `gorm:"ForeignKey:UserID" json:"user,omitempty"`
  Music     []Music  `gorm:"ForeignKey:SheetID" json:"music, omiempty"`
}

func (Sheet) TableName() string{
  return "sheets"
}