package entity

type User struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Password  string

	Accounts []*Account
}
