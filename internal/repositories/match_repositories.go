package repositories

import (
	"context"

	"github.com/yurichandra/gunners/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MatchRepository :nodoc:
type MatchRepository struct {
	db *mongo.Database
}

// NewMatchRepository :nodoc:
func NewMatchRepository(database *mongo.Database) *MatchRepository {
	return &MatchRepository{
		db: database,
	}
}

// Get :nodoc
func (repo *MatchRepository) Get(ctx context.Context) ([]models.Match, error) {
	matches := make([]models.Match, 0)

	cursor, err := repo.db.Collection("matches").Find(ctx, bson.M{})
	if err != nil {
		return []models.Match{}, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var row models.Match
		err := cursor.Decode(&row)
		if err != nil {
			return []models.Match{}, err
		}

		matches = append(matches, row)
	}

	return matches, nil
}

// Store :nodoc
func (repo *MatchRepository) Store(ctx context.Context, data models.Match) (models.Match, error) {
	_, err := repo.db.Collection("matches").InsertOne(ctx, data)
	if err != nil {
		return models.Match{}, err
	}

	return data, nil
}
