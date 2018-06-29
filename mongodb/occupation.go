package mongodb

import (
	"github.com/jianhan/ms-sui-ideas/db"
	poccupation "github.com/jianhan/ms-sui-ideas/proto/occupation"
	"gopkg.in/mgo.v2"
)

type occupation struct {
	base
}

func NewOccupation(session *mgo.Session, db, collection string) db.Occupation {
	c := session.DB(db).C(collection)
	nameIndex := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
	if err := c.EnsureIndex(nameIndex); err != nil {
		panic(err)
	}
	b := base{
		session:    session,
		db:         db,
		collection: collection,
	}
	return &occupation{
		base: b,
	}
}

func (d *occupation) NewOccupations(occupations []*poccupation.Occupation) ([]*poccupation.Occupation, error) {
	return nil, nil
}

func (d *occupation) UpdateOccupations(occupations []*poccupation.Occupation) ([]*poccupation.Occupation, error) {
	return nil, nil
}

func (d *occupation) HideOccupations(ids []string) error {
	return nil
}
