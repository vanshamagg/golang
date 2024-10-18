package structures

type User struct {
	FirstName string
	LastName  string
	Age       int
	Address   string
}

type Employee struct {
	User
	EmployeeId string
}

func CreateUser(u *User) User {

	user := new(User)
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Address = u.Address
	user.Age = u.Age

	return *user
}
