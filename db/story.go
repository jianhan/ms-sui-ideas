package db

import (
	"github.com/jianhan/ms-sui-ideas/proto/story"
)

type Story interface {
	NewStories(stories []*story.Story) (*story.Story, error)
	UpdateStories(stories []*story.Story) ([]*story.Story, error)
	DeleteStories(ids []string) error
}
