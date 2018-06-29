package db

import (
	"github.com/jianhan/ms-sui-ideas/proto/occupation"
)

type Occupation interface {
	NewOccupations(occupations []*occupations.Occupation) (*occupations.Occupation, error)
	UpdateOccupations(occupations []*occupations.Occupation) (*occupations.Occupation, error)
	HideOccupations(ids []string) ([]*occupations.Occupation, error)
}
