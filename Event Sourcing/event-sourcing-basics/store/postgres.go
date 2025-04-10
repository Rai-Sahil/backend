package store

import (
	"github.com/Rai-Sahil/backend/events"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type EventStore struct {
	DB *gorm.DB
}

func NewEventStore(dsn string) (*EventStore, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&events.Event{})
	if err != nil {
		return nil, err
	}

	return &EventStore{DB: db}, nil
}

func (s *EventStore) Append(event events.Event) error {
	return s.DB.Create(&event).Error
}

func (s *EventStore) GetEventsForAccount(accountId string) ([]events.Event, error) {
	var eventsList []events.Event
	err := s.DB.Where("account_id = ?", accountId).Order("created_at ASC").Find(&eventsList).Error
	return eventsList, err
}
