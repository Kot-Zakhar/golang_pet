package repository

import (
	"time"

	// "golang.org/x/exp/slices"

	"golang_pet/app/model"
)

var inMemoryEvents []model.Event = []model.Event{
	{
		Id:        0,
		Title:     "First",
		CreatedAt: time.Now(),
	},
	{
		Id:        1,
		Title:     "Second",
		CreatedAt: time.Now(),
	},
	{
		Id:        2,
		Title:     "Third",
		CreatedAt: time.Now(),
	},
}

func GetAllEvents() []model.Event {
	return inMemoryEvents
}
