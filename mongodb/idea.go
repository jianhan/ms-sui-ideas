package mongodb

import (
	"strings"

	"github.com/jianhan/ms-sui-ideas/db"
	pidea "github.com/jianhan/ms-sui-ideas/proto/idea"
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
