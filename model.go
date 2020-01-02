package main

import "time"

type User struct {
	FirstName       string              `json:"first_name" bson:"first_name"`
	LastName        string              `json:"last_name" bson:"last_name"`
	Email           string              `json:"email" bson:"email"`
	Password        string              `json:"password" bson:"password"`
	ContactNumbers  []string            `json:"contact_numbers" bson:"contact_numbers"`
	Address         map[string]string   `json:"address" bson:"address"`
	Devices         []map[string]string `json:"devices" bson:"devices"`
	EmailVerifiedAt time.Time           `json:"email_verified_at" bson:"email_verified_at"`
	CreatedAt       time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at" bson:"updated_at"`
}
