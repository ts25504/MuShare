package music

type Sheet struct {
  UserID    int
  Token     string
  Name      string
  Privilege string
}

type Audio struct {
  Token      string
  Name       string
  AudioUrl   string
  ImageUrl   string
  ArtistID   int
  ArtistName int
  SheetID    int
}
