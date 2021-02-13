package models

import "time"

// Match :nodoc:
type Match struct {
	HomeTeam string    `bson:"homeTeam"`
	AwayTeam string    `bson:"awayTeam"`
	Score    []uint    `bson:"score"`
	Tag      string    `bson:"tag"`
	Date     time.Time `bson:"date"`
}
