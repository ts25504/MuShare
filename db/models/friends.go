package models

type Friends struct {
  ID        uint    `gorm:"type:int(10);unsigned;primary_key;auto_increment"`
  FromID    uint    `gorm:"type:int(10);unsigned"`
  ToID      uint    `gorm:"type:int(10);unsigned"`
  Privilege string  `gorm:"type:varchar(100)"`
  User      User    `gorm:"ForeignKey:ToID"`
}

func (Friends) TableName() string{
  return "friends"
}
