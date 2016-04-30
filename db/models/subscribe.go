package models

type Subscribe struct {
  ID       uint  `gorm:""`
  CreateAt uint  `gorm:""`
  SheetID  uint  `gorm:""`
  UserID   uint  `gorm:""`
  Sheet    Sheet `gorm:"ForeignKey: SheetID"`
}

func (Subscribe) TableName() string{
  return "subscribe"
}