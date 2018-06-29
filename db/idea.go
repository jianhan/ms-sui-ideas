package db

import (
	"github.com/jianhan/ms-sui-ideas/proto/idea"
)

type Idea interface {
	Ideas(ideas []*idea.Idea) ([]*idea.Idea, error)
	UpdateIdeas(ideas []*idea.Idea) ([]*idea.Idea, error)
	DeleteIdeas(ids []string) error
	HideIdeas(ids []string) ([]*idea.Idea, error)
}
