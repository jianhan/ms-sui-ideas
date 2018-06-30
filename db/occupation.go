package db

import (
	"github.com/jianhan/ms-sui-ideas/proto/occupation"
)

type Occupation interface {
	NewOccupations(occupations []*occupation.Occupation) (int, int, []*occupation.Occupation, error)
	UpdateOccupations(occupations []*occupation.Occupation) (int, int, []*occupation.Occupation, error)
	HideOccupations(ids []string) (int, int, error)
	ShowOccupations(ids []string) (int, int, error)
}
