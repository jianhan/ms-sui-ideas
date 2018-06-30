package mongodb

import (
	"github.com/jianhan/ms-sui-ideas/db"
	"gopkg.in/mgo.v2"

	pstory "github.com/jianhan/ms-sui-ideas/proto/story"
	"gopkg.in/mgo.v2/bson"
)

type story struct {
	base
}

func Story(session *mgo.Session, db, collection string) db.Story {
	c := session.DB(db).C(collection)
	titleIdeaIndex := mgo.Index{
		Key:    []string{"title", "idea_id"},
		Unique: true,
	}
	if err := c.EnsureIndex(titleIdeaIndex); err != nil {
		panic(err)
	}

	hiddenIndex := mgo.Index{
		Key: []string{"hidden"},
	}
	if err := c.EnsureIndex(hiddenIndex); err != nil {
		panic(err)
	}

	b := base{
		session:    session,
		db:         db,
		collection: collection,
	}

	return &story{
		base: b,
	}
}

func (d *story) CreateStories(stories []*pstory.Story) (int64, int64, []*pstory.Story, error) {
	var retVal []*pstory.Story
	bulk := d.session.DB(d.db).C(d.collection).Bulk()
	for _, s := range stories {
		bulk.Insert(s)
		retVal = append(retVal, s)
	}
	matchResult, err := bulk.Run()
	if err != nil {
		return 0, 0, nil, err
	}

	return int64(matchResult.Matched), int64(matchResult.Modified), retVal, nil
}

func (d *story) UpdateStories(stories []*pstory.Story) (int64, int64, []*pstory.Story, error) {
	var retVal []*pstory.Story
	bulk := d.session.DB(d.db).C(d.collection).Bulk()
	for _, s := range stories {
		bulk.Update(
			bson.M{"_id": s.ID},
			bson.M{"$set": s},
		)
		retVal = append(retVal, s)
	}
	matchResult, err := bulk.Run()
	if err != nil {
		return 0, 0, nil, err
	}

	return int64(matchResult.Matched), int64(matchResult.Modified), retVal, nil
}

func (d *story) HideStories(ids []string) (int64, int64, error) {
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

func (d *story) ShowStories(ids []string) (int64, int64, error) {
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

func (d *story) DeleteStories(ids []string) (int64, int64, error) {
	c := d.session.DB(d.db).C(d.collection)
	info, err := c.RemoveAll(bson.M{"$in": ids})
	if err != nil {
		return 0, 0, err
	}

	return int64(info.Matched), int64(info.Removed), nil
}
