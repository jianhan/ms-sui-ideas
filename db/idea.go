package db

import (
	"github.com/jianhan/ms-sui-ideas/proto/idea"
)

type Idea interface {
	Ideas(ideas []*ideas.Idea) (*ideas.Idea, error)
	UpdateIdeas(ideas []*ideas.Idea) (*ideas.Idea, error)
	DeleteIdeas(ids []string) error
	HideIdeas(ids []string) error
}
