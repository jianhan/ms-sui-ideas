package idea

import (
	"github.com/jianhan/pkg/proto"
	"github.com/leebenson/conform"
)

func (r *CreateIdeasRequest) Preprocess() error {
	for k := range r.Ideas {
		conform.Strings(r.Ideas[k])
		if err := proto.Fix(r.Ideas[k]); err != nil {
			return err
		}
	}

	return nil
}
