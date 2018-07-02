package idea

import (
	"strings"

	"github.com/jianhan/pkg/proto"
	"github.com/leebenson/conform"
)

func (r *UpsertIdeasRequest) Preprocess() error {
	for k := range r.Ideas {
		conform.Strings(r.Ideas[k])
		if err := proto.Fix(r.Ideas[k]); err != nil {
			return err
		}
	}

	return nil
}

func (r *IdeaFilter) Preprocess() error {
	for k := range r.Names {
		r.Names[k] = strings.Trim(r.Names[k], " ")
	}
	for k := range r.Ids {
		r.Ids[k] = strings.Trim(r.Ids[k], " ")
	}

	return nil
}
