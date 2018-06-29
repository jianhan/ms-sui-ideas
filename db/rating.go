package db

import (
	"github.com/jianhan/ms-sui-ideas/proto/rating"
)

type Rating interface {
	NewRatings(ratings []*rating.Rating) ([]*rating.Rating, error)
	UpdateRatings(ratings []*rating.Rating) ([]*rating.Rating, error)
	DeleteRatings(ids []string) error
}
