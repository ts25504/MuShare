package models

type Friends struct {
  Model
  FromID int    `gorm:"" json:"fromId"`
  ToID   int    `gorm:"" json:"toId"`
  State  string  `gorm:"" json:"-"`
  User   User    `gorm:"ForeignKey:ToID" json:"user,omitempty"`
}

func (Friends) TableName() string {
  return "friends"
}
