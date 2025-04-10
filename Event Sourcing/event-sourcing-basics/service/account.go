package service

import (
	"encoding/json"
	"time"

	"github.com/Rai-Sahil/backend/events"
	"github.com/Rai-Sahil/backend/store"
	"github.com/google/uuid"
)

type AccountService struct {
	Store	store.EventStore
}

func (s *AccountService) OpenAccount(accountId string) error {
	event := events.Event {
		ID:			uuid.New().String(),
		Type:		events.AccountOpened,
		AccountID: 	accountId,
		Payload:	[]byte(`{"initialBalance: : 0}`),
		CreatedAt: 	time.Now().Unix(),
	}
	return s.Store.Append(event)
}

func (s *AccountService) Deposit(accountId string, amount int64) error {
	payload, _ := json.Marshal(map[string]int64{"amount": amount})
	event := events.Event {
		ID:			uuid.New().String(),
		Type:		events.MoneyDeposited,
		AccountID: 	accountId,
		Payload: 	payload,
		CreatedAt: 	time.Now().Unix(),
	}
	return s.Store.Append(event)
}


func (s *AccountService) Withdraw(accountId string, amount int64) error {
	payload, _ := json.Marshal(map[string]int64{"amount": amount})
	event := events.Event {
		ID:			uuid.New().String(),
		Type:		events.MoneyWithdrawn,
		AccountID: 	accountId,
		Payload: 	payload,
		CreatedAt: 	time.Now().Unix(),
	}
	return s.Store.Append(event)
}

func (s *AccountService) RebuildBalance(accountId string) (int64, error) {
	var balance int64 = 0

	eventStream, err := s.Store.GetEventsForAccount(accountId)
	if err != nil {
		return balance, err 
	}

	for _, evt := range eventStream {
		switch evt.Type {
		case events.MoneyDeposited:
			var payload map[string]int64
			_ = json.Unmarshal(evt.Payload, &payload)
			balance += payload["amount"]
		case events.MoneyWithdrawn:
			var payload map[string]int64
			_ = json.Unmarshal(evt.Payload, &payload)
			balance -= payload["amount"]
		}
	}
	return balance, nil
}

