package models

import (
  "github.com/jinzhu/gorm"
  "time"
)

type Friends struct {
  ID        int            `gorm:"" json:"-"`
  CreatedAt int64          `gorm:"" json:"-"`
  UpdatedAt int64          `gorm:"" json:"-"`
  UserID    int            `gorm:"" json:"-"`
  FriendID  int            `gorm:"" json:"friendId"`
  State     string         `gorm:"" json:"-"`
  User      User           `gorm:"ForeignKey:UserID" json:"user,omitempty"`
  Friend    User           `gorm:"ForeignKey:FriendID" json:"friend,omitempty"`
}

func (Friends) TableName() string {
  return "friends"
}

func (friend *Friends) BeforeCreate(scope *gorm.Scope) (err error) {
  scope.SetColumn("created_at", time.Now().Unix())
  scope.SetColumn("updated_at", time.Now().Unix())
  return
}

func (friend *Friends) BeforeUpdate(scope *gorm.Scope) (err error) {
  scope.SetColumn("updated_at", time.Now().Unix())
  return
}
