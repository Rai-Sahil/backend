package events

type EventType string

const (
	AccountOpened	EventType = "AccountOpened"
	MoneyDeposited	EventType = "MoneyDeposited"
	MoneyWithdrawn	EventType = "MoneyWithdrawn"
)

type Event struct {
	ID			string 		`json:"id" gorm:"primaryKey"`
	Type		EventType	`json:"type"`
	AccountID	string		`json:"account_id"`
	Payload		[]byte		`json:"payload"`
	CreatedAt	int64		`json:"created_at" gorm:"autoCreateTime"`
}
