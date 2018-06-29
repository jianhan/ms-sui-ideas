package occupation

import "github.com/micro/go-micro/errors"

func (r *NewOccupationsRequest) Validate() error {
	if len(r.Occupations) == 0 {
		return errors.BadRequest("NewOccupationsRequest", "Empty occupations")
	}

	return nil
}
