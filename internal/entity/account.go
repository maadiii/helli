package entity

type Account struct {
	ID       int
	Username string
	Fund     float64

	User *User
}
