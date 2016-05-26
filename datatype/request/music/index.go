package music

type Sheet struct {
  Name        string
  UpdateName  string
  Privilege   string
  UserID      int
  RequestToID int
}

type Audio struct {
  Name       string
  UpdateName string
  AudioUrl   string
  ImageUrl   string
  ArtistID   int
  ArtistName int
  SheetID    int
  UserID     int
}

