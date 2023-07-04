package users

type User struct {
	Email     string
	Firstname string
	Lastname  string
	Password  string
	Phone     string
	Username  string
}

func UserConstruct(email, firstname, lastname, password, phone, username string) User {
	return User{
		Email:     email,
		Firstname: firstname,
		Lastname:  lastname,
		Password:  password,
		Phone:     phone,
		Username:  username,
	}
}
