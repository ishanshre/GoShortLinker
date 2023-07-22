package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	LongUrl string `json:"long_url"`
}

type Url struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UrlCode   string             `json:"url_code,omitempty" bson:"url_code,omitempty"`
	ShortUrl  string             `json:"short_url,omitempty" bson:"short_url,omitempty"`
	LongUrl   string             `json:"long_url,omitempty" bson:"long_url,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	ExpiredAt time.Time          `json:"expired_at,omitempty" bson:"expired_at,omitempty"`
}
