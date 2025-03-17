package vo

type UserLogin struct {
	Email    string
	PassWord string
}

type UserRegister struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=0,lte=25"`
	FirstName string `json:"first_name" validate:"required,gte=0,lte=25"`
	LastName  string `json:"last_name" validate:"required,gte=0,lte=25"`
}

type UserUpdate struct {
	FirstName *string `json:"first_name" validate:"omitempty,gte=1,lte=25"`
	LastName  *string `json:"last_name" validate:"omitempty,gte=1,lte=25"`
	NickName  *string `json:"nickname" validate:"omitempty,gte=1,lte=25"`
	Status    *string `json:"status" validate:"omitempty,oneof=active blocked pending"`
}
