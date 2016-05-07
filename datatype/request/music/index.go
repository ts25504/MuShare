package music

type sheet struct {
  UserId    int
  Token     string
  SheetId   int
  Name      string
  Privilege string
}

type audio struct {
  UserId     int
  Token      string
  Name       string
  AudioUrl   string
  CoverUrl   string
  ArtistId   int
  ArtistName int
  SheetId    int
}
