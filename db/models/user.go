package models

type User struct {
  Model
  Name        string      `gorm:"" json:"name"`
  Mail        string      `gorm:"" json:"mail"`
  Phone       string      `gorm:"" json:"phone"`
  Avatar      string      `gorm:"" json:"avatar"`
  Gender      int         `gorm:"" json:"gender"`
  Birth       uint        `gorm:"" json:"birth"`
  Description string      `gorm:"" json:"description"`
  Password    string      `gorm:"" json:"-"`
  LastLoginAt uint        `gorm:"" json:"-"`
  Friends     []Friends   `gorm:"ForeignKey:FromID" json:"friends,omitempty"`
  Albums      []Album     `gorm:"ForeignKey:UserID" json:"albums,omitempty"`
  Sheets      []Sheet     `gorm:"ForeignKey:UserID" json:"sheets,omitempty"`
  Subscribe   []Subscribe `gorm:"ForeignKey:UserID" json:"token,omitempty"`
  Token       string      `gorm:"-"`
}

func (User) TableName() string {
  return "users"
}