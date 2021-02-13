package models

import "time"

// Match :nodoc:
type Match struct {
	HomeTeam string    `bson:"homeTeam" json:"homeTeam"`
	AwayTeam string    `bson:"awayTeam" json:"awayTeam"`
	Score    []uint    `bson:"score" json:"score"`
	Tag      string    `bson:"tag" json:"tag"`
	Date     time.Time `bson:"date" json:"date"`
}
