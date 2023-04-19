package user

type User struct {
	ID        int
	Firstname string
	Lastname  string
}

func NewUser(id int, firstname, lastname string) User {
	return User{
		ID:        id,
		Firstname: firstname,
		Lastname:  lastname,
	}
}

type UserRepository interface {
	Register(user User, hash, salt string) error
}
