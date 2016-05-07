package models

type Audio struct {
  Model
  Name       string  `gorm:"" json:"name"`
  Duration   uint    `gorm:"" json:"duration"`
  AudioUrl   string  `gorm:""`
  CoverUrl   string  `gorm:""`
  ArtistID   uint    `gorm:"" json:"artistId"`
  SheetID    uint    `gorm:""`
  Artist     Artist  `gorm:"ForeignKey:ArtistID" json:"artist,omitempty"`
  Sheet      Sheet   `gorm:"ForeignKey:SheetID" json:"sheet,omitempty"`
}

func (Audio) TableName() string {
  return "audio"
}
