package story

import (
	fmt "fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro/go-micro/errors"
)

func (r *CreateStoriesRequest) Validate() error {
	if len(r.Stories) == 0 {
		return errors.BadRequest("CreateStoriesRequest", "Empty stories")
	}

	for k := range r.Stories {
		if _, err := govalidator.ValidateStruct(r.Stories[k]); err != nil {
			return errors.BadRequest("CreateStoriesRequest", fmt.Sprintf("Validation failed, %s", err.Error()))
		}
	}
}
