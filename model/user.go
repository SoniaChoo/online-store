package model

// User is the user register info
type User struct {
	email  string
	passwd string
}

func (u User) GetEmail() string {
	return u.email
}

func (u User) GetPasswd() string {
	return u.passwd
}
