package models

type Friends struct {
  ID     uint    `gorm:""`
  FromID uint    `gorm:""`
  ToID   uint    `gorm:""`
  State  string  `gorm:""`
  User   User    `gorm:"ForeignKey:ToID"`
}

func (Friends) TableName() string {
  return "friends"
}
