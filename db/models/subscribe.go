package models

type Subscribe struct {
  Model
  SheetID  uint  `gorm:""`
  UserID   uint  `gorm:""`
  Sheet    Sheet `gorm:"ForeignKey: SheetID"`
}

func (Subscribe) TableName() string{
  return "subscribe"
}