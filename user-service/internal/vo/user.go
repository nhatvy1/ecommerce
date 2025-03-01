package vo

type UserLogin struct {
	Email    string
	PassWord string
}

type UserRegister struct {
	Email     string
	PassWord  string
	FirstName string
	LastName  string
}

type UserUpdate struct {
	FirstName string
	LastName  string
	NickName  string
}
