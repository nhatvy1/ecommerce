package models

import "database/sql/driver"

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
	Pending  Status = "pending"
)

func (p *Status) Scan(value interface{}) error {
	*p = Status(value.([]byte))
	return nil
}

func (p Status) Value() (driver.Value, error) {
	return string(p), nil
}

type User struct {
	BaseModel
	FirstName string `json:"first_name" gorm:"column:first_name;not null;" validate:"required,min=1,max=50"`
	LastName  string `json:"last_name" gorm:"column:last_name;not null;" validate:"required,min=1,max=50"`
	Email     string `json:"email" gorm:"column:email;not null;unique;" validate:"email,required"`
	Password  string `json:"password" gorm:"column:password;not null;" validate:"required"`
	Status    Status `json:"status" gorm:"column:status;type:ENUM('active', 'blocked', 'pending');default:'active';"`
	NickName  string `json:"nickname" gorm:"column:nickname;default:null;"`
}

type APIUser struct {
	FirstName string
	LastName  string
}
