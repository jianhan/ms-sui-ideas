package occupation

import (
	fmt "fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro/go-micro/errors"
)

func (r *CreateOccupationsRequest) Validate() error {
	if len(r.Occupations) == 0 {
		return errors.BadRequest("NewOccupationsRequest", "Empty occupations")
	}
	for k := range r.Occupations {
		if _, err := govalidator.ValidateStruct(r.Occupations[k]); err != nil {
			return errors.BadRequest("CreateOccupationsRequest", fmt.Sprintf("Validation failed, %s", err.Error()))
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
			return errors.BadRequest("UpdateOccupationsRequest", fmt.Sprintf("Validation failed, %s", err.Error()))
		}
	}

	return nil
}

func (r *HideOccupationsRequest) Validate() error {
	if len(r.Ids) == 0 {
		return errors.BadRequest("HideOccupationsRequest", "Empty ids")
	}
	for k := range r.Ids {
		if isUUID := govalidator.IsUUID(r.Ids[k]); !isUUID {
			return errors.BadRequest("HideOccupationsRequest", fmt.Sprintf("'%s' is not a valid UUID", r.Ids[k]))
		}
	}

	return nil
}

func (r *ShowOccupationsRequest) Validate() error {
	if len(r.Ids) == 0 {
		return errors.BadRequest("ShowOccupationsRequest", "Empty ids")
	}
	for k := range r.Ids {
		if isUUID := govalidator.IsUUID(r.Ids[k]); !isUUID {
			return errors.BadRequest("ShowOccupationsRequest", fmt.Sprintf("'%s' is not a valid UUID", r.Ids[k]))
		}
	}

	return nil
}
