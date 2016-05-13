package music

type Sheet struct {
  Token         string
  Name          string
  UpdateName    string
  Privilege     string
  RequestFromID int
  RequestToID   int
}

type Audio struct {
  Token         string
  Name          string
  UpdateName    string
  AudioUrl      string
  ImageUrl      string
  ArtistID      int
  ArtistName    int
  SheetID       int
  RequestFromID int
}

