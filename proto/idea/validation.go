package idea

import (
	"fmt"

	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/micro/go-micro/errors"
)

func isValidAppType(t int32) bool {
	return AppType_name[t] != ""
}

func (r *UpsertIdeasRequest) Validate() error {
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
		for aK := range r.Ideas[k].AppTypes {
			if int(r.Ideas[k].AppTypes[aK]) >= len(AppType_name) {
				return errors.BadRequest("CreateIdeasRequest", fmt.Sprintf("Invalid app type %s", r.Ideas[k].AppTypes[aK]))
			}
		}
	}

	return nil
}

func (r *IdeaFilter) Validate() error {
	for k := range r.Ids {
		if !govalidator.IsUUID(r.Ids[k]) {
			return errors.BadRequest("IdeaFilter", fmt.Sprintf("'%s' must be a valid UUID", r.Ids[k]))
		}
	}
	for k := range r.Names {
		if strings.Trim(r.Names[k], " ") == "" {
			return errors.BadRequest("Validate", "Name must not be blank")
		}
	}

	return nil
}

func (r *AddRatingsRequest) Validate() error {
	if !govalidator.IsUUID(r.IdeaId) {
		return errors.BadRequest("AddRatingsRequest", fmt.Sprintf("Idea ID '%s' must be a valid UUID", r.IdeaId))
	}
	if len(r.Ratings) == 0 {
		return errors.BadRequest("AddRatingsRequest", "Ratings can not be empty")
	}
	for k := range r.Ratings {
		if _, err := govalidator.ValidateStruct(r.Ratings[k]); err != nil {
			return errors.BadRequest("AddRatingsRequest", err.Error())
		}
	}

	return nil
}
