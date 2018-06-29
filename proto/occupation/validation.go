package occupation

import (
	"github.com/asaskevich/govalidator"
	"github.com/micro/go-micro/errors"
)

func (r *NewOccupationsRequest) Validate() error {
	if len(r.Occupations) == 0 {
		return errors.BadRequest("NewOccupationsRequest", "Empty occupations")
	}
	for k := range r.Occupations {
		if _, err := govalidator.ValidateStruct(r.Occupations[k]); err != nil {
			return err
		}
	}

	return nil
}

func (r *UpdateOccupationsRequest) Validate() error {
	if len(r.Occupations) == 0 {
		return errors.BadRequest("UpdateOccupationsRequest", "Empty occupations")
	}
	for k := range r.Occupations {
		if _, err := govalidator.ValidateStruct(r.Occupations[k]); err != nil {
			return err
		}
	}

	return nil
}
