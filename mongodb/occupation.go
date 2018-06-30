package mongodb

import (
	"github.com/jianhan/ms-sui-ideas/db"
	poccupation "github.com/jianhan/ms-sui-ideas/proto/occupation"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type occupation struct {
	base
}

func Occupation(session *mgo.Session, db, collection string) db.Occupation {
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
	retVal := []*poccupation.Occupation{}
	bulk := d.session.DB(d.db).C(d.collection).Bulk()
	for _, occupation := range occupations {
		bulk.Insert(occupation)
		retVal = append(retVal, occupation)
	}
	_, err := bulk.Run()
	if err != nil {
		return nil, err
	}

	return retVal, nil
}

func (d *occupation) UpdateOccupations(occupations []*poccupation.Occupation) ([]*poccupation.Occupation, error) {
	retVal := []*poccupation.Occupation{}
	bulk := d.session.DB(d.db).C(d.collection).Bulk()
	for _, occupation := range occupations {
		bulk.Update(
			bson.M{"_id": occupation.ID},
			bson.M{"$set": occupation},
		)
		retVal = append(retVal, occupation)
	}
	_, err := bulk.Run()
	if err != nil {
		return nil, err
	}

	return retVal, nil
}

func (d *occupation) HideOccupations(ids []string) error {
	return nil
}
