package service

import (
	"golang_pet/app/model"
	"golang_pet/app/repository"
)

func GetEvents() []model.Event {
	events := repository.GetAllEvents()
	// some logic
	return events
}
