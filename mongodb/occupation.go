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

func (d *occupation) CreateOccupations(occupations []*poccupation.Occupation) (int64, int64, []*poccupation.Occupation, error) {
	var retVal []*poccupation.Occupation
	bulk := d.session.DB(d.db).C(d.collection).Bulk()
	for _, occupation := range occupations {
		bulk.Insert(occupation)
		retVal = append(retVal, occupation)
	}
	matchResult, err := bulk.Run()
	if err != nil {
		return 0, 0, nil, err
	}

	return int64(matchResult.Matched), int64(matchResult.Modified), retVal, nil
}

func (d *occupation) UpdateOccupations(occupations []*poccupation.Occupation) (int64, int64, []*poccupation.Occupation, error) {
	var retVal []*poccupation.Occupation
	bulk := d.session.DB(d.db).C(d.collection).Bulk()
	for _, occupation := range occupations {
		bulk.Update(
			bson.M{"_id": occupation.ID},
			bson.M{"$set": occupation},
		)
		retVal = append(retVal, occupation)
	}
	matchResult, err := bulk.Run()
	if err != nil {
		return 0, 0, nil, err
	}

	return int64(matchResult.Matched), int64(matchResult.Modified), retVal, nil
}

func (d *occupation) HideOccupations(ids []string) (int64, int64, error) {
	bulk := d.session.DB(d.db).C(d.collection).Bulk()
	for _, id := range ids {
		bulk.Update(
			bson.M{"_id": id},
			bson.M{"$set": bson.M{"hidden": true}},
		)
	}
	matchResult, err := bulk.Run()
	if err != nil {
		return 0, 0, err
	}

	return int64(matchResult.Matched), int64(matchResult.Modified), nil
}

func (d *occupation) ShowOccupations(ids []string) (int64, int64, error) {
	bulk := d.session.DB(d.db).C(d.collection).Bulk()
	for _, id := range ids {
		bulk.Update(
			bson.M{"_id": id},
			bson.M{"$set": bson.M{"hidden": false}},
		)
	}
	matchResult, err := bulk.Run()
	if err != nil {
		return 0, 0, err
	}

	return int64(matchResult.Matched), int64(matchResult.Modified), nil
}
