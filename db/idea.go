package db

import (
	"github.com/jianhan/ms-sui-ideas/proto/idea"
)

type Idea interface {
	ListIdeas(filter *idea.IdeaFilter) ([]*idea.Idea, error)
	CreateIdeas(ideas []*idea.Idea) (int64, int64, []*idea.Idea, error)
	UpdateIdeas(ideas []*idea.Idea) (int64, int64, []*idea.Idea, error)
	DeleteIdeas(filter *idea.IdeaFilter) (int64, error)
	HideIdeas(filter *idea.IdeaFilter) error
	ShowIdeas(filter *idea.IdeaFilter) error
	AddRatings(ideaID string, ratings []*idea.Rating) (*idea.Idea, error)
}
