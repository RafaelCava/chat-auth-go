package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Room struct {
	ID          string    `json:"id" valid:"uuid" gorm:"type:uuid;primary_key;"`
	Name        string    `json:"name" valid:"stringlength(5|50),required" gorm:"size:50"`
	OwnerId     string    `json:"owner_id" valid:"uuid,required" gorm:"type:uuid;"`
	Owner       User      `json:"owner,omitempty" gorm:"foreignKey:OwnerId" valid:"-"`
	ActiveUsers []User    `json:"active_users,omitempty" gorm:"many2many:room_users;"`
	CreatedAt   time.Time `json:"created_at" gorm:"index" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"index" valid:"-"`
}

func NewRoom(name, ownerId string) (*Room, error) {
	room := Room{
		Name:    name,
		OwnerId: ownerId,
	}
	room.prepare()

	err := room.Validate()
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (room *Room) prepare() {
	room.ID = uuid.NewV4().String()
	room.CreatedAt = time.Now()
	room.UpdatedAt = time.Now()
}

func (room *Room) Validate() error {
	_, err := govalidator.ValidateStruct(room)

	if err != nil {
		return err
	}

	return nil
}
