package music

type Sheet struct {
  UserID        int
  Token         string
  Name          string
  Privilege     string
  RequestFromID int
  RequestToID   int
}

type Audio struct {
  Token         string
  Name          string
  AudioUrl      string
  ImageUrl      string
  ArtistID      int
  ArtistName    int
  SheetID       int
  RequestUserID int
}

