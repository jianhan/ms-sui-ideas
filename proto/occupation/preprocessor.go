package occupation

import (
	"github.com/jianhan/pkg/proto"
	"github.com/leebenson/conform"
)

func (r *CreateOccupationsRequest) Preprocess() error {
	for k := range r.Occupations {
		conform.Strings(r.Occupations[k])
		if err := proto.Fix(r.Occupations[k]); err != nil {
			return err
		}
	}

	return nil
}

func (r *UpdateOccupationsRequest) Preprocess() error {
	for k := range r.Occupations {
		conform.Strings(r.Occupations[k])
		if err := proto.Fix(r.Occupations[k]); err != nil {
			return err
		}
	}

	return nil
}
