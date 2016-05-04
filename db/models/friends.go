package models

type Friends struct {
  Model
  FromID int    `gorm:""`
  ToID   int    `gorm:""`
  State  string  `gorm:""`
  User   User    `gorm:"ForeignKey:ToID"`
}

func (Friends) TableName() string {
  return "friends"
}
