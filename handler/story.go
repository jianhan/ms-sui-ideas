package handler

import (
	"context"

	"github.com/jianhan/ms-sui-ideas/db"
	pstory "github.com/jianhan/ms-sui-ideas/proto/story"
	"github.com/nats-io/go-nats-streaming"
)

type Story struct {
	db       db.Story
	stanConn stan.Conn
}

func NewStory(db db.Story, stanConn stan.Conn) *Story {
	return &Story{db: db, stanConn: stanConn}
}

func (h *Story) CreateStories(ctx context.Context, req *pstory.CreateStoriesRequest, rsp *pstory.CreateStoriesResponse) (err error) {
	if rsp.Matched, rsp.Modified, rsp.Stories, err = h.db.CreateStories(req.Stories); err != nil {
		return
	}

	return
}

func (h *Story) UpdateStories(ctx context.Context, req *pstory.UpdateStoriesRequest, rsp *pstory.UpdateStoriesResponse) (err error) {
	if rsp.Matched, rsp.Modified, rsp.Stories, err = h.db.UpdateStories(req.Stories); err != nil {
		return
	}

	return
}

func (h *Story) HideStories(ctx context.Context, req *pstory.HideStoriesRequest, rsp *pstory.HideStoriesResponse) (err error) {
	if rsp.Matched, rsp.Modified, err = h.db.HideStories(req.Ids); err != nil {
		return
	}

	return
}

func (h *Story) ShowStories(ctx context.Context, req *pstory.ShowStoriesRequest, rsp *pstory.ShowIdeasResponse) (err error) {
	if rsp.Matched, rsp.Modified, err = h.db.ShowStories(req.Ids); err != nil {
		return
	}

	return
}

func (h *Story) DeleteStories(ctx context.Context, req *pstory.DeleteStoriesRequest, rsp *pstory.DeleteStoriesResponse) (err error) {
	if rsp.Matched, rsp.Removed, err = h.db.DeleteStories(req.Ids); err != nil {
		return
	}

	return
}
