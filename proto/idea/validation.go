package idea

import (
	fmt "fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro/go-micro/errors"
)

func (r *CreateIdeasRequest) Validate() error {
	if len(r.Ideas) == 0 {
		return errors.BadRequest("CreateIdeasRequest", "Empty ideas")
	}

	// validate ideas
	for k := range r.Ideas {
		if _, err := govalidator.ValidateStruct(r.Ideas[k]); err != nil {
			return errors.BadRequest("CreateIdeasRequest", fmt.Sprintf("Validation failed, %s", err.Error()))
		}

		// validate stories
		for sK := range r.Ideas[k].Stories {
			if _, err := govalidator.ValidateStruct(r.Ideas[k].Stories[sK]); err != nil {
				return errors.BadRequest("CreateIdeasRequest", fmt.Sprintf("Validation failed, %s", err.Error()))
			}
		}
	}

	return nil
}
