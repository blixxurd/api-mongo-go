package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MinecraftServer struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `bson:"name" json:"name"`
	Address       string             `bson:"address" json:"address"`
	Port          int                `bson:"port" json:"port"`
	Version       string             `bson:"version" json:"version"`
	PlayersOnline int                `bson:"playersOnline" json:"playersOnline"`
	MaxPlayers    int                `bson:"maxPlayers" json:"maxPlayers"`
	Status        string             `bson:"status" json:"status"`
	MOTD          string             `bson:"motd" json:"motd"`
	LastChecked   time.Time          `bson:"lastChecked" json:"lastChecked"`
	Tags          []string           `bson:"tags" json:"tags"`
}
