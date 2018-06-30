package db

import (
	"github.com/jianhan/ms-sui-ideas/proto/occupation"
)

type Occupation interface {
	CreateOccupations(occupations []*occupation.Occupation) (int64, int64, []*occupation.Occupation, error)
	UpdateOccupations(occupations []*occupation.Occupation) (int64, int64, []*occupation.Occupation, error)
	HideOccupations(ids []string) (int64, int64, error)
	ShowOccupations(ids []string) (int64, int64, error)
}
