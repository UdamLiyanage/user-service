package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName       string              `json:"first_name" bson:"first_name"`
	LastName        string              `json:"last_name" bson:"last_name"`
	Email           string              `json:"email" bson:"email"`
	Password        string              `json:"password" bson:"password"`
	ContactNumber   string              `json:"contact_numbers,omitempty" bson:"contact_numbers,omitempty"`
	Address         map[string]string   `json:"address,omitempty" bson:"address,omitempty"`
	Devices         []map[string]string `json:"devices,omitempty" bson:"devices,omitempty"`
	EmailVerifiedAt time.Time           `json:"email_verified_at,omitempty" bson:"email_verified_at,omitempty"`
	CreatedAt       time.Time           `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt       time.Time           `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type AttachDevice struct {
	UserID     string `json:"user_id" bson:"user_id"`
	DeviceID   string `json:"device_id" bson:"device_id"`
	DeviceName string `json:"device_name" bson:"device_name"`
}

type AttachOrganisation struct {
	UserID           primitive.ObjectID `json:"user_id" bson:"user_id"`
	OrganisationID   string             `json:"organisation_id" bson:"organisation_id"`
	OrganisationName string             `json:"organisation_name" bson:"organisation_name"`
}
