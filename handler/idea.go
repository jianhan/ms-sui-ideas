package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jianhan/ms-sui-ideas/db"
	"github.com/jianhan/ms-sui-ideas/proto/idea"
	"github.com/nats-io/go-nats-streaming"
)

type Idea struct {
	db       db.Idea
	stanConn stan.Conn
}

func NewIdea(db db.Idea, stanConn stan.Conn) *Idea {
	return &Idea{db: db, stanConn: stanConn}
}
func (h *Idea) ListIdeas(ctx context.Context, req *idea.IdeaFilter, rsp *idea.ListIdeasResponse) (err error) {
	if rsp.Ideas, err = h.db.ListIdeas(req); err != nil {
		return err
	}

	return
}

func (h *Idea) CreateIdeas(ctx context.Context, req *idea.UpsertIdeasRequest, rsp *idea.CreateIdeasResponse) (err error) {
	if rsp.Matched, rsp.Modified, rsp.Ideas, err = h.db.CreateIdeas(req.Ideas); err != nil {
		return
	}

	return
}

func (h *Idea) UpdateIdeas(ctx context.Context, req *idea.UpsertIdeasRequest, rsp *idea.UpdateIdeasResponse) (err error) {
	if rsp.Matched, rsp.Modified, rsp.Ideas, err = h.db.UpdateIdeas(req.Ideas); err != nil {
		return
	}

	return
}

func (h *Idea) DeleteIdeas(ctx context.Context, req *idea.IdeaFilter, rsp *idea.DeleteIdeasResponse) (err error) {
	if rsp.Deleted, err = h.db.DeleteIdeas(req); err != nil {
		return
	}

	return
}

func (h *Idea) HideIdeas(ctx context.Context, req *idea.IdeaFilter, rsp *empty.Empty) error {
	return h.db.HideIdeas(req)
}

func (h *Idea) ShowIdeas(ctx context.Context, req *idea.IdeaFilter, rsp *empty.Empty) error {
	return h.db.ShowIdeas(req)
}
