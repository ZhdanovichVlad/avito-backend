package models

import (
	"time"
)

type FeedBack struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"bidId,omitempty"`
	Status      string    `json:"status,omitempty"`
	BidId       string    `json:"bidId,omitempty"`
	BidFeedback string    `json:"BidFeedback,omitempty"`
	UserName    string    `json:"UserName,omitempty"`
	Version     int       `json:"version,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

//id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
//bidId UUID REFERENCES bids(id) ON DELETE CASCADE,
//name VARCHAR(50) NOT NULL,
//status bids_status NOT NULL,
//authorType author_types NOT NULL,
//authorId VARCHAR(50) NOT NULL,
//bidFeedback VARCHAR(1000) NOT NULL,
//username VARCHAR(50) NOT NULL,
//version INT NOT NULL,
//createdAt TIMESTAMP NOT NULL
