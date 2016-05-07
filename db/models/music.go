package models

type Music struct {
  Model
  Name         string  `gorm:"" json:"name"`
  Duration     uint    `gorm:"" json:"duration"`
  OriginID     uint    `gorm:""`
  CompressedID uint    `gorm:""`
  AlbumID      uint    `gorm:""`
  ArtistID     uint    `gorm:"" json:"artistId"`
  SheetID      uint    `gorm:""`
  Origin       Audio   `gorm:"ForeignKey:OriginID"`
  Compressed   Audio   `gorm:"ForeignKey:CompressedID"`
  Album        Album   `gorm:"ForeignKey:AlbumID"`
  Artist       Artist  `gorm:"ForeignKey:ArtistID" json:"artist,omitempty"`
  Sheet        Sheet   `gorm:"ForeignKey:SheetID" json:"sheet,omitempty"`
}

func (Music) TableName() string{
  return "music"
}
