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

// FindByTag :nodoc
func (repo *MatchRepository) FindByTag(ctx context.Context, tag string) (models.Match, error) {
	cursor, err := repo.db.Collection("matches").Find(ctx, bson.M{"tag": tag})
	if err != nil {
		return models.Match{}, err
	}

	var record models.Match
	for cursor.Next(ctx) {
		cursor.Decode(&record)
	}

	return record, nil
}

// Update :nodoc
func (repo *MatchRepository) Update(ctx context.Context, data models.Match) (models.Match, error) {
	_, err := repo.
		db.Collection("matches").UpdateOne(
		ctx,
		bson.M{"tag": data.Tag},
		bson.D{
			{"$set", bson.D{{"score", data.Score}}},
		},
	)
	if err != nil {
		return models.Match{}, err
	}

	newMatch, _ := repo.FindByTag(ctx, data.Tag)
	return newMatch, nil
}

// Store :nodoc
func (repo *MatchRepository) Store(ctx context.Context, data models.Match) (models.Match, error) {
	_, err := repo.db.Collection("matches").InsertOne(ctx, data)
	if err != nil {
		return models.Match{}, err
	}

	return data, nil
}
