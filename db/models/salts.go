package models

type Salts struct {
  Model
  UserID    int      `gorm:"" json:"userId"`
	Salt      string   `gorm:""`
}

func (Salts) TableName() string{
  return "salts"
}
