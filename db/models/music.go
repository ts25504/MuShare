package models

type Music struct {
  Model
  Name         string  `gorm:""`
  Duration     uint    `gorm:""`
  OriginID     uint    `gorm:""`
  CompressedID uint    `gorm:""`
  AlbumID      uint    `gorm:""`
  ArtistID     uint    `gorm:""`
  SheetID      uint    `gorm:""`
  Origin       Audio   `gorm:"ForeignKey:OriginID"`
  Compressed   Audio   `gorm:"ForeignKey:CompressedID"`
  Album        Album   `gorm:"ForeignKey:AlbumID"`
  Artist       Artist  `gorm:"ForeignKey:ArtistID"`
  Sheet        Sheet   `gorm:"ForeignKey:SheetID"`
}

func (Music) TableName() string{
  return "music"
}
