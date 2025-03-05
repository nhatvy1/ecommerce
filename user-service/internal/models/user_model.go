package models

import (
	"database/sql/driver"
)

type Status string

const (
	Active  Status = "active"
	Blocked Status = "blocked"
	Pending Status = "pending"
)

func (p *Status) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		*p = Status(b)
	}
	return nil
}

func (p Status) Value() (driver.Value, error) {
	return string(p), nil
}

func ToStatus(s string) Status {
	switch s {
	case string(Active):
		return Active
	case string(Blocked):
		return Blocked
	case string(Pending):
		return Pending
	default:
		return Pending
	}
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

type APIUserEmail struct {
	Email string
}
