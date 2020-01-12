package model

// User is the user register info
type User struct {
	Email  string
	Passwd string
}

type ChangeUser struct {
	ConflictTest string
	Passwd string
	NewEmail string
	NewPasswd string
}
