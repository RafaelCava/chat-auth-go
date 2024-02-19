package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string    `json:"id" valid:"uuid" gorm:"type:uuid;primary_key;"`
	Username  string    `json:"username,omitempty" gorm:"size:50" valid:"stringlength(5|50),optional"`
	Password  string    `json:"password,omitempty" valid:"required" gorm:"not null;"`
	Email     string    `json:"email" valid:"email,required" gorm:"unique;not null;size:100;"`
	Logo      string    `json:"logo,omitempty" valid:"url,optional"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"index" valid:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"index" valid:"-"`
}

func NewUser(username, password, email, logo string) (*User, error) {
	user := User{
		Username: username,
		Password: password,
		Email:    email,
		Logo:     logo,
	}
	user.prepare()

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (user *User) prepare() {
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) Validate() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}
