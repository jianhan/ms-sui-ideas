package db

import (
	"github.com/jianhan/ms-sui-ideas/proto/occupation"
)

type Occupation interface {
	NewOccupations(occupations []*occupation.Occupation) (*occupation.Occupation, error)
	UpdateOccupations(occupations []*occupation.Occupation) (*occupation.Occupation, error)
	HideOccupations(ids []string) ([]*occupation.Occupation, error)
}
