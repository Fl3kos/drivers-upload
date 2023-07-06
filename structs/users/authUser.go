package users

type AuthUser struct {
	Email     string
	Firstname string
	Lastname  string
	Password  string
	Phone     string
	Username  string
}

func UserConstruct(email, firstname, lastname, password, phone, username string) AuthUser {
	return AuthUser{
		Email:     email,
		Firstname: firstname,
		Lastname:  lastname,
		Password:  password,
		Phone:     phone,
		Username:  username,
	}
}