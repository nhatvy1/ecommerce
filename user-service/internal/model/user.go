package model

type UserRegister struct {
	UserEmail    string `json:"email" validate:"required,email"`
	UserPassword string `json:"password" validate:"required,gte=0,lte=25"`
}

type UserLogin struct {
	UserEmail    string `json:"email" validate:"required,email"`
	UserPassword string `json:"password" validate:"required"`
}

type UserUpdate struct {
	FirstName *string `json:"first_name" validate:"omitempty,gte=1,lte=25"`
	LastName  *string `json:"last_name" validate:"omitempty,gte=1,lte=25"`
	NickName  *string `json:"nickname" validate:"omitempty,gte=1,lte=25"`
	Status    *string `json:"status" validate:"omitempty,oneof=active blocked pending"`
}
