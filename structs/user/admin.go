package user

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "Admin",
			LastName:  "Admin",
			birthDate: "BD",
		},
	}
}
