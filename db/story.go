package db

import (
	"github.com/jianhan/ms-sui-ideas/proto/story"
)

type Story interface {
	CreateStories(stories []*story.Story) (int64, int64, []*story.Story, error)
	UpdateStories(stories []*story.Story) (int64, int64, []*story.Story, error)
	HideStories(ids []string) (int64, int64, error)
	ShowStories(ids []string) (int64, int64, error)
	DeleteStories(ids []string) (int64, int64, error)
}
