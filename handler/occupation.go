package handler

import (
	"context"

	"github.com/jianhan/ms-sui-ideas/db"
	"github.com/jianhan/ms-sui-ideas/proto/occupation"
	"github.com/nats-io/go-nats-streaming"
)

type Occupation struct {
	db       db.Occupation
	stanConn stan.Conn
}

func NewOccupation(db db.Occupation, stanConn stan.Conn) *Occupation {
	return &Occupation{db: db, stanConn: stanConn}
}

func (h *Occupation) NewOccupations(ctx context.Context, req *occupation.NewOccupationsRequest, rsp *occupation.Occupations) (err error) {
	if rsp.Occupations, err = h.db.NewOccupations(req.Occupations); err != nil {
		return
	}

	return
}

func (h *Occupation) UpdateOccupations(ctx context.Context, req *occupation.UpdateOccupationsRequest, rsp *occupation.Occupations) (err error) {
	if rsp.Occupations, err = h.db.UpdateOccupations(req.Occupations); err != nil {
		return
	}

	return
}

func (h *Occupation) HideOccupations(ctx context.Context, req *occupation.HideOccupationsRequest, rsp *occupation.Occupations) (err error) {
	return
}
