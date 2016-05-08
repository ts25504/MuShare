package music

type Sheet struct {
  UserID    int
  Token     string
  SheetId   int
  Name      string
  Privilege string
}

type Audio struct {
  UserID     int
  Token      string
  Name       string
  AudioUrl   string
  CoverUrl   string
  ArtistID   int
  ArtistName int
  SheetID    int
}
