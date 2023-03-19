package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	// ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"password" bson:"password"`
	IsAdmin      bool               `json:"isAdmin" bson:"isAdmin"`
	RefreshToken string             `json:"refreshToken" bson:"refreshToken"`
	Organization string             `json:"organization" bson:"organization"`
}

type Tokens struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Token     string             `json:"token" bson:"token"`
	UserID    string             `json:"userID" bson:"userID"`
	Type      string             `json:"type" bson:"type"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	ExpiresAt time.Time          `json:"expiresAt" bson:"UpdatedAt"`
}

type Organization struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Users []UserOrganization `json:"users" bson:"users"`
}

type UserOrganization struct {
	UserID  primitive.ObjectID `json:"_id" bson:"_id"`
	IsAdmin bool               `json:"isAdmin" bson:"isAdmin"`
}