package user

type Account struct {
  Mail     string
  Phone    string
  Name     string
  Password string
}

type Friend struct {
  FromID int
  ToID   int
  Token  string
}
