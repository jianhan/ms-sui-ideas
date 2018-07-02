package mongodb

import (
	"strings"

	"github.com/jianhan/ms-sui-ideas/db"
	pidea "github.com/jianhan/ms-sui-ideas/proto/idea"
	"github.com/micro/go-micro/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type idea struct {
	base
}

func Idea(session *mgo.Session, db, collection string) db.Idea {
	c := session.DB(db).C(collection)
	nameIndex := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
	if err := c.EnsureIndex(nameIndex); err != nil {
		panic(err)
	}
	slugIndex := mgo.Index{
		Key:    []string{"slug"},
		Unique: true,
	}
	if err := c.EnsureIndex(slugIndex); err != nil {
		panic(err)
	}
	hiddenIndex := mgo.Index{
		Key: []string{"hidden"},
	}
	if err := c.EnsureIndex(hiddenIndex); err != nil {
		panic(err)
	}
	ageIndex := mgo.Index{
		Key: []string{"min_age", "max_age"},
	}
	if err := c.EnsureIndex(ageIndex); err != nil {
		panic(err)
	}
	textIndex := mgo.Index{
		Key: []string{"$text:name", "$text:description"},
	}
	if err := c.EnsureIndex(textIndex); err != nil {
		panic(err)
	}
	b := base{
		session:    session,
		db:         db,
		collection: collection,
	}
	return &idea{
		base: b,
	}
}

func (d *idea) ListIdeas(filter pidea.IdeaFilter) ([]*pidea.Idea, error) {
	// define query
	query := bson.M{}
	var r []*pidea.Idea

	// set IDs query condition
	if len(filter.Ids) > 0 {
		query["_id"] = bson.M{"$in": filter.Ids}
	}

	// set hidden condition
	query["hidden"] = bson.M{"$eq": filter.Hidden}

	// set search
	if strings.TrimSpace(filter.Query) != "" {
		query["$text"] = bson.M{"$search": filter.Query}
	}

	// sorts
	sorts := []string{"name"}
	if filter.Sorts != nil {
		sorts = filter.Sorts
	}

	// set pagination
	currentPage := 1
	if filter.CurrentPage > 0 {
		currentPage = int(filter.CurrentPage)
	}
	perPage := 20
	if filter.PerPage > 0 {
		perPage = int(filter.PerPage)
	}

	// get results
	if err := d.session.DB(d.db).C(d.collection).Find(query).Sort(sorts...).Skip(perPage * (currentPage - 1)).Limit(perPage).All(&r); err != nil {
		return nil, err
	}

	return r, nil
}

func (d *idea) CreateIdeas(ideas []*pidea.Idea) (int64, int64, []*pidea.Idea, error) {
	bulk := d.session.DB(d.db).C(d.collection).Bulk()
	var retVal []*pidea.Idea
	for _, v := range ideas {
		bulk.Insert(v)
		retVal = append(retVal, v)
	}
	r, err := bulk.Run()
	if err != nil {
		return 0, 0, nil, err
	}

	return int64(r.Matched), int64(r.Modified), retVal, nil
}

func (d *idea) UpdateIdeas(ideas []*pidea.Idea) (int64, int64, []*pidea.Idea, error) {
	bulk := d.session.DB(d.db).C(d.collection).Bulk()
	var retVal []*pidea.Idea
	for _, v := range ideas {
		bulk.Update(bson.M{"_id": v.ID}, v)
		retVal = append(retVal, v)
	}
	r, err := bulk.Run()
	if err != nil {
		return 0, 0, nil, err
	}

	return int64(r.Matched), int64(r.Modified), retVal, nil
}

func (d *idea) DeleteIdeas(filter pidea.IdeaFilter) (deleted int64, err error) {
	ideas, err := d.ListIdeas(filter)
	if err != nil {
		return
	}

	// construct ids
	var ids []string
	for k := range ideas {
		ids = append(ids, ideas[k].ID)
	}
	if len(ids) == 0 {
		return 0, errors.BadRequest("DeleteIdeas", "can not find any records to delete")
	}

	// delete and generate results
	deleted = int64(len(ids))
	if err = d.session.DB(d.db).C(d.collection).Remove(bson.M{"_id": bson.M{"$in": ids}}); err != nil {
		return 0, err
	}

	return
}
