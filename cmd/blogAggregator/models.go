package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"github.com/AxterDoesCode/blogAggregator/internal/database"
)

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedtoFeed(dbf database.Feed) Feed {
	return Feed{
		ID:            dbf.ID,
		CreatedAt:     dbf.CreatedAt,
		UpdatedAt:     dbf.UpdatedAt,
		Name:          dbf.Name,
		Url:           dbf.Url,
		UserID:        dbf.UserID,
		LastFetchedAt: convertNullTime(dbf.LastFetchedAt),
	}
}

func convertNullTime(sqlTime sql.NullTime) *time.Time {
	if sqlTime.Valid {
		return &sqlTime.Time
	}
	return nil
}
