package models

type Subscribe struct {
  Model
  SheetID  int  `gorm:""`
  UserID   int  `gorm:""`
  Sheet    Sheet `gorm:"ForeignKey: SheetID"`
}

func (Subscribe) TableName() string{
  return "subscribe"
}