package user

type Account struct {
  Mail     string
  Name     string
  Phone    string
  Password string
}

type Friend struct {
  UserID   int
  FriendID int
}
